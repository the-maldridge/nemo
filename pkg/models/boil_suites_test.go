// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	t.Run("Events", testEvents)
	t.Run("Streams", testStreams)
}

func TestDelete(t *testing.T) {
	t.Run("Events", testEventsDelete)
	t.Run("Streams", testStreamsDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Events", testEventsQueryDeleteAll)
	t.Run("Streams", testStreamsQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Events", testEventsSliceDeleteAll)
	t.Run("Streams", testStreamsSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Events", testEventsExists)
	t.Run("Streams", testStreamsExists)
}

func TestFind(t *testing.T) {
	t.Run("Events", testEventsFind)
	t.Run("Streams", testStreamsFind)
}

func TestBind(t *testing.T) {
	t.Run("Events", testEventsBind)
	t.Run("Streams", testStreamsBind)
}

func TestOne(t *testing.T) {
	t.Run("Events", testEventsOne)
	t.Run("Streams", testStreamsOne)
}

func TestAll(t *testing.T) {
	t.Run("Events", testEventsAll)
	t.Run("Streams", testStreamsAll)
}

func TestCount(t *testing.T) {
	t.Run("Events", testEventsCount)
	t.Run("Streams", testStreamsCount)
}

func TestHooks(t *testing.T) {
	t.Run("Events", testEventsHooks)
	t.Run("Streams", testStreamsHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Events", testEventsInsert)
	t.Run("Events", testEventsInsertWhitelist)
	t.Run("Streams", testStreamsInsert)
	t.Run("Streams", testStreamsInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("EventToStreamUsingStream", testEventToOneStreamUsingStream)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("StreamToEvents", testStreamToManyEvents)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("EventToStreamUsingEvents", testEventToOneSetOpStreamUsingStream)
}

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
func TestToManyAdd(t *testing.T) {
	t.Run("StreamToEvents", testStreamToManyAddOpEvents)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {}

func TestReload(t *testing.T) {
	t.Run("Events", testEventsReload)
	t.Run("Streams", testStreamsReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Events", testEventsReloadAll)
	t.Run("Streams", testStreamsReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Events", testEventsSelect)
	t.Run("Streams", testStreamsSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Events", testEventsUpdate)
	t.Run("Streams", testStreamsUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Events", testEventsSliceUpdateAll)
	t.Run("Streams", testStreamsSliceUpdateAll)
}