package archiver_service

import (
	"ms/sun/base"
	"ms/sun/helper"
	"ms/sun2/shared/x"
	"time"
)

func gc_trigger() {
	time.Sleep(time.Minute)
	for {
		time.Sleep(time.Second * 10)
		x.NewTriggerLog_Deleter().CreatedSe_LT(helper.TimeNow()).Delete(base.DB)
	}
}
