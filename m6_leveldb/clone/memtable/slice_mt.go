package memtable

import (
	"bytes"
	"errors"
)

type Entry struct {
	key   []byte
	value []byte
}

// Initial version, in-memory and no persistence
type SliceMemtable struct {
	data []Entry
	len  int
}

func (db *SliceMemtable) Get(key []byte) (value []byte, err error) {
	for i := 0; i < db.len; i++ {
		res := bytes.Compare(key, db.data[i].key)
		if res == 0 {
			return db.data[i].value, nil
		} else if res > 0 {
			return nil, errors.New("search key not found")
		}
	}
	return nil, errors.New("search key not found")
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
		newEntry := Entry{key: key, value: value}
		db.data = append(db.data, newEntry)
		db.len++
		return nil
	}

	for i := 0; i < db.len; i++ {
		res := bytes.Compare(db.data[i].key, key)
		if res == 0 {
			db.data[i].value = value
			return nil
		} else if res < 0 {
			newEntry := Entry{key: key, value: value}
			db.data = append(db.data[:i+1], db.data[i:]...)
			db.data[i] = newEntry
			db.len++
			return nil
		}
	}
	newEntry := Entry{key: key, value: value}
	db.data = append(db.data, newEntry)
	db.len++
	return nil
}

func (db *SliceMemtable) Delete(key []byte) error {
	for i := 0; i < db.len; i++ {
		res := bytes.Compare(db.data[i].key, key)
		if res == 0 {
			db.data = append(db.data[:i], db.data[i+1:]...)
			db.len--
			return nil
		}
	}
	return errors.New("search key not found")
}

func (db *SliceMemtable) RangeScan(start, limit []byte) (Iterator, error) {
	return nil, nil
}

func NewSliceMT() *SliceMemtable {
	return &SliceMemtable{
		data: []Entry{},
	}
}
