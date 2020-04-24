// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// AccountKey is an object representing the database table.
type AccountKey struct {
	ID                    uint64      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Coin                  null.String `boil:"coin" json:"coin,omitempty" toml:"coin" yaml:"coin,omitempty"`
	Account               null.String `boil:"account" json:"account,omitempty" toml:"account" yaml:"account,omitempty"`
	WalletAddress         string      `boil:"wallet_address" json:"wallet_address" toml:"wallet_address" yaml:"wallet_address"`
	P2SHSegwitAddress     string      `boil:"p2sh_segwit_address" json:"p2sh_segwit_address" toml:"p2sh_segwit_address" yaml:"p2sh_segwit_address"`
	FullPublicKey         string      `boil:"full_public_key" json:"full_public_key" toml:"full_public_key" yaml:"full_public_key"`
	WalletMultisigAddress string      `boil:"wallet_multisig_address" json:"wallet_multisig_address" toml:"wallet_multisig_address" yaml:"wallet_multisig_address"`
	RedeemScript          string      `boil:"redeem_script" json:"redeem_script" toml:"redeem_script" yaml:"redeem_script"`
	WalletImportFormat    string      `boil:"wallet_import_format" json:"wallet_import_format" toml:"wallet_import_format" yaml:"wallet_import_format"`
	Idx                   uint64      `boil:"idx" json:"idx" toml:"idx" yaml:"idx"`
	AddrStatus            uint8       `boil:"addr_status" json:"addr_status" toml:"addr_status" yaml:"addr_status"`
	UpdatedAt             null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *accountKeyR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L accountKeyL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AccountKeyColumns = struct {
	ID                    string
	Coin                  string
	Account               string
	WalletAddress         string
	P2SHSegwitAddress     string
	FullPublicKey         string
	WalletMultisigAddress string
	RedeemScript          string
	WalletImportFormat    string
	Idx                   string
	AddrStatus            string
	UpdatedAt             string
}{
	ID:                    "id",
	Coin:                  "coin",
	Account:               "account",
	WalletAddress:         "wallet_address",
	P2SHSegwitAddress:     "p2sh_segwit_address",
	FullPublicKey:         "full_public_key",
	WalletMultisigAddress: "wallet_multisig_address",
	RedeemScript:          "redeem_script",
	WalletImportFormat:    "wallet_import_format",
	Idx:                   "idx",
	AddrStatus:            "addr_status",
	UpdatedAt:             "updated_at",
}

// Generated where

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint64) IN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelperuint8 struct{ field string }

func (w whereHelperuint8) EQ(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint8) NEQ(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint8) LT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint8) LTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint8) GT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint8) GTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint8) IN(slice []uint8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AccountKeyWhere = struct {
	ID                    whereHelperuint64
	Coin                  whereHelpernull_String
	Account               whereHelpernull_String
	WalletAddress         whereHelperstring
	P2SHSegwitAddress     whereHelperstring
	FullPublicKey         whereHelperstring
	WalletMultisigAddress whereHelperstring
	RedeemScript          whereHelperstring
	WalletImportFormat    whereHelperstring
	Idx                   whereHelperuint64
	AddrStatus            whereHelperuint8
	UpdatedAt             whereHelpernull_Time
}{
	ID:                    whereHelperuint64{field: "`account_key`.`id`"},
	Coin:                  whereHelpernull_String{field: "`account_key`.`coin`"},
	Account:               whereHelpernull_String{field: "`account_key`.`account`"},
	WalletAddress:         whereHelperstring{field: "`account_key`.`wallet_address`"},
	P2SHSegwitAddress:     whereHelperstring{field: "`account_key`.`p2sh_segwit_address`"},
	FullPublicKey:         whereHelperstring{field: "`account_key`.`full_public_key`"},
	WalletMultisigAddress: whereHelperstring{field: "`account_key`.`wallet_multisig_address`"},
	RedeemScript:          whereHelperstring{field: "`account_key`.`redeem_script`"},
	WalletImportFormat:    whereHelperstring{field: "`account_key`.`wallet_import_format`"},
	Idx:                   whereHelperuint64{field: "`account_key`.`idx`"},
	AddrStatus:            whereHelperuint8{field: "`account_key`.`addr_status`"},
	UpdatedAt:             whereHelpernull_Time{field: "`account_key`.`updated_at`"},
}

// AccountKeyRels is where relationship names are stored.
var AccountKeyRels = struct {
}{}

// accountKeyR is where relationships are stored.
type accountKeyR struct {
}

// NewStruct creates a new relationship struct
func (*accountKeyR) NewStruct() *accountKeyR {
	return &accountKeyR{}
}

// accountKeyL is where Load methods for each relationship are stored.
type accountKeyL struct{}

var (
	accountKeyAllColumns            = []string{"id", "coin", "account", "wallet_address", "p2sh_segwit_address", "full_public_key", "wallet_multisig_address", "redeem_script", "wallet_import_format", "idx", "addr_status", "updated_at"}
	accountKeyColumnsWithoutDefault = []string{"coin", "account", "wallet_address", "p2sh_segwit_address", "full_public_key", "wallet_multisig_address", "redeem_script", "wallet_import_format", "idx"}
	accountKeyColumnsWithDefault    = []string{"id", "addr_status", "updated_at"}
	accountKeyPrimaryKeyColumns     = []string{"id"}
)

type (
	// AccountKeySlice is an alias for a slice of pointers to AccountKey.
	// This should generally be used opposed to []AccountKey.
	AccountKeySlice []*AccountKey

	accountKeyQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	accountKeyType                 = reflect.TypeOf(&AccountKey{})
	accountKeyMapping              = queries.MakeStructMapping(accountKeyType)
	accountKeyPrimaryKeyMapping, _ = queries.BindMapping(accountKeyType, accountKeyMapping, accountKeyPrimaryKeyColumns)
	accountKeyInsertCacheMut       sync.RWMutex
	accountKeyInsertCache          = make(map[string]insertCache)
	accountKeyUpdateCacheMut       sync.RWMutex
	accountKeyUpdateCache          = make(map[string]updateCache)
	accountKeyUpsertCacheMut       sync.RWMutex
	accountKeyUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single accountKey record from the query.
func (q accountKeyQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AccountKey, error) {
	o := &AccountKey{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for account_key")
	}

	return o, nil
}

// All returns all AccountKey records from the query.
func (q accountKeyQuery) All(ctx context.Context, exec boil.ContextExecutor) (AccountKeySlice, error) {
	var o []*AccountKey

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AccountKey slice")
	}

	return o, nil
}

// Count returns the count of all AccountKey records in the query.
func (q accountKeyQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count account_key rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q accountKeyQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if account_key exists")
	}

	return count > 0, nil
}

// AccountKeys retrieves all the records using an executor.
func AccountKeys(mods ...qm.QueryMod) accountKeyQuery {
	mods = append(mods, qm.From("`account_key`"))
	return accountKeyQuery{NewQuery(mods...)}
}

// FindAccountKey retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAccountKey(ctx context.Context, exec boil.ContextExecutor, iD uint64, selectCols ...string) (*AccountKey, error) {
	accountKeyObj := &AccountKey{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `account_key` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, accountKeyObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from account_key")
	}

	return accountKeyObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AccountKey) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no account_key provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(accountKeyColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	accountKeyInsertCacheMut.RLock()
	cache, cached := accountKeyInsertCache[key]
	accountKeyInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			accountKeyAllColumns,
			accountKeyColumnsWithDefault,
			accountKeyColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(accountKeyType, accountKeyMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(accountKeyType, accountKeyMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `account_key` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `account_key` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `account_key` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, accountKeyPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into account_key")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == accountKeyMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for account_key")
	}

CacheNoHooks:
	if !cached {
		accountKeyInsertCacheMut.Lock()
		accountKeyInsertCache[key] = cache
		accountKeyInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the AccountKey.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AccountKey) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	key := makeCacheKey(columns, nil)
	accountKeyUpdateCacheMut.RLock()
	cache, cached := accountKeyUpdateCache[key]
	accountKeyUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			accountKeyAllColumns,
			accountKeyPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update account_key, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `account_key` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, accountKeyPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(accountKeyType, accountKeyMapping, append(wl, accountKeyPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update account_key row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for account_key")
	}

	if !cached {
		accountKeyUpdateCacheMut.Lock()
		accountKeyUpdateCache[key] = cache
		accountKeyUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q accountKeyQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for account_key")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for account_key")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AccountKeySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `account_key` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, accountKeyPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in accountKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all accountKey")
	}
	return rowsAff, nil
}

var mySQLAccountKeyUniqueColumns = []string{
	"id",
	"wallet_address",
	"p2sh_segwit_address",
	"wallet_import_format",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AccountKey) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no account_key provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	nzDefaults := queries.NonZeroDefaultSet(accountKeyColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAccountKeyUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	accountKeyUpsertCacheMut.RLock()
	cache, cached := accountKeyUpsertCache[key]
	accountKeyUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			accountKeyAllColumns,
			accountKeyColumnsWithDefault,
			accountKeyColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			accountKeyAllColumns,
			accountKeyPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert account_key, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "account_key", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `account_key` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(accountKeyType, accountKeyMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(accountKeyType, accountKeyMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for account_key")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == accountKeyMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(accountKeyType, accountKeyMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for account_key")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for account_key")
	}

CacheNoHooks:
	if !cached {
		accountKeyUpsertCacheMut.Lock()
		accountKeyUpsertCache[key] = cache
		accountKeyUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single AccountKey record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AccountKey) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AccountKey provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), accountKeyPrimaryKeyMapping)
	sql := "DELETE FROM `account_key` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from account_key")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for account_key")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q accountKeyQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no accountKeyQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from account_key")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for account_key")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AccountKeySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `account_key` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, accountKeyPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from accountKey slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for account_key")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *AccountKey) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAccountKey(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountKeySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AccountKeySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountKeyPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `account_key`.* FROM `account_key` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, accountKeyPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AccountKeySlice")
	}

	*o = slice

	return nil
}

// AccountKeyExists checks if the AccountKey row exists.
func AccountKeyExists(ctx context.Context, exec boil.ContextExecutor, iD uint64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `account_key` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if account_key exists")
	}

	return exists, nil
}