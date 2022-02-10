package uri

import (
	"testing"
)

func TestFromPath(t *testing.T) {
	path := `C:\Users\With Space\file.vault-server`
	uri := FromPath(path)

	expectedURI := "file:///C:/Users/With%20Space/file.vault-server"
	if uri != expectedURI {
		t.Fatalf("URI doesn't match.\nExpected: %q\nGiven: %q",
			expectedURI, uri)
	}
}

func TestPathFromURI_valid_windowsFile(t *testing.T) {
	uri := "file:///C:/Users/With%20Space/vault-test/file.vault-server"
	if !IsURIValid(uri) {
		t.Fatalf("Expected %q to be valid", uri)
	}

	expectedPath := `C:\Users\With Space\vault-test\file.vault-server`
	path, err := PathFromURI(uri)
	if err != nil {
		t.Fatal(err)
	}
	if path != expectedPath {
		t.Fatalf("Expected full path: %q, given: %q",
			expectedPath, path)
	}
}

func TestPathFromURI_valid_windowsDir(t *testing.T) {
	uri := "file:///C:/Users/With%20Space/vault-test"
	if !IsURIValid(uri) {
		t.Fatalf("Expected %q to be valid", uri)
	}

	expectedPath := `C:\Users\With Space\vault-test`
	path, err := PathFromURI(uri)
	if err != nil {
		t.Fatal(err)
	}
	if path != expectedPath {
		t.Fatalf("Expected full path: %q, given: %q",
			expectedPath, path)
	}
}
