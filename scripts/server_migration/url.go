package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 180; i++ {
		fmt.Printf("http://mardomsara.com:8000/file.z%02d\n", i)
	}

}
