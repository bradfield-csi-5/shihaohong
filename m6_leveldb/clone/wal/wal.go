package wal

import (
	"errors"
	"io"
	"os"
)

const (
	OP_PUT = iota
	OP_DELETE
)

type Log struct {
	path string
}

func (wal *Log) ClearLog() error {
	if err := os.Truncate(wal.path, 0); err != nil {
		return err
	}
	return nil
}

func (wal *Log) Put(key, value []byte) error {
	err := wal.appendToLog(OP_PUT, key, value)
	if err != nil {
		return err
	}
	return nil
}

func (wal *Log) Delete(key []byte) error {
	err := wal.appendToLog(OP_DELETE, key, nil)
	if err != nil {
		return err
	}
	return nil
}

func (wal *Log) appendToLog(operator byte, key, value []byte) error {
	f, err := os.OpenFile(wal.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	entry := Entry{
		operator: operator,
		key:      key,
		keyLen:   byte(len(key)),
		value:    value,
		valueLen: byte(len(value)),
	}
	data := entry.encode()
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func NewLog(path string) Log {
	return Log{
		path: path,
	}
}

func (wal *Log) Read() ([]Entry, error) {
	res := make([]Entry, 0)

	f, err := os.Open(wal.path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	for {
		opSlice := make([]byte, 1)
		_, err := f.Read(opSlice)
		// special err check here to check if end of WAL is reached
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		} else if err != nil && errors.Is(err, io.EOF) {
			return res, nil
		}

		op := opSlice[0]
		keyLenSlice := make([]byte, 1)
		_, err = f.Read(keyLenSlice)
		if err != nil {
			return nil, err
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

		entry := Entry{
			operator: op,
			keyLen:   keyLen,
			key:      keySlice,
			valueLen: valLen,
			value:    valSlice,
		}

		res = append(res, entry)
	}
}

// for now, a delete just stores val len 0, val empty
type Entry struct {
	operator byte
	keyLen   byte
	key      []byte
	valueLen byte
	value    []byte
}

func (e *Entry) encode() []byte {
	bs := []byte{
		e.operator,
		e.keyLen,
	}

	bs = append(bs, e.key...)
	bs = append(bs, e.valueLen)
	bs = append(bs, e.value...)
	return bs
}
