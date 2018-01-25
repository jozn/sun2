package model_service

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/shared/x"
)

func Action_OnPostCommentAdd(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	not := x.Action{
		ActionId:       helper.NextRowsSeqId(),
		ActorUserId:    comment.UserId,
		ActionTypeEnum: int(x.ActionEnum_ACTION_POST_COMMENTED),
		PostId:         post.PostId,
		PeerUserId:     post.UserId,
		CommentId:      comment.CommentId,
		Murmur64Hash:   hashPostCommented(post.PostId, comment.CommentId),
		CreatedTime:    helper.TimeNow(),
	}
	not.Save(base.DB)
}

func Action_OnPostCommentDeleted(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	refId := hashPostCommented(comment.PostId, comment.CommentId)
	x.NewAction_Deleter().
		ActorUserId_Eq(comment.UserId).
		Murmur64Hash_Eq(refId).
		Delete(base.DB)

}

////////// For Follows ///////////
func Action_OnFollowed(UserId, FollowedPeerUserId, FLId int) {

	hash := hashFollowed(UserId, FollowedPeerUserId)
	not := x.Action{
		ActionId:       helper.NextRowsSeqId(),
		ActorUserId:    UserId,
		ActionTypeEnum: int(x.ActionEnum_ACTION_FOLLOWED_USER),
		PostId:         0,
		PeerUserId:     FollowedPeerUserId,
		CommentId:      0,
		Murmur64Hash:   hash,
		CreatedTime:    helper.TimeNow(),
	}

	not.Save(base.DB)
}

func Action_OnUnFollowed(UserId, FollowedPeerUserId, FLId int) {
	hash := hashFollowed(UserId, FollowedPeerUserId)

	x.NewAction_Deleter().
		ActorUserId_Eq(UserId).
		Murmur64Hash_Eq(hash).
		Delete(base.DB)
}

////////////// For Like ///////////////
func Action_OnPostLiked(lk *x.Like) {
	if lk == nil {
		return
	}
	hash := hashPostLiked(lk.UserId, lk.PostId)
	not := x.Action{
		ActionId:       helper.NextRowsSeqId(),
		ActorUserId:    lk.UserId,
		ActionTypeEnum: int(x.ActionEnum_ACTION_POST_LIKED),
		PostId:         lk.PostId,
		PeerUserId:     0,
		CommentId:      0,
		Murmur64Hash:   hash,
		CreatedTime:    helper.TimeNow(),
	}

	not.Save(base.DB)
}

func Action_OnPostUnLiked(lk *x.Like) {
	if lk == nil {
		return
	}
	hash := hashPostLiked(lk.UserId, lk.PostId)

	x.NewAction_Deleter().
		ActorUserId_Eq(lk.UserId).
		Murmur64Hash_Eq(hash).
		Delete(base.DB)
}
