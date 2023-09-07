package v4

import (
	"sync"
	"time"
)

type Set struct {
	items map[int]bool
	addCh chan int
	delCh chan int

	readCopy map[int]bool

	close chan struct{}
	wg    sync.WaitGroup
}

func NewSet() *Set {
	set := &Set{
		items: make(map[int]bool),
		addCh: make(chan int, 100),
		delCh: make(chan int, 100),
		close: make(chan struct{}), // Channel to signal closing
	}
	set.wg.Add(1)
	go set.process(1000) // 1 second refresh interval
	return set
}

func (s *Set) process(refreshIntervalMs int) {
	ticker := time.NewTicker(time.Duration(refreshIntervalMs) * time.Millisecond)

	defer s.wg.Done()
	for {
		select {
		case <-ticker.C:
			s.readCopy = make(map[int]bool)
			for k, v := range s.items {
				s.readCopy[k] = v
			}
		case key := <-s.addCh:
			s.items[key] = true
		case key := <-s.delCh:
			delete(s.items, key)
		case <-s.close:
			return
		}
	}
}

func (s *Set) Add(key int) bool {
	if s.Contains(key) {
		return false
	}
	s.addCh <- key
	return true
}

func (s *Set) Delete(key int) bool {
	if !s.Contains(key) {
		return false
	}
	s.delCh <- key
	return true
}

func (s *Set) Contains(key int) bool {
	_, found := s.readCopy[key]
	return found
}

func (s *Set) Close() {
	close(s.close)
	s.wg.Wait()
}
