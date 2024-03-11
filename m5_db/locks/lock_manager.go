package main

import (
	"fmt"
	"sync"
	"time"
)

type LocksManager struct {
	locks map[string]Lock
}

type Lock struct {
	sLockCount int
	xLockCount int
	mutex      *sync.RWMutex
}

func (lm *LocksManager) lockRowS(table string, idx int) {
	// mutex r lock
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		lock = Lock{
			sLockCount: 0,
			xLockCount: 0,
			mutex:      &sync.RWMutex{},
		}
		lm.locks[key] = lock
	}

	lock.mutex.RLock()
	lock.sLockCount++
}

func (lm *LocksManager) lockRowX(table string, idx int) {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		lock = Lock{}
		lm.locks[key] = lock
	}

	lock.mutex.Lock()
	lock.xLockCount++
}

func (lm *LocksManager) unlockRowS(table string, idx int) {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		lock = Lock{}
		lm.locks[key] = lock
	}

	lock.mutex.RUnlock()
	lock.sLockCount--
	if lock.sLockCount == 0 && lock.xLockCount == 0 {
		delete(lm.locks, key)
	}
}

func (lm *LocksManager) unlockRowX(table string, idx int) {
	key := getLockHashKey(table, idx)
	lock, exists := lm.locks[key]
	if !exists {
		lock = Lock{}
		lm.locks[key] = lock
	}

	lock.mutex.Unlock()
	lock.sLockCount--
	if lock.sLockCount == 0 && lock.xLockCount == 0 {
		delete(lm.locks, key)
	}
}

func (lm *LocksManager) checkForDeadlock() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("one second elapsed")
		// check deadlock
	}
}

func getLockHashKey(table string, idx int) string {
	return fmt.Sprintf("%s%d", table, idx)
}
