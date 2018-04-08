package home_fanout_service

import (
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/servises/mem_user_service"
	"ms/sun/shared/x"
)

func resetHomeFanoutForUser(userId int) {
	var toSaveArr []x.HomeFanout
	x.NewHomeFanout_Deleter().ForUserId_Eq(userId).Delete(base.DB)

	um, ok := mem_user_service.GetForUser(userId)
	if !ok {
		return
	}

	followedIds := um.GetFollowed()
	for _, uid := range followedIds {
		fum, ok := mem_user_service.GetForUser(uid)
		if ok {
			posts := fum.GetLastPosts()
			for _, post := range posts {
				if post == nil {
					continue
				}
				r := x.HomeFanout{
					OrderId:    helper.NanoRowIdAtSec(post.CreatedTime),
					ForUserId:  userId,
					PostId:     post.PostId,
					PostUserId: post.UserId,
				}
				toSaveArr = append(toSaveArr, r)
			}
		}
	}

	x.MassReplace_HomeFanout(toSaveArr, base.DB)
}

func ResetHomeFanoutAll() {
	uids, err := x.NewUser_Selector().Select_UserId().GetIntSlice(base.DB)
	if err != nil {
		return
	}
	for _, uid := range uids {
		resetHomeFanoutForUser(uid)
	}
}
