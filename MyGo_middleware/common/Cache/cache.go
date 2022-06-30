package caches

import (
	"fmt"
	"time"
)

// Cache interface contains all behaviors for cache adapter.
// usage:
//	cache.Register("file",cache.NewFileCache) // this operation is run in init method of file.go.
//	c,err := cache.NewCache("file","{....}")
//	c.Put("key",value, 3600 * time.Second)
//	v := c.Get("key")
//
//	c.Incr("counter")  // now is 1
//	c.Incr("counter")  // now is 2
//	count := c.Get("counter").(int)
type Cache interface {
	// get cached value by key.
	Get(key string) (string, error)
	// get cached value by key.
	Expire(key string, expire int) error
	// get cached value by key.
	HGet(hashKey string, key string) (string, error)
	// hset cached value by key.
	HSet(hashKey string, key string, value string) error
	// get cached value by key.
	HGetAll(hashKey string) ([]string, error)
	// GetMulti is a batch version of Get.
	GetMulti(keys []string) ([]string, error)
	// set cached value with key and expire time.
	Put(key string, val interface{}, timeout time.Duration) error
	// set cached value with key
	Set(key string, val interface{}) error
	// increase cached int value by key, as a counter.
	HPut(hashKey string, key string, val interface{}, timeout time.Duration) error
	// delete cached value by key.
	Delete(key string) error
	// increase cached int value by key, as a counter.
	HDelete(hashKey string, key string) error
	// increase cached int value by key, as a counter.
	Incr(key string) error
	// increase cached int value by key, as a counter.
	HIncr(hashKey string, key string) error
	// decrease cached int value by key, as a counter.
	Decr(key string) error
	// increase cached int value by key, as a counter.
	HDecr(hashKey string, key string) error
	// check if cached value exists or not.
	IsExist(key string) bool
	// check if cached value exists or not.
	HIsExist(hashKey string, key string) bool
	// clear all cache.
	ClearAll() error
	// start gc routine based on config string settings.
	StartAndGC(config string) error
	//increase count int value by key, as a counter.
	HInDecrBatch(hashKey string, key string, count int64) error
}

// Instance is a function create a new Cache Instance
type Instance func() Cache

var adapters = make(map[string]Instance)

// Register makes a cache adapter available by the adapter name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, adapter Instance) {
	if adapter == nil {
		panic("cache: Register adapter is nil")
	}
	if _, ok := adapters[name]; ok {
		panic("cache: Register called twice for adapter " + name)
	}
	adapters[name] = adapter
}

// NewCache Create a new cache driver by adapter name and config string.
// config need to be correct JSON as string: {"interval":360}.
// it will start gc automatically.
func NewCache(adapterName, config string) (adapter Cache, err error) {
	instanceFunc, ok := adapters[adapterName]
	if !ok {
		err = fmt.Errorf("cache: unknown adapter name %q (forgot to import?)", adapterName)
		return
	}
	adapter = instanceFunc()
	err = adapter.StartAndGC(config)
	if err != nil {
		adapter = nil
	}
	return
}
