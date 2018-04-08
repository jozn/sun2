package model_service_bk

import (
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/servises/memcache_service"
	"ms/sun/shared/x"
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
		Notify_OnPostLiked(l)
		Action_OnPostLiked(l)
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
		Notify_OnPostUnLiked(l)
		Action_OnPostUnLiked(l)
	}
}
