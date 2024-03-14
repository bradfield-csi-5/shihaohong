package main

import (
	// "errors"
	"fmt"
	"time"
)

type LocksManager struct {
	locks map[string]*Lock
}

func NewLocksManager() *LocksManager {
	lm := &LocksManager{
		locks: map[string]*Lock{},
	}
	fmt.Println("starting deadlock checker")
	go lm.checkForDeadlock()
	return lm
}

type LockTransaction struct {
	txn        Transaction
	lockStatus chan LockStatus
}

type LockStatus struct {
	success bool
	err     error
}

type Lock struct {
	name    string
	owners  map[int]LockTransaction
	waiting []LockTransaction
}

func NewLock(name string) *Lock {
	return &Lock{
		name:    name,
		owners:  map[int]LockTransaction{},
		waiting: []LockTransaction{},
	}
}

// func (lm *LocksManager) lockRowS(table string, idx int, txn Transaction) {
// 	// mutex r lock
// 	key := getLockHashKey(table, idx)
// 	lock, exists := lm.locks[key]
// 	if !exists {
// 		lock = NewLock()
// 		lm.locks[key] = lock
// 	}

// 	lock.mutex.RLock()
// 	lock.owners[idx] = txn
// 	lock.sLockCount++
// }

// func (lm *LocksManager) unlockRowS(table string, idx int, txn Transaction) {
// 	key := getLockHashKey(table, idx)
// 	lock, exists := lm.locks[key]
// 	if !exists {
// 		panic("attempting to unlock from nonexistent transaction")
// 	}

// 	lock.mutex.RUnlock()
// 	delete(lock.owners, idx)
// 	lock.sLockCount--
// }

// func (lm *LocksManager) unlockRowX(table string, idx int, txn Transaction) {
// 	key := getLockHashKey(table, idx)
// 	lock, exists := lm.locks[key]
// 	if !exists {
// 		panic("attempting to unlock from nonexistent transaction")
// 	}

// 	lock.mutex.Unlock()
// 	delete(lock.owners, idx)
// 	lock.xLockCount = 0
// }

func (lm *LocksManager) checkForDeadlock() {
	for {
		time.Sleep(3 * time.Second)
		fmt.Println("checking for deadlocks...")
		wfg := NewWaitForGraph()

		for _, lock := range lm.locks {
			// add owners as nodes
			for _, owner := range lock.owners {
				wfg.AddNode(owner.txn.id)
			}
		}

		for _, lock := range lm.locks {
			for _, owner := range lock.owners {
				for _, waiter := range lock.waiting {
					wfg.AddEdge(waiter.txn.id, owner.txn.id, *lock)
				}
			}
		}

		// check deadlock
		for i := range wfg.nodes {
			if wfg.hasCycles(i, map[int]bool{i: true}) {
				fmt.Println("DEADLOCK DETECTED OMG")
				// for now just kill first node, should clear the cycle-causing node in reality
				lock := wfg.nodes[i].OutNeighbors[0].associatedLock
				fmt.Println("lock to clean out:")
				fmt.Println(lock.name)

				fmt.Println("txn to vanquish")
				fmt.Println(lock.owners[2].txn.id)

				lock.owners = make(map[int]LockTransaction)
				fmt.Println("waiting txns")
				fmt.Println(lock.waiting)
				if len(lock.waiting) != 0 {
					nt := lock.waiting[0]
					lock.owners[nt.txn.id] = nt
					lock.waiting = lock.waiting[1:]
					fmt.Println("next txn given lock:")
					fmt.Println(nt.txn.id)
					nt.lockStatus <- LockStatus{
						success: true,
						err:     nil,
					}
				}
			}
		}
	}
}

func (lm *LocksManager) lockRow(table string, idx int, txn Transaction) {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		lock = NewLock(key)
		lm.locks[key] = lock
	}

	lt := LockTransaction{
		txn:        txn,
		lockStatus: make(chan LockStatus, 1),
	}
	if len(lock.owners) != 0 {
		wt := LockTransaction{
			txn:        txn,
			lockStatus: make(chan LockStatus, 1),
		}
		lock.waiting = append(lock.waiting, wt)
		fmt.Println("waiting for lock")
		res := <-wt.lockStatus
		fmt.Println("stopped waiting")
		if res.err != nil {
			fmt.Println("received lock successfully")
		} else {
			fmt.Println("error receiving lock, canceled request due to deadlock")
		}
	}
	lock.owners[txn.id] = lt
}

func (lm *LocksManager) unlockRow(table string, idx int, txn Transaction) {
	fmt.Println("unlocking row for")
	key := getLockHashKey(table, idx)
	fmt.Println(txn.id)
	fmt.Println(key)
	lock, exists := lm.locks[key]
	if !exists {
		fmt.Println("lock does not exist")
		panic("attempting to unlock from nonexistent lock")
	}

	_, exists = lock.owners[txn.id]
	fmt.Println("exists")
	fmt.Println(exists)

	if !exists {
		fmt.Println("trying to unlock with an owner that does not exist")
		fmt.Println(txn.id)
		fmt.Println(table)
		fmt.Println(idx)
		return
	}
	delete(lock.owners, idx)

	// no more waiting transactions, nothing to do
	if len(lock.waiting) == 0 {
		return
	}

	wt := lock.waiting[0]
	lock.waiting = lock.waiting[1:]
	lock.owners[wt.txn.id] = wt
	wt.lockStatus <- LockStatus{success: true, err: nil}
}

func getLockHashKey(table string, idx int) string {
	return fmt.Sprintf("%s%d", table, idx)
}
