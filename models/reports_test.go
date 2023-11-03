// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testReports(t *testing.T) {
	t.Parallel()

	query := Reports()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testReportsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
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

	count, err := Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReportsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Reports().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReportsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReportSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReportsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ReportExists(ctx, tx, o.ReportID)
	if err != nil {
		t.Errorf("Unable to check if Report exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ReportExists to return true, but got false.")
	}
}

func testReportsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	reportFound, err := FindReport(ctx, tx, o.ReportID)
	if err != nil {
		t.Error(err)
	}

	if reportFound == nil {
		t.Error("want a record, got nil")
	}
}

func testReportsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Reports().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testReportsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Reports().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testReportsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reportOne := &Report{}
	reportTwo := &Report{}
	if err = randomize.Struct(seed, reportOne, reportDBTypes, false, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}
	if err = randomize.Struct(seed, reportTwo, reportDBTypes, false, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = reportOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = reportTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Reports().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testReportsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	reportOne := &Report{}
	reportTwo := &Report{}
	if err = randomize.Struct(seed, reportOne, reportDBTypes, false, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}
	if err = randomize.Struct(seed, reportTwo, reportDBTypes, false, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = reportOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = reportTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testReportsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReportsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(reportColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReportsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
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

func testReportsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ReportSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testReportsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Reports().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	reportDBTypes = map[string]string{`ReportID`: `integer`, `UID`: `bigint`, `Description`: `text`, `FileType`: `smallint`, `FileID`: `text`, `SubmitBy`: `text`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_             = bytes.MinRead
)

func testReportsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(reportPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(reportAllColumns) == len(reportPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, reportDBTypes, true, reportPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testReportsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(reportAllColumns) == len(reportPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Report{}
	if err = randomize.Struct(seed, o, reportDBTypes, true, reportColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, reportDBTypes, true, reportPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(reportAllColumns, reportPrimaryKeyColumns) {
		fields = reportAllColumns
	} else {
		fields = strmangle.SetComplement(
			reportAllColumns,
			reportPrimaryKeyColumns,
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

	slice := ReportSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testReportsUpsert(t *testing.T) {
	t.Parallel()

	if len(reportAllColumns) == len(reportPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Report{}
	if err = randomize.Struct(seed, &o, reportDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Report: %s", err)
	}

	count, err := Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, reportDBTypes, false, reportPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Report struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Report: %s", err)
	}

	count, err = Reports().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
