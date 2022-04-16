package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"syscall"

	lsctx "github.com/tomhjp/vault-ls/internal/context"
	"github.com/tomhjp/vault-ls/internal/filesystem"
	"github.com/tomhjp/vault-ls/internal/langserver"
	"github.com/tomhjp/vault-ls/internal/langserver/handlers"
	"github.com/tomhjp/vault-ls/internal/logging"
	ilsp "github.com/tomhjp/vault-ls/internal/lsp"
	lsp "github.com/tomhjp/vault-ls/internal/protocol"
)

type WASMCommand struct {
	Version string
}

func (c *WASMCommand) flags() *flag.FlagSet {
	fs := defaultFlagSet("wasm")

	fs.Usage = func() {
		_, _ = fmt.Fprint(os.Stdout, c.Help())
	}

	return fs
}

func (c *WASMCommand) Run(_ []string) error {
	// f := c.flags()
	// if err := f.Parse(args); err != nil {
	// 	return fmt.Errorf("Error parsing command-line flags: %s", err)
	// }

	logger := logging.NewLogger(os.Stderr)
	ctx, cancelFunc := lsctx.WithSignalCancel(context.Background(), logger, syscall.SIGINT, syscall.SIGTERM)
	defer cancelFunc()

	logger.Printf("Starting vault-ls %s", c.Version)

	cc := &lsp.ClientCapabilities{}
	cc.TextDocument.Completion.CompletionItem.SnippetSupport = true
	fs := filesystem.NewFilesystem()

	ctx = lsctx.WithLanguageServerVersion(ctx, c.Version)
	ctx = lsctx.WithDocumentStorage(ctx, fs)
	ctx = ilsp.WithClientCapabilities(ctx, cc)

	srv := langserver.NewLangServer(ctx, handlers.NewSession)
	srv.SetLogger(logger)

	// err := srv.StartAndWait(os.Stdin, os.Stdout)
	// if err != nil {
	// 	return fmt.Errorf("Failed to start server: %s", err)
	// }
	w := wasmService{
		ctx: ctx,
	}

	registerGlobals(&w)

	<-make(chan struct{})

	return nil
}

type wasmService struct {
	ctx context.Context
}

func (w *wasmService) complete(doc lsp.TextDocumentItem, position lsp.Position) (lsp.CompletionList, error) {
	emptyList := lsp.CompletionList{
		IsIncomplete: true,
		Items:        []lsp.CompletionItem{},
	}

	items, err := handlers.TextDocumentComplete(w.ctx, lsp.CompletionParams{
		TextDocumentPositionParams: lsp.TextDocumentPositionParams{
			TextDocument: lsp.TextDocumentIdentifier{
				URI: lsp.DocumentURI(doc.URI),
			},
			Position: position,
		},
	})

	if err != nil {
		return emptyList, err
	}
	return items, nil
}

func (w *wasmService) didOpen(doc lsp.TextDocumentItem) error {
	return handlers.TextDocumentDidOpen(w.ctx, lsp.DidOpenTextDocumentParams{
		TextDocument: doc,
	})
}

func (w *wasmService) didChange(uri lsp.DocumentURI, version int32, changes []lsp.TextDocumentContentChangeEvent) error {
	return handlers.TextDocumentDidChange(w.ctx, lsp.DidChangeTextDocumentParams{
		TextDocument: lsp.VersionedTextDocumentIdentifier{
			TextDocumentIdentifier: lsp.TextDocumentIdentifier{
				URI: uri,
			},
			Version: version,
		},
		ContentChanges: changes,
	})
}

func (w *wasmService) didClose(uri lsp.DocumentURI) error {
	return handlers.TextDocumentDidClose(w.ctx, lsp.DidCloseTextDocumentParams{
		TextDocument: lsp.TextDocumentIdentifier{
			URI: uri,
		},
	})
}

func (c *WASMCommand) Help() string {
	helpText := `
Usage: vault-ls wasm [options]

` + c.Synopsis() + "\n\n" + helpForFlags(c.flags())

	return strings.TrimSpace(helpText)
}

func (c *WASMCommand) Synopsis() string {
	return "Starts the Language Server in WASM mode"
}
