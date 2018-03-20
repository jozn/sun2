package home_fanout_service

import (
    "ms/sun2/shared/x"
    "ms/sun/base"
    "ms/sun/helper"
)

func resetHomeFanoutForUser(userId int)  {
    x.NewHomeFanout_Deleter().ForUserId_Eq(userId).Delete(base.DB)
    helper.NanoTimeBeforeNowSeconds()
}


