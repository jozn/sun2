package main

import (
	"ms/sun_old/helper"
	"ms/sun/shared/x"
    "time"
    "fmt"
    "runtime"
)

func main() {

	arr := make([]*x.Post,0, 1000000)
	mp := make(map[int]*x.Post,1000000)
	for i := 1; i < 1000000; i++ {
        p := &x.Post{
            Id:             helper.NextRowsSeqId(),
            UserId:         10,
            TypeId:         0,
            Text:           "ljgsdg slgjsdg sdglsdjg sgjsg sgsdjg ",
            FormatedText:   "",
            MediaCount:     10,
            SharedTo:       20,
            DisableComment: 0,
            HasTag:         10,
            LikesCount:     0,
            CommentsCount:  0,
            EditedTime:     165410,
            CreatedTime:    450,
        }
        arr = append(arr,p)
        mp[p.Id] = p
	}
    fmt.Println(len(arr))
    runtime.GC()
	time.Sleep(time.Second*1000)

}
