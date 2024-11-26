package main

import (
	"fmt"
	"log"
	"net/http"
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

type User struct {
	ID       int
	FistName string
	Age      int
	Gender   string
}

type Server struct {
	Storage    Storer[int, *User]
	ListenAddr string
}

func NewServer(ListenAddr string) *Server {
	return &Server{
		Storage:    NewKVStore[int, *User](),
		ListenAddr: ListenAddr,
	}
}

func (s *Server) handlePut(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Foo"))
}

func (s *Server) Start() {
	fmt.Printf("HTTP server is running on post %s", s.ListenAddr)

	http.HandleFunc("/put", s.handlePut)

	log.Fatal(http.ListenAndServe(s.ListenAddr, nil))
}

func main() {
	s := NewServer(":3000")
	s.Start()
}
