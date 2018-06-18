package memcache_service

import (
	"fmt"
	c "github.com/patrickmn/go-cache"
	"ms/sun/shared/golib/go_map"
	"time"
)

const _CACHE_LIKE = "LIKES_"

var likeCache = c.New(time.Hour*24, time.Minute*1)

func DidUserLikedPost(UserId, PostId int) bool {
	if UserId < 1 || PostId < 1 {
		return false
	}
	mp, ok := likeCache.Get(getLikeKey(UserId))
	var m *go_map.ConcurrentIntMap
	if ok {
		m, ok = mp.(*go_map.ConcurrentIntMap)
		if ok { //should never happens
			return false
		}
	} else {
		m = fillLikeCacheForUser(UserId)
	}

	val := m.GetVal(PostId)
	return val > 0
}

func AddToUserLikedPosts(UserId, PostId int) {
	if UserId < 1 || PostId < 1 {
		return
	}
	mp, ok := likeCache.Get(getLikeKey(UserId))
	if !ok {
		return
	}
	m, ok := mp.(*go_map.ConcurrentIntMap)
	if !ok {
		return
	}
	m.SetVal(PostId, 1)
}

func DeleteFromUserLikedPosts(UserId, PostId int) {
	if UserId < 1 || PostId < 1 {
		return
	}
	mp, ok := likeCache.Get(getLikeKey(UserId))
	if !ok {
		return
	}
	m, ok := mp.(*go_map.ConcurrentIntMap)
	if !ok {
		return
	}
	m.Delete(PostId)
}

//////// Internals ////////
func getLikeKey(UserId int) string {
	return fmt.Sprintf("_LIKES_%d", UserId)
}

func fillLikeCacheForUser(userId int) *go_map.ConcurrentIntMap {
	/*postIds, err := x.NewLike_Selector().Select_PostId().UserId_Eq(userId).GetIntSlice(base.DB)
	if err != nil {
		//just return something not to painc
		return go_map.NewConcurrentIntMap(0)
	}
	m := go_map.NewConcurrentIntMap(len(postIds))
	m.SetKeys(postIds, 1)
	likeCache.Set(getLikeKey(userId), m, 0)
	return m*/
	return nil
}
