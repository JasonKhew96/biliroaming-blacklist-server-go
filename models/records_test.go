// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRecords(t *testing.T) {
	t.Parallel()

	query := Records()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRecordsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecordsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Records().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecordsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecordSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRecordsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RecordExists(ctx, tx, o.RecordID)
	if err != nil {
		t.Errorf("Unable to check if Record exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RecordExists to return true, but got false.")
	}
}

func testRecordsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	recordFound, err := FindRecord(ctx, tx, o.RecordID)
	if err != nil {
		t.Error(err)
	}

	if recordFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRecordsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Records().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRecordsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Records().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRecordsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	recordOne := &Record{}
	recordTwo := &Record{}
	if err = randomize.Struct(seed, recordOne, recordDBTypes, false, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	if err = randomize.Struct(seed, recordTwo, recordDBTypes, false, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = recordOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = recordTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Records().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRecordsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	recordOne := &Record{}
	recordTwo := &Record{}
	if err = randomize.Struct(seed, recordOne, recordDBTypes, false, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	if err = randomize.Struct(seed, recordTwo, recordDBTypes, false, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = recordOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = recordTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func recordBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Record) error {
	*o = Record{}
	return nil
}

func recordAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Record) error {
	*o = Record{}
	return nil
}

func recordAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Record) error {
	*o = Record{}
	return nil
}

func recordBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Record) error {
	*o = Record{}
	return nil
}

func recordAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Record) error {
	*o = Record{}
	return nil
}

func recordBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Record) error {
	*o = Record{}
	return nil
}

func recordAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Record) error {
	*o = Record{}
	return nil
}

func recordBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Record) error {
	*o = Record{}
	return nil
}

func recordAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Record) error {
	*o = Record{}
	return nil
}

func testRecordsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Record{}
	o := &Record{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, recordDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Record object: %s", err)
	}

	AddRecordHook(boil.BeforeInsertHook, recordBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	recordBeforeInsertHooks = []RecordHook{}

	AddRecordHook(boil.AfterInsertHook, recordAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	recordAfterInsertHooks = []RecordHook{}

	AddRecordHook(boil.AfterSelectHook, recordAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	recordAfterSelectHooks = []RecordHook{}

	AddRecordHook(boil.BeforeUpdateHook, recordBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	recordBeforeUpdateHooks = []RecordHook{}

	AddRecordHook(boil.AfterUpdateHook, recordAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	recordAfterUpdateHooks = []RecordHook{}

	AddRecordHook(boil.BeforeDeleteHook, recordBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	recordBeforeDeleteHooks = []RecordHook{}

	AddRecordHook(boil.AfterDeleteHook, recordAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	recordAfterDeleteHooks = []RecordHook{}

	AddRecordHook(boil.BeforeUpsertHook, recordBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	recordBeforeUpsertHooks = []RecordHook{}

	AddRecordHook(boil.AfterUpsertHook, recordAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	recordAfterUpsertHooks = []RecordHook{}
}

func testRecordsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecordsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(recordColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRecordsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRecordsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RecordSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRecordsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Records().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	recordDBTypes = map[string]string{`RecordID`: `integer`, `UID`: `bigint`, `Description`: `text`, `ChatID`: `bigint`, `MessageID`: `bigint`, `ApprovedBy`: `bigint`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_             = bytes.MinRead
)

func testRecordsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(recordPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(recordAllColumns) == len(recordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, recordDBTypes, true, recordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRecordsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(recordAllColumns) == len(recordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Record{}
	if err = randomize.Struct(seed, o, recordDBTypes, true, recordColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, recordDBTypes, true, recordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(recordAllColumns, recordPrimaryKeyColumns) {
		fields = recordAllColumns
	} else {
		fields = strmangle.SetComplement(
			recordAllColumns,
			recordPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RecordSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRecordsUpsert(t *testing.T) {
	t.Parallel()

	if len(recordAllColumns) == len(recordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Record{}
	if err = randomize.Struct(seed, &o, recordDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Record: %s", err)
	}

	count, err := Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, recordDBTypes, false, recordPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Record: %s", err)
	}

	count, err = Records().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
