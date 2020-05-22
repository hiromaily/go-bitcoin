// Code generated by SQLBoiler 3.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"strconv"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// M type is for providing columns and column values to UpdateAll.
type M map[string]interface{}

// ErrSyncFail occurs during insert when the record could not be retrieved in
// order to populate default value information. This usually happens when LastInsertId
// fails or there was a primary key configuration that was not resolvable.
var ErrSyncFail = errors.New("models: failed to synchronize data after insert")

type insertCache struct {
	query        string
	retQuery     string
	valueMapping []uint64
	retMapping   []uint64
}

type updateCache struct {
	query        string
	valueMapping []uint64
}

func makeCacheKey(cols boil.Columns, nzDefaults []string) string {
	buf := strmangle.GetBuffer()

	buf.WriteString(strconv.Itoa(cols.Kind))
	for _, w := range cols.Cols {
		buf.WriteString(w)
	}

	if len(nzDefaults) != 0 {
		buf.WriteByte('.')
	}
	for _, nz := range nzDefaults {
		buf.WriteString(nz)
	}

	str := buf.String()
	strmangle.PutBuffer(buf)
	return str
}

// Enum values for account_key.coin
const (
	AccountKeyCoinBTC = "btc"
	AccountKeyCoinBCH = "bch"
	AccountKeyCoinEth = "eth"
)

// Enum values for account_key.account
const (
	AccountKeyAccountClient  = "client"
	AccountKeyAccountDeposit = "deposit"
	AccountKeyAccountPayment = "payment"
	AccountKeyAccountStored  = "stored"
)

// Enum values for address.coin
const (
	AddressCoinBTC = "btc"
	AddressCoinBCH = "bch"
	AddressCoinEth = "eth"
)

// Enum values for address.account
const (
	AddressAccountClient  = "client"
	AddressAccountDeposit = "deposit"
	AddressAccountPayment = "payment"
	AddressAccountStored  = "stored"
)

// Enum values for auth_account_key.coin
const (
	AuthAccountKeyCoinBTC = "btc"
	AuthAccountKeyCoinBCH = "bch"
)

// Enum values for auth_fullpubkey.coin
const (
	AuthFullpubkeyCoinBTC = "btc"
	AuthFullpubkeyCoinBCH = "bch"
)

// Enum values for btc_tx.coin
const (
	BTCTXCoinBTC = "btc"
	BTCTXCoinBCH = "bch"
)

// Enum values for btc_tx.action
const (
	BTCTXActionDeposit  = "deposit"
	BTCTXActionPayment  = "payment"
	BTCTXActionTransfer = "transfer"
)

// Enum values for eth_tx.coin
const (
	EthTXCoinEth = "eth"
)

// Enum values for eth_tx.action
const (
	EthTXActionDeposit  = "deposit"
	EthTXActionPayment  = "payment"
	EthTXActionTransfer = "transfer"
)

// Enum values for payment_request.coin
const (
	PaymentRequestCoinBTC = "btc"
	PaymentRequestCoinBCH = "bch"
	PaymentRequestCoinEth = "eth"
)

// Enum values for seed.coin
const (
	SeedCoinBTC = "btc"
	SeedCoinBCH = "bch"
)
