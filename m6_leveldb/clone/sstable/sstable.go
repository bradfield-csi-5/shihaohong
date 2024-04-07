package sstable

import (
	"bytes"
	"errors"

	"github.com/shihaohong/leveldb_clone/iterator"
)

type SSTable struct {
	entries []Entry
}

type Entry struct {
	key   []byte
	value []byte
}

func (t *SSTable) Get(key []byte) (value []byte, err error) {
	for _, entry := range t.entries {
		res := bytes.Compare(key, entry.key)
		if res == 0 {
			return entry.value, nil
		} else if res > 0 {
			return nil, errors.New("search key not found")
		}
	}
	return nil, errors.New("search key not found")
}

func (t *SSTable) Has(key []byte) (ret bool, err error) {
	_, err = t.Get(key)
	if err != nil {
		return true, nil
	}
	return false, errors.New("search key not found")
}

// TODO
func (t *SSTable) RangeScan(start, limit []byte) (iterator.Iterator, error) {
	return nil, nil
}

func NewSSTable() *SSTable {
	return &SSTable{}
}
