package main

import "fmt"

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
	db := Database{tables}
	lm := LocksManager{
		locks: map[string]Lock{},
	}

	fmt.Println(db)
	txn1 := Transaction{
		id:           1,
		locksManager: lm,
		database:     db,
	}
	txn2 := Transaction{
		id:           2,
		locksManager: lm,
		database:     db,
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
	// txn1.locksManager.lockRowS("a", 1)
	// txn2.locksManager.lockRowS("a", 1)
	// txn1.locksManager.unlockRowS("a", 1)
	// txn2.locksManager.unlockRowS("a", 1)
	// println("shared lock test success")

	// deadlock!
	// txn1.locksManager.lockRowX("a", 1)
	// txn2.locksManager.lockRowS("a", 1)

	// exclusive lock test
	// txn1.locksManager.lockRowX("a", 1)
	// txn1.locksManager.unlockRowX("a", 1)
	// txn2.locksManager.lockRowS("a", 1)
	// txn2.locksManager.unlockRowS("a", 1)
	// println("exclusive lock test success")

	// txn1.locksManager.unlockRowS("a", 1)
	// txn2.locksManager.unlockRowS("a", 1)

	/*
	   void sLockDoesNotBlockSLock() throws Throwable {
	       final Waiter waiter = new Waiter();

	       lm.lock(lockName, txn, Lock.LockMode.SHARED);

	       new Thread(() -> {
	           try {
	               lm.lock(lockName, 2, Lock.LockMode.SHARED);
	           } catch (DeadlockException e) {
	           }

	           waiter.resume();
	       }).start();

	       try {
	           waiter.await(10);
	       } catch (TimeoutException e) {
	           fail("Lock was not acquired, but it should have been");
	       }
	   }
	*/

	// test for single writer
}
