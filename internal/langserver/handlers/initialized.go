package handlers

import (
	"context"

	lsp "github.com/tomhjp/vault-ls/internal/protocol"
)

func Initialized(ctx context.Context, params lsp.InitializedParams) error {
	return nil
}
