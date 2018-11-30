package main

import (
	"fmt"
	"math/rand"
	"ms/sun/shared/x"
	"ms/sun_old/base"
	"time"
)

var i2 = 0

func main() {
	base.DefultConnectToMysql()

	for i := 0; i < 6; i++ {

		go f2()
	}

	time.Sleep(time.Hour)
}

func f2() {
	for {
		arr := make([]x.Session, 0, 10000)
		for i := 0; i < 1000; i++ {
			p := x.Session{
				UserId:                rand.Intn(1522000),
				SessionUuid:           "ljgsldg sgksglkjsg sdjglsdkg ",
				ClientUuid:            "sljgk sgklsdjg sgksdg sdgsdgj",
				DeviceUuid:            "sgjlsdkgj slgkjs gskgjs lgskgj s",
				LastActivityTime:      0,
				LastIpAddress:         "slgjskgjsdg sgkjs",
				LastWifiMacAddress:    "",
				LastNetworkType:       "",
				LastNetworkTypeEnumId: 0,
				AppVersion:            0,
				UpdatedTime:           0,
				CreatedTime:           0,
			}

			arr = append(arr, p)
		}
		i2++
		fmt.Println(i2)
		x.MassReplace_Session(arr, base.DB)
	}

}
