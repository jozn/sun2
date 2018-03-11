package mem_user_service

import (
	"ms/sun2/shared/golib/go_map"
	"sync"
)

type allMp struct {
	mp map[int]*MemUser
	sync.RWMutex
}

func (s *allMp) GetForUser(userId int) *MemUser {
	s.RLock()
	mu := s.mp[userId]
	s.RUnlock()
	if mu == nil {
		mu = &MemUser{
			userId:           userId,
			followedUserIds:  go_map.NewConcurrentIntMap(100),
			followersUserIds: go_map.NewConcurrentIntMap(100),
		}
		s.Lock()
		s.mp[userId] = mu
		s.Unlock()
	}
	return mu
}

var allMemUserMap = &allMp{
	mp: make(map[int]*MemUser, 1000),
}

func GetForUser(userId int) (mu *MemUser, ok bool) {
	mu = allMemUserMap.GetForUser(userId)
	if mu != nil {
		return mu, true
	}
	return nil, false
}
