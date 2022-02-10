package lsp

import (
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/tomhjp/vault-ls/internal/mdplain"
	lsp "github.com/tomhjp/vault-ls/internal/protocol"
)

func markupContent(content lang.MarkupContent, mdSupported bool) lsp.MarkupContent {
	value := content.Value

	kind := lsp.PlainText
	if content.Kind == lang.MarkdownKind {
		if mdSupported {
			kind = lsp.Markdown
		} else {
			value = mdplain.Clean(value)
		}
	}

	return lsp.MarkupContent{
		Kind:  kind,
		Value: value,
	}
}
