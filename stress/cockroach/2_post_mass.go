package main

import (
	"fmt"
	"ms/sun/shared/conns"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"sync/atomic"
)

func main() {
	x.LogTableSqlReq.PostCdb = false

	i := int64(0)
	fn := func() {
		var arr []x.PostCdb
		for ; i < 100000000; i++ {
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&i, 1)
				n := int(i)
				p := x.PostCdb{
					PostId:           helper.NextRowsSeqId(),
					UserId:           j,
					PostTypeEnum:     0,
					PostCategoryEnum: 0,
					MediaId:          n,
					PostKey:          "",
					Text:             helper.FactRandStrEmoji(50, true),
					RichText:         "",
					MediaCount:       0,
					SharedTo:         0,
					DisableComment:   0,
					Source:           0,
					HasTag:           0,
					Seq:              0,
					CommentsCount:    0,
					LikesCount:       0,
					ViewsCount:       0,
					EditedTime:       0,
					CreatedTime:      0,
					ReSharedPostId:   0,
				}
				arr = append(arr, p)
			}
			err := x.MassInsert_PostCdb(arr, conns.DB_PG)
			arr = arr[:0]
			if i%100 == 0 {
				fmt.Println(i)
			}
			helper.NoErr(err)
		}
	}

	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	fn()

}
