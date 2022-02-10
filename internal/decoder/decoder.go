package decoder

import (
	"context"

	"github.com/hashicorp/hcl-lang/decoder"
	ilsp "github.com/tomhjp/vault-ls/internal/lsp"
)

func NewDecoder(ctx context.Context, pathReader decoder.PathReader) *decoder.Decoder {
	d := decoder.NewDecoder(pathReader)
	d.SetContext(decoderContext(ctx))
	return d
}

func decoderContext(ctx context.Context) decoder.DecoderContext {
	dCtx := decoder.DecoderContext{
		UtmSource:     "vault-ls",
		UseUtmContent: true,
	}

	clientName, ok := ilsp.ClientName(ctx)
	if ok {
		dCtx.UtmMedium = clientName
	}
	return dCtx
}
