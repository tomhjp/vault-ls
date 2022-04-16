package handlers

import (
	"context"
	"fmt"

	lsctx "github.com/tomhjp/vault-ls/internal/context"
	ilsp "github.com/tomhjp/vault-ls/internal/lsp"
	lsp "github.com/tomhjp/vault-ls/internal/protocol"
)

func TextDocumentComplete(ctx context.Context, params lsp.CompletionParams) (lsp.CompletionList, error) {
	var list lsp.CompletionList

	fs, err := lsctx.DocumentStorage(ctx)
	if err != nil {
		return list, err
	}

	clientCaps, err := ilsp.ClientCapabilities(ctx)
	if err != nil {
		return list, err
	}

	doc, err := fs.GetDocument(ilsp.FileHandlerFromDocumentURI(params.TextDocument.URI))
	if err != nil {
		return list, err
	}

	decoder, err := decoderForDocument(ctx, doc)
	if err != nil {
		return list, err
	}

	fPos, err := ilsp.FilePositionFromDocumentPosition(params.TextDocumentPositionParams, doc)
	if err != nil {
		return list, err
	}

	fmt.Printf("Looking for candidates at %q -> %#v\n", doc.Filename(), fPos.Position())
	candidates, err := decoder.CandidatesAtPos(doc.Filename(), fPos.Position())
	fmt.Printf("received candidates: %#v\n", candidates)
	return ilsp.ToCompletionList(candidates, clientCaps.TextDocument), err
}
