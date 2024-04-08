package sstable

import (
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/shihaohong/leveldb_clone/consts"
	"github.com/shihaohong/leveldb_clone/entry"
	"github.com/shihaohong/leveldb_clone/iterator"
)

type SSTable struct {
	Entries []entry.Entry
}

func (t *SSTable) Get(key []byte) (value []byte, err error) {
	for _, entry := range t.Entries {
		res := bytes.Compare(entry.Key, key)
		if res == 0 {
			return entry.Value, nil
		} else if res > 0 {
			return nil, consts.ErrSearchKeyNotFound
		}
	}
	return nil, consts.ErrSearchKeyNotFound
}

func (t *SSTable) Has(key []byte) (ret bool, err error) {
	_, err = t.Get(key)
	if err != nil {
		return true, nil
	}
	return false, consts.ErrSearchKeyNotFound
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

func NewSSTableFromPath(sstPath string) (*SSTable, error) {
	sst := &SSTable{
		Entries: make([]entry.Entry, 0),
	}
	f, err := os.Open(sstPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	for {
		keyLenSlice := make([]byte, 1)
		_, err = f.Read(keyLenSlice)
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		} else if err != nil && errors.Is(err, io.EOF) {
			return sst, nil
		}

		keyLen := uint8(keyLenSlice[0])
		keySlice := make([]byte, keyLen)
		_, err = f.Read(keySlice)
		if err != nil {
			return nil, err
		}
		valLenSlice := make([]byte, 1)
		_, err = f.Read(valLenSlice)
		if err != nil {
			return nil, err
		}
		valLen := uint8(valLenSlice[0])

		valSlice := make([]byte, valLen)
		_, err = f.Read(valSlice)
		if err != nil {
			return nil, err
		}

		entry := entry.Entry{
			Key:   keySlice,
			Value: valSlice,
		}

		sst.Entries = append(sst.Entries, entry)
	}
}

func NewSSTable(entries []entry.Entry) *SSTable {
	return &SSTable{
		Entries: entries,
	}
}
