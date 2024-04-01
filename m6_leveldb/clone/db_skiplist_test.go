package main

import (
	"fmt"
	"testing"
)

func TestInsertToTail(t *testing.T) {
	db := NewSkipListDB()

	newNode := &Node{
		key:   []byte("key1"),
		value: []byte("item1"),
		level: 0,
	}
	db.InsertToTail(newNode)

	// check root and subsequent
	fmt.Println("first node")
	fmt.Println(string(db.root.next[0].key))
	fmt.Println(string(db.root.next[0].value))
	fmt.Println(db.root.next[0].level)

	newNode = &Node{
		key:   []byte("key2"),
		value: []byte("item2"),
		level: 2,
	}
	db.InsertToTail(newNode)
	fmt.Println("second node")
	fmt.Println(string(db.root.next[0].next[0].key))
	fmt.Println(string(db.root.next[0].next[0].value))
	fmt.Println(db.root.next[0].next[0].level)

	fmt.Println(string(db.root.next[1].key))
	fmt.Println(string(db.root.next[1].value))
	fmt.Println(db.root.next[1].level)

	fmt.Println(string(db.root.next[2].key))
	fmt.Println(string(db.root.next[2].value))
	fmt.Println(db.root.next[2].level)

	newNode = &Node{
		key:   []byte("key3"),
		value: []byte("item3"),
		level: 1,
	}
	db.InsertToTail(newNode)

	fmt.Println("third node")
	fmt.Println(string(db.root.next[0].next[0].next[0].key))
	fmt.Println(string(db.root.next[0].next[0].next[0].value))
	fmt.Println(db.root.next[0].next[0].next[0].level)

	fmt.Println(string(db.root.next[1].next[1].key))
	fmt.Println(string(db.root.next[1].next[1].value))
	fmt.Println(db.root.next[1].next[1].level)
}
