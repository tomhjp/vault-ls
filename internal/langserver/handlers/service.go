package handlers

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/code"
	rpch "github.com/creachadair/jrpc2/handler"
	"github.com/hashicorp/hcl-lang/decoder"
	"github.com/hashicorp/hcl-lang/lang"
	lsctx "github.com/tomhjp/vault-ls/internal/context"
	idecoder "github.com/tomhjp/vault-ls/internal/decoder"
	"github.com/tomhjp/vault-ls/internal/filesystem"
	"github.com/tomhjp/vault-ls/internal/langserver/session"
	ilsp "github.com/tomhjp/vault-ls/internal/lsp"
	lsp "github.com/tomhjp/vault-ls/internal/protocol"
)

type service struct {
	logger *log.Logger

	srvCtx context.Context

	sessCtx     context.Context
	stopSession context.CancelFunc

	fs      filesystem.Filesystem
	decoder *decoder.Decoder

	additionalHandlers map[string]rpch.Func
}

var discardLogs = log.New(ioutil.Discard, "", 0)

func NewSession(srvCtx context.Context) session.Session {
	fs := filesystem.NewFilesystem()

	sessCtx, stopSession := context.WithCancel(srvCtx)
	return &service{
		logger:      discardLogs,
		fs:          fs,
		srvCtx:      srvCtx,
		sessCtx:     sessCtx,
		stopSession: stopSession,
	}
}

func (svc *service) SetLogger(logger *log.Logger) {
	svc.logger = logger
}

// Assigner builds out the jrpc2.Map according to the LSP protocol
// and passes related dependencies to handlers via context
func (svc *service) Assigner() (jrpc2.Assigner, error) {
	svc.logger.Println("Preparing new session ...")

	session := session.NewSession(svc.stopSession)

	err := session.Prepare()
	if err != nil {
		return nil, fmt.Errorf("Unable to prepare session: %w", err)
	}

	svc.fs.SetLogger(svc.logger)

	cc := &lsp.ClientCapabilities{}

	clientName := ""

	m := map[string]rpch.Func{
		"initialize": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := session.Initialize(req)
			if err != nil {
				return nil, err
			}

			ctx = ilsp.WithClientCapabilities(ctx, cc)
			ctx = ilsp.ContextWithClientName(ctx, &clientName)
			ctx = lsctx.WithDocumentStorage(ctx, svc.fs)

			version, ok := lsctx.LanguageServerVersion(svc.srvCtx)
			if ok {
				ctx = lsctx.WithLanguageServerVersion(ctx, version)
			}

			return handle(ctx, req, Initialize)
		},
		"initialized": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := session.ConfirmInitialization(req)
			if err != nil {
				return nil, err
			}

			return handle(ctx, req, Initialized)
		},
		"textDocument/didChange": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := session.CheckInitializationIsConfirmed()
			if err != nil {
				return nil, err
			}
			ctx = lsctx.WithDocumentStorage(ctx, svc.fs)
			return handle(ctx, req, TextDocumentDidChange)
		},
		"textDocument/didOpen": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := session.CheckInitializationIsConfirmed()
			if err != nil {
				return nil, err
			}
			ctx = lsctx.WithDocumentStorage(ctx, svc.fs)
			return handle(ctx, req, TextDocumentDidOpen)
		},
		"textDocument/didClose": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := session.CheckInitializationIsConfirmed()
			if err != nil {
				return nil, err
			}
			ctx = lsctx.WithDocumentStorage(ctx, svc.fs)
			return handle(ctx, req, TextDocumentDidClose)
		},
		"textDocument/completion": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := session.CheckInitializationIsConfirmed()
			if err != nil {
				return nil, err
			}

			ctx = lsctx.WithDocumentStorage(ctx, svc.fs)
			ctx = ilsp.WithClientCapabilities(ctx, cc)

			return handle(ctx, req, TextDocumentComplete)
		},
		"shutdown": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := session.Shutdown(req)
			if err != nil {
				return nil, err
			}
			ctx = lsctx.WithDocumentStorage(ctx, svc.fs)
			return handle(ctx, req, Shutdown)
		},
		"exit": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := session.Exit()
			if err != nil {
				return nil, err
			}

			svc.stopSession()

			return nil, nil
		},
		"$/cancelRequest": func(ctx context.Context, req *jrpc2.Request) (interface{}, error) {
			err := session.CheckInitializationIsConfirmed()
			if err != nil {
				return nil, err
			}

			return handle(ctx, req, CancelRequest)
		},
	}

	// For use in tests, e.g. to test request cancellation
	if len(svc.additionalHandlers) > 0 {
		for methodName, handlerFunc := range svc.additionalHandlers {
			m[methodName] = handlerFunc
		}
	}

	return convertMap(m), nil
}

func (svc *service) Finish(_ jrpc2.Assigner, status jrpc2.ServerStatus) {
	if status.Closed || status.Err != nil {
		svc.logger.Printf("session stopped unexpectedly (err: %v)", status.Err)
	}

	svc.stopSession()
}

// convertMap is a helper function allowing us to omit the jrpc2.Func
// signature from the method definitions
func convertMap(m map[string]rpch.Func) rpch.Map {
	hm := make(rpch.Map, len(m))

	for method, fun := range m {
		hm[method] = rpch.New(fun)
	}

	return hm
}

const requestCancelled code.Code = -32800

// handle calls a jrpc2.Func compatible function
func handle(ctx context.Context, req *jrpc2.Request, fn interface{}) (interface{}, error) {
	f := rpch.New(fn)
	result, err := f.Handle(ctx, req)
	if ctx.Err() != nil && errors.Is(ctx.Err(), context.Canceled) {
		err = fmt.Errorf("%w: %s", requestCancelled.Err(), err)
	}
	return result, err
}

func decoderForDocument(ctx context.Context, doc filesystem.Document) (*decoder.PathDecoder, error) {
	fs, err := lsctx.DocumentStorage(ctx)
	if err != nil {
		return nil, err
	}
	decoder := idecoder.NewDecoder(ctx, &idecoder.PathReader{
		FS: fs,
	})

	return decoder.Path(lang.Path{
		Path:       doc.FullPath(),
		LanguageID: doc.LanguageID(),
	})
}
