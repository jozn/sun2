package action_fanout_service

import (
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/servises/mem_user_service"
	"ms/sun/shared/x"
)

func resetActionFanoutForUser(userId int) {
	var toSaveArr []x.ActionFanout
	x.NewActionFanout_Deleter().ForUserId_Eq(userId).Delete(base.DB)

	um, ok := mem_user_service.GetForUser(userId)
	if !ok {
		return
	}

	followedIds := um.GetFollowed()
	for _, uid := range followedIds {
		fum, ok := mem_user_service.GetForUser(uid)
		if ok {
			acts := fum.GetLastActions()
			for _, act := range acts {
				if act == nil {
					continue
				}
				r := x.ActionFanout{
					OrderId:      helper.NanoRowIdAtSec(act.CreatedTime),
					ForUserId:    userId,
					ActionId:     act.ActionId,
					ActorUserId:  act.ActorUserId,
				}
				toSaveArr = append(toSaveArr, r)
			}
		}
	}

	x.MassReplace_ActionFanout(toSaveArr, base.DB)
}

func ResetActionFanoutAll() {
	uids, err := x.NewUser_Selector().Select_UserId().GetIntSlice(base.DB)
	if err != nil {
		return
	}
	for _, uid := range uids {
		resetActionFanoutForUser(uid)
	}
}
