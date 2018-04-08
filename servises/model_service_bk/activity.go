package model_service_bk

import (
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/servises/event_service"
	"ms/sun/shared/x"
)

func Action_OnPostCommentAdd(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	hash := hashPostCommented(post.PostId, comment.CommentId)
	not := x.Action{
		ActionId:       helper.NextRowsSeqId(),
		ActorUserId:    comment.UserId,
		ActionTypeEnum: int(x.ActionEnum_ACTION_POST_COMMENTED),
		PostId:         post.PostId,
		PeerUserId:     post.UserId,
		CommentId:      comment.CommentId,
		Murmur64Hash:   hash,
		CreatedTime:    helper.TimeNow(),
	}
	not.Save(base.DB)

	event := x.Event{
		ByUserId:     comment.UserId,
		PeerUserId:   post.UserId,
		PostId:       post.PostId,
		CommentId:    comment.CommentId,
		ActionId:     not.ActionId,
		Murmur64Hash: hash,
	}
	event_service.SaveEvent(event_service.COMMENTED_POST_EVENT, event)
}

func Action_OnPostCommentDeleted(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	hash := hashPostCommented(comment.PostId, comment.CommentId)
	x.NewAction_Deleter().
		ActorUserId_Eq(comment.UserId).
		Murmur64Hash_Eq(hash).
		Delete(base.DB)

	event := x.Event{
		ByUserId:   comment.UserId,
		PeerUserId: post.UserId,
		PostId:     post.PostId,
		CommentId:  comment.CommentId,
		//ActionId:     not.ActionId,
		Murmur64Hash: hash,
	}
	event_service.SaveEvent(event_service.UNCOMMENTED_POST_EVENT, event)

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

	event := x.Event{
		ByUserId:     UserId,
		PeerUserId:   FollowedPeerUserId,
		PostId:       0,
		CommentId:    0,
		ActionId:     not.ActionId,
		Murmur64Hash: hash,
	}
	event_service.SaveEvent(event_service.FOLLOWED_USER_EVENT, event)
}

func Action_OnUnFollowed(UserId, FollowedPeerUserId, FLId int) {
	hash := hashFollowed(UserId, FollowedPeerUserId)

	x.NewAction_Deleter().
		ActorUserId_Eq(UserId).
		Murmur64Hash_Eq(hash).
		Delete(base.DB)

	event := x.Event{
		ByUserId:   UserId,
		PeerUserId: FollowedPeerUserId,
		PostId:     0,
		CommentId:  0,
		//ActionId:     not.ActionId,
		Murmur64Hash: hash,
	}
	event_service.SaveEvent(event_service.UNFOLLOWED_USER_EVENT, event)
}

////////////// For Like ///////////////
func Action_OnPostLiked(lk *x.Like) {
	if lk == nil {
		return
	}
	hash := hashPostLiked(lk.UserId, lk.PostId)
	act := x.Action{
		ActionId:       helper.NextRowsSeqId(),
		ActorUserId:    lk.UserId,
		ActionTypeEnum: int(x.ActionEnum_ACTION_POST_LIKED),
		PostId:         lk.PostId,
		PeerUserId:     0,
		CommentId:      0,
		Murmur64Hash:   hash,
		CreatedTime:    helper.TimeNow(),
	}

	act.Save(base.DB)

	event := x.Event{
		ByUserId:     lk.UserId,
		PeerUserId:   0,
		PostId:       lk.PostId,
		CommentId:    0,
		ActionId:     act.ActionId,
		Murmur64Hash: hash,
	}
	event_service.SaveEvent(event_service.LIKED_POST_EVENT, event)
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

	event := x.Event{
		ByUserId:   lk.UserId,
		PeerUserId: 0,
		PostId:     lk.PostId,
		CommentId:  0,
		//ActionId:     act.ActionId,
		Murmur64Hash: hash,
	}
	event_service.SaveEvent(event_service.UNLIKED_POST_EVENT, event)
}
