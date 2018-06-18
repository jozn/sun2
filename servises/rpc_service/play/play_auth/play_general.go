package main

import (
	"fmt"
	"ms/sun/servises/rpc_http_service"
	"net/http"
    "ms/sun/shared/x"
    "ms/sun/servises/rpc_http_service/client_http"
)

func main() {
	go func() {
		http.HandleFunc("/rpc", rpc_http_service.ServeHttpRpc)
		http.ListenAndServe("localhost:9999", nil)
	}()
	i := 0
	fmt.Println("1")
	i++
	fmt.Println("------------- send sms -------------------")

    /*checkUsernam := func(uname string) {
        fmt.Println("============ sendSms ===============")
        d := &x.RPC_General_Types_CheckUserName_Param{
            UserName: uname,
        }

        res, err := x.RPC_AllClinetsPlay.RPC_General.CheckUserName(d, client_http.FN)
        fmt.Printf("result of rpc final: %#+v  %s \n", res, err)
    }
    _ = checkUsernam
*/

    for _,un := range []string{"nn","mmm","jgjjh","user_kiaqfk"} {
        checkUsernam(un)
    }
}

func checkUsernam (uname string) {
    fmt.Println("============ checkUsernam ===============")
    d := &x.RPC_General_Types_CheckUserName_Param{
        UserName: uname,
    }
    fmt.Println("req:", d)

    res, err := x.RPC_AllClinetsPlay.RPC_General.CheckUserName(d, client_http.FN)
    fmt.Printf("result of rpc final: %#+v  %s \n", res, err)
}


