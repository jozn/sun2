package model_service_bk

import (
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/shared/x"
)
func Comment_Add(UserId, PostId int, Text string) x.Comment {
	cmt := x.Comment{
		CommentId:   helper.NextRowsSeqId(),
		UserId:      UserId,
		PostId:      PostId,
		CreatedTime: helper.TimeNow(),
		LikesCount:  0,
		Text:        Text,
	}

	err := cmt.Insert(base.DB)

	if err == nil {
		Counter.IncerPostCommentsCount(PostId, 1)
		post, ok := x.Store.GetPostByPostId(PostId)
		if ok {
			Notify_OnPostCommented(&cmt, post)
			Action_OnPostCommentAdd(&cmt, post)
		}
	}

	return cmt
}

func Comment_Delete(UserId, PostId, CommentId int) bool {
	post, _ := x.Store.GetPostByPostId(PostId)

	com, err := x.NewComment_Selector().CommentId_Eq(CommentId).UserId_Eq(UserId).PostId_Eq(PostId).GetRow(base.DB)
	if err != nil {
		com.Delete(base.DB)
		Counter.IncerPostCommentsCount(PostId, -1)
		Notify_OnPostCommentedDeleted(com, post)
		Action_OnPostCommentDeleted(com, post)
		return true
	}

	return false
}
