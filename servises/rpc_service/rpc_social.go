package rpc_service

import (
	"ms/sun/base"
	"ms/sun2/servises/model_service"
	"ms/sun2/servises/view_service"
	"ms/sun2/shared/x"
)

type rpc_socila int

func (rpc_socila) AddComment(param *x.PB_SocialParam_AddComment, userParam x.RPC_UserParam) (res x.PB_SocialResponse_AddComment, err error) {
	if len(param.Text) == 0 || param.PostId < 1 {
		res.Extra = nil //todo implement this
		return
	}
	cmt := model_service.Comment_Add(userParam.GetUserId(), int(param.PostId), param.Text)
	res.Comment = view_service.SingleCommentView(cmt.CommentId)
	return
}

func (rpc_socila) DeleteComment(param *x.PB_SocialParam_DeleteComment, userParam x.RPC_UserParam) (res x.PB_SocialResponse_DeleteComment, err error) {
	ok := model_service.Comment_Delete(userParam.GetUserId(), int(param.PostId), int(param.CommentId))
	res.Deleted = ok
	return
}

func (rpc_socila) AddPost(param *x.PB_SocialParam_AddPost, userParam x.RPC_UserParam) (res x.PB_SocialResponse_AddPost, err error) {
	postParam := model_service.PostAddParam{
		UserId:     userParam.GetUserId(),
		Text:       param.Text,
		MediaBytes: param.ImageBlob,
	}
	postId, err := model_service.AddPost(postParam)
	if err != nil {
		return
	}
	res.PostView, _ = view_service.PostSingleViewForPostId(postId, userParam.GetUserId())

	return
}

func (rpc_socila) EditPost(param *x.PB_SocialParam_EditPost, userParam x.RPC_UserParam) (res x.PB_SocialResponse_EditPost, err error) {
	x.NewPost_Updater().PostId(int(param.PostId)).UserId_Eq(int(userParam.GetUserId())).Text(param.Text).Update(base.DB)
	return
}

func (rpc_socila) DeletePost(param *x.PB_SocialParam_DeletePost, userParam x.RPC_UserParam) (res x.PB_SocialResponse_DeletePost, err error) {
	pan("implement me")
	return
}

func (rpc_socila) ArchivePost(param *x.PB_SocialParam_ArchivePost, userParam x.RPC_UserParam) (res x.PB_SocialResponse_ArchivePost, err error) {
	pan("implement me")
	return
}

func (rpc_socila) LikePost(param *x.PB_SocialParam_LikePost, userParam x.RPC_UserParam) (res x.PB_SocialResponse_LikePost, err error) {
	model_service.Like_LikePost(userParam.GetUserId(), int(param.PostId))
	return
}

func (rpc_socila) UnLikePost(param *x.PB_SocialParam_UnLikePost, userParam x.RPC_UserParam) (res x.PB_SocialResponse_UnLikePost, err error) {
	model_service.Like_UnlikePost(userParam.GetUserId(), int(param.PostId))
	return
}

//todo fix me
func (rpc_socila) FollowUser(param *x.PB_SocialParam_FollowUser, userParam x.RPC_UserParam) (res x.PB_SocialResponse_FollowUser, err error) {
	model_service.Follow(userParam.GetUserId(), int(param.UserId))

	return
}

func (rpc_socila) UnFollowUser(param *x.PB_SocialParam_UnFollowUser, userParam x.RPC_UserParam) (res x.PB_SocialResponse_UnFollowUser, err error) {
	model_service.UnFollow(userParam.GetUserId(), int(param.UserId))
	return
}
