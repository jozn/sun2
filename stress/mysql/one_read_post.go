package main

import (
	"fmt"
	"math/rand"
	"ms/sun/shared/x"
	"ms/sun_old/base"
)

func main() {
	n := 0
	next := func() int {
		n++
		return n
	}
	ids, _ := x.NewPost_Selector().Select_PostId().GetIntSlice(base.DB)
	_ = ids
	work := func() {
		for {
			//_, err := x.HomeFanoutByOrderId(base.DB, rand.Intn(40000))
			p, err := x.PostByPostId(base.DB, ids[rand.Intn(len(ids))])
			if n := next(); n%10000 == 0 {
				fmt.Println(n, p.PostId, err)
			}
		}
	}

	for i := 0; i < 9; i++ {
		go work()
	}
	work()

}
