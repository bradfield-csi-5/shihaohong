package sstable

import (
	"bytes"
	"errors"
	"os"

	"github.com/shihaohong/leveldb_clone/entry"
	"github.com/shihaohong/leveldb_clone/iterator"
)

type SSTable struct {
	Entries []entry.Entry
}

func (t *SSTable) Get(key []byte) (value []byte, err error) {
	for _, entry := range t.Entries {
		res := bytes.Compare(key, entry.Key)
		if res == 0 {
			return entry.Value, nil
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

func (t *SSTable) WriteToDisk(sstPath string) error {
	f, err := os.OpenFile(sstPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, e := range t.Entries {
		bs := make([]byte, 0)
		bs = append(bs, byte(len(e.Key)))
		bs = append(bs, e.Key...)
		bs = append(bs, byte(len(e.Value)))
		bs = append(bs, e.Value...)
		_, err = f.Write(bs)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewSSTable(entries []entry.Entry) *SSTable {
	return &SSTable{
		Entries: entries,
	}
}
