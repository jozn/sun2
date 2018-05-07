package file_service

import (
	"fmt"
	"os/exec"
)

type Resizer struct {
	InputFullPath  string
	OutputFullPath string
	Quality        int
	Width          int
}

func (f *Resizer) ResizeFFMPEG() {
	if f.Width <= 0 {
		f.Width = 100
	}
	scale := fmt.Sprintf("scale=%d:-1", f.Width)
	arg := []string{
		"-i",
		f.InputFullPath,
		"-vf",
		scale,
		f.OutputFullPath,
	}
	cmd := exec.Command("ffmpeg", arg...)
	err := cmd.Run()
	if err != nil {
		println(err)
	}
}
