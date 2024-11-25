package main

import "sync"

type Storer interface {
	Put(string, []byte) error
	Get(string )([]byte, error)
    Update(string, []byte) error
    Delete(string)([]byte, error)
}

type KVStore struct {
    mu sync.RWMutex
    data map[string][]byte
}

func NewKVStore() *KVStore {
    return &KVStore{
        data: make(map[string][]byte),

    }
}

func putData(s Storer) error {
    return s.Put("Name", []byte("tango terry"))
}

func main(){
    kv := NewKVStore()

    putData(kv)
}
