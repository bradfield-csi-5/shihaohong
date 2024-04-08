package leveldb

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/shihaohong/leveldb_clone/consts"
	"github.com/shihaohong/leveldb_clone/memtable"
	"github.com/shihaohong/leveldb_clone/sstable"
	"github.com/shihaohong/leveldb_clone/wal"
)

const manifestPath = "CURRENT"
const byteSizeMax = 2048

/*
TODOs:
- db.Store: Create new sstables, enumerate them? [done]
- db.Get: check SSTable for result if not in Memtable
- db.Put: call db.Store when byte estimate exceeds 1k
*/

type LevelDB struct {
	wal     wal.Log
	mt      memtable.Memtable
	sstPath string
}

func NewLevelDB(walPath string, sstPath string) LevelDB {
	ldb := LevelDB{
		wal:     wal.NewLog(walPath),
		mt:      memtable.NewSkipListMT(),
		sstPath: sstPath,
	}
	ldb.init()
	return ldb
}

func (db *LevelDB) init() error {
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
	res, err := db.mt.Get(key)
	// fmt.Println("GET")
	// fmt.Println(string(res))
	// fmt.Println(err)
	if err == nil && res != nil {
		return res, nil
	}

	if !errors.Is(err, consts.ErrSearchKeyNotFound) {
		return nil, err
	}
	currentEntry, err := db.getCurrentManifestEntry()
	if err != nil {
		return nil, err
	}
	manifestInt, err := getManifestIndexAsInt(currentEntry)
	if err != nil {
		return nil, err
	}
	for i := manifestInt; i >= 0; i-- {
		sstPath := fmt.Sprintf("%s-%06d", db.sstPath, i)
		sst, err := sstable.NewSSTableFromPath(sstPath)
		if err != nil {
			return nil, err
		}
		// fmt.Println("got here")
		// fmt.Println(i)
		res, err := sst.Get(key)
		// fmt.Println("srch key")
		// fmt.Println(err)
		// fmt.Println(errors.Is(err, consts.ErrSearchKeyNotFound))
		if !errors.Is(err, consts.ErrSearchKeyNotFound) && err != nil {
			return nil, err
		} else if res != nil && err == nil {
			return res, nil
		}
	}

	return nil, consts.ErrSearchKeyNotFound
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

	if db.mt.GetByteEstimate() > byteSizeMax {
		fmt.Println("byte est:")
		fmt.Println(db.mt.GetByteEstimate())
		err = db.Store()
		if err != nil {
			return err
		}
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

	// write to the correct manifest file
	manifest, err := db.getNextManifestEntry()
	if err != nil {
		return err
	}
	err = sst.WriteToDisk(manifest)
	if err != nil {
		return err
	}
	err = db.updateManifestEntry(manifest)
	if err != nil {
		return err
	}

	// empty wal
	// empty mt
	db.wal.Clear()
	db.mt.Clear()

	return nil
}

// Manifests 0-indexed
func (db *LevelDB) getNextManifestEntry() (string, error) {
	f, err := os.OpenFile(manifestPath, os.O_RDONLY, 0644)
	if os.IsNotExist(err) {
		manifestZero := fmt.Sprintf("%s-%06d", db.sstPath, 0)
		// handle the case where the file doesn't exist
		err = os.WriteFile(manifestPath, []byte(manifestZero), 0666)
		if err != nil {
			return "", nil
		}
		return manifestZero, nil
	} else if err != nil {
		return "", err
	}
	defer f.Close()

	res, err := os.ReadFile(manifestPath)
	if err != nil {
		return "", err
	}

	resString := string(res)
	resString, err = incrementManifestCount(resString)
	if err != nil {
		return "", err
	}
	return resString, nil
}

func (db *LevelDB) getCurrentManifestEntry() (string, error) {
	res, err := os.ReadFile(manifestPath)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func incrementManifestCount(input string) (string, error) {
	// Using regular expression to split the input string into non-numeric and numeric parts
	re := regexp.MustCompile(`^(.*?)(\d+)$`)
	matches := re.FindStringSubmatch(input)
	if len(matches) != 3 {
		// Return original input if pattern doesn't match
		return "", errors.New("unexpected file name format")
	}

	// Extracting the non-numeric and numeric parts
	prefix := matches[1]
	numStr := matches[2]

	// Converting the numeric part to integer
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return "", err
	}

	// Incrementing the number
	num++
	updatedNumStr := fmt.Sprintf("%06d", num)
	updatedString := prefix + updatedNumStr

	return updatedString, nil
}

func (db *LevelDB) updateManifestEntry(manifest string) error {
	err := os.WriteFile(manifestPath, []byte(manifest), 0666)
	if err != nil {
		return err
	}
	return nil
}

func getManifestIndexAsInt(manifestPath string) (int, error) {
	// Using regular expression to split the input string into non-numeric and numeric parts
	re := regexp.MustCompile(`^(.*?)(\d+)$`)
	matches := re.FindStringSubmatch(manifestPath)
	if len(matches) != 3 {
		// Return original input if pattern doesn't match
		return -1, errors.New("unexpected file name format")
	}

	// Extracting the non-numeric and numeric parts
	numStr := matches[2]

	// Converting the numeric part to integer
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return -1, err
	}
	return num, nil
}
