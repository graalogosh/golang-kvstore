package kvstore

import (
	"context"
	"sync"
)

type KVStorage interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Put(ctx context.Context, key string, val interface{}) error
	Delete(ctx context.Context, key string) error
}

type KVStorageImpl struct {
	mx sync.RWMutex
	m  map[string]interface{}
}

func NewKVStorage() *KVStorageImpl {
	return &KVStorageImpl{
		m: make(map[string]interface{}),
	}
}

func (store *KVStorageImpl) Get(ctx context.Context, key string) (interface{}, error) {
	store.mx.RLock()
	val := store.m[key]
	store.mx.RUnlock()
	return val, nil
}

func (store *KVStorageImpl) Put(ctx context.Context, key string, val interface{}) error {
	store.mx.Lock()
	store.m[key] = val
	store.mx.Unlock()
	return nil
}

func (store *KVStorageImpl) Delete(ctx context.Context, key string) error {
	store.mx.Lock()
	delete(store.m, key)
	store.mx.Unlock()
	return nil
}
