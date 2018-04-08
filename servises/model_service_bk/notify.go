package model_service_bk

import (
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
    "ms/sun/servises/view_service"
)

//////// For Comment //////////
func Notify_OnPostCommented(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	if comment.UserId == post.UserId {
		return
	}

	key := hashPostCommented(post.PostId, comment.CommentId)

	not := x.Notify{
		NotifyId:       helper.NextRowsSeqId(),
		ForUserId:      post.UserId,
		ActorUserId:    comment.UserId,
		NotifyTypeEnum: int(x.NotifyEnum_NOTIFY_POST_COMMENTED),
		PostId:         post.PostId,
		CommentId:      comment.CommentId,
		PeerUserId:     0,
		Murmur64Hash:   key,
		SeenStatus:     0,
		CreatedTime:    helper.TimeNow(),
	}
	not.Save(base.DB)

	Notify_PushToUserPipe(not)
}

func Notify_OnPostCommentedDeleted(comment *x.Comment, post *x.Post) {
	if comment == nil || post == nil {
		return
	}

	hash := hashPostCommented(post.PostId, comment.CommentId)

	nr := x.NotifyRemoved{
		Murmur64Hash: hash,
		ForUserId:    comment.UserId,
	}

	nr.Save(base.DB)

	Notify_PushToUserPipeRemoved(hash)
}

////////// For Follows ///////////
func Notify_OnFollowed(UserId, FollowedPeerUserId, FLId int) {
	if UserId == FollowedPeerUserId { //must never reach here at all
		return
	}

	hash := hashFollowed(UserId, FollowedPeerUserId)
	nf := x.Notify{
		NotifyId:       helper.NextRowsSeqId(),
		ForUserId:      FollowedPeerUserId,
		ActorUserId:    UserId,
		NotifyTypeEnum: int(x.NotifyEnum_NOTIFY_FOLLOWED_YOU),
		PostId:         0,
		CommentId:      0,
		PeerUserId:     0,
		Murmur64Hash:   hash,
		SeenStatus:     0,
		CreatedTime:    helper.TimeNow(),
	}

	nf.Save(base.DB)

	Notify_PushToUserPipe(nf)
}

func Notify_OnUnFollowed(UserId, FollowedPeerUserId int) {

	hash := hashFollowed(UserId, FollowedPeerUserId)

	x.NewNotify_Deleter().
		ActorUserId_Eq(UserId).
		Murmur64Hash_Eq(hash).
		Delete(base.DB)

	nr := &x.NotifyRemoved{
		Murmur64Hash: hash,
		ForUserId:    UserId,
	}

	nr.Save(base.DB)

	Notify_PushToUserPipeRemoved(hash)

}

////////////// For Likes ///////////////
func Notify_OnPostLiked(lk *x.Like) {
	post, ok := x.Store.GetPostByPostId(lk.PostId)
	if !ok {
		return
	}

	if lk.UserId == post.UserId {
		return
	}

	hash := hashPostLiked(lk.UserId, lk.PostId)

	nf := x.Notify{
		NotifyId:       helper.NextRowsSeqId(),
		ForUserId:      post.UserId,
		ActorUserId:    lk.UserId,
		NotifyTypeEnum: int(x.NotifyEnum_NOTIFY_POST_LIKED),
		PostId:         post.PostId,
		CommentId:      0,
		PeerUserId:     0,
		Murmur64Hash:   hash,
		SeenStatus:     0,
		CreatedTime:    helper.TimeNow(),
	}

	nf.Save(base.DB)
}

func Notify_OnPostUnLiked(lk *x.Like) {
	post, ok := x.Store.GetPostByPostId(lk.PostId)
	if !ok {
		return
	}

	hash := hashPostLiked(lk.UserId, lk.PostId)

	x.NewNotify_Deleter().
		ActorUserId_Eq(lk.UserId).
		Murmur64Hash_Eq(hash).
		Delete(base.DB)

	nr := x.NotifyRemoved{
		Murmur64Hash: hash,
		ForUserId:    post.UserId,
	}

	nr.Save(base.DB)
}

//fix: must be NotifyView
//TODO add x.PB_PushNotify to proto
func Notify_PushToUserPipe(nf x.Notify) {
	nv := view_service.Notify_notifyToView(&nf, nf.ForUserId)
	_ = nv
	//call := base.NewCallWithData("NotifyAddOne", nv)
	//call := NewPB_CommandToClient_WithData("NotifyAddOne", nil)
	//AllPipesMap.SendToUser(nf.ForUserId, call)
}

func Notify_PushToUserPipeRemoved(id int) {

}

func Notify_ListOfRemovedAndEmptyIt(UserId int) []int {
	res, _ := x.NewNotifyRemoved_Selector().Select_Murmur64Hash().ForUserId_Eq(UserId).GetIntSlice(base.DB)
	if res != nil && len(res) > 0 {
		x.NewNotifyRemoved_Deleter().ForUserId_Eq(UserId).Delete(base.DB)
	}
	return res
}
