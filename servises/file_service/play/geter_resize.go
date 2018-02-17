package main

import (
	"fmt"
	"ms/sun/base"
	"ms/sun2/shared/x"
	"net/http"
	"time"
)

func main() {
	i := 0
	ids, _ := x.NewFilePost_Selector().Select_Id().GetIntSlice(base.DB)
	ch := make(chan int,100)
	for _, id := range ids {
		go func(c int) {
			ch <- 1
			u := fmt.Sprintf("http://localhost:5151/post_file/%d_1080.jpg", c)
            i++
			fmt.Println(i, u)
			http.Get(u)
			<-ch
		}(id)
	}

	time.Sleep(time.Hour)
	fmt.Println(i)

}
