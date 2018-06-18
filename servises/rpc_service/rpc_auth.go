package rpc_service

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"ms/sun/servises/sun_utils"
	"ms/sun/servises/view_service"
	"ms/sun/shared/base"
	"ms/sun/shared/helper"
	"ms/sun/shared/x"
	"net/http"
	"regexp"
	"strings"
    url2 "net/url"
)

var sms_numbers = []string{"10008713774983", "10009007504275", "20009007504275"}

type rpc_auth int

var iranPhoneReg = regexp.MustCompile(`989[\d]{9}`)

func (rpc_auth) SendConfirmCode(param *x.RPC_Auth_Types_SendConfirmCode_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_SendConfirmCode_Response, errRes error) {

	res.SmsNumbers = sms_numbers

	if len(param.Hash) < 20 {
		res.Done = false
		res.ErrorMessage = "hash must be set."
		return
	}

	sms := &x.Sms{
		Hash:            param.Hash,
		ClientPhone:     param.Phone,
		GenratedCode:    0,
		SmsSenderNumber: "",
		SmsSendStatues:  "wrong",
		Carrier:         "",
		Country:         []byte{},
		IsValidPhone:    0,
		IsConfirmed:     0,
		IsLogin:         0,
		IsRegister:      0,
		RetriedCount:    0,
		TTL:             helper.TimeNow() + 24*3600,
	}

	valid := sun_utils.IsValidIranPhoneNumber(param.Phone)
	if !valid {
		res.Done = false
		res.ErrorMessage = "شماره معتبر نمی باشد. در حال حاظر فقط شماره های ایران پشتیبان می شود."
		sms.Save(base.DB)
		return
	}
	sp := &sendSmsParam{
		phone: param.Phone,
	}

	sendIranSms(sp)
	sms.SmsSenderNumber = sp.smsNumber
	sms.SmsHttpBody = sp.sentBody
	sms.GenratedCode = sp.code
	sms.SmsSendStatues = sp.sentResult
	if sp.err != nil {
		sms.Err = sp.err.Error()
	}
    sms.IsValidPhone = 1

	sms.Save(base.DB)

	if sp.err != nil {
		res.Done = false
		res.ErrorMessage = "خطایی در فرستادن sms رخ داد."
		return
	} else {
		res.Done = true
		res.ErrorMessage = ""
	}

	return
}

func (rpc_auth) ConfirmCode(param *x.RPC_Auth_Types_ConfirmCode_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_ConfirmCode_Response, errRes error) {

    fmt.Println("pppp: ",param)

	rows, err := x.NewSms_Selector().Hash_Eq(param.Hash).ClientPhone_Eq(param.Phone).GetRows(base.DB)
	if err != nil || len(param.Phone) == 0 {
		res.Done = false
		res.ErrorMessage = "پیدا نشد."
		return
	}

	fmt.Println("pppp: 2 ",param)

	for _, sms := range rows {
		if sms.GenratedCode == int(param.Code) {
			sms.IsConfirmed = 1
			user, ok := x.Store.User_ByPhone(param.Phone)
			if ok { //login
				sms.IsLogin = 1
				sms.Save(base.DB)
				res.Done = true
				res.ErrorMessage = ""
				res.SelfUserView = view_service.SelfUserView(user.UserId)
				return
			}
			//register

			sms.IsLogin = 0
			sms.IsRegister = 1
			sms.Save(base.DB)
			res.Done = true
			res.ErrorMessage = ""
			return
		}
	}

	res.Done = false
	res.ErrorMessage = "خطایی رخ داد، پیدا نشد."
	return
}

func (rpc_auth) SingUp(param *x.RPC_Auth_Types_SingUp_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_SingUp_Response, errRes error) {
	username := param.UserName
	if len(username) < 3 {
		username = fmt.Sprintf("user_%v", helper.RandAlphaNumbericString(6))
	}

	user := x.User{
		UserName:      username,
		UserNameLower: strings.ToLower(username),
		FirstName:     param.FirstName,
		LastName:      param.LastName,
		Phone:         param.Phone,
		Email:         param.Email,
		About:         "",
		PasswordHash:  "",
		PasswordSalt:  "",
		PostSeq:       0,
		CreatedTime:   helper.TimeNow(),
		VersionTime:   0,
	}

	//check email and phone validation
	sel := x.NewUser_Selector()
	if len(param.Email) > 0 {
	    sel.Email_Eq(param.Email)
    }
    if len(param.Phone) > 0 {
        sel.Phone_Eq(param.Phone)
    }
    us,err := sel.GetRows(base.DB)
    if err != nil || len(us) != 0 {
        res.Done = false
        res.ErrorMessage = "این ابمیل یا شماره قبلا ثیت شده است."
        return
    }

	err = user.Save(base.DB)
	if err != nil {
		res.Done = false
		res.ErrorMessage = "خطا در عضویت."
		return
	}

	res.Done = true
	res.SelfUserView = view_service.SelfUserView(user.UserId)

	return
}

func (rpc_auth) SingIn(param *x.RPC_Auth_Types_SingIn_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_SingIn_Response, errRes error) {
	res.Done = false
    res.ErrorMessage = "implement me"
	return
}

func (rpc_auth) LogOut(param *x.RPC_Auth_Types_LogOut_Param, userParam x.RPC_UserParam) (res x.RPC_Auth_Types_LogOut_Response, errRes error) {
    res.Done = false
    res.ErrorMessage = "implement me"
    return
}

//phone : 989....
type sendSmsParam struct {
	phone      string
	smsNumber  string
	code       int
	err        error
	sentBody   string
	sentResult string
}

func sendIranSms(p *sendSmsParam) {
	if len(p.phone) < 5 {

	}
    smsNumber := "10008713774983"
    username := "9164150318"
    passweord := "4067787000"
    code := 1000 + rand.Intn(9000)
    to := helper.StrToInt(string(p.phone[2:]), 0)
    txt := fmt.Sprintf(SEND_TEXT, code)

    p.sentResult = "fail"
    p.code = code
    p.sentResult = "not proceed"
    p.smsNumber = smsNumber

	url := fmt.Sprintf(`http://onlinepanel.ir/post/sendsms.ashx?from=%s&to=%d&text=%s&password=%s&username=%s`,
		smsNumber, to, url2.PathEscape(txt), passweord, username)
	res, err := http.Get(url)
	fmt.Println(url,err)
	if err != nil {
		p.err = err
		return
	}
	bts, err := ioutil.ReadAll(res.Body)

	if err != nil {
		p.err = err
		return
	}
	body := string(bts)
	p.sentBody = body

	if len(body) == 3 {
		parts := strings.Split(body, "-")
		if len(parts) == 2 && parts[1] == "0" {
			p.sentResult = "done"
		}
	}else {
	    fmt.Println("wrong sms url",url)
        p.sentResult = "fail send sms"
    }

}

const SEND_TEXT = `کد تایید: %d
با تشکر مردم سرا
mardomsara.com
`
const SEND_TEXT2 = `کد تایید: %d
با تشکر مردم\u2009سرا
`

/*
func getCodebasedOnClientHsh(hash string) int {

    // everyday just one code
    s := time.Now().Format("Jan _2 ") + hash
    digest
}

*/
