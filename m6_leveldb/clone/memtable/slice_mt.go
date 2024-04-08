package memtable

import (
	"bytes"

	"github.com/shihaohong/leveldb_clone/consts"
	"github.com/shihaohong/leveldb_clone/entry"
	"github.com/shihaohong/leveldb_clone/iterator"
)

// Initial version, in-memory and no persistence
type SliceMemtable struct {
	data []entry.Entry
	len  int
}

func (db *SliceMemtable) Get(key []byte) (value []byte, err error) {
	for i := 0; i < db.len; i++ {
		res := bytes.Compare(db.data[i].Key, key)
		if res == 0 {
			return db.data[i].Value, nil
		} else if res < 0 {
			return nil, consts.ErrSearchKeyNotFound
		}
	}
	return nil, consts.ErrSearchKeyNotFound
}

func (db *SliceMemtable) Has(key []byte) (ret bool, err error) {
	_, err = db.Get(key)
	if err != nil {
		return true, nil
	}
	return false, err
}

func (db *SliceMemtable) Put(key, value []byte) error {
	if db.len == 0 {
		newEntry := entry.Entry{Key: key, Value: value}
		db.data = append(db.data, newEntry)
		db.len++
		return nil
	}

	for i := 0; i < db.len; i++ {
		res := bytes.Compare(db.data[i].Key, key)
		if res == 0 {
			db.data[i].Value = value
			return nil
		} else if res < 0 {
			newEntry := entry.Entry{Key: key, Value: value}
			db.data = append(db.data[:i+1], db.data[i:]...)
			db.data[i] = newEntry
			db.len++
			return nil
		}
	}
	newEntry := entry.Entry{Key: key, Value: value}
	db.data = append(db.data, newEntry)
	db.len++
	return nil
}

func (db *SliceMemtable) Delete(key []byte) error {
	for i := 0; i < db.len; i++ {
		res := bytes.Compare(db.data[i].Key, key)
		if res == 0 {
			db.data = append(db.data[:i], db.data[i+1:]...)
			db.len--
			return nil
		}
	}
	return consts.ErrSearchKeyNotFound
}

func (mt *SliceMemtable) GetAll() error {
	return nil
}

func (mt *SliceMemtable) Clear() error {
	return nil
}

func (db *SliceMemtable) RangeScan(start, limit []byte) (iterator.Iterator, error) {
	return nil, nil
}

func NewSliceMT() *SliceMemtable {
	return &SliceMemtable{
		data: []entry.Entry{},
	}
}
