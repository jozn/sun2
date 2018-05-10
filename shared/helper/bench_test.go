package helper_test

import (
    "testing"
    "ms/sun/shared/helper"
    "fmt"
)

func BenchmarkSqlManyDollars(b *testing.B) {
    s:=helper.SqlManyDollars(4, 10000, false)
    fmt.Println(s)
	for i := 0; i < b.N; i++ {
		helper.SqlManyDollars(4, 10000, true)
	}
}
