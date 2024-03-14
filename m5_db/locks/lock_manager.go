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
	lockType   LockType
}

type LockStatus struct {
	err error
}

type Lock struct {
	name     string
	owners   map[int]LockTransaction
	waiting  []LockTransaction
	lockType LockType
}

type LockType string

const (
	LockX   = "X"
	LockS   = "S"
	LockNil = "N"
)

func NewLock(name string) *Lock {
	return &Lock{
		name:     name,
		owners:   map[int]LockTransaction{},
		waiting:  []LockTransaction{},
		lockType: LockNil,
	}
}

func (lm *LocksManager) lockRowS(table string, idx int, txn Transaction) error {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		lock = NewLock(key)
		lm.locks[key] = lock
	}

	lt := LockTransaction{
		txn:        txn,
		lockStatus: make(chan LockStatus),
		lockType:   LockS,
	}

	if len(lock.owners) != 0 && lock.lockType != LockS {
		lock.waiting = append(lock.waiting, lt)
		res := <-lt.lockStatus
		if res.err != nil {
			return res.err
		}
	}
	lock.owners[txn.id] = lt
	lock.lockType = LockS
	return nil
}

func (lm *LocksManager) unlockRowS(table string, idx int, txn Transaction) {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		panic("attempting to unlock from nonexistent lock")
	}

	_, exists = lock.owners[txn.id]
	if !exists {
		return
	}
	delete(lock.owners, txn.id)
	if len(lock.owners) == 0 {
		lock.lockType = LockNil
	}

	// no more waiting transactions, nothing to do
	if len(lock.waiting) == 0 {
		return
	}

	wt := lock.waiting[0]
	lock.waiting = lock.waiting[1:]
	lock.owners[wt.txn.id] = wt
	lock.lockType = wt.lockType
	wt.lockStatus <- LockStatus{err: nil}
}

func (lm *LocksManager) lockRowX(table string, idx int, txn Transaction) error {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		lock = NewLock(key)
		lm.locks[key] = lock
	}

	lt := LockTransaction{
		txn:        txn,
		lockStatus: make(chan LockStatus, 1),
		lockType:   LockX,
	}
	if len(lock.owners) != 0 {
		lock.waiting = append(lock.waiting, lt)
		res := <-lt.lockStatus
		if res.err != nil {
			return res.err
		}
	}
	lock.owners[txn.id] = lt
	return nil
}

func (lm *LocksManager) unlockRowX(table string, idx int, txn Transaction) {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		panic("attempting to unlock from nonexistent lock")
	}

	_, exists = lock.owners[txn.id]
	if !exists {
		return
	}
	delete(lock.owners, txn.id)

	if len(lock.owners) == 0 {
		lock.lockType = LockNil
	}
	// no more waiting transactions, nothing to do
	if len(lock.waiting) == 0 {
		return
	}

	wt := lock.waiting[0]
	lock.waiting = lock.waiting[1:]
	lock.owners[wt.txn.id] = wt
	lock.lockType = wt.lockType
	wt.lockStatus <- LockStatus{err: nil}
}

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
				fmt.Println("====== DEADLOCK DETECTED =====")
				// TODO: for now just kill first node, should clear the cycle-causing node in reality
				lock := wfg.nodes[i].OutNeighbors[0].associatedLock
				lock.owners = make(map[int]LockTransaction)
				if len(lock.waiting) != 0 {
					nt := lock.waiting[0]
					lock.owners[nt.txn.id] = nt
					lock.waiting = lock.waiting[1:]
					nt.lockStatus <- LockStatus{
						err: nil,
					}
				}
			}
		}
	}
}

func getLockHashKey(table string, idx int) string {
	return fmt.Sprintf("%s%d", table, idx)
}
