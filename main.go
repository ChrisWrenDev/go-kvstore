package main

import "sync"

type Storer[k comparable, v any] interface {
	Put(k, v) error
	Get(k)(v, error)
    Update(k, v) error
    Delete(k)(v, error)
}

type KVStore[k comparable, v any] struct {
    mu sync.RWMutex
    data map[k]v
}

func (s *KVStore[k, v]) Put(k, v) error {}

func (s *KVStore[k, v]) Get(k) (v, error) {}

func (s *KVStore[k, v]) Update(k, v) error {}

func (s *KVStore[k, v]) Delete(k) (v, error) {}

func NewKVStore[k comparable, v any]() *KVStore[k, v] {
    return &KVStore[k, v]{
        data: make(map[k]v),

    }
}

func putData(s Storer[string, int]) error {
    return s.Put("Name", 1)
}

type Block struct {}
type Transaction struct {}

func main(){
    _ = NewKVStore[string, *Block]()
    _ = NewKVStore[string, *Transaction]()

    // putData(kv)
}
