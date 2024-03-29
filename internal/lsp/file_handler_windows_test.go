package lsp

import (
	"testing"

	lsp "github.com/tomhjp/vault-ls/internal/protocol"
)

var (
	validWindowsPath = "file:///C:/Users/With%20Space/vault-test/file.vault-server"
)

func TestFileHandler_valid_windows(t *testing.T) {
	path := "file:///C:/Users/With%20Space/vault-test/file.vault-server"
	fh := FileHandlerFromDocumentURI(lsp.DocumentURI(path))
	if !fh.Valid() {
		t.Fatalf("Expected %q to be valid", path)
	}

	expectedDir := `C:\Users\With Space\vault-test`
	if fh.Dir() != expectedDir {
		t.Fatalf("Expected dir: %q, given: %q",
			expectedDir, fh.Dir())
	}

	expectedFilename := "file.vault-server"
	if fh.Filename() != expectedFilename {
		t.Fatalf("Expected filename: %q, given: %q",
			expectedFilename, fh.Filename())
	}

	expectedFullPath := `C:\Users\With Space\vault-test\file.vault-server`
	if fh.FullPath() != expectedFullPath {
		t.Fatalf("Expected full path: %q, given: %q",
			expectedFullPath, fh.FullPath())
	}

	if fh.URI() != validWindowsPath {
		t.Fatalf("Expected document URI: %q, given: %q",
			validWindowsPath, fh.URI())
	}
}
