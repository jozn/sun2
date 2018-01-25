package websocket_service

import (
	"fmt"
	"github.com/jozn/protobuf/proto"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"ms/sun/helper"
	"ms/sun2/servises/rpc_service"
	"ms/sun2/shared/config"
	"ms/sun2/shared/x"
	"net/http"
	"strings"
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

type rpcHttpResponseHandler struct {
	w http.ResponseWriter
	r *http.Request
}

var cnt = 0

func (h *rpcHttpResponseHandler) HandleOfflineResult(resOut x.RpcResponseOutput) {
	//fmt.Println("rpcHttpResponseHandler")
	//helper.PertyPrint(resOut)
	cnt++
	logHttpRpc.Println("s;" ,cnt)
	msg, ok := resOut.ResponseData.(proto.Message)
	if ok {
		bts, err := proto.Marshal(msg)
		if err != nil {
			if config.IS_DEBUG {
				log.Panic("ms- http rpc handling proto", err)
			}
			return
		}
		resToClient := &x.PB_ResponseToClient{
			ClientCallId: int64(resOut.CommandToServer.ClientCallId),
			PBClass:      resOut.PBClassName,
			RpcFullName:  resOut.RpcName,
			Data:         bts,
		}

		resToClientBts, err := proto.Marshal(resToClient)
		if err != nil {
			if config.IS_DEBUG {
				log.Panic("ms- http rpc handling proto clinet", err)
			}
			return
		}
		if cnt%500 == 0 {
			fmt.Println(cnt)
		}
		h.w.Write(resToClientBts)
	} else {
		if config.IS_DEBUG {
			log.Panic("ms- http HandleOfflineResult ResponseData is not of proto message ")
		}
	}
}

func (rpcHttpResponseHandler) IsUserOnlineResult(interface{}, error) {
	panic("implement me")
}

func (rpcHttpResponseHandler) HandelError(error) {
	panic("implement me")
}

//dep
func ServeHttpRpc_Path_old(w http.ResponseWriter, r *http.Request) {
	//defer recover()
	err := r.ParseForm()
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

	method := strings.Replace(r.URL.Path, "/", "", 10)
	logRpc.Println("WS serving http rpc: ", method)
	if config.IS_DEBUG {
		//fmt.Println("WS serving http rpc: ", method)
	}
	cmd := x.PB_CommandToServer{
		ClientCallId: 2155,
		Command:      method,
		Data:         []byte{},
	}
	rpcHttpResponseHandler2 := &rpcHttpResponseHandler{}

	x.HandleRpcs(cmd, userParam, rpc_service.RpcAll, rpcHttpResponseHandler2)
}
