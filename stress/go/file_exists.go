package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 0; i < 1000000; i++ {
		//p := fmt.Sprintf("E:/%d/file_%d.jpg", i)
		//p := fmt.Sprintf("E:/WEB_MS/files/pics/%.jpg", i)
		//p := fmt.Sprintf("E:/WEB_MS/files/file/dl/file/pic/photo/%.jpg", i)
		p := fmt.Sprintf("E:/Intro to Parallel Programming Videos.zip")
		ok := Exists(p)
		if i%10000 == 0 {
			fmt.Println(i, ok)
		}
	}
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
