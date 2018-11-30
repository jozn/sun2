package main

import (
	"fmt"
	"math/rand"
	"ms/sun/shared/helper"
	"ms/sun_old/base"
	//"ms/sun/shared/x"
	"ms/sun/shared/x"
	"time"
)

var write = 0
var read = 0

func main() {
	base.DefultConnectToMysql()

	fn := func() {
		x, err := x.NewHomeFanout_Selector().ForUserId_Eq(rand.Intn(10000)).GetRows(base.DB)
		read++

		if read%100 == 0 {
			fmt.Println("read ", read, len(x), err)
		}

		/*//witr=e
		  arr := make([]x.HomeFanout, 0, 5000)
		  for i := 0; i < 1000; i++ {
		      p := x.HomeFanout{
		          OrderId:        helper.NextRowsSeqId(),
		          ForUserId: rand.Intn(10000)+1,
		          PostId:    helper.NextRowsSeqId(),
		      }
		      write++
		      arr = append(arr, p)
		  }

		  if write%1 == 0{

		      fmt.Println("write: ", write)
		  }
		  x.MassReplace_HomeFanout(arr, base.DB)*/
	}
	//
	//go f4()
	//go f4()
	//go f4()
	//
	//go f5()
	//go f5()
	//go f5()

	for i := 0; i < 6; i++ {
		go fn()

	}

	time.Sleep(time.Hour)
}

func all() {
	x, err := x.NewHomeFanout_Selector().ForUserId_Eq(rand.Intn(10000)).GetRows(base.DB)
	read++

	if read%100 == 0 {
		fmt.Println("read ", read, len(x), err)
	}

	//witr=e
	arr := make([]x.HomeFanout, 0, 5000)
	for i := 0; i < 1000; i++ {
		p := x.HomeFanout{
			OrderId:   helper.NextRowsSeqId(),
			ForUserId: rand.Intn(10000) + 1,
			PostId:    helper.NextRowsSeqId(),
		}
		write++
		arr = append(arr, p)
	}

	if write%1 == 0 {

		fmt.Println("write: ", write)
	}
	x.MassReplace_HomeFanout(arr, base.DB)

}
func f5() {

	for {
		x, err := x.NewHomeFanout_Selector().ForUserId_Eq(rand.Intn(10000)).GetRows(base.DB)
		read++

		if read%100 == 0 {
			fmt.Println("read ", read, len(x), err)
		}
		time.Sleep(time.Millisecond * 10)
	}
}

func f4() {
	for {
		arr := make([]x.HomeFanout, 0, 5000)
		for i := 0; i < 1000; i++ {
			p := x.HomeFanout{
				OrderId:   helper.NextRowsSeqId(),
				ForUserId: rand.Intn(10000) + 1,
				PostId:    helper.NextRowsSeqId(),
			}
			write++
			arr = append(arr, p)
		}

		if write%1 == 0 {

			fmt.Println("write: ", write)
		}
		x.MassReplace_HomeFanout(arr, base.DB)
		time.Sleep(time.Millisecond * 10)
	}

}
