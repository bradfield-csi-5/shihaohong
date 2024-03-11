package main

import (
	"fmt"
	"sync"
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

type Lock struct {
	sLockCount int
	xLockCount int
	owners     map[int]Transaction
	waiting    []Transaction
	mutex      *sync.RWMutex
}

func NewLock() *Lock {
	return &Lock{
		sLockCount: 0,
		xLockCount: 0,
		owners:     map[int]Transaction{},
		waiting:    []Transaction{},
		mutex:      &sync.RWMutex{},
	}
}

func (lm *LocksManager) lockRowS(table string, idx int, txn Transaction) {
	// mutex r lock
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		lock = NewLock()
		lm.locks[key] = lock
	}

	lock.mutex.RLock()
	lock.owners[idx] = txn
	lock.sLockCount++
}

func (lm *LocksManager) lockRowX(table string, idx int, txn Transaction) {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		lock = NewLock()
		lm.locks[key] = lock
	}

	if len(lock.owners) != 0 {
		lock.waiting = append(lock.waiting, txn)
	}
	lock.mutex.Lock()
	// only one exclusive lock allowed, so clear and move txn to owner
	lock.xLockCount = 1
	lock.sLockCount = 0
	lock.owners = make(map[int]Transaction)
	lock.owners[txn.id] = txn

	// if lock was in waiting, move it to owner
	index := -1
	for i, txn := range lock.waiting {
		if txn.id == idx {
			index = i
		}
	}
	if index != -1 {
		lock.waiting = append(lock.waiting[:index], lock.waiting[index+1:]...)
	}
}

func (lm *LocksManager) unlockRowS(table string, idx int, txn Transaction) {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		panic("attempting to unlock from nonexistent transaction")
	}

	lock.mutex.RUnlock()
	delete(lock.owners, idx)
	lock.sLockCount--
}

func (lm *LocksManager) unlockRowX(table string, idx int, txn Transaction) {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		panic("attempting to unlock from nonexistent transaction")
	}

	lock.mutex.Unlock()
	delete(lock.owners, idx)
	lock.xLockCount = 0
}

func (lm *LocksManager) checkForDeadlock() {
	for {
		time.Sleep(3 * time.Second)
		fmt.Println("checking for deadlocks...")
		wfg := NewWaitForGraph()

		for _, lock := range lm.locks {
			// add owners as nodes
			for _, owner := range lock.owners {
				wfg.AddNode(owner.id)
			}
		}

		for _, lock := range lm.locks {
			for _, owner := range lock.owners {
				for _, waiter := range lock.waiting {
					wfg.AddEdge(waiter.id, owner.id)
				}
			}
		}

		// check deadlock
		for i := range wfg.nodes {
			if wfg.hasCycles(i, map[int]bool{i: true}) {
				fmt.Println("DEADLOCK DETECTED OMG")
				panic("TODO: handle cycle gracefully")
			}
		}
	}
}

func getLockHashKey(table string, idx int) string {
	return fmt.Sprintf("%s%d", table, idx)
}
