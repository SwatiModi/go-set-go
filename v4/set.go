package v4

import (
	"gosetgo/ds"
	"sync"
)

var (
	shardLocker = sync.Map{}
	numShards   = 3
)

type set struct {
	shards []map[interface{}]bool
}

func NewSet() ds.Set {
	return &set{
		shards: make([]map[interface{}]bool, numShards),
	}
}

func (s *set) Add(item interface{}) bool {

	lockKey := item.(int) % numShards

	// get or set shard lock
	if _, ok := shardLocker.Load(lockKey); !ok {
		shardLocker.Store(lockKey, &sync.RWMutex{})
		s.shards[lockKey] = make(map[interface{}]bool)
	}
	v, _ := shardLocker.Load(lockKey)
	shardLock := v.(*sync.RWMutex)

	shardLock.RLock()
	if _, ok := s.shards[lockKey][item]; ok {
		shardLock.RUnlock()
		return false
	}
	shardLock.RUnlock()

	shardLock.Lock()
	defer shardLock.Unlock()
	s.shards[lockKey][item] = true
	return true
}

func (s *set) Remove(item interface{}) bool {
	lockKey := item.(int) % numShards

	// get or set shard lock
	if _, ok := shardLocker.Load(lockKey); !ok {
		shardLocker.Store(lockKey, &sync.RWMutex{})
		s.shards[lockKey] = make(map[interface{}]bool)
	}
	v, _ := shardLocker.Load(lockKey)
	shardLock := v.(*sync.RWMutex)

	shardLock.RLock()
	if _, ok := s.shards[lockKey][item]; !ok {
		shardLock.RUnlock()
		return false
	}
	shardLock.RUnlock()

	shardLock.Lock()
	defer shardLock.Unlock()
	delete(s.shards[lockKey], item)
	return true
}

func (s *set) Contains(item interface{}) bool {
	lockKey := item.(int) % numShards

	// get or set shard lock
	if _, ok := shardLocker.Load(lockKey); !ok {
		shardLocker.Store(lockKey, &sync.RWMutex{})
		s.shards[lockKey] = make(map[interface{}]bool)
	}
	v, _ := shardLocker.Load(lockKey)
	shardLock := v.(*sync.RWMutex)

	shardLock.RLock()
	defer shardLock.RUnlock()
	_, ok := s.shards[lockKey][item]
	return ok
}
