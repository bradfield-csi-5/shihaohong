package wal

import (
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
	wal.appendToLog(OP_PUT, key, value)
	return nil
}

func (wal *Log) Delete(key []byte) error {
	wal.appendToLog(OP_DELETE, key, nil)
	return nil
}

func (wal *Log) appendToLog(operator byte, key, value []byte) error {
	f, err := os.OpenFile(wal.path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	entry := Entry{
		operator:    operator,
		key:         key,
		keyLength:   byte(len(key)),
		value:       value,
		valueLength: byte(len(value)),
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

func (wal *Log) ReadFromLog() ([]byte, error) {

	// read operator
	// read key
	// read value
	//

	return nil, nil
}

// for now, a delete just stores val len 0, val empty
type Entry struct {
	operator    byte
	keyLength   byte
	key         []byte
	valueLength byte
	value       []byte
}

func (e *Entry) encode() []byte {
	bs := []byte{
		e.operator,
		e.keyLength,
	}

	bs = append(bs, e.key...)
	bs = append(bs, e.valueLength)
	bs = append(bs, e.value...)
	return bs
}
