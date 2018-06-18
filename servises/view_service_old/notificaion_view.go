package view_service

import (
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
)

func Notify_GetLastsViews(UserId, last int) (res []*x.PB_NotifyView) {
	selector := x.NewNotify_Selector().ForUserId_Eq(UserId).Limit(100).OrderBy_NotifyId_Desc()
	if last > 0 {
		selector.NotifyId_LE(last)
	}

	nots, err := selector.GetRows(base.DB)

	if err != nil {
		helper.DebugPrintln(err)
		return
	}

	res = make([]*x.PB_NotifyView, 0, len(nots))

	notify_fillCaches(nots)

	for _, nf := range nots {
		res = append(res, Notify_notifyToView(nf, UserId))
	}

	return res
}

func Notify_notifyToView(nf *x.Notify, MeId int) *x.PB_NotifyView {
	nv := &x.PB_NotifyView{
		NotifyId:      int64(nf.NotifyId),
		ForUserId:     int32(nf.ForUserId),
		ActorUserId:   int32(nf.ActorUserId),
		NotiyTypeEnum: int32(nf.NotifyTypeEnum),
		PostId:        int64(nf.PostId),
		CommentId:     int64(nf.CommentId),
		PeerUserId:    int32(nf.PeerUserId),
		Murmur64Hash:  int64(nf.Murmur64Hash),
		SeenStatus:    int32(nf.SeenStatus),
		CreatedTime:   int32(nf.CreatedTime),
	}

	nv.ActorUserView = UserViewAndMe(nf.ActorUserId, MeId)

	switch x.NotifyEnum(nf.NotifyTypeEnum) {

	case x.NotifyEnum_NOTIFY_FOLLOWED_YOU:

	case x.NotifyEnum_NOTIFY_POST_LIKED:
		nv.PostView, _ = PostSingleViewForPostId(nf.PostId, MeId)

	case x.NotifyEnum_NOTIFY_POST_COMMENTED:
		nv.PostView, _ = PostSingleViewForPostId(nf.PostId, MeId)
		nv.CommentView = SingleCommentView(nf.CommentId)
	}

	return nv
}

//copy of Activity_fillCaches with modificaion - make in sync
func notify_fillCaches(nots []*x.Notify) {
	var pre_posts []int
	var pre_comments []int

	for _, nf := range nots {
		switch x.NotifyEnum(nf.NotifyTypeEnum) {

		case x.NotifyEnum_NOTIFY_FOLLOWED_YOU:

		case x.NotifyEnum_NOTIFY_POST_LIKED:
			pre_posts = append(pre_posts, nf.PostId)

		case x.NotifyEnum_NOTIFY_POST_COMMENTED:
			pre_comments = append(pre_comments, nf.CommentId)
			pre_posts = append(pre_posts, nf.PostId)
		}
	}

	x.Store.PreLoadCommentByCommentIds(pre_comments)
	x.Store.PreLoadPostByPostIds(pre_posts)
}
