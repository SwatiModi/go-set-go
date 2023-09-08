package v5

import (
	"sync"
	"time"
)

type Set struct {
	items map[interface{}]bool
	addCh chan interface{}
	delCh chan interface{}

	readCopy  map[interface{}]bool
	refreshMU sync.RWMutex

	close chan struct{}
	wg    sync.WaitGroup
}

func NewSet() *Set {
	set := &Set{
		items: make(map[interface{}]bool),
		addCh: make(chan interface{}, 100),
		delCh: make(chan interface{}, 100),
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
			s.refreshMU.Lock()
			s.readCopy = make(map[interface{}]bool)
			for k, v := range s.items {
				s.readCopy[k] = v
			}
			s.refreshMU.Unlock()
		case key := <-s.addCh:
			s.items[key] = true
		case key := <-s.delCh:
			delete(s.items, key)
		case <-s.close:
			return
		}
	}
}

func (s *Set) Add(key interface{}) bool {
	if s.Contains(key) {
		return false
	}
	s.addCh <- key
	return true
}

func (s *Set) Remove(key interface{}) bool {
	if !s.Contains(key) {
		return false
	}
	s.delCh <- key
	return true
}

func (s *Set) Contains(key interface{}) bool {
	s.refreshMU.RLock()
	_, found := s.readCopy[key]
	s.refreshMU.RUnlock()
	return found
}

func (s *Set) Close() {
	close(s.close)
	s.wg.Wait()
}
