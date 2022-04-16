package context

import (
	"context"

	"github.com/tomhjp/vault-ls/internal/filesystem"
	lsp "github.com/tomhjp/vault-ls/internal/protocol"
)

type contextKey struct {
	Name string
}

func (k *contextKey) String() string {
	return k.Name
}

var (
	ctxDs                   = &contextKey{"document storage"}
	ctxLsVersion            = &contextKey{"language server version"}
	ctxProgressToken        = &contextKey{"progress token"}
	ctxExperimentalFeatures = &contextKey{"experimental features"}
)

func missingContextErr(ctxKey *contextKey) *MissingContextErr {
	return &MissingContextErr{ctxKey}
}

func WithDocumentStorage(ctx context.Context, fs filesystem.DocumentStorage) context.Context {
	return context.WithValue(ctx, ctxDs, fs)
}

func DocumentStorage(ctx context.Context) (filesystem.DocumentStorage, error) {
	fs, ok := ctx.Value(ctxDs).(filesystem.DocumentStorage)
	if !ok {
		return nil, missingContextErr(ctxDs)
	}

	return fs, nil
}

func WithLanguageServerVersion(ctx context.Context, version string) context.Context {
	return context.WithValue(ctx, ctxLsVersion, version)
}

func LanguageServerVersion(ctx context.Context) (string, bool) {
	version, ok := ctx.Value(ctxLsVersion).(string)
	if !ok {
		return "", false
	}
	return version, true
}

func WithProgressToken(ctx context.Context, pt lsp.ProgressToken) context.Context {
	return context.WithValue(ctx, ctxProgressToken, pt)
}

func ProgressToken(ctx context.Context) (lsp.ProgressToken, bool) {
	pt, ok := ctx.Value(ctxProgressToken).(lsp.ProgressToken)
	if !ok {
		return "", false
	}
	return pt, true
}
