package main

const maxLevel = 3

type Node struct {
	key   []byte
	value []byte

	// Array of pointers to next nodes
	// next[0] is bottom level ("slow lane")
	// next[maxLevel - 1] is top level ("express lane")
	next [maxLevel]*Node

	// Max valid level for this node
	level int
}

type SkipListDB struct {
	root *Node
}

func NewSkipListDB() SkipListDB {
	// root node should have no value, need to start somewhere
	rootNode := &Node{}

	// there should also be a "nil node" to signal termination for that level
	// should be impossible to have a "nil intermediate node"
	nilNode := &Node{}

	for lvl := maxLevel - 1; lvl >= 0; lvl-- {
		rootNode.next[lvl] = nilNode
	}

	return SkipListDB{
		root: rootNode,
	}
}

func (db *SkipListDB) Get(key []byte) (value []byte, err error) {
	return nil, nil
}

func (db *SkipListDB) Has(key []byte) (ret bool, err error) {
	return true, nil
}

func (db *SkipListDB) Put(key, value []byte) error {
	return nil
}

func (db *SkipListDB) Delete(key []byte) error {
	return nil
}

// Use to seed the skip list
func (db *SkipListDB) InsertToTail(node *Node) error {
	for lvl := node.level; lvl >= 0; lvl-- {
		currNode := db.root
		// assume nil key means nil node
		for currNode.next[lvl].key != nil {
			currNode = currNode.next[lvl]
		}
		nilNode := currNode.next[lvl]
		currNode.next[lvl] = node
		node.next[lvl] = nilNode
	}
	return nil
}
