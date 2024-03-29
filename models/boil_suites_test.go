// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Admins", testAdmins)
	t.Run("BilibiliUsers", testBilibiliUsers)
	t.Run("Records", testRecords)
	t.Run("Reports", testReports)
}

func TestDelete(t *testing.T) {
	t.Run("Admins", testAdminsDelete)
	t.Run("BilibiliUsers", testBilibiliUsersDelete)
	t.Run("Records", testRecordsDelete)
	t.Run("Reports", testReportsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Admins", testAdminsQueryDeleteAll)
	t.Run("BilibiliUsers", testBilibiliUsersQueryDeleteAll)
	t.Run("Records", testRecordsQueryDeleteAll)
	t.Run("Reports", testReportsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Admins", testAdminsSliceDeleteAll)
	t.Run("BilibiliUsers", testBilibiliUsersSliceDeleteAll)
	t.Run("Records", testRecordsSliceDeleteAll)
	t.Run("Reports", testReportsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Admins", testAdminsExists)
	t.Run("BilibiliUsers", testBilibiliUsersExists)
	t.Run("Records", testRecordsExists)
	t.Run("Reports", testReportsExists)
}

func TestFind(t *testing.T) {
	t.Run("Admins", testAdminsFind)
	t.Run("BilibiliUsers", testBilibiliUsersFind)
	t.Run("Records", testRecordsFind)
	t.Run("Reports", testReportsFind)
}

func TestBind(t *testing.T) {
	t.Run("Admins", testAdminsBind)
	t.Run("BilibiliUsers", testBilibiliUsersBind)
	t.Run("Records", testRecordsBind)
	t.Run("Reports", testReportsBind)
}

func TestOne(t *testing.T) {
	t.Run("Admins", testAdminsOne)
	t.Run("BilibiliUsers", testBilibiliUsersOne)
	t.Run("Records", testRecordsOne)
	t.Run("Reports", testReportsOne)
}

func TestAll(t *testing.T) {
	t.Run("Admins", testAdminsAll)
	t.Run("BilibiliUsers", testBilibiliUsersAll)
	t.Run("Records", testRecordsAll)
	t.Run("Reports", testReportsAll)
}

func TestCount(t *testing.T) {
	t.Run("Admins", testAdminsCount)
	t.Run("BilibiliUsers", testBilibiliUsersCount)
	t.Run("Records", testRecordsCount)
	t.Run("Reports", testReportsCount)
}

func TestInsert(t *testing.T) {
	t.Run("Admins", testAdminsInsert)
	t.Run("Admins", testAdminsInsertWhitelist)
	t.Run("BilibiliUsers", testBilibiliUsersInsert)
	t.Run("BilibiliUsers", testBilibiliUsersInsertWhitelist)
	t.Run("Records", testRecordsInsert)
	t.Run("Records", testRecordsInsertWhitelist)
	t.Run("Reports", testReportsInsert)
	t.Run("Reports", testReportsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Admins", testAdminsReload)
	t.Run("BilibiliUsers", testBilibiliUsersReload)
	t.Run("Records", testRecordsReload)
	t.Run("Reports", testReportsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Admins", testAdminsReloadAll)
	t.Run("BilibiliUsers", testBilibiliUsersReloadAll)
	t.Run("Records", testRecordsReloadAll)
	t.Run("Reports", testReportsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Admins", testAdminsSelect)
	t.Run("BilibiliUsers", testBilibiliUsersSelect)
	t.Run("Records", testRecordsSelect)
	t.Run("Reports", testReportsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Admins", testAdminsUpdate)
	t.Run("BilibiliUsers", testBilibiliUsersUpdate)
	t.Run("Records", testRecordsUpdate)
	t.Run("Reports", testReportsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Admins", testAdminsSliceUpdateAll)
	t.Run("BilibiliUsers", testBilibiliUsersSliceUpdateAll)
	t.Run("Records", testRecordsSliceUpdateAll)
	t.Run("Reports", testReportsSliceUpdateAll)
}
