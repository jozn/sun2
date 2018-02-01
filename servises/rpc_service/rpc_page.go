package rpc_service

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/servises/memcache_service"
	"ms/sun2/servises/view_service"
	"ms/sun2/shared/x"
	"strings"
)

type rpc_page int

func (rpc_page) GetCommentsPage(param *x.PB_PageParam_GetCommentsPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetCommentsPage, err error) {
	res.CommentViewList = view_service.GetCommentsOfPost(int(param.PostId), int(param.Limit), int(param.Last), userParam.GetUserId())
	return
}

func (rpc_page) GetHomePage(param *x.PB_PageParam_GetHomePage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetHomePage, err error) {
	fids := memcache_service.GetUserFollowsSlice(userParam.GetUserId())
	ins := append(fids, userParam.GetUserId())
	selctor := x.NewPost_Selector().UserId_In(ins).OrderBy_PostId_Desc().Limit(int(param.Limit))

	if param.Last > 0 {
		selctor.PostId_LT(int(param.Last))
	}

	posts, err := selctor.GetRows(base.DB)
	if err != nil {
		return
	}
	res.PostViewList = view_service.PostsViews(posts, userParam.GetUserId())
	return
}

func (rpc_page) GetProfilePage(param *x.PB_PageParam_GetProfilePage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetProfilePage, err error) {
	panic("implement me")
}

func (rpc_page) GetLikesPage(param *x.PB_PageParam_GetLikesPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetLikesPage, err error) {
	pid := int(param.PostId)
	limit := 50
	if param.Limit != 0 {
		limit = int(param.Limit)
	}
	if pid < 1 {
		return
	}

	selector := x.NewLike_Selector().
		PostId_Eq(pid).
		OrderBy_Id_Desc().
		Limit(limit)

	if param.Last > 0 {
		selector.Id_LT(int(param.Last))
	}
	rows, err := selector.GetRows(base.DB)
	if err != nil {
		helper.DebugErr(err)
		return
	}
	uids := make([]int, 0, len(rows))
	for _, row := range rows {
		uids = append(uids, row.UserId)
	}
	x.Store.PreLoadUserByUserIds(uids)

	r := make([]*x.PB_UserViewRowify, 0)
	for _, row := range rows {
		v := &x.PB_UserViewRowify{
			Id:          int64(row.Id),
			CreatedTime: int32(row.CreatedTime),
			UserView:    view_service.UserViewAndMe(row.UserId, userParam.GetUserId()),
		}
		r = append(r, v)
	}
	res.UserViewRowifyList = r

	return
	/*pid := int(param.PostId)
	if pid < 1 {
		return
	}

	selector := x.NewLike_Selector().Select_UserId().PostId_Eq(pid)
	if param.Last > 0 {
		selector.Id_LT(int(param.Last))
	}
	UserIds, err := selector.OrderBy_Id_Desc().Limit(int(param.Limit)).GetIntSlice(base.DB)
	if err != nil {
		helper.DebugErr(err)
		return
	}
	res.UserViewList = view_service.UserSliceViewAndMe(UserIds, userParam.GetUserId())*/

	return
}

func (rpc_page) GetFollowersPage(param *x.PB_PageParam_GetFollowersPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetFollowersPage, err error) {
	uid := int(param.UserId)
	limit := 50
	if param.Limit != 0 {
		limit = int(param.Limit)
	}
	if uid < 1 {
		return
	}

	selector := x.NewFollowingListMember_Selector().
		FollowedUserId_Eq(uid).
		OrderBy_Id_Desc().
		Limit(limit)

	if param.Last > 0 {
		selector.Id_LT(int(param.Last))
	}
	rows, err := selector.GetRows(base.DB)
	if err != nil {
		helper.DebugErr(err)
		return
	}
	uids := make([]int, 0, len(rows))
	for _, row := range rows {
		uids = append(uids, row.UserId)
	}
	x.Store.PreLoadUserByUserIds(uids)

	r := make([]*x.PB_UserViewRowify, 0)
	for _, row := range rows {
		v := &x.PB_UserViewRowify{
			Id:          int64(row.Id),
			CreatedTime: int32(row.CreatedTime),
			UserView:    view_service.UserViewAndMe(row.UserId, userParam.GetUserId()),
		}
		r = append(r, v)
	}
	res.UserViewRowifyList = r

	return
}

func (rpc_page) GetFollowingsPage(param *x.PB_PageParam_GetFollowingsPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetFollowingsPage, err error) {
	//copy of above with adjustments
	uid := int(param.UserId)
	limit := 50
	if param.Limit != 0 {
		limit = int(param.Limit)
	}
	if uid < 1 {
		return
	}

	selector := x.NewFollowingListMember_Selector().
		UserId_Eq(uid).
		OrderBy_Id_Desc().
		Limit(limit)

	if param.Last > 0 {
		selector.Id_LT(int(param.Last))
	}
	rows, err := selector.GetRows(base.DB)
	if err != nil {
		helper.DebugErr(err)
		return
	}
	uids := make([]int, 0, len(rows))
	for _, row := range rows {
		uids = append(uids, row.FollowedUserId)
	}
	x.Store.PreLoadUserByUserIds(uids)

	r := make([]*x.PB_UserViewRowify, 0)
	for _, row := range rows {
		v := &x.PB_UserViewRowify{
			Id:          int64(row.Id),
			CreatedTime: int32(row.CreatedTime),
			UserView:    view_service.UserViewAndMe(row.FollowedUserId, userParam.GetUserId()),
		}
		r = append(r, v)
	}
	res.UserViewRowifyList = r

	return
}

func (rpc_page) GetNotifiesPage(param *x.PB_PageParam_GetNotifiesPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetNotifiesPage, err error) {

	res.NotifyViewList = view_service.Notify_GetLastsViews(userParam.GetUserId(), int(param.Last))
	return
}

func (rpc_page) GetUserActionsPage(param *x.PB_PageParam_GetUserActionsPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetUserActionsPage, err error) {
	res.ActionViewList = view_service.Action_GetLastsViews(userParam.GetUserId(), 100, int(param.Last))
	return
}

func (rpc_page) GetSuggestedPostsPage(param *x.PB_PageParam_GetSuggestedPostsPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetSuggestedPostsPage, err error) {
	panic("implement me")
}

func (rpc_page) GetSuggestedUsrsPage(param *x.PB_PageParam_GetSuggestedUsrsPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetSuggestedUsrsPage, err error) {
	panic("implement me")
}

func (rpc_page) GetSuggestedTagsPage(param *x.PB_PageParam_GetSuggestedTagsPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetSuggestedTagsPage, err error) {
	panic("implement me")
}

func (rpc_page) GetLastPostsPage(param *x.PB_PageParam_GetLastPostsPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetLastPostsPage, err error) {
	limit := 30
	if param.Limit != 0 {
		limit = int(param.Limit)
	}
	selector := x.NewPost_Selector().
		Select_PostId().
		OrderBy_PostId_Desc().
		Limit(limit)

	if param.Last > 0 {
		selector.PostId_LT(int(param.Last))
	}

	posts, err := selector.GetIntSlice(base.DB)
	if err != nil {
		return
	}

	res.PostViewList = view_service.PostsViewsForPostIds(posts, userParam.GetUserId())
	return
}

func (rpc_page) GetTagPage(param *x.PB_PageParam_GetTagPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_GetTagPage, err error) {
	tagName := param.Tag
	tagName = strings.Replace(tagName, "#", "", -1)
	limit := 30
	if param.Limit != 0 {
		limit = int(param.Limit)
	}

	tag, err := x.NewTag_Selector().Name_Eq(tagName).GetRow(base.DB)

	if err != nil {
		return
	}

	selector := x.NewTagsPost_Selector().Select_PostId().TagId_Eq(tag.TagId).OrderBy_Id_Desc().Limit(limit)
	if param.Last > 0 {
		selector.Id_LT(int(param.Last))
	}
	postIds, err := selector.GetIntSlice(base.DB)
	res.PostViewList = view_service.PostsViewsForPostIds(postIds, userParam.GetUserId())
	return
}

func (rpc_page) SearchTagsPage(param *x.PB_PageParam_SearchTagsPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_SearchTagsPage, err error) {
	panic("implement me")
}

func (rpc_page) SearchUsersPage(param *x.PB_PageParam_SearchUsersPage, userParam x.RPC_UserParam) (res x.PB_PageResponse_SearchUsersPage, err error) {
	panic("implement me")
}
