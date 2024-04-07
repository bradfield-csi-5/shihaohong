package memtable

import (
	"github.com/shihaohong/leveldb_clone/sstable"
)

type Memtable interface {
	sstable.ImmutableDB

	// Put sets the value for the given key. It overwrites any previous value
	// for that key; a mt is not a multi-map.
	Put(key, value []byte) error

	// Delete deletes the value for the given key.
	Delete(key []byte) error
}
