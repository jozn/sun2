package main

import (
    "ms/sun2/shared/x"
    "ms/sun/base"
    "net/http"
    "fmt"
)

func main()  {
    ids,_ := x.NewFilePost_Selector().Select_Id().GetIntSlice(base.DB)
    for _, id := range ids {
        u :=fmt.Sprintf("http://localhost:5151/getfile/%d.jpg",id)
        fmt.Println(u)
        http.Get(u)
    }
}
