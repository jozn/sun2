package go_map

import (
	"sync"
)

func NewConcurrentIntMap(size int) *ConcurrentIntMap {
	return &ConcurrentIntMap{
		mp: make(map[int]int, size),
	}
}

type ConcurrentIntMap struct {
	mp  map[int]int
	mux sync.RWMutex
}

func (s *ConcurrentIntMap) GetVal(key int) int {
	s.mux.RLock()
	val := s.mp[key]
	s.mux.RUnlock()
	return val
}

func (s *ConcurrentIntMap) GetAllKeys() []int {
	arr := make([]int, len(s.mp))
	s.mux.RLock()
	for k, _ := range s.mp {
		arr = append(arr, k)
	}
	s.mux.RUnlock()
	return arr
}

func (s *ConcurrentIntMap) SetVal(key, val int) {
	s.mux.Lock()
	s.mp[key] = val
	s.mux.Unlock()
}

func (s *ConcurrentIntMap) Delete(key int) {
	s.mux.Lock()
	delete(s.mp, key)
	s.mux.Unlock()
}

//this is used in some where like all the posts a users liked: keys is posts , value is just a random number
func (s *ConcurrentIntMap) SetKeys(keys []int, val int) {
	s.mux.Lock()
	for _, key := range keys {
		s.mp[key] = val
	}
	s.mux.Unlock()
}
