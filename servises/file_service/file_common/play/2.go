package main

import (
    "fmt"
    "ms/sun/servises/file_service/file_common"
    "ms/sun/shared/x"
)

func main()  {
    p:=x.PostMedia{
        MediaId: 1525756746191010004,
        UserId: 0,
        PostId: 0,
        AlbumId: 0,
        MediaTypeEnum: 0,
        Width: 400,
        Height: 0,
        Size: 0,
        Duration: 0,
        Extension: ".jpg",
        Md5Hash: "",
        Color: "",
        CreatedTime: 0,
        ViewCount: 0,
        Extra: "",
    }
    fmt.Println(file_common.SelectPostSharedPostMedia(p))
}
