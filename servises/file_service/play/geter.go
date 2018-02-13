package main

import (
	"fmt"
	"ms/sun/base"
	"ms/sun2/shared/x"
	"net/http"
)

func main() {
	i := 0
	ids, _ := x.NewFilePost_Selector().Select_Id().GetIntSlice(base.DB)
	for _, id := range ids {
		u := fmt.Sprintf("http://localhost:5151/post_file/%d.jpg", id)
		fmt.Println(u)
		http.Get(u)
		i++
	}

	for _, id := range ids {
		u := fmt.Sprintf("http://localhost:5151/post_file/%d.jpg", id)
		fmt.Println(u)
		http.Get(u)
		i++
	}

    ids, _ = x.NewFileMsg_Selector().Select_Id().GetIntSlice(base.DB)
    for _, id := range ids {
        u := fmt.Sprintf("http://localhost:5151/msg_file/%d.jpg", id)
        fmt.Println(u)
        http.Get(u)
        i++
    }

    fmt.Println(i)

}
