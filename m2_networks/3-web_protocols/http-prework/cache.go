package main

import "sync"

type Cache struct {
	hm map[string][]byte
	m  sync.Mutex
}

func NewCache() *Cache {
	c := &Cache{
		hm: make(map[string][]byte),
	}
	return c
}

func (c *Cache) Put(k string, v []byte) {
	c.m.Lock()
	_, ok := c.hm[k]
	if !ok {
		c.hm[k] = v
	}
	c.m.Unlock()
}

func (c *Cache) Get(k string) (v []byte) {
	c.m.Lock()
	if cv, ok := c.hm[k]; ok {
		v = cv
	}
	c.m.Unlock()
	return
}
