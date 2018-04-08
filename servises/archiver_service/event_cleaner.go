package archiver_service

import (
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"time"
)

func cleanEvents() {
	go func() {
		for {
			time.Sleep(time.Second * 100)
			old := helper.NanoTimeBeforeNowSeconds(10)
			x.NewEvent_Deleter().EventId_LE(old).Delete(base.DB)
		}
	}()
}
