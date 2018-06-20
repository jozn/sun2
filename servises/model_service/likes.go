package model_service

import (
	"ms/sun/servises/event_service"
	"ms/sun/servises/memcache_service"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"ms/sun_old/base"
)

func Like_LikePost(UserId, PostId int) {
	p, ok := x.Store.GetPostByPostId(PostId)
	if !ok {
		return
	}
	l := &x.Likes{
		Id:          helper.NanoRowIdSeq(),
		UserId:      UserId,
		PostType:    p.PostType,
		PostId:      PostId,
		CreatedTime: helper.TimeNow(),
	}

	err := l.Insert(base.DB)
	if err == nil {
		//x.NewPost_Updater().LikesCount_Increment(1).PostId_Eq(PostId).Update(base.DB)
		x.NewPostCount_Updater().LikesCount_Increment(1).PostId_Eq(PostId).Update(base.DB)
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
	l, err := x.NewLikes_Selector().UserId_Eq(UserId).PostId_Eq(PostId).GetRow(base.DB)
	if err != nil {
		return
	}

	err = l.Delete(base.DB)
	if err == nil {
		memcache_service.DeleteFromUserLikedPosts(UserId, PostId)
		//x.NewPost_Updater().LikesCount_Increment(-1).PostId_Eq(PostId).Update(base.DB)
		x.NewPostCount_Updater().LikesCount_Increment(-1).PostId_Eq(PostId).Update(base.DB)

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
