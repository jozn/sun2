package event_service

import "ms/sun2/shared/x"

func preloadevents(events []x.Event) (res []GeneralEvent) {
	/*postsMp := make(map[int]*x.Post, 100)
	commentMp := make(map[int]*x.Comment, 100)
	actionsMp := make(map[int]*x.Action, 100)
	usersMp := make(map[int]*x.User, 100)*/

	postsArr := make([]int, 0, 100)
	commentsArr := make([]int, 0, 100)
	actionArr := make([]int, 0, 100)
	usersArr := make([]int, 0, 100)

	for _, e := range events {
		if e.PostId > 0 {
			postsArr = append(postsArr, e.PostId)
		}
		if e.CommentId > 0 {
			commentsArr = append(commentsArr, e.CommentId)
		}
		if e.ActionId > 0 {
			actionArr = append(actionArr, e.ActionId)
		}
		if e.ByUserId > 0 {
			usersArr = append(usersArr, e.ByUserId)
		}
		if e.PeerUserId > 0 {
			usersArr = append(usersArr, e.PeerUserId)
		}
	}

	x.Store.PreLoadPostByPostIds(postsArr)
	x.Store.PreLoadCommentByCommentIds(commentsArr)
	x.Store.PreLoadActionByActionIds(actionArr)
	x.Store.PreLoadUserByUserIds(usersArr)

	for _, e := range events {
		eg := GeneralEvent{
			Event: e,
		}
		if e.PostId > 0 {
			eg.Post, _ = x.Store.GetPostByPostId_JustCache(e.PostId)
		}
		if e.CommentId > 0 {
			eg.Comment, _ = x.Store.GetCommentByCommentId_JustCache(e.CommentId)
		}
		if e.ActionId > 0 {
			eg.Action, _ = x.Store.GetActionByActionId_JustCache(e.ActionId)
		}
		if e.ByUserId > 0 {
			eg.ByUser, _ = x.Store.GetUserByUserId_JustCache(e.ByUserId)
		}
		if e.PeerUserId > 0 {
			eg.PeerUser, _ = x.Store.GetUserByUserId_JustCache(e.PeerUserId)
		}

		res = append(res, eg)
	}
	return
}
