// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// BilibiliUser is an object representing the database table.
type BilibiliUser struct {
	UID         int64     `boil:"uid" json:"uid" toml:"uid" yaml:"uid"`
	Counter     int64     `boil:"counter" json:"counter" toml:"counter" yaml:"counter"`
	IsWhitelist bool      `boil:"is_whitelist" json:"is_whitelist" toml:"is_whitelist" yaml:"is_whitelist"`
	BanUntil    null.Time `boil:"ban_until" json:"ban_until,omitempty" toml:"ban_until" yaml:"ban_until,omitempty"`
	RequestedAt time.Time `boil:"requested_at" json:"requested_at" toml:"requested_at" yaml:"requested_at"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *bilibiliUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bilibiliUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BilibiliUserColumns = struct {
	UID         string
	Counter     string
	IsWhitelist string
	BanUntil    string
	RequestedAt string
	CreatedAt   string
	UpdatedAt   string
}{
	UID:         "uid",
	Counter:     "counter",
	IsWhitelist: "is_whitelist",
	BanUntil:    "ban_until",
	RequestedAt: "requested_at",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var BilibiliUserTableColumns = struct {
	UID         string
	Counter     string
	IsWhitelist string
	BanUntil    string
	RequestedAt string
	CreatedAt   string
	UpdatedAt   string
}{
	UID:         "bilibili_users.uid",
	Counter:     "bilibili_users.counter",
	IsWhitelist: "bilibili_users.is_whitelist",
	BanUntil:    "bilibili_users.ban_until",
	RequestedAt: "bilibili_users.requested_at",
	CreatedAt:   "bilibili_users.created_at",
	UpdatedAt:   "bilibili_users.updated_at",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
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

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var BilibiliUserWhere = struct {
	UID         whereHelperint64
	Counter     whereHelperint64
	IsWhitelist whereHelperbool
	BanUntil    whereHelpernull_Time
	RequestedAt whereHelpertime_Time
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	UID:         whereHelperint64{field: "\"bilibili_users\".\"uid\""},
	Counter:     whereHelperint64{field: "\"bilibili_users\".\"counter\""},
	IsWhitelist: whereHelperbool{field: "\"bilibili_users\".\"is_whitelist\""},
	BanUntil:    whereHelpernull_Time{field: "\"bilibili_users\".\"ban_until\""},
	RequestedAt: whereHelpertime_Time{field: "\"bilibili_users\".\"requested_at\""},
	CreatedAt:   whereHelpertime_Time{field: "\"bilibili_users\".\"created_at\""},
	UpdatedAt:   whereHelpertime_Time{field: "\"bilibili_users\".\"updated_at\""},
}

// BilibiliUserRels is where relationship names are stored.
var BilibiliUserRels = struct {
}{}

// bilibiliUserR is where relationships are stored.
type bilibiliUserR struct {
}

// NewStruct creates a new relationship struct
func (*bilibiliUserR) NewStruct() *bilibiliUserR {
	return &bilibiliUserR{}
}

// bilibiliUserL is where Load methods for each relationship are stored.
type bilibiliUserL struct{}

var (
	bilibiliUserAllColumns            = []string{"uid", "counter", "is_whitelist", "ban_until", "requested_at", "created_at", "updated_at"}
	bilibiliUserColumnsWithoutDefault = []string{"uid"}
	bilibiliUserColumnsWithDefault    = []string{"counter", "is_whitelist", "ban_until", "requested_at", "created_at", "updated_at"}
	bilibiliUserPrimaryKeyColumns     = []string{"uid"}
	bilibiliUserGeneratedColumns      = []string{}
)

type (
	// BilibiliUserSlice is an alias for a slice of pointers to BilibiliUser.
	// This should almost always be used instead of []BilibiliUser.
	BilibiliUserSlice []*BilibiliUser

	bilibiliUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bilibiliUserType                 = reflect.TypeOf(&BilibiliUser{})
	bilibiliUserMapping              = queries.MakeStructMapping(bilibiliUserType)
	bilibiliUserPrimaryKeyMapping, _ = queries.BindMapping(bilibiliUserType, bilibiliUserMapping, bilibiliUserPrimaryKeyColumns)
	bilibiliUserInsertCacheMut       sync.RWMutex
	bilibiliUserInsertCache          = make(map[string]insertCache)
	bilibiliUserUpdateCacheMut       sync.RWMutex
	bilibiliUserUpdateCache          = make(map[string]updateCache)
	bilibiliUserUpsertCacheMut       sync.RWMutex
	bilibiliUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single bilibiliUser record from the query.
func (q bilibiliUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*BilibiliUser, error) {
	o := &BilibiliUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for bilibili_users")
	}

	return o, nil
}

// All returns all BilibiliUser records from the query.
func (q bilibiliUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (BilibiliUserSlice, error) {
	var o []*BilibiliUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BilibiliUser slice")
	}

	return o, nil
}

// Count returns the count of all BilibiliUser records in the query.
func (q bilibiliUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count bilibili_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q bilibiliUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if bilibili_users exists")
	}

	return count > 0, nil
}

// BilibiliUsers retrieves all the records using an executor.
func BilibiliUsers(mods ...qm.QueryMod) bilibiliUserQuery {
	mods = append(mods, qm.From("\"bilibili_users\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"bilibili_users\".*"})
	}

	return bilibiliUserQuery{q}
}

// FindBilibiliUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBilibiliUser(ctx context.Context, exec boil.ContextExecutor, uID int64, selectCols ...string) (*BilibiliUser, error) {
	bilibiliUserObj := &BilibiliUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"bilibili_users\" where \"uid\"=$1", sel,
	)

	q := queries.Raw(query, uID)

	err := q.Bind(ctx, exec, bilibiliUserObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from bilibili_users")
	}

	return bilibiliUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BilibiliUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bilibili_users provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(bilibiliUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bilibiliUserInsertCacheMut.RLock()
	cache, cached := bilibiliUserInsertCache[key]
	bilibiliUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bilibiliUserAllColumns,
			bilibiliUserColumnsWithDefault,
			bilibiliUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bilibiliUserType, bilibiliUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bilibiliUserType, bilibiliUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"bilibili_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"bilibili_users\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into bilibili_users")
	}

	if !cached {
		bilibiliUserInsertCacheMut.Lock()
		bilibiliUserInsertCache[key] = cache
		bilibiliUserInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the BilibiliUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BilibiliUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	bilibiliUserUpdateCacheMut.RLock()
	cache, cached := bilibiliUserUpdateCache[key]
	bilibiliUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bilibiliUserAllColumns,
			bilibiliUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update bilibili_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"bilibili_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bilibiliUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bilibiliUserType, bilibiliUserMapping, append(wl, bilibiliUserPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update bilibili_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for bilibili_users")
	}

	if !cached {
		bilibiliUserUpdateCacheMut.Lock()
		bilibiliUserUpdateCache[key] = cache
		bilibiliUserUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q bilibiliUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for bilibili_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for bilibili_users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BilibiliUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bilibiliUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"bilibili_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bilibiliUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bilibiliUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bilibiliUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BilibiliUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no bilibili_users provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(bilibiliUserColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	bilibiliUserUpsertCacheMut.RLock()
	cache, cached := bilibiliUserUpsertCache[key]
	bilibiliUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bilibiliUserAllColumns,
			bilibiliUserColumnsWithDefault,
			bilibiliUserColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			bilibiliUserAllColumns,
			bilibiliUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert bilibili_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bilibiliUserPrimaryKeyColumns))
			copy(conflict, bilibiliUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"bilibili_users\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bilibiliUserType, bilibiliUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bilibiliUserType, bilibiliUserMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert bilibili_users")
	}

	if !cached {
		bilibiliUserUpsertCacheMut.Lock()
		bilibiliUserUpsertCache[key] = cache
		bilibiliUserUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single BilibiliUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BilibiliUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BilibiliUser provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bilibiliUserPrimaryKeyMapping)
	sql := "DELETE FROM \"bilibili_users\" WHERE \"uid\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from bilibili_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for bilibili_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bilibiliUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bilibiliUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bilibili_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bilibili_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BilibiliUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bilibiliUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"bilibili_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bilibiliUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bilibiliUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for bilibili_users")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BilibiliUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindBilibiliUser(ctx, exec, o.UID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BilibiliUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BilibiliUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bilibiliUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"bilibili_users\".* FROM \"bilibili_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bilibiliUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BilibiliUserSlice")
	}

	*o = slice

	return nil
}

// BilibiliUserExists checks if the BilibiliUser row exists.
func BilibiliUserExists(ctx context.Context, exec boil.ContextExecutor, uID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"bilibili_users\" where \"uid\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, uID)
	}
	row := exec.QueryRowContext(ctx, sql, uID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if bilibili_users exists")
	}

	return exists, nil
}

// Exists checks if the BilibiliUser row exists.
func (o *BilibiliUser) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return BilibiliUserExists(ctx, exec, o.UID)
}
