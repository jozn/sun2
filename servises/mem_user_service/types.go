package mem_user_service

import (
	"ms/sun/base"
	"ms/sun2/shared/golib/go_map"
	"ms/sun2/shared/x"
)

/////////////////////////// MemUser ////////////////////////////
type MemUser struct {
	userId            int
	lastPosts         []*x.Post
	lastActions       []*x.Action
	followedUserIds   *go_map.ConcurrentIntMap
	followersUserIds  *go_map.ConcurrentIntMap
	isPostsLoaded     bool
	isActivityLoaded  bool
	isFollowersLoaded bool
	isFollowedLoaded  bool
}

func newMemUser() MemUser {
	return MemUser{
		followedUserIds:  go_map.NewConcurrentIntMap(10),
		followersUserIds: go_map.NewConcurrentIntMap(10),
	}
}

func (m *MemUser) GetLastPost() []*x.Post {
	if !m.isPostsLoaded {
		m.lastPosts, _ = x.NewPost_Selector().UserId_Eq(m.userId).OrderBy_PostId_Desc().Limit(50).GetRows(base.DB)
		m.isPostsLoaded = true
	}
	return m.lastPosts
}

func (m *MemUser) GetLastActions() []*x.Action {
	if !m.isActivityLoaded {
		m.lastActions, _ = x.NewAction_Selector().ActorUserId_Eq(m.userId).OrderBy_ActionId_Desc().Limit(100).GetRows(base.DB)
		m.isActivityLoaded = true
	}
	return m.lastActions
}

func (m *MemUser) GetFollowers() []int {
	if !m.isFollowersLoaded {
		ids, _ := x.NewFollowingListMember_Selector().
			Select_UserId().
			FollowedUserId_Eq(m.userId).
			GetIntSlice(base.DB)
		m.followersUserIds.SetKeys(ids, 1)
		m.isFollowersLoaded = true
	}
	return m.followersUserIds.GetAllKeys()
}

func (m *MemUser) GetFollowed() []int {
	if !m.isFollowedLoaded {
		ids, _ := x.NewFollowingListMember_Selector().
			Select_FollowedUserId().
			UserId_Eq(m.userId).
			GetIntSlice(base.DB)
		m.followedUserIds.SetKeys(ids, 1)
		m.isFollowedLoaded = true
	}
	return m.followedUserIds.GetAllKeys()
}
