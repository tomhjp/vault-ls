package lsp

// LanguageID represents the coding language
// of a file
type LanguageID string

const (
	Vault       LanguageID = "vault-server"
	VaultAgent  LanguageID = "vault-agent"
	VaultPolicy LanguageID = "vault-policy"
)

func (l LanguageID) String() string {
	return string(l)
}
