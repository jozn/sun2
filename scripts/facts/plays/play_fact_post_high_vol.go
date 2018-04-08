package main

import (
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
    "fmt"
)

func main() {

	for j := 0; j < 10000; j++ {
        var posts []x.PostCopy
        for i := 0; i < 1000; i++ {
			p := x.PostCopy{
				PostId:         helper.NextRowsSeqId(),
				UserId:         654565,
				PostTypeEnum:   0,
				MediaId:        2,
				PostKey:        helper.RandAlphaNumbericString(10),
				Text:           helper.FactRandStrEmoji(200, true),
				RichText:       "",
				MediaCount:     0,
				SharedTo:       0,
				DisableComment: 0,
				Source:         0,
				HasTag:         0,
				Seq:            0,
				CommentsCount:  0,
				LikesCount:     0,
				ViewsCount:     55,
				EditedTime:     0,
				CreatedTime:    0,
				ReSharedPostId: 0,
			}
			posts = append(posts, p)
		}
		x.MassReplace_PostCopy(posts, base.DB)
		fmt.Println(j*1000)
	}

}
