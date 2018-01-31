package model_service

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/shared/x"
	"ms/sun2/servises/memcache_service"
)

/*
spec: 1: following
      2: requested
      0: not following or could not
*/
func Follow(UserId, FollowedPeerUserId int) int {
	if UserId == FollowedPeerUserId || UserId < 1 || FollowedPeerUserId < 1 {
		return 0
	}

	flm := x.FollowingListMember{
		ListId:         UserId,
		UserId:         UserId,
		FollowedUserId: FollowedPeerUserId,
		CreatedTime:  helper.TimeNowMs(),
	}

	err := flm.Insert(base.DB)

	if err == nil {
		memcache_service.AddToUserFollwings(UserId, FollowedPeerUserId)
		Counter.UpdateUserFollowingCounts(UserId, 1)
		Counter.UpdateUserFollowersCounts(FollowedPeerUserId, 1)

		Notify_OnFollowed(UserId, FollowedPeerUserId, flm.Id)
		Action_OnFollowed(UserId, FollowedPeerUserId, flm.Id)
		return 1
	}
	return 0
}

func UnFollow(UserId, FollowedPeerUserId int) {
	if UserId == FollowedPeerUserId || UserId < 1 || FollowedPeerUserId < 1 {
		return
	}

	memcache_service.DeleteFromUserFollwings(UserId, FollowedPeerUserId)
	flm, err := x.NewFollowingListMember_Selector().UserId_Eq(UserId).FollowedUserId_Eq(FollowedPeerUserId).GetRow(base.DB)
	if err != nil {
		return
	}

	err = flm.Delete(base.DB)

	if err == nil {
		Counter.UpdateUserFollowingCounts(UserId, -1)
		Counter.UpdateUserFollowersCounts(FollowedPeerUserId, -1)

		Notify_OnUnFollowed(UserId, FollowedPeerUserId)
		Action_OnUnFollowed(UserId, FollowedPeerUserId, flm.Id)
	}
	//n, err := NewFollowingListMember_Deleter().UserId_Eq(UserId).FollowedUserId_Eq(FollowedPeerUserId).Delete(base.DB)
}
