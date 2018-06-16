package sun_utils_test

import (
	"ms/sun/servises/sun_utils"
	"testing"
)

func TestIranRexex(t *testing.T) {
	var tests = []struct {
		phone string
		pass  bool
	}{
		{"989015132328", true},
		{"989315139328", true},
		{"989175132328", true},
		{"989905132328", true},

		{"98 9905132328", false},
		{"979015132328", false},
		{"98901532328", false},
		{"982101532328", false},
		{"9890151323282", false},
		{"9890151323282222", false},
		{"9890151sss222", false},
		{"9015132328", false},
		{"9871152782283", false},
	}

	for i, test := range tests {
		r := sun_utils.IranPhoneReg.MatchString(test.phone)
		if r != test.pass {
			t.Error("Iran regex failed", test, r, i)
		}
	}

}
