package main

import (
    "ms/sun/shared/x"
    "ms/sun/shared/helper"
    "ms/sun/shared/dbs"
)

func main()  {

    p := x.PostCdb{
        PostId: helper.NextRowsSeqId(),
        UserId: 0,
        PostTypeEnum: 0,
        PostCategoryEnum: 0,
        MediaId: 0,
        PostKey: "",
        Text: helper.FactRandStrEmoji(10,true),
        RichText: "",
        MediaCount: 0,
        SharedTo: 0,
        DisableComment: 0,
        Source: 0,
        HasTag: 0,
        Seq: 0,
        CommentsCount: 0,
        LikesCount: 0,
        ViewsCount: 0,
        EditedTime: 0,
        CreatedTime: 0,
        ReSharedPostId: 0,
    }

    err := p.Save(dbs.DB_PG)
    helper.NoErr(err)
}
