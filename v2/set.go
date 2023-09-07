package v2

import (
	"gosetgo/ds"
	"sync"
)

type set struct {
	items map[interface{}]bool
	mu    sync.Mutex
}

func NewSet() ds.Set {
	return &set{
		items: make(map[interface{}]bool),
		mu:    sync.Mutex{},
	}
}

func (s *set) Add(item interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[item]; ok {
		return false
	}
	s.items[item] = true
	return true
}

func (s *set) Remove(item interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[item]; !ok {
		return false
	}
	delete(s.items, item)
	return true
}

func (s *set) Contains(item interface{}) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.items[item]
	return ok
}
