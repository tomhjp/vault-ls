package decoder

import (
	"context"
	"fmt"

	"github.com/hashicorp/hcl-lang/decoder"
	"github.com/hashicorp/hcl-lang/lang"
	"github.com/hashicorp/hcl-lang/reference"
	"github.com/hashicorp/hcl-lang/schema"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/tomhjp/vault-ls/internal/filesystem"
	ilsp "github.com/tomhjp/vault-ls/internal/lsp"
	"github.com/tomhjp/vault-ls/internal/schemas"
)

type PathReader struct {
	FS filesystem.DocumentStorage
}

var _ decoder.PathReader = &PathReader{}

func (pr *PathReader) Paths(_ context.Context) []lang.Path {
	paths := make([]lang.Path, 0)
	return paths
}

func (pr *PathReader) PathContext(path lang.Path) (*decoder.PathContext, error) {
	var schema *schema.BodySchema
	switch path.LanguageID {
	case ilsp.Vault.String():
		schema = schemas.ServerSchema
	case ilsp.VaultAgent.String():
		schema = schemas.AgentSchema
	case ilsp.VaultPolicy.String():
		schema = schemas.PolicySchema
	default:
		return nil, fmt.Errorf("unknown language ID: %q", path.LanguageID)
	}

	doc, err := pr.FS.GetDocument(ilsp.FileHandlerFromPath(path.Path))
	if err != nil {
		return nil, err
	}

	docBytes, err := doc.Text()
	if err != nil {
		return nil, err
	}
	hclFile, hclDiagnostics := hclparse.NewParser().ParseHCL(docBytes, doc.Filename())
	if hclDiagnostics.HasErrors() {
		return nil, hclDiagnostics.Errs()[0]
	}

	return &decoder.PathContext{
		Schema:           schema,
		ReferenceOrigins: make(reference.Origins, 0),
		ReferenceTargets: make(reference.Targets, 0),
		Files: map[string]*hcl.File{
			doc.Filename(): hclFile,
		},
	}, nil
}
