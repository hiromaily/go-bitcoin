package key

import (
	"github.com/hiromaily/go-crypto-wallet/pkg/account"
)

// Generator is key generator interface
type Generator interface {
	CreateKey(seed []byte, actType account.AccountType, idxFrom, count uint32) ([]WalletKey, error)
}

// WalletKey keys
// - [BTC] P2PKHAddr is not used anywhere, P2SHSegWitAddr should be used.
// - [BCH] P2SHSegWitAddr is invalid. P2PKHAddr should be used.
type WalletKey struct {
	WIF            string
	P2PKHAddr      string
	P2SHSegWitAddr string
	Bech32Addr     string
	FullPubKey     string
	RedeemScript   string
}
