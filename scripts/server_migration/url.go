package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 300; i++ {
		fmt.Printf("http://mardomsara.com:8080/photos.z%02d\n", i)
	}

}
