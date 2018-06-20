package rpc_service

import (
	"ms/sun/servises/model_service"
	"ms/sun/servises/view_service"
	"ms/sun/shared/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
    "ms/sun/servises/post_counter_service"
)

type rpc_socila int

func (rpc_socila) AddComment(param *x.RPC_Social_Types_AddComment_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_AddComment_Response, errRes error) {
	cmt, ok := model_service.Comment_Add(userParam.GetUserId(), int(param.PostId), param.Text)
	if !ok {
		res.Done = false
		res.Error = &x.PB_Error{
			ShowError:    true,
			ErrorMessage: "خطایی رخ داد",
		}
		return
	}
	res.Done = true
	res.Comment = view_service.SingleCommentView(cmt.CommentId)
	return
}

func (rpc_socila) DeleteComment(param *x.RPC_Social_Types_DeleteComment_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_DeleteComment_Response, errRes error) {
	ok := model_service.Comment_Delete(userParam.GetUserId(), int(param.PostId), int(param.CommentId))
	if !ok {
		res.Done = false
		res.Error = &x.PB_Error{
			ShowError:    true,
			ErrorMessage: "خطایی رخ داد",
		}
		return
	}
	res.Done = true

	return
}

func (rpc_socila) EditComment(param *x.RPC_Social_Types_EditComment_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_EditComment_Response, errRes error) {

	comment, ok := x.Store.GetCommentByCommentId(int(param.CommentId))
	if !ok || comment.UserId != userParam.GetUserId() || len(param.Text) == 0 {
		res.Done = false
		res.Error = &x.PB_Error{
			ShowError:    true,
			ErrorMessage: "خطایی رخ داد",
		}
		return
	}
	comment.Text = param.Text
	comment.IsEdited = 1
	comment.Save(base.DB)
	res.Done = false
	return
}

func (rpc_socila) LikeComment(param *x.RPC_Social_Types_LikeComment_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_LikeComment_Response, errRes error) {

	return
}

func (rpc_socila) AddPost(param *x.RPC_Social_Types_AddPost_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_AddPost_Response, errRes error) {
	postParam := model_service.PostAddParam{
		UserId:     userParam.GetUserId(),
		Text:       param.Text,
		MediaBytes: param.ImageBlob,
	}
	postId, err := model_service.AddPost(postParam)
	if err != nil {
		res.Done = false
		res.Error = &x.PB_Error{
			ShowError:    true,
			ErrorMessage: "خطایی رخ داد",
		}
		return
	}
	res.Done = true
	res.PostView, _ = view_service.PostSingleViewForPostId(postId, userParam.GetUserId())
	return
}

func (rpc_socila) EditPost(param *x.RPC_Social_Types_EditPost_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_EditPost_Response, errRes error) {
	x.NewPost_Updater().PostId(int(param.PostId)).UserId_Eq(int(userParam.GetUserId())).Text(param.Text).Update(base.DB)
	res.PostView, _ = view_service.PostSingleViewForPostId(int(param.PostId), userParam.GetUserId())
	return
}

func (rpc_socila) DeletePost(param *x.RPC_Social_Types_DeletePost_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_DeletePost_Response, errRes error) {
	post, ok := x.Store.GetPostByPostId(int(param.PostId))
	if !ok || post.PostId != userParam.GetUserId() {
		res.Done = false
		res.Error = someErr()
		return
	}
	model_service.DeletePostFully(userParam.GetUserId(), post.PostId)

	res.Done = true
	return
}

func (rpc_socila) ArchivePost(param *x.RPC_Social_Types_ArchivePost_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_ArchivePost_Response, errRes error) {
	return
}

func (rpc_socila) PromotePost(param *x.RPC_Social_Types_PromotePost_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_PromotePost_Response, errRes error) {
	pro := &x.PostPromoted{
		PostId:      int(param.PostId),
		ByUserId:    userParam.GetUserId(),
		PostUserId:  0,
		BazzarUuid:  "---",
		Package:     "sample",
		EndTime:     helper.TimeNow() + 86400,
		CreatedTime: helper.TimeNow(),
	}
	pro.Save(base.DB)

	res.Done = true

	return
}

func (rpc_socila) LikePost(param *x.RPC_Social_Types_LikePost_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_LikePost_Response, errRes error) {
	model_service.Like_LikePost(userParam.GetUserId(), int(param.PostId))
	res.Done = true
	return
}

func (rpc_socila) UnLikePost(param *x.RPC_Social_Types_UnLikePost_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_UnLikePost_Response, errRes error) {
	model_service.Like_UnlikePost(userParam.GetUserId(), int(param.PostId))
	res.Done = true
	return
}

func (rpc_socila) FollowUser(param *x.RPC_Social_Types_FollowUser_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_FollowUser_Response, errRes error) {
	model_service.Follow(userParam.GetUserId(), int(param.UserId))
	res.Done = true
	return
}

func (rpc_socila) UnFollowUser(param *x.RPC_Social_Types_UnFollowUser_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_UnFollowUser_Response, errRes error) {
	model_service.UnFollow(userParam.GetUserId(), int(param.UserId))
	res.Done = true
	return
}

func (rpc_socila) BlockUser(param *x.RPC_Social_Types_BlockUser_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_BlockUser_Response, errRes error) {
	b := &x.Blocked{
		UserId:        userParam.GetUserId(),
		BlockedUserId: int(param.UserId),
		CreatedTime:   helper.TimeNow(),
	}
	b.Save(base.DB)

	res.Done = true
	return
}

func (rpc_socila) UnBlockUser(param *x.RPC_Social_Types_UnBlockUser_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_UnBlockUser_Response, errRes error) {
	x.NewBlocked_Deleter().UserId_Eq(userParam.GetUserId()).BlockedUserId_Eq(int(param.UserId)).Delete(base.DB)
	res.Done = true
	return
}

func (rpc_socila) AddSeenPosts(param *x.RPC_Social_Types_AddSeenPosts_Param, userParam x.RPC_UserParam) (res x.RPC_Social_Types_AddSeenPosts_Response, errRes error) {
	for id := range param.PostIds {
        post_counter_service.AddSeen(int(id))
    }
    res.Done = true
    res.RandHash = param.RandHash
    return
}

func someErr() *x.PB_Error {
	return &x.PB_Error{
		ShowError:    true,
		ErrorMessage: "خطایی رخ داد",
	}
}
