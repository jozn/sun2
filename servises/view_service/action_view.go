package view_service

import (
	"ms/sun/base"
	"ms/sun2/servises/memcache_service"
	"ms/sun2/shared/x"
)

//page >= 1
func Action_GetLastsViews(MeId, Page, Limit, Last int) []x.PB_ActionView {
	uids := memcache_service.GetUserFollowsSlice(MeId)

	selector := x.NewAction_Selector().ActorUserId_In(uids).OrderBy_ActionId_Desc().Limit(Limit)
	if Last > 0 {
		selector.ActionId_LT(Last)
	} else if Page >= 1 {
		selector.Offset((Page - 1) * Limit)
	}

	nots, err := selector.GetRows(base.DB)

	res := make([]x.PB_ActionView, 0, len(nots))
	if err != nil {
		return res
	}

	//fill caches
	action_fillCaches(nots)

	for _, act := range nots {
		av := x.PB_ActionView{
			ActionId:       int64(act.ActionId),
			ActorUserId:    int32(act.ActorUserId),
			ActionTypeEnum: int32(act.ActionTypeEnum),
			PeerUserId:     int32(act.PeerUserId),
			PostId:         int64(act.PostId),
			CommentId:      int64(act.CommentId),
			Murmur64Hash:   int64(act.Murmur64Hash),
			CreatedTime:    int32(act.CreatedTime),
		}
		av.ActorUser = UserViewAndMe(act.ActorUserId, MeId)

		switch x.ActionEnum(act.ActionTypeEnum) {
		case x.ActionEnum_ACTION_POST_LIKED:
			av.Post, _ = PostSingleViewForPostId(act.PostId, MeId)

		case x.ActionEnum_ACTION_FOLLOWED_USER:
			av.FollowedUser = UserViewAndMe(act.PeerUserId, MeId)

		case x.ActionEnum_ACTION_POST_COMMENTED:
			av.Comment = SingleCommentView(act.CommentId)
		}
		res = append(res, av)
	}
	return res
}

func action_fillCaches(nots []*x.Action) {
	var pre_posts []int
	var pre_comments []int

	for _, nf := range nots {
		switch x.ActionEnum(nf.ActionTypeEnum) {
		case x.ActionEnum_ACTION_FOLLOWED_USER:

		case x.ActionEnum_ACTION_POST_LIKED:
			pre_posts = append(pre_posts, nf.PostId)

		case x.ActionEnum_ACTION_POST_COMMENTED:
			pre_comments = append(pre_comments, nf.PostId)
			pre_posts = append(pre_posts, nf.CommentId)
		}
	}

	x.Store.PreLoadCommentByCommentIds(pre_comments)
	x.Store.PreLoadPostByPostIds(pre_posts)
}
