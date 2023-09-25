package main

import (
	"fmt"
	"sync"
	"time"
)

// 10 seconds TTL
const maxTTL = 10

type item struct {
	value      []byte
	lastAccess int64
}

// cache from https://stackoverflow.com/a/25487392
type Cache struct {
	hm map[string]*item
	m  sync.Mutex
}

func NewCache() *Cache {
	c := &Cache{
		hm: make(map[string]*item),
	}
	go func() {
		for now := range time.Tick(time.Second) {
			c.m.Lock()
			for k, v := range c.hm {
				if now.Unix()-v.lastAccess > int64(maxTTL) {
					fmt.Println("deleting: ", c.hm)
					delete(c.hm, k)
				}
			}
			c.m.Unlock()
		}
	}()
	return c
}

func (c *Cache) Put(k string, v []byte) {
	c.m.Lock()
	it, ok := c.hm[k]
	if !ok {
		it = &item{value: v}
		c.hm[k] = it
	}
	it.lastAccess = time.Now().Unix()
	c.m.Unlock()
}

func (c *Cache) Get(k string) (v []byte) {
	c.m.Lock()
	if it, ok := c.hm[k]; ok {
		v = it.value
		it.lastAccess = time.Now().Unix()
	}
	c.m.Unlock()
	return
}
