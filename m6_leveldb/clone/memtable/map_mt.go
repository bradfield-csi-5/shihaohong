package memtable

import (
	"sort"
)

// Initial version, in-memory and no persistence
type MapMemtable struct {
	data map[string][]byte
}

func (mt *MapMemtable) Get(key []byte) (value []byte, err error) {
	return mt.data[string(key)], nil
}

func (mt *MapMemtable) Has(key []byte) (ret bool, err error) {
	_, ok := mt.data[string(key)]
	return ok, nil
}

func (mt *MapMemtable) Put(key, value []byte) error {
	mt.data[string(key)] = value
	return nil
}

func (mt *MapMemtable) Delete(key []byte) error {
	delete(mt.data, string(key))
	return nil
}

func (mt *MapMemtable) RangeScan(start, limit []byte) (Iterator, error) {
	startKey := string(start)
	limitKey := string(limit)

	keys := make([]string, 0)
	for k := range mt.data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	startFound := false
	iterator := NewIter()
	for k := range keys {
		// when start is found, toggle
		if !startFound && keys[k] == startKey {
			startFound = true
		}

		// continue iterating if we haven't found the starting key
		if !startFound {
			continue
		}

		// append to iterator tuple list
		tuple := Tuple{
			key:   []byte(keys[k]),
			value: mt.data[keys[k]],
		}
		iterator.tuples = append(iterator.tuples, tuple)

		// reached limit key, break out
		if keys[k] == limitKey {
			break
		}
	}

	// if len(iterator.tuples) == 0 {
	// 	return nil, errors.New("no tuples within the range found")
	// }

	return iterator, nil
}

func NewMapMT() *MapMemtable {
	return &MapMemtable{
		data: make(map[string][]byte),
	}
}

type Iter struct {
	tuples []Tuple
	index  int
}

type Tuple struct {
	key   []byte
	value []byte
}

func (it *Iter) Next() bool {
	if it.index < len(it.tuples) {
		it.index++
	}
	return it.index < len(it.tuples)
}

// TODO: is an error possible for a memory store?
func (it *Iter) Error() error {
	return nil
}

func (it *Iter) Key() []byte {
	if it.index >= len(it.tuples) {
		return nil
	}
	return it.tuples[it.index].key
}

func (it *Iter) Value() []byte {
	if it.index >= len(it.tuples) {
		return nil
	}
	return it.tuples[it.index].value
}

func NewIter() *Iter {
	return &Iter{
		tuples: make([]Tuple, 0),
		index:  0,
	}
}
