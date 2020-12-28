package gkv

import (
	"sync"
	"time"
)

type Gkv struct {
	storage map[string]*item
	mu      *sync.RWMutex
}

type item struct {
	Value          interface{}
	ExpirationTime time.Duration
}

func New() *Gkv {
	obj := &Gkv{
		storage: make(map[string]*item),
		mu:      new(sync.RWMutex),
	}
	return obj
}

// 设置一个key/value数据
func (g *Gkv) Set(k string, v interface{}, ex time.Duration) (bool, error) {
	g.mu.Lock()
	g.storage[k] = &item{
		Value:          v,
		ExpirationTime: ex,
	}
	g.mu.Unlock()
	return true, nil
}

// 根据key获取数据
func (g *Gkv) Get(k string) (interface{}, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	if v, ok := g.storage[k]; ok {
		return v.Value, nil
	}
	return nil, nil
}

// 根据key清除对应数据
func (g *Gkv) Del(k ...string) (bool, error) {
	g.mu.Lock()
	for i := 0; i < len(k); i++ {
		delete(g.storage, k[i])
	}
	g.mu.Unlock()
	return true, nil
}

// 检查key是否存在
func (g *Gkv) Exists(k string) (bool, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	_, ok := g.storage[k]
	return ok, nil
}

// 获取当前所有数据
func (g *Gkv) GetAll() map[string]interface{} {
	g.mu.RLock()
	defer g.mu.RUnlock()
	m := make(map[string]interface{}, len(g.storage))
	for k, v := range g.storage {
		m[k] = v.Value
	}
	return m
}

// 清除当前所有数据
func (g *Gkv) FlushCache() (bool, error) {
	g.mu.Lock()
	g.storage = make(map[string]*item)
	g.mu.Unlock()
	return true, nil
}
