package sstable

import "github.com/shihaohong/leveldb_clone/iterator"

type ImmutableDB interface {
	// Get gets the value for the given key. It returns an error if the
	// mt does not contain the key.
	Get(key []byte) (value []byte, err error)

	// Has returns true if the mt contains the given key.
	Has(key []byte) (ret bool, err error)

	// RangeScan returns an Iterator (see below) for scanning through all
	// key-value pairs in the given range, ordered by key ascending.
	RangeScan(start, limit []byte) (iterator.Iterator, error)
}
