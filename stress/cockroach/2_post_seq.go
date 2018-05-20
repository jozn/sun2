package main

import (
	"ms/sun/shared/dbs"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"sync/atomic"
    "time"
    "fmt"
)

func main() {
    x.LogTableSqlReq.PostCdb = false
	i := int64(0)
	fn := func() {
		for {
			atomic.AddInt64(&i, 1)

			p := x.PostCdb{
				PostId:           helper.NextRowsSeqId(),
				UserId:           0,
				PostTypeEnum:     0,
				PostCategoryEnum: 0,
				MediaId:          int(i),
				PostKey:          "",
				Text:             helper.FactRandStrEmoji(10, true),
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

			err := p.Insert(dbs.DB_PG)
			helper.NoErr(err)
		}

	}

	go func() {
        for  {
            time.Sleep(time.Second)
            fmt.Println(i)
        }
    }()






	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
	go fn()
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
