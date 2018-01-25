package rpc_service

import "ms/sun2/shared/x"

type rpc_other int

func (rpc_other) Echo(param *x.PB_OtherParam_Echo, userParam x.RPC_UserParam) (res x.PB_OtherResponse_Echo, err error) {
	res.Text = " +++ " + param.Text + " Hamid ++++"
	return
}
