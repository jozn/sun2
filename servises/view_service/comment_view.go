package view_service

import (
	"ms/sun2/shared/x"
)

func SingleCommentView(CommentId int) *x.PB_CommentView {
	com, ok := x.Store.GetCommentByCommentId(CommentId)
	if ok {
		return &x.PB_CommentView{
			CommentId:   int64(com.CommentId),
			UserId:      int32(com.UserId),
			PostId:      int32(com.PostId),
			Text:        com.Text,
			LikesCount:  int32(com.LikesCount),
			CreatedTime: int32(com.CreatedTime),
		}
	}
	return nil
}
