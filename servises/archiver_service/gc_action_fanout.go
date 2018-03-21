package archiver_service

import (
	"ms/sun/base"
	"ms/sun2/shared/x"
	"time"
)

const ACTION_FANOUT_LIMIT_PER_USER = 1000

func gc_action_fanout() {
	time.Sleep(time.Minute)
	for {
		time.Sleep(time.Second)
		userIds, err := x.NewUser_Selector().Select_UserId().Order_Rand().GetIntSlice(base.DB)
		if err != nil {
			continue
		}

		for _, uid := range userIds {
			time.Sleep(time.Millisecond * 1000)
			act, err := x.NewActionFanout_Selector().ForUserId_Eq(uid).OrderBy_OrderId_Desc().Limit(1).Offset(ACTION_FANOUT_LIMIT_PER_USER).GetRow(base.DB)
			if err != nil || act == nil {
				continue
			}
			x.NewActionFanout_Deleter().ForUserId_Eq(uid).OrderId_LT(act.OrderId).Delete(base.DB)
		}
	}
}
