package main

import (
	"fmt"
	"os/exec"
)

type resizer struct {
	inputFullPath  string
	outputFullPath string
	quality        int
	width          int
}

func (f *resizer) resizeFFMPEG() {
	if f.width <= 0 {
		f.width = 100
	}
	scale := fmt.Sprintf("scale=%d:-1", f.width)
	arg := []string{
		"-i",
		f.inputFullPath,
		"-vf",
		scale,
		f.outputFullPath,
	}
	cmd := exec.Command("ffmpeg", arg...)
	err := cmd.Run()
	if err != nil {
		println(err)
	}
}

func main() {
	s := resizer{
		inputFullPath:  `C:\Go\_gopath\src\ms\sun\servises\file_service\play\img\1.jpg`,
		outputFullPath: `C:\Go\_gopath\src\ms\sun\servises\file_service\play\img\1_ot3.jpg`,
		width:          500,
	}

	s.resizeFFMPEG()
}
