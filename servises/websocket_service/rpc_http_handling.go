package websocket_service

import (
	"github.com/jozn/protobuf/proto"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"ms/sun/helper"
	"ms/sun2/servises/rpc_service"
	"ms/sun2/shared/config"
	"ms/sun2/shared/x"
	"net/http"
)

func ServeHttpRpc(w http.ResponseWriter, r *http.Request) {
	//defer recover()
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
	x.HandleRpcs(cmd, userParam, rpc_service.RpcAll, handler)
}

