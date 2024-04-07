package memtable

import (
	"github.com/shihaohong/leveldb_clone/entry"
	"github.com/shihaohong/leveldb_clone/sstable"
)

type Memtable interface {
	sstable.ImmutableDB

	// Gets all entries in the memtable
	GetAll() (vals []entry.Entry, err error)

	// Put sets the value for the given key. It overwrites any previous value
	// for that key; a mt is not a multi-map.
	Put(key, value []byte) error

	// Delete deletes the value for the given key.
	Delete(key []byte) error

	// Clears the in-memory data store
	Clear() error
}
