package leveldb

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
		mt:  memtable.NewSkipListMT(),
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

func (db *LevelDB) ClearWAL() error {
	return db.wal.Clear()
}

func (db *LevelDB) Get(key []byte) ([]byte, error) {
	return db.mt.Get(key)
}

func (db *LevelDB) Put(key, value []byte) error {
	err := db.wal.Put(key, value)
	if err != nil {
		return err
	}
	err = db.mt.Put(key, value)
	if err != nil {
		return err
	}

	return nil
}

func (db *LevelDB) Delete(key []byte) error {
	err := db.wal.Delete(key)
	if err != nil {
		return err
	}
	err = db.mt.Delete(key)
	if err != nil {
		return err
	}

	return nil
}

func (db *LevelDB) store() error {
	// jump through all values in mt
	// for every value, create an sstable entry

	// create sstable file

	// empty wal

	// empty mt

	return nil
}
