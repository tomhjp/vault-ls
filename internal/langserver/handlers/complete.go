package handlers

import (
	"context"

	lsctx "github.com/tomhjp/vault-ls/internal/context"
	ilsp "github.com/tomhjp/vault-ls/internal/lsp"
	lsp "github.com/tomhjp/vault-ls/internal/protocol"
)

func (svc *service) TextDocumentComplete(ctx context.Context, params lsp.CompletionParams) (lsp.CompletionList, error) {
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

	decoder, err := svc.decoderForDocument(doc)
	if err != nil {
		return list, err
	}

	fPos, err := ilsp.FilePositionFromDocumentPosition(params.TextDocumentPositionParams, doc)
	if err != nil {
		return list, err
	}

	svc.logger.Printf("Looking for candidates at %q -> %#v", doc.Filename(), fPos.Position())
	candidates, err := decoder.CandidatesAtPos(doc.Filename(), fPos.Position())
	svc.logger.Printf("received candidates: %#v", candidates)
	return ilsp.ToCompletionList(candidates, clientCaps.TextDocument), err
}
