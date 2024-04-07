package memtable

import (
	"bytes"
	"errors"
	"math/rand"

	"github.com/shihaohong/leveldb_clone/iterator"
)

// probability constant for when to increment the level when determining random level
const p = 0.25
const maxLevel = 5

type Node struct {
	key   []byte
	value []byte

	// Array of pointers to next nodes
	// next[0] is bottom level ("slow lane")
	// next[maxLevel - 1] is top level ("express lane")
	next [maxLevel]*Node

	// Max valid level for this node
	level int

	isLastNode bool
}

type SkipListMemtable struct {
	root *Node
}

func NewSkipListMT() *SkipListMemtable {
	// root node should have no value, need to start somewhere
	rootNode := &Node{}

	// there should also be a "nil node" to signal termination for that level
	// should be impossible to have a "nil intermediate node"
	nilNode := &Node{isLastNode: true}

	for lvl := maxLevel - 1; lvl >= 0; lvl-- {
		rootNode.next[lvl] = nilNode
	}

	return &SkipListMemtable{
		root: rootNode,
	}
}

func (db *SkipListMemtable) Get(key []byte) (value []byte, err error) {
	currNode := db.root

	for i := maxLevel - 1; i >= 0; i-- {
		for !currNode.next[i].isLastNode && bytes.Compare(currNode.next[i].key, key) < 0 {
			currNode = currNode.next[i]
		}
	}

	currNode = currNode.next[0]
	if bytes.Equal(currNode.key, key) {
		return currNode.value, nil
	}

	return nil, errors.New("search key not found")
}

func (db *SkipListMemtable) Has(key []byte) (ret bool, err error) {
	_, err = db.Get(key)
	if err != nil {
		return true, nil
	}
	return false, errors.New("search key not found")
}

func (db *SkipListMemtable) Put(key, value []byte) error {
	var update [maxLevel]*Node
	currNode := db.root

	// for every level, find the node to update
	for i := maxLevel - 1; i >= 0; i-- {
		for !currNode.next[i].isLastNode && bytes.Compare(currNode.next[i].key, key) < 0 {
			currNode = currNode.next[i]
		}
		update[i] = currNode
	}
	currNode = currNode.next[0]

	// if the search key is found, just replace it
	if bytes.Equal(currNode.key, key) {
		currNode.value = value
	} else {
		// else
		lvl := randomLevel()
		newNode := &Node{
			key:   key,
			value: value,
			level: lvl,
		}
		for i := 0; i < maxLevel; i++ {
			newNode.next[i] = update[i].next[i]
			update[i].next[i] = newNode
		}
	}
	return nil
}

func (db *SkipListMemtable) Delete(key []byte) error {
	var update [maxLevel]*Node
	currNode := db.root

	for i := maxLevel - 1; i >= 0; i-- {
		for !currNode.next[i].isLastNode && bytes.Compare(currNode.next[i].key, key) < 0 {
			currNode = currNode.next[i]
		}
		update[i] = currNode
	}
	currNode = currNode.next[0]

	if !bytes.Equal(currNode.key, key) {
		return errors.New("search key not found")
	}

	for i := 0; i < maxLevel; i++ {
		if update[i].next[i] != currNode {
			break
		}
		update[i].next[i] = currNode.next[i]
	}
	return nil
}

func (db *SkipListMemtable) RangeScan(start, limit []byte) (iterator.Iterator, error) {
	return nil, nil
}

// Use to seed the skip list, for testing only since its assumed that
// the passed in values are sorted by key
func (db *SkipListMemtable) InsertToTail(node *Node) error {
	currNode := db.root
	for lvl := maxLevel - 1; lvl >= 0; lvl-- {
		// assume nil key means nil node
		for !currNode.next[lvl].isLastNode {
			currNode = currNode.next[lvl]
		}
		node.next[lvl] = currNode.next[lvl]
		currNode.next[lvl] = node
	}
	return nil
}

func randomLevel() int {
	lvl := 0
	for rand.Float32() > p && lvl < maxLevel-1 {
		lvl += 1
	}
	return lvl
}
