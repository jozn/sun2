package rpc_service

import (
	"ms/sun/shared/base"
	"ms/sun/shared/x"
)

type rpc_general int

func (rpc_general) Echo(param *x.RPC_General_Types_Echo_Param, userParam x.RPC_UserParam) (res x.RPC_General_Types_Echo_Response, errRes error) {
	res.Done = true
	res.Text = param.Text
	return
}

func (rpc_general) CheckUserName(param *x.RPC_General_Types_CheckUserName_Param, userParam x.RPC_UserParam) (res x.RPC_General_Types_CheckUserName_Response, errRes error) {
	res.UserName = param.UserName
	if len(param.UserName) < 3 {
        res.IsAvailable = false
		return
	}
	username, err := x.NewUser_Selector().Select_UserNameLower().UserNameLower_Eq(param.UserName).GetString(base.DB)
	if err != nil || len(username) == 0 {
		res.IsAvailable = true
	} else {
		res.IsAvailable = false
	}
	//fmt.Println("3", res)
	return
}
