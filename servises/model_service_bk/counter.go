package model_service_bk

import (
	"ms/sun/base"
	"ms/sun2/shared/x"
)

type _cntImpl int

var Counter _cntImpl

func (t _cntImpl) UpdateUserFollowingCounts(UserId int, cnt int) {
	x.NewUser_Updater().UserId_Eq(UserId).FollowingCount_Increment(cnt).Update(base.DB)
}

func (t _cntImpl) UpdateUserFollowersCounts(UserId int, cnt int) {
	x.NewUser_Updater().UserId_Eq(UserId).FollowersCount_Increment(cnt).Update(base.DB)
}

func (t _cntImpl) UpdateUserPostsCounts(UserId int, cnt int) {
	x.NewUser_Updater().UserId_Eq(UserId).PostsCount_Increment(cnt).Update(base.DB)
}

func (t _cntImpl) IncerPostCommentsCount(PostId, CountDiff int) {
	x.NewPost_Updater().PostId_Eq(PostId).CommentsCount_Increment(CountDiff).Update(base.DB)
}
