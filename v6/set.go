package v6

import (
	"sync/atomic"
)

type Set struct {
	items    []int32
	capacity int
}

func NewSet() *Set {
	return &Set{
		items:    make([]int32, 5),
		capacity: 5,
	}
}

func (s *Set) Add(item int) bool {
	if item < 0 || item >= s.capacity {
		return false
	}

	atomic.StoreInt32(&s.items[item], 1)
	return true
}

func (s *Set) Remove(item int) bool {
	if item < 0 || item >= s.capacity {
		return false
	}

	atomic.StoreInt32(&s.items[item], 0)
	return true
}

func (s *Set) Contains(item int) bool {
	if item < 0 || item >= s.capacity {
		return false
	}

	return atomic.LoadInt32(&s.items[item]) != 0
}
