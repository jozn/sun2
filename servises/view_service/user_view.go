package view_service

import (
	"ms/sun/servises/memcache_service"
	"ms/sun/shared/x"
)

func UserViewAndMe(UserId, MeId int) *x.PB_UserView {
	user, ok := x.Store.GetUserByUserId(UserId)
	if !ok {
		return &x.PB_UserView{
			UserId:    int32(0),
			UserName:  "",
			FirstName: "کاربر حذف شده",
			LastName:  "",
		}
	}

	view := &x.PB_UserView{
		UserId:             int32(user.UserId),
		UserName:           user.UserName,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		AvatarId:           int64(user.AvatarId),
		ProfilePrivacyEnum: int32(user.ProfilePrivacyEnum),
		Phone:              int64(user.Phone),
		About:              user.About,
		FollowersCount:     int32(user.FollowersCount),
		FollowingCount:     int32(user.FollowingCount),
		PostsCount:         int32(user.PostsCount),
		MediaCount:         int32(user.MediaCount),
		LikesCount:         int32(user.LikesCount),
		ResharedCount:      int32(user.ResharedCount),
		//LastActiveTime: int32(user.LastActiveTime),
	}

	view.MyFollwing = memcache_service.DoesUserFollows(MeId, UserId)

	return view
}

func UserSliceViewAndMe(UserIds []int, MeId int) (res []*x.PB_UserView) {
	x.Store.PreLoadUserByUserIds(UserIds)
	x.Store.PreLoadUserByUserIds(UserIds)

	for _, uid := range UserIds {
		res = append(res, UserViewAndMe(uid, MeId))
	}
	return
}
