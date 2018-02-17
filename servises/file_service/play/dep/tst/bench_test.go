package tst

import (
    "testing"
    "regexp"
)

func BenchmarkF22(b *testing.B) {
    for i := 0; i < b.N; i++ {
        f25()
        //helper.StrToInt("500",0)
    }
}

var size = regexp.MustCompile(`/([0-9]+(_[[:alpha:]]+)?\.\w+)`)
//var size = regexp.MustCompile(`.+_([[:alpha:]]+)\..+`)
//var size = regexp.MustCompile(`.+_(\d+)\..+`)
var thumb = regexp.MustCompile(`.+_thumb\..+`)
func f25()  {
    url := "http://localhost:5151/msg_file/1518506476136010007_thumb.jpg"

    size.FindStringSubmatch(url)
    //thumb.FindStringSubmatch(url)
}
