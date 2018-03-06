package archive

import (
    "ms/sun2/shared/golib/go_map"
    "sync"
)

// collections os user
type masterCache struct {
    mp  map[int]*go_map.ConcurrentIntMap
    mux sync.RWMutex
}

func newMasterCache() *masterCache {
    return &masterCache{
        mp: make(map[int]*go_map.ConcurrentIntMap, 10000),
    }
}

func (s *masterCache) GetVal(key int) (*go_map.ConcurrentIntMap, bool) {
    s.mux.RLocker()
    val, ok := s.mp[key]
    s.mux.RUnlock()
    return val, ok
}

func (s *masterCache) SetVal(key int, val *go_map.ConcurrentIntMap) {
    s.mux.Lock()
    s.mp[key] = val
    s.mux.Unlock()
}

func (s *masterCache) Delete(key int) {
    s.mux.Lock()
    delete(s.mp, key)
    s.mux.Unlock()
}

