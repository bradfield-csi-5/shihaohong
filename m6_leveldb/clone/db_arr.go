package main

import (
	"bytes"
	"errors"
)

type Entry struct {
	key   []byte
	value []byte
}

// Initial version, in-memory and no persistence
type ArrDB struct {
	data []Entry
	len  int
}

func (db *ArrDB) Get(key []byte) (value []byte, err error) {
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

func (db *ArrDB) Has(key []byte) (ret bool, err error) {
	_, err = db.Get(key)
	if err != nil {
		return true, nil
	}
	return false, err
}

func (db *ArrDB) Put(key, value []byte) error {
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
		} else if res > 0 {
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

func (db *ArrDB) Delete(key []byte) error {
	return nil
}

func (db *ArrDB) RangeScan(start, limit []byte) (Iterator, error) {
	return nil, nil
}

func NewArrDB() ArrDB {
	return ArrDB{
		data: []Entry{},
	}
}
