package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	rowA := []Row{
		{5},
		{10},
		{15},
	}
	rowB := []Row{
		{50},
		{55},
		{60},
	}
	tables := map[string]Table{
		"a": {rows: rowA},
		"b": {rows: rowB},
	}
	db := NewDatabase(tables)
	lm := NewLocksManager(db)

	txn1 := Transaction{
		id: 1,
	}
	txn2 := Transaction{
		id: 2,
	}

	// SIMPLE WAITFORGRAPH TESTS
	// waitForGraph := WaitForGraph{
	// 	nodes: map[int]*Node{},
	// }

	// waitForGraph.AddNode(1)
	// waitForGraph.AddNode(2)
	// waitForGraph.AddEdge(1, 2)
	// fmt.Println("node 1 inneighbors:")
	// fmt.Println(waitForGraph.nodes[1].InNeighbors)
	// fmt.Println("node 1 outneighbors:")
	// fmt.Println(waitForGraph.nodes[1].OutNeighbors)

	// fmt.Println("node 2 inneighbors:")
	// fmt.Println(waitForGraph.nodes[2].InNeighbors)
	// fmt.Println("node 2 outneighbors:")
	// fmt.Println(waitForGraph.nodes[2].OutNeighbors)

	// waitForGraph.AddEdge(2, 1)

	// fmt.Println("node 1 inneighbors:")
	// fmt.Println(waitForGraph.nodes[1].InNeighbors)
	// fmt.Println("node 1 outneighbors:")
	// fmt.Println(waitForGraph.nodes[1].OutNeighbors)

	// fmt.Println("node 2 inneighbors:")
	// fmt.Println(waitForGraph.nodes[2].InNeighbors)
	// fmt.Println("node 2 outneighbors:")
	// fmt.Println(waitForGraph.nodes[2].OutNeighbors)
	// waitForGraph.RemoveEdge(waitForGraph.nodes[1], waitForGraph.nodes[2])
	// waitForGraph.RemoveEdge(waitForGraph.nodes[2], waitForGraph.nodes[1])

	// fmt.Println("node 1 inneighbors:")
	// fmt.Println(waitForGraph.nodes[1].InNeighbors)
	// fmt.Println("node 1 outneighbors:")
	// fmt.Println(waitForGraph.nodes[1].OutNeighbors)

	// fmt.Println("node 2 inneighbors:")
	// fmt.Println(waitForGraph.nodes[2].InNeighbors)
	// fmt.Println("node 2 outneighbors:")
	// fmt.Println(waitForGraph.nodes[2].OutNeighbors)

	// SIMPLE SYNCHRONOUS LOCK TESTS
	// shared lock test
	println("start shared lock test")
	lm.lockRowS("a", 1, txn1)
	lm.lockRowS("a", 1, txn2)
	lm.unlockRowS("a", 1, txn1)
	lm.unlockRowS("a", 1, txn2)
	println("shared lock test success")

	// deadlock!
	// lm.lockRow("a", 1, txn1)
	// lm.lockRowS("a", 1, txn2)

	// exclusive lock test
	println("start exclusive lock test")
	lm.lockRowX("a", 1, txn1)
	lm.unlockRowX("a", 1, txn1)
	lm.lockRowS("a", 1, txn2)
	lm.unlockRowS("a", 1, txn2)
	println("exclusive lock test success")

	// cleanable deadlock
	fmt.Println("deadlock test running")
	var wg sync.WaitGroup
	wg.Add(1)
	go func(txn Transaction) {
		defer wg.Done()
		// 1
		lm.lockRowX("a", 1, txn)
		time.Sleep(2 * time.Second)
		// 3
		lm.lockRowX("b", 1, txn)
		lm.unlockRowX("a", 1, txn)
		lm.unlockRowX("b", 1, txn)
		fmt.Println("txn 1 complete!")
	}(txn1)

	wg.Add(1)
	go func(txn Transaction) {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		// 2
		err := lm.lockRowX("b", 1, txn)
		if err != nil {
			panic(err)
		}

		time.Sleep(3 * time.Second)
		// 4
		err = lm.lockRowX("a", 1, txn)
		if err != nil {
			panic(err)
		}

		time.Sleep(2 * time.Second)
		lm.unlockRowX("a", 1, txn)
		lm.unlockRowX("b", 1, txn)
	}(txn2)
	wg.Wait()
	fmt.Println("deadlock test success")

	// lm.unlockRowS("a", 1)
	// lm.unlockRowS("a", 1)
}
