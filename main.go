package main

import (
	"fmt"
	"log"
	"sync"
)

type Storer[k comparable, v any] interface {
	Put(k, v) error
	Get(k) (v, error)
	Update(k, v) error
	Delete(k) (v, error)
}

type KVStore[k comparable, v any] struct {
	mu   sync.RWMutex
	data map[k]v
}

func (s *KVStore[k, v]) Put(key k, value v) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return nil
}

func (s *KVStore[k, v]) Get(key k) (v, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]

	if !ok {
		return value, fmt.Errorf("the key(%v) does not exist", key)
	}

	return value, nil
}

func (s *KVStore[k, v]) Exists(key k) bool {
	_, ok := s.data[key]

	return ok
}

func (s *KVStore[k, v]) Update(key k, value v) error {
	s.mu.RLock() // Read lock
	defer s.mu.RUnlock()

	if !s.Exists(key) {
		return fmt.Errorf("the key(%v) does not exist", key)
	}

	s.data[key] = value

	return nil
}

func (s *KVStore[k, v]) Delete(key k) (v, error) {
	s.mu.Lock() // Write lock
	defer s.mu.Unlock()

	value, ok := s.data[key]

	if !ok {
		return value, fmt.Errorf("the key(%v) does not exist", key)
	}

	delete(s.data, key)

	return value, nil
}

func NewKVStore[k comparable, v any]() *KVStore[k, v] {
	return &KVStore[k, v]{
		data: make(map[k]v),
	}
}

type Server struct {
	Store Storer[string, string]
}

func (s *Server) getUserByName(name string) (string, error) {
	return s.Store.Get(name)
}

func main() {
	store := NewKVStore[string, string]()

	if err := store.Put("foo", "barr"); err != nil {
		log.Fatal(err)
	}

	value, err := store.Get("foo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)

	if err := store.Update("foo", "barr"); err != nil {
		log.Fatal(err)
	}

	value, err = store.Get("foo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)

	// putData(store)
}
