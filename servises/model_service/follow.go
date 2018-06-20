package model_service

import (
	"ms/sun/servises/event_service"
	"ms/sun/servises/memcache_service"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"ms/sun_old/base"
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

	//todo add if blocked

	flm := x.Followed{
		Id:             helper.NanoRowIdSeq(),
		UserId:         UserId,
		FollowedUserId: FollowedPeerUserId,
		CreatedTime:    helper.TimeNow(),
	}

	err := flm.Insert(base.DB)

	if err == nil {
		memcache_service.AddToUserFollwings(UserId, FollowedPeerUserId)
		Counter.UpdateUserFollowingCounts(UserId, 1)
		Counter.UpdateUserFollowersCounts(FollowedPeerUserId, 1)

		hash := hashFollowed(UserId, FollowedPeerUserId)
		event := x.Event{
			ByUserId:     UserId,
			PeerUserId:   FollowedPeerUserId,
			ActionId:     helper.NextRowsSeqId(),
			Murmur64Hash: hash,
		}
		event_service.SaveEvent(event_service.FOLLOWED_USER_EVENT, event)

		//Notify_OnFollowed(ProfileUserId, FollowedPeerUserId, flm.Id)
		//Action_OnFollowed(ProfileUserId, FollowedPeerUserId, flm.Id)
		return 1
	}
	return 0
}

func UnFollow(UserId, FollowedPeerUserId int) {
	if UserId == FollowedPeerUserId || UserId < 1 || FollowedPeerUserId < 1 {
		return
	}

	memcache_service.DeleteFromUserFollwings(UserId, FollowedPeerUserId)
	flm, err := x.NewFollowed_Selector().UserId_Eq(UserId).FollowedUserId_Eq(FollowedPeerUserId).GetRow(base.DB)
	if err != nil {
		return
	}

	err = flm.Delete(base.DB)

	if err == nil {
		Counter.UpdateUserFollowingCounts(UserId, -1)
		Counter.UpdateUserFollowersCounts(FollowedPeerUserId, -1)

		hash := hashFollowed(UserId, FollowedPeerUserId)
		event := x.Event{
			ByUserId:     UserId,
			PeerUserId:   FollowedPeerUserId,
			ActionId:     helper.NextRowsSeqId(),
			Murmur64Hash: hash,
		}
		event_service.SaveEvent(event_service.UNFOLLOWED_USER_EVENT, event)

		//Notify_OnUnFollowed(ProfileUserId, FollowedPeerUserId)
		//Action_OnUnFollowed(ProfileUserId, FollowedPeerUserId, flm.Id)
	}
	//n, err := NewFollowingListMember_Deleter().UserId_Eq(ProfileUserId).FollowedUserId_Eq(FollowedPeerUserId).Delete(base.DB)
}
