package main

import (
	"errors"
	"sort"
)

type DB interface {
	// Get gets the value for the given key. It returns an error if the
	// DB does not contain the key.
	Get(key []byte) (value []byte, err error)

	// Has returns true if the DB contains the given key.
	Has(key []byte) (ret bool, err error)

	// Put sets the value for the given key. It overwrites any previous value
	// for that key; a DB is not a multi-map.
	Put(key, value []byte) error

	// Delete deletes the value for the given key.
	Delete(key []byte) error

	// RangeScan returns an Iterator (see below) for scanning through all
	// key-value pairs in the given range, ordered by key ascending.
	RangeScan(start, limit []byte) (Iterator, error)
}

type Iterator interface {
	// Next moves the iterator to the next key/value pair.
	// It returns false if the iterator is exhausted.
	Next() bool

	// Error returns any accumulated error. Exhausting all the key/value pairs
	// is not considered to be an error.
	Error() error

	// Key returns the key of the current key/value pair, or nil if done.
	Key() []byte

	// Value returns the value of the current key/value pair, or nil if done.
	Value() []byte
}

// Initial version, in-memory and no persistence
type MemoryDB struct {
	data map[string][]byte
}

func (db *MemoryDB) Get(key []byte) (value []byte, err error) {
	return db.data[string(key)], nil
}

func (db *MemoryDB) Has(key []byte) (ret bool, err error) {
	_, ok := db.data[string(key)]
	return ok, nil
}

func (db *MemoryDB) Put(key, value []byte) error {
	db.data[string(key)] = value
	return nil
}

func (db *MemoryDB) Delete(key []byte) error {
	delete(db.data, string(key))
	return nil
}

func (db *MemoryDB) RangeScan(start, limit []byte) (Iterator, error) {
	startKey := string(start)
	limitKey := string(limit)

	keys := make([]string, 0)
	for k := range db.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	startFound := false
	iterator := NewIter()
	for k := range keys {
		// when start is found, toggle
		if !startFound && keys[k] == startKey {
			startFound = true
		}

		// continue iterating if we haven't found the starting key
		if !startFound {
			continue
		}

		// append to iterator tuple list
		tuple := Tuple{
			key:   []byte(keys[k]),
			value: db.data[keys[k]],
		}
		iterator.tuples = append(iterator.tuples, tuple)

		// reached limit key, break out
		if keys[k] == limitKey {
			break
		}
	}

	if len(iterator.tuples) == 0 {
		return nil, errors.New("no tuples within the range found")
	}

	return iterator, nil
}

func NewMemoryDB() MemoryDB {
	return MemoryDB{
		data: make(map[string][]byte),
	}
}

type Iter struct {
	tuples []Tuple
	index  int
}

type Tuple struct {
	key   []byte
	value []byte
}

func (it *Iter) Next() bool {
	it.index++
	return it.index < len(it.tuples)
}

// TODO: is an error possible for a memory store?
func (it *Iter) Error() error {
	return nil
}

func (it *Iter) Key() []byte {
	return it.tuples[it.index].key
}

func (it *Iter) Value() []byte {
	return it.tuples[it.index].value
}

func NewIter() *Iter {
	return &Iter{
		tuples: make([]Tuple, 0),
		index:  0,
	}
}
