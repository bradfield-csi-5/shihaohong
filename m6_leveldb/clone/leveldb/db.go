package main

import (
	"errors"

	"github.com/shihaohong/leveldb_clone/memtable"
	"github.com/shihaohong/leveldb_clone/wal"
)

type LevelDB struct {
	wal wal.Log
	mt  memtable.Memtable
}

func NewLevelDB(path string) LevelDB {
	return LevelDB{
		wal: wal.NewLog(path),
		mt:  memtable.NewMapMT(),
	}
}

func (db *LevelDB) Init() error {
	entries, err := db.wal.Read()
	if err != nil {
		return err
	}

	for _, entry := range entries {
		switch entry.Operator {
		case wal.OP_PUT:
			db.mt.Put(entry.Key, entry.Value)
		case wal.OP_DELETE:
			db.mt.Delete(entry.Key)
		default:
			return errors.New("invalid op code")
		}
	}
	return nil
}

func (db *LevelDB) Get(key []byte) ([]byte, error) {
	return db.mt.Get(key)
}

func (db *LevelDB) Put(key, value []byte) error {
	err := db.wal.Put(key, value)
	if err != nil {
		return err
	}
	db.mt.Put(key, value)

	return nil
}
