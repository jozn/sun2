package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(`cassandra_walker`)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Errorf("Error: %s", err)
	}

}
