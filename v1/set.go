package v1

import "gosetgo/ds"

type set struct {
	items map[interface{}]bool
}

func NewSet() ds.Set {
	return &set{
		items: make(map[interface{}]bool),
	}
}

func (s *set) Add(item interface{}) bool {
	if s.Contains(item) {
		return false
	}
	s.items[item] = true
	return true
}

func (s *set) Remove(item interface{}) bool {
	if !s.Contains(item) {
		return false
	}
	delete(s.items, item)
	return true
}

func (s *set) Contains(item interface{}) bool {
	_, ok := s.items[item]
	return ok
}
