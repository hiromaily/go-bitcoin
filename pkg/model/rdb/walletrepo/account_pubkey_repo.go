package walletrepo

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/hiromaily/go-bitcoin/pkg/account"
)

// AccountPublicKeyTable account_key_clientテーブル
type AccountPublicKeyTable struct {
	ID            int64      `db:"id"`
	WalletAddress string     `db:"wallet_address"`
	Account       string     `db:"account"`
	IsAllocated   bool       `db:"is_allocated"`
	UpdatedAt     *time.Time `db:"updated_at"`
}

var accountPubKeyTableName = map[account.AccountType]string{
	account.AccountTypeClient:  "account_pubkey_client",
	account.AccountTypeReceipt: "account_pubkey_receipt",
	account.AccountTypePayment: "account_pubkey_payment",
	account.AccountTypeQuoine:  "account_pubkey_quoine",
	account.AccountTypeFee:     "account_pubkey_fee",
	account.AccountTypeStored:  "account_pubkey_stored",
}

//getAllAccountPubKeyTable
func (r *WalletRepository) getAllAccountPubKeyTable(tbl string) ([]AccountPublicKeyTable, error) {
	sql := "SELECT * FROM %s;"
	sql = fmt.Sprintf(sql, tbl)
	//r.logger.Debugf("sql: %s", sql)

	var accountKeyTable []AccountPublicKeyTable
	err := r.db.Select(&accountKeyTable, sql)
	if err != nil {
		return nil, err
	}

	return accountKeyTable, nil
}

// GetAllAccountPubKeyTable account_pubkey_table(client, payment, receipt...)テーブルから全レコードを取得
func (r *WalletRepository) GetAllAccountPubKeyTable(accountType account.AccountType) ([]AccountPublicKeyTable, error) {
	return r.getAllAccountPubKeyTable(accountPubKeyTableName[accountType])
}

//getOneUnAllocatedAccountPubKeyTable account_pubkey_table(client, payment, receipt...)テーブルからis_allocated=falseの1レコードを取得
func (r *WalletRepository) getOneUnAllocatedAccountPubKeyTable(tbl string) (*AccountPublicKeyTable, error) {
	sql := "SELECT * FROM %s WHERE is_allocated=false ORDER BY id LIMIT 1;"
	sql = fmt.Sprintf(sql, tbl)
	//r.logger.Debugf("sql: %s", sql)

	var accountKeyTable AccountPublicKeyTable
	err := r.db.Get(&accountKeyTable, sql)
	if err != nil {
		return nil, err
	}

	return &accountKeyTable, nil
}

// GetOneUnAllocatedAccountPubKeyTable account_pubkey_table(client, payment, receipt...)テーブルからis_allocated=falseの1レコードを取得
func (r *WalletRepository) GetOneUnAllocatedAccountPubKeyTable(accountType account.AccountType) (*AccountPublicKeyTable, error) {
	return r.getOneUnAllocatedAccountPubKeyTable(accountPubKeyTableName[accountType])
}

// insertAccountPubKeyTable account_key_table(client, payment, receipt...)テーブルにレコードを作成する
//TODO:BulkInsertがやりたい
func (r *WalletRepository) insertAccountPubKeyTable(tbl string, accountPubKeyTables []AccountPublicKeyTable, tx *sqlx.Tx, isCommit bool) error {

	sql := `
INSERT INTO %s (wallet_address, account) 
VALUES (:wallet_address, :account)
`
	sql = fmt.Sprintf(sql, tbl)
	//r.logger.Debugf("sql: %s", sql)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	for _, accountPubKeyTable := range accountPubKeyTables {
		_, err := tx.NamedExec(sql, accountPubKeyTable)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if isCommit {
		tx.Commit()
	}

	return nil
}

// InsertAccountPubKeyTable account_pubkey_table(client, payment, receipt...)テーブルにレコードを作成する
func (r *WalletRepository) InsertAccountPubKeyTable(accountType account.AccountType, accountPubKeyTables []AccountPublicKeyTable, tx *sqlx.Tx, isCommit bool) error {
	return r.insertAccountPubKeyTable(accountPubKeyTableName[accountType], accountPubKeyTables, tx, isCommit)
}

// updateAccountOnAccountPubKeyTable Accountを更新する
func (r *WalletRepository) updateAccountOnAccountPubKeyTable(tbl string, accountKeyTable []AccountPublicKeyTable, tx *sqlx.Tx, isCommit bool) error {
	sql := `
UPDATE %s SET account=:account, updated_at=:updated_at 
WHERE id=:id
`
	sql = fmt.Sprintf(sql, tbl)
	//r.logger.Debugf("sql: %s", sql)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	for _, accountKey := range accountKeyTable {
		_, err := tx.NamedExec(sql, accountKey)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if isCommit {
		tx.Commit()
	}

	return nil
}

// UpdateAccountOnAccountPubKeyTable Accountを更新する
func (r *WalletRepository) UpdateAccountOnAccountPubKeyTable(accountType account.AccountType, accountKeyTable []AccountPublicKeyTable, tx *sqlx.Tx, isCommit bool) error {
	return r.updateAccountOnAccountPubKeyTable(accountPubKeyTableName[accountType], accountKeyTable, tx, isCommit)
}

// updateIsAllocatedOnAccountPubKeyTable IsAllocatedを更新する
func (r *WalletRepository) updateIsAllocatedOnAccountPubKeyTable(tbl string, accountKeyTable []AccountPublicKeyTable, tx *sqlx.Tx, isCommit bool) error {
	sql := `
UPDATE %s SET is_allocated=:is_allocated, updated_at=:updated_at 
WHERE wallet_address=:wallet_address
`
	sql = fmt.Sprintf(sql, tbl)
	//r.logger.Debugf("sql: %s", sql)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	for _, accountKey := range accountKeyTable {
		_, err := tx.NamedExec(sql, accountKey)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if isCommit {
		tx.Commit()
	}

	return nil
}

// UpdateIsAllocatedOnAccountPubKeyTable IsAllocatedを更新する
func (r *WalletRepository) UpdateIsAllocatedOnAccountPubKeyTable(accountType account.AccountType, accountKeyTable []AccountPublicKeyTable, tx *sqlx.Tx, isCommit bool) error {
	return r.updateIsAllocatedOnAccountPubKeyTable(accountPubKeyTableName[accountType], accountKeyTable, tx, isCommit)
}