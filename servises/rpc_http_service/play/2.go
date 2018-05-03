package main

import (
	"fmt"
	"github.com/hokaccha/go-prettyjson"
	"ms/sun/shared/x"
)

func main() {
	x := x.User{}
	v := map[string]interface{}{
		"str":   "foo",
		"num":   100,
		"bool":  false,
		"null":  nil,
		"array": []string{"foo", "bar", "baz"},
		"map": map[string]interface{}{
			"foo": "bar",
		},
	}

	s, _ := prettyjson.Marshal(v)
	fmt.Println(string(s))
	s, _ = prettyjson.Marshal(x)
	fmt.Println(string(s))
}
