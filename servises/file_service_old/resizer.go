package file_service_old

import (
    "os/exec"
    "fmt"
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