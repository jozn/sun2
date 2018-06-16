package sun_utils

import "regexp"

var IranPhoneReg = regexp.MustCompile(`^989[\d]{9}$`)

func IsValidIranPhoneNumber(phone string) bool  {
    return IranPhoneReg.MatchString(phone)
}
