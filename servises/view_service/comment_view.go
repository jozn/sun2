package view_service

import (
	"ms/sun/base"
	"ms/sun2/shared/x"
)

func SingleCommentView(CommentId int) *x.PB_CommentView {
	com, ok := x.Store.GetCommentByCommentId(CommentId)
	if ok {
		return &x.PB_CommentView{
			CommentId:   int64(com.CommentId),
			UserId:      int32(com.UserId),
			PostId:      int64(com.PostId),
			Text:        com.Text,
			LikesCount:  int32(com.LikesCount),
			CreatedTime: int32(com.CreatedTime),
		}
	}
	return nil
}

func GetCommentsOfPost(PostId, Limit, Last, MeId int) (views []*x.PB_CommentView) {
	selector := x.NewComment_Selector().PostId_Eq(PostId).OrderBy_CommentId_Desc().Limit(int(Limit))

	if Last > 0 {
		selector.CommentId_LT(int(Last))
	}

	comments, err := selector.GetRows2(base.DB)
	if err != nil {
		return
	}

	var userIds = make([]int, 0, len(comments))
	for _, com := range comments {
		userIds = append(userIds, com.UserId)
	}

	x.Store.PreLoadUserByUserIds(userIds)
	views = make([]*x.PB_CommentView, 0, len(comments))
	for _, com := range comments {
		v := &x.PB_CommentView{
			CommentId:   int64(com.CommentId),
			UserId:      int32(com.UserId),
			PostId:      int64(com.PostId),
			Text:        com.Text,
			LikesCount:  int32(com.LikesCount),
			CreatedTime: int32(com.CreatedTime),
		}
		v.SenderUserView = UserViewAndMe(com.UserId, MeId)
		views = append(views, v)
	}
	return
}
