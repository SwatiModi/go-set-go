package v25

import (
	"gosetgo/ds"
	"sync"
)

type set struct {
	items sync.Map
}

func NewSet() ds.Set {
	return &set{
		items: sync.Map{},
	}
}

func (s *set) Add(item interface{}) bool {
	if _, ok := s.items.Load(item); ok {
		return false
	}
	s.items.Store(item, true)
	return true
}

func (s *set) Remove(item interface{}) bool {
	if _, ok := s.items.Load(item); !ok {
		return false
	}
	s.items.Delete(item)
	return true
}

func (s *set) Contains(item interface{}) bool {
	_, ok := s.items.Load(item)
	return ok
}
