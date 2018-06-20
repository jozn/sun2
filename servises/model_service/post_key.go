package model_service

import (
	"ms/sun_old/base"
	"ms/sun/shared/x"
	"time"
    "math/rand"
)

var nextPostKey = make(chan string)

func loadPostKeysLoop() {
	for {
	    rnd := rand.Intn(10000)
		//keys, err := x.NewPostKeys_Selector().Used_Eq(0).Order_Rand().Limit(2).GetRows(base.DB)
		keys, err := x.NewPostKeys_Selector().RandShard_GE(rnd).OrderBy_Id_Asc().Limit(2).GetRows(base.DB)
		if err != nil {
			time.Sleep(time.Millisecond*100)
			continue
		}
		var ids []int
		for _, postKey := range keys {
			ids = append(ids, postKey.Id)
		}
		x.NewPostKeys_Deleter().Id_In(ids).Delete(base.DB)
		for _, postKey := range keys {
			nextPostKey <- postKey.PostKeyStr
		}
	}
}
