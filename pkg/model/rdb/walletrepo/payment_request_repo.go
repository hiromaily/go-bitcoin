package walletrepo

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//PaymentRequest payment_requestテーブル
type PaymentRequest struct {
	ID          int64      `db:"id"`
	PaymentID   *int64     `db:"payment_id"`
	AddressFrom string     `db:"address_from"`
	AccountFrom string     `db:"account_from"`
	AddressTo   string     `db:"address_to"`
	Amount      string     `db:"amount"`
	IsDone      bool       `db:"is_done"`
	UpdatedAt   *time.Time `db:"updated_at"`
}

// GetPaymentRequestAll PaymentRequestテーブル全体を返す
func (r *WalletRepository) GetPaymentRequestAll() ([]PaymentRequest, error) {
	//sql := "SELECT * FROM payment_request WHERE is_done=false"
	sql := "SELECT * FROM payment_request WHERE payment_id IS NULL"
	//logger.Debugf("sql: %s", sql)

	var paymentRequests []PaymentRequest
	err := r.db.Select(&paymentRequests, sql)

	return paymentRequests, err
}

// GetPaymentRequestByPaymentID PaymentRequestテーブル全体を返す
func (r *WalletRepository) GetPaymentRequestByPaymentID(paymentID int64) ([]PaymentRequest, error) {
	sql := "SELECT * FROM payment_request WHERE payment_id=?"
	//logger.Debugf("sql: %s", sql)

	var paymentRequests []PaymentRequest
	err := r.db.Select(&paymentRequests, sql, paymentID)

	return paymentRequests, err
}

// InsertPaymentRequest PaymentRequestテーブルに出金依頼レコードを作成する
//TODO:BulkInsertがやりたい
func (r *WalletRepository) InsertPaymentRequest(paymentRequests []PaymentRequest, tx *sqlx.Tx, isCommit bool) error {

	sql := `
INSERT INTO payment_request (address_from, account_from, address_to, amount) 
VALUES (:address_from, :account_from, :address_to, :amount)
`
	//logger.Debugf("sql: %s", sql)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	for _, paymentRequest := range paymentRequests {
		_, err := tx.NamedExec(sql, paymentRequest)
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

// UpdateIsDoneOnPaymentRequest is_doneフィールドをtrueに更新する
//TODO:暫定で追加したのみ、実際の仕様に合わせて修正が必要
//TODO:payment_idレコードを追加したので、is_doneフィールドはいらないかもしれない
//TODO:一応、通知まで終わったレコードはdoneにしておく
func (r *WalletRepository) UpdateIsDoneOnPaymentRequest(paymentID int64, tx *sqlx.Tx, isCommit bool) (int64, error) {
	sql := `
UPDATE payment_request SET is_done=true WHERE payment_id=? 
`
	//logger.Debugf("sql: %s", sql)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	//res, err := tx.Exec(sql)
	res, err := tx.Exec(sql, paymentID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if isCommit {
		tx.Commit()
	}
	affectedNum, _ := res.RowsAffected()

	return affectedNum, nil
}

// UpdatePaymentIDOnPaymentRequest 出金トランザクション作成済のレコードのpayment_idを更新する
func (r *WalletRepository) UpdatePaymentIDOnPaymentRequest(paymentID int64, ids []int64, tx *sqlx.Tx, isCommit bool) (int64, error) {
	sql := `
UPDATE payment_request SET payment_id=? WHERE id IN (?) 
`

	//In対応
	query, args, err := sqlx.In(sql, paymentID, ids)
	if err != nil {
		return 0, errors.Errorf("sqlx.In() error: %v", err)
	}
	query = r.db.Rebind(query)
	//logger.Debugf("sql: %s", query)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	//res, err := tx.Exec(sql, paymentID, ids)
	res, err := tx.Exec(query, args...)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if isCommit {
		tx.Commit()
	}
	affectedNum, _ := res.RowsAffected()

	return affectedNum, nil

}

// ResetAnyFlagOnPaymentRequestForTestOnly テーブルを初期化する(テストでしか使用することはない)
func (r *WalletRepository) ResetAnyFlagOnPaymentRequestForTestOnly(tx *sqlx.Tx, isCommit bool) (int64, error) {
	sql := "UPDATE payment_request SET is_done=false, payment_id=NULL"
	//logger.Debugf("sql: %s", sql)

	if tx == nil {
		tx = r.db.MustBegin()
	}

	res, err := tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if isCommit {
		tx.Commit()
	}
	affectedNum, _ := res.RowsAffected()

	return affectedNum, nil
}