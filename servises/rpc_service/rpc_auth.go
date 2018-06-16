package rpc_service

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"ms/sun/servises/sun_utils"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"net/http"
	"regexp"
)

type rpc_auth int

var iranPhoneReg = regexp.MustCompile(`989[\d]{9}`)

func (rpc_auth) SendConfirmCode(param *x.RPC_Auth_Types_SendConfirmCode_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_SendConfirmCode_Response, err error) {

	if len(param.Hash) < 20 {
		res.Done = false
		res.ErrorMessage = "hash must be set."
		return
	}

	valid := sun_utils.IsValidIranPhoneNumber(param.Phone)
	if !valid {
		res.Done = false
		res.ErrorMessage = "شماره معتبر نمی باشد. در حال حاظر فقط شماره های ایران پشتیبان می شود."
		return
	}

	return
}

func (rpc_auth) ConfirmCode(param *x.RPC_Auth_Types_ConfirmCode_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_ConfirmCode_Response, err error) {

	return
}

func (rpc_auth) SingUp(param *x.RPC_Auth_Types_SingUp_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_SingUp_Response, err error) {

	return
}

func (rpc_auth) SingIn(param *x.RPC_Auth_Types_SingIn_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_SingIn_Response, err error) {
	res.Done = false
	return
}

func (rpc_auth) LogOut(param *x.RPC_Auth_Types_LogOut_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_LogOut_Response, err error) {
	panic("implement me")
}

//phone : 989....
func sendIranSms(phone string) {
	if len(phone) < 5 {

	}
	smsNumber := 10008713774983
	username := "9164150318"
	passweord := "4067787"
	code := 1000 + rand.Intn(9000)
	to := helper.StrToInt(string(phone[2:]), 0)
	txt := fmt.Sprintf(SEND_TEXT, code)
	url := fmt.Sprintf(`http://onlinepanel.ir/post/sendsms.ashx?from=%d&to=%d&text=%s&password=%s&username=%s`,
		smsNumber, to, txt, passweord, username)
	res, err := http.Get(url)
	if err != nil {

	}
	bts, err := ioutil.ReadAll(res.Body)

	if err != nil {

	}
	body := string(bts)

}

const SEND_TEXT = `کد تایید: %d
با تشکر مردم\u2009سرا
`

/*
func getCodebasedOnClientHsh(hash string) int {

    // everyday just one code
    s := time.Now().Format("Jan _2 ") + hash
    digest
}

*/
