package main

import (
	"fmt"
	"ms/sun/shared/x"
	"ms/sun_old/base"
	"net/http"
)

func main() {
	i := 0
	fn := func() {
        ids, _ := x.NewFilePost_Selector().Select_Id().GetIntSlice(base.DB)
        for _, id := range ids {
            u := fmt.Sprintf("http://localhost:5151/post_file/%d.jpg", id)
            fmt.Println(i, u)
            http.Get(u)
            i++
        }

        for _, id := range ids {
            u := fmt.Sprintf("http://localhost:5151/avatar/%d_1080.jpg", id)
            fmt.Println(i, u)
            http.Get(u)
            i++
        }

        for _, id := range ids {
            u := fmt.Sprintf("http://localhost:5151/post_file/%d_thumb.jpg", id)
            fmt.Println(i, u)
            http.Get(u)
            i++
        }

        ids, _ = x.NewFileMsg_Selector().Select_Id().GetIntSlice(base.DB)
        for _, id := range ids {
            u := fmt.Sprintf("http://localhost:5151/msg_file/%d.jpg", id)
            fmt.Println(i, u)
            http.Get(u)
            i++
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

	fmt.Println(i)

}
