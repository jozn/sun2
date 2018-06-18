package memcache_service

import (
	"fmt"
	c "github.com/patrickmn/go-cache"
	"ms/sun/shared/golib/go_map"
	"ms/sun/shared/x"
	"time"
)

var followCache = c.New(time.Hour*24, time.Minute*1)

func DoesUserFollows(MeId, PeerId int) x.FollowingEnum {
    //fmt.Println("mmm", MeId,PeerId)
	if MeId < 1 || PeerId < 1 {
		return x.FollowingEnum(0)
	}
	mp, ok := followCache.Get(getFollowKey(MeId))
	var m *go_map.ConcurrentIntMap
	if ok {
		m, ok = mp.(*go_map.ConcurrentIntMap)
		if !ok { //should never happens
			return x.FollowingEnum(0)
		}
	} else {
		m = fillfollowCacheForUser(MeId)
	}

	val := m.GetVal(PeerId)
	return x.FollowingEnum(val)
}

func GetUserFollowsSlice(UserId int) []int {
	if UserId < 1 {
		return []int{}
	}
	mp, ok := followCache.Get(getFollowKey(UserId))
	var m *go_map.ConcurrentIntMap
	if ok {
		m, ok = mp.(*go_map.ConcurrentIntMap)
		if ok { //should never happens
			return []int{}
		}
	} else {
		m = fillfollowCacheForUser(UserId)
	}

	return m.GetAllKeys()
}

func AddToUserFollwings(UserId, PeerId int) {
	if UserId < 1 || PeerId < 1 {
		return
	}
	mp, ok := followCache.Get(getFollowKey(UserId))
	if !ok {
		return
	}
	m, ok := mp.(*go_map.ConcurrentIntMap)
	if !ok {
		return
	}
	m.SetVal(PeerId, int(x.FollowingEnum_FOLLOWING)) //todo fix me add enums
}

func DeleteFromUserFollwings(UserId, PeerId int) {
	if UserId < 1 || PeerId < 1 {
		return
	}
	mp, ok := followCache.Get(getFollowKey(UserId))
	if !ok {
		return
	}
	m, ok := mp.(*go_map.ConcurrentIntMap)
	if !ok {
		return
	}
	m.Delete(PeerId)
}

//////// Internals ////////
func getFollowKey(UserId int) string {
	return fmt.Sprintf("_Follwings_%d", UserId)
}

func fillfollowCacheForUser(userId int) *go_map.ConcurrentIntMap {
	/*postIds, err := x.NewFollowingListMember_Selector().Select_FollowedUserId().UserId_Eq(userId).GetIntSlice(base.DB)
	if err != nil {
		//just return something not to painc
		return go_map.NewConcurrentIntMap(0)
	}
	m := go_map.NewConcurrentIntMap(len(postIds))
	m.SetKeys(postIds, 1)
	followCache.Set(getFollowKey(userId), m, 0)
	return m*/
	return nil
}
