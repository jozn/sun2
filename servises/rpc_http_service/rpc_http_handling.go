package rpc_http_service

import (
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"ms/sun/shared/helper"
	"ms/sun/servises/rpc_service"
	"ms/sun/shared/config"
	"ms/sun/shared/x"
	"net/http"
    "fmt"
    "github.com/golang/protobuf/proto"
    "ms/sun/servises/monitor_service"
)

func ServeHttpRpc(w http.ResponseWriter, r *http.Request) {
    //fmt.Println(r)
	//defer recover()
    fmt.Println(")))))))))))))))))")
	err := r.ParseMultipartForm(1280000)
	helper.NoErr(err)
	session := r.Form.Get("SessionUid")
	_ = session

	uid := 6
	uid = helper.StrToInt(r.Form.Get("user_id"), 6)
	//TODO add session check functiality
	if false {
		//e(session)
		return //not Autorized user
	}

	userParam := RPC_UserParam_Imple{
		userId: uid,
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		if config.IS_DEBUG {
			log.Panic("ms- http can not get file for pb", err)
		}
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		if config.IS_DEBUG {
			log.Panic("ms- can read bytes", err)
		}
	}

	cmd := x.PB_CommandToServer{}

	err = proto.Unmarshal(bytes, &cmd)
	if err != nil {
		if config.IS_DEBUG {
			log.Panic("ms- http rpc handling proto", err)
		}
	}
	handler := &rpcHttpResponseHandler{
		w: w,
		r: r,
	}

	//log.Printf(" serving &s command via http rpc \n", cmd.Command)
	monitor_service.AddRpc(cmd.Command)
    fmt.Println("+++++++++++++++++++=")
	x.HandleRpcs(cmd, userParam, rpc_service.RpcAll, handler)
}


type RPC_UserParam_Imple struct {
    userId int
}

func (s RPC_UserParam_Imple) GetUserId() int {
    return s.userId
}

func (s RPC_UserParam_Imple) IsUser() bool {
    return s.userId > 0
}