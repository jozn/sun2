package model_service

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/servises/event_service"
	"ms/sun2/shared/x"
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
			hash := hashPostCommented(cmt.PostId, cmt.CommentId)
			event := x.Event{
				ByUserId:     UserId,
				PeerUserId:   post.UserId,
				PostId:       cmt.PostId,
				CommentId:    cmt.CommentId,
				ActionId:     helper.NextRowsSeqId(),
				Murmur64Hash: hash,
			}

			event_service.SaveEvent(event_service.COMMENTED_POST_EVENT, event)
		}
	}

	return cmt
}

func Comment_Delete(UserId, PostId, CommentId int) bool {
	post, _ := x.Store.GetPostByPostId(PostId)

	cmt, err := x.NewComment_Selector().CommentId_Eq(CommentId).UserId_Eq(UserId).PostId_Eq(PostId).GetRow(base.DB)
	if err != nil {
		cmt.Delete(base.DB)
		Counter.IncerPostCommentsCount(PostId, -1)

		hash := hashPostCommented(post.PostId, cmt.CommentId)
		event := x.Event{
			ByUserId:     UserId,
			PeerUserId:   post.UserId,
			PostId:       cmt.PostId,
			CommentId:    cmt.CommentId,
			ActionId:     helper.NextRowsSeqId(),
			Murmur64Hash: hash,
		}

		event_service.SaveEvent(event_service.UNCOMMENTED_POST_EVENT, event)
		return true
	}

	return false
}
