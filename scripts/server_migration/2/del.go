package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	cnt := 0
	done := 0
	doneErr := 0
	var delChan = make(chan string, 1)
	file, err := os.Open("./ignore.txt")
	if err != nil {
		log.Panic(err)
	}
	lines := bufio.NewReader(file)

	for i := 0; i < 1; i++ {
		go func() {
			for st := range delChan {
				/*cmd := exec.Command("rm", strings.TrimSpace(st))
				err = cmd.Start()
				output, err := cmd.CombinedOutput()*/
				//print(cmd)
				err = os.RemoveAll(st)
				if err != nil {
					fmt.Println("del err : ", err)
					doneErr++
				} else {
					done++
				}
				cnt++
				if cnt%10000 == 0 {
					fmt.Println("del: ", st)
				}
			}
		}()
	}

	go func() {
		for {
			time.Sleep(time.Second * 30)
			fmt.Println("delete done: ", done)
			fmt.Println("delete err: ", doneErr)
		}
	}()

	for {
		st, err := lines.ReadString('\n')
		if err != nil {
			break
			fmt.Println("err reading file : ", err)
		}
		delChan <- st
	}

	close(delChan)
	fmt.Println("finished ")
}
