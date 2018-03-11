package model_service

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/servises/event_service"
	"ms/sun2/servises/memcache_service"
	"ms/sun2/shared/x"
)

func Like_LikePost(UserId, PostId int) {
	p, ok := x.Store.GetPostByPostId(PostId)
	if !ok {
		return
	}
	l := &x.Like{
		UserId:       UserId,
		PostTypeEnum: p.PostTypeEnum,
		PostId:       PostId,
		LikeEnum:     0, //emotions
		CreatedTime:  helper.TimeNow(),
	}

	err := l.Insert(base.DB)
	if err == nil {
		x.NewPost_Updater().LikesCount_Increment(1).PostId_Eq(PostId).Update(base.DB)
		memcache_service.AddToUserLikedPosts(UserId, PostId)

		hash := hashPostLiked(UserId, PostId)
		event := x.Event{
			ByUserId:     UserId,
			PeerUserId:   p.UserId,
			ActionId:     helper.NextRowsSeqId(),
			Murmur64Hash: hash,
		}
		event_service.SaveEvent(event_service.LIKED_POST_EVENT, event)
		//Notify_OnPostLiked(l)
		//Action_OnPostLiked(l)
	}
}

func Like_UnlikePost(UserId, PostId int) {
	l, err := x.NewLike_Selector().UserId_Eq(UserId).PostId_Eq(PostId).GetRow(base.DB)
	if err != nil {
		return
	}

	err = l.Delete(base.DB)
	if err == nil {
		memcache_service.DeleteFromUserLikedPosts(UserId, PostId)
		x.NewPost_Updater().LikesCount_Increment(-1).PostId_Eq(PostId).Update(base.DB)

		hash := hashPostLiked(UserId, PostId)
		event := x.Event{
			ByUserId:     UserId,
			PeerUserId:   0,
			ActionId:     helper.NextRowsSeqId(),
			Murmur64Hash: hash,
		}
		event_service.SaveEvent(event_service.UNLIKED_POST_EVENT, event)
		//Notify_OnPostUnLiked(l)
		//Action_OnPostUnLiked(l)
	}
}
