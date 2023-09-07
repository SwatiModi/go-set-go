package v3

import (
	"gosetgo/ds"
	"sync"
)

type set struct {
	items map[interface{}]bool
	mu    sync.RWMutex
}

func NewSet() ds.Set {
	return &set{
		items: make(map[interface{}]bool),
		mu:    sync.RWMutex{},
	}
}

func (s *set) Add(item interface{}) bool {
	s.mu.RLock()
	if _, ok := s.items[item]; ok {
		s.mu.RUnlock()
		return false
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[item] = true
	return true
}

func (s *set) Remove(item interface{}) bool {
	s.mu.RLock()
	if _, ok := s.items[item]; !ok {
		s.mu.RUnlock()
		return false
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, item)
	return true
}

func (s *set) Contains(item interface{}) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.items[item]
	return ok
}
