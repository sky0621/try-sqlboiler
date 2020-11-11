// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package boiled

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Items", testItems)
	t.Run("Migrations", testMigrations)
}

func TestDelete(t *testing.T) {
	t.Run("Items", testItemsDelete)
	t.Run("Migrations", testMigrationsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Items", testItemsQueryDeleteAll)
	t.Run("Migrations", testMigrationsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Items", testItemsSliceDeleteAll)
	t.Run("Migrations", testMigrationsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Items", testItemsExists)
	t.Run("Migrations", testMigrationsExists)
}

func TestFind(t *testing.T) {
	t.Run("Items", testItemsFind)
	t.Run("Migrations", testMigrationsFind)
}

func TestBind(t *testing.T) {
	t.Run("Items", testItemsBind)
	t.Run("Migrations", testMigrationsBind)
}

func TestOne(t *testing.T) {
	t.Run("Items", testItemsOne)
	t.Run("Migrations", testMigrationsOne)
}

func TestAll(t *testing.T) {
	t.Run("Items", testItemsAll)
	t.Run("Migrations", testMigrationsAll)
}

func TestCount(t *testing.T) {
	t.Run("Items", testItemsCount)
	t.Run("Migrations", testMigrationsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Items", testItemsHooks)
	t.Run("Migrations", testMigrationsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Items", testItemsInsert)
	t.Run("Items", testItemsInsertWhitelist)
	t.Run("Migrations", testMigrationsInsert)
	t.Run("Migrations", testMigrationsInsertWhitelist)
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
	t.Run("Items", testItemsReload)
	t.Run("Migrations", testMigrationsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Items", testItemsReloadAll)
	t.Run("Migrations", testMigrationsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Items", testItemsSelect)
	t.Run("Migrations", testMigrationsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Items", testItemsUpdate)
	t.Run("Migrations", testMigrationsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Items", testItemsSliceUpdateAll)
	t.Run("Migrations", testMigrationsSliceUpdateAll)
}
