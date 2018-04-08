package archiver_service

import (
	"ms/sun_old/base"
	"ms/sun_old/helper"
	"ms/sun/shared/x"
	"time"
)

func gc_trigger() {
	time.Sleep(time.Minute)
	for {
		time.Sleep(time.Second * 10)
		x.NewTriggerLog_Deleter().CreatedSe_LT(helper.TimeNow()).Delete(base.DB)
	}
}
