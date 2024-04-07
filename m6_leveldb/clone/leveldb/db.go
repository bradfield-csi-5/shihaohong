package leveldb

import (
	"errors"

	"github.com/shihaohong/leveldb_clone/memtable"
	"github.com/shihaohong/leveldb_clone/sstable"
	"github.com/shihaohong/leveldb_clone/wal"
)

/*
TODOs:
- db.Store: Merge existing sstable if it exists. Assume only one file
- db.Get: check SSTable for result if not in Memtable
- db.Put: call db.Store when byte estimate exceeds 1k
*/

type LevelDB struct {
	wal     wal.Log
	mt      memtable.Memtable
	sstPath string
}

func NewLevelDB(walPath string, sstPath string) LevelDB {
	return LevelDB{
		wal:     wal.NewLog(walPath),
		mt:      memtable.NewSkipListMT(),
		sstPath: sstPath,
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

func (db *LevelDB) ClearMT() error {
	return db.mt.Clear()
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

func (db *LevelDB) Store() error {
	// jump through all values in mt
	// for every value, create an sstable entry
	vals, err := db.mt.GetAll()
	if err != nil {
		return err
	}

	// create sstable file
	sst := sstable.NewSSTable(vals)
	err = sst.WriteToDisk(db.sstPath)
	if err != nil {
		return err
	}

	// empty wal
	// empty mt
	db.wal.Clear()
	db.mt.Clear()

	return nil
}
