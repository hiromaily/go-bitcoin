package wallets

import (
	"github.com/hiromaily/go-bitcoin/pkg/account"
	"github.com/hiromaily/go-bitcoin/pkg/wallet/key"
)

// Keygener is for keygen wallet service interface
type Keygener interface {
	GenerateSeed() ([]byte, error)
	StoreSeed(strSeed string) ([]byte, error)
	GenerateAccountKey(accountType account.AccountType, seed []byte, count uint32) ([]key.WalletKey, error)
	ImportPrivKey(accountType account.AccountType) error
	ImportFullPubKey(fileName string) error
	CreateMultisigAddress(accountType account.AccountType) error
	ExportAddress(accountType account.AccountType) (string, error)
	SignTx(filePath string) (string, bool, string, error)

	Done()
}
