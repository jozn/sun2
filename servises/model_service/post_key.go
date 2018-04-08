package model_service

import (
	"ms/sun_old/base"
	"ms/sun/shared/x"
	"time"
)

var nextPostKey = make(chan string)

func loadPostKeysLoop() {
	for {
		keys, err := x.NewPostKey_Selector().Used_Eq(0).Order_Rand().Limit(2).GetRows(base.DB)
		if err != nil {
			time.Sleep(time.Second)
			continue
		}
		var ids []int
		for _, postKey := range keys {
			ids = append(ids, postKey.Id)
		}
		x.NewPostKey_Updater().Id_In(ids).Used(1).Update(base.DB)
		for _, postKey := range keys {
			nextPostKey <- postKey.PostKeyStr
		}
	}
}
