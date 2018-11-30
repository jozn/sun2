package main

import (
	"fmt"
	"math/rand"
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	//"ms/sun/shared/x"
	"time"
    "ms/sun/shared/x"
)

var write = 0
var read =0

func main() {
	base.DefultConnectToMysql()

	for i := 0; i < 10; i++ {
		go all()
	}

	time.Sleep(time.Hour)
}

func reader()  {
    for  {
        rows,err:=x.NewHomeFanout_Selector().ForUserId_Eq(rand.Intn(10000)).GetRows(base.DB)
        read++

        if read%1 == 0 {
            fmt.Println("read ", read, len(rows),err)
        }

        time.Sleep(time.Millisecond*10)
    }

}
func writeer()  {
    for  {
        //witr=e
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
        x.MassReplace_HomeFanout(arr, base.DB)

        time.Sleep(time.Millisecond*10)
    }

}

func all()  {
    for  {
        rows,err:=x.NewHomeFanout_Selector().ForUserId_Eq(rand.Intn(10000)).OrderBy_OrderId_Desc().Limit(100).GetRows(base.DB)
        read++

        if read%1 == 0 {
            fmt.Println("read ", read, len(rows),err)
        }

        //witr=e
        arr := make([]x.HomeFanout, 0, 5000)
        for i := 0; i < 100; i++ {
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
        x.MassReplace_HomeFanout(arr, base.DB)

        time.Sleep(time.Millisecond*10)
    }

}
