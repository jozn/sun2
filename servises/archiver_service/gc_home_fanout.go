package archiver_service

import (
	"ms/sun/base"
	"ms/sun2/shared/x"
	"time"
)

const HOME_FANOUT_LIMIT_PER_USER = 500

func gc_home_fanout() {
	time.Sleep(time.Minute)
	for {
		time.Sleep(time.Second)
		userIds, err := x.NewUser_Selector().Select_UserId().Order_Rand().GetIntSlice(base.DB)
		if err != nil {
			continue
		}

		for _, uid := range userIds {
			time.Sleep(time.Millisecond * 1000)
			home, err := x.NewHomeFanout_Selector().ForUserId_Eq(uid).OrderBy_OrderId_Desc().Limit(1).Offset(HOME_FANOUT_LIMIT_PER_USER).GetRow(base.DB)
			if err != nil || home == nil {
				continue
			}
			x.NewHomeFanout_Deleter().ForUserId_Eq(uid).OrderId_LT(home.OrderId).Delete(base.DB)
		}
	}
}
