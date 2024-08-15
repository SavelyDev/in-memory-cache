package cache

import (
	"errors"
	"sync"
	"time"
)

type cache struct {
	locker  sync.RWMutex
	storage map[string]any
}

func New() *cache {
	return &cache{
		storage: make(map[string]any),
	}
}

func (c *cache) Set(key string, value any, ttl time.Duration) {
	c.locker.Lock()
	defer c.locker.Unlock()

	c.storage[key] = value

	time.AfterFunc(ttl, func() {
		c.Delete(key)
	})
}

func (c *cache) Get(key string) (any, error) {
	c.locker.RLock()
	defer c.locker.RUnlock()

	val, ok := c.storage[key]
	if !ok {
		return nil, errors.New("there is no key")
	}

	return val, nil
}

func (c *cache) Delete(key string) {
	c.locker.Lock()
	defer c.locker.Unlock()

	delete(c.storage, key)
}
