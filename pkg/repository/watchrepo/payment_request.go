package watchrepo

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.uber.org/zap"

	models "github.com/hiromaily/go-crypto-wallet/pkg/models/rdb"
	"github.com/hiromaily/go-crypto-wallet/pkg/wallet/coin"
)

// PaymentRequestRepositorier is PaymentRequestRepository interface
type PaymentRequestRepositorier interface {
	GetAll() ([]*models.PaymentRequest, error)
	GetAllByPaymentID(paymentID int64) ([]*models.PaymentRequest, error)
	InsertBulk(items []*models.PaymentRequest) error
	UpdatePaymentID(paymentID int64, ids []int64) (int64, error)
	UpdateIsDone(paymentID int64) (int64, error)
	DeleteAll() (int64, error)
}

// PaymentRequestRepository is repository for payment_request table
type PaymentRequestRepository struct {
	dbConn       *sql.DB
	tableName    string
	coinTypeCode coin.CoinTypeCode
	logger       *zap.Logger
}

// NewPaymentRequestRepository returns PaymentRequestRepository object
func NewPaymentRequestRepository(dbConn *sql.DB, coinTypeCode coin.CoinTypeCode, logger *zap.Logger) *PaymentRequestRepository {
	return &PaymentRequestRepository{
		dbConn:       dbConn,
		tableName:    "payment_request",
		coinTypeCode: coinTypeCode,
		logger:       logger,
	}
}

// GetAll returns all records whose payment_id is null
func (r *PaymentRequestRepository) GetAll() ([]*models.PaymentRequest, error) {
	// sql := "SELECT * FROM payment_request WHERE payment_id IS NULL"
	ctx := context.Background()

	prItems, err := models.PaymentRequests(
		qm.Where("coin=?", r.coinTypeCode.String()),
		qm.And("payment_id IS NULL"),
	).All(ctx, r.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.PaymentRequests().All()")
	}
	return prItems, nil
}

// GetAllByPaymentID returns all records searched by payment_id
func (r *PaymentRequestRepository) GetAllByPaymentID(paymentID int64) ([]*models.PaymentRequest, error) {
	// sql := "SELECT * FROM payment_request WHERE payment_id=?"
	ctx := context.Background()

	prItems, err := models.PaymentRequests(
		qm.Where("coin=?", r.coinTypeCode.String()),
		qm.And("payment_id=?", paymentID),
	).All(ctx, r.dbConn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to call models.PaymentRequests().All()")
	}
	return prItems, nil
}

// InsertBulk inserts multiple records
func (r *PaymentRequestRepository) InsertBulk(items []*models.PaymentRequest) error {
	ctx := context.Background()
	return models.PaymentRequestSlice(items).InsertAll(ctx, r.dbConn, boil.Infer())
}

// UpdatePaymentID updates isDone
func (r *PaymentRequestRepository) UpdatePaymentID(paymentID int64, ids []int64) (int64, error) {
	// sql := `UPDATE payment_request SET payment_id=? WHERE id IN (?)`
	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{
		models.PaymentRequestColumns.PaymentID: paymentID,
	}

	// change []int64 to []interface
	targetIDs := make([]interface{}, len(ids))
	for i, v := range ids {
		targetIDs[i] = v
	}

	return models.PaymentRequests(
		qm.WhereIn("id IN ?", targetIDs...), // unique
	).UpdateAll(ctx, r.dbConn, updCols)
}

// UpdateIsDone updates isDone
func (r *PaymentRequestRepository) UpdateIsDone(paymentID int64) (int64, error) {
	// sql := `UPDATE payment_request SET is_done=true WHERE payment_id=?`
	ctx := context.Background()

	// Set updating columns
	updCols := map[string]interface{}{
		models.PaymentRequestColumns.IsDone: true,
	}
	return models.PaymentRequests(
		qm.Where("coin=?", r.coinTypeCode.String()),
		qm.And("payment_id=?", paymentID),
	).UpdateAll(ctx, r.dbConn, updCols)
}

// DeleteAll deletes all records
func (r *PaymentRequestRepository) DeleteAll() (int64, error) {
	ctx := context.Background()
	return models.PaymentRequests(
		qm.Where("coin=?", r.coinTypeCode.String()),
	).DeleteAll(ctx, r.dbConn)
}
