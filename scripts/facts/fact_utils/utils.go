package fact_utils

import (
    "math/rand"
    "ms/sun2/shared/x"
    "ms/sun/base"
    "io/ioutil"
    "log"
)

func GetRandUserId() int {
    return rand.Intn(80)+1
}

var pids []int
func GetRandPostId() int {
    if len(pids) == 0 {
        pids, _ = x.NewPost_Selector().Select_PostId().GetIntSlice(base.DB)
    }
    return pids[rand.Intn(len(pids))]
}

func RandImage() (fn string, bs []byte, err error) {
    const dir = `C:\Go\_gopath\src\ms\sun\upload\samples`
    imageFiles, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }
    fn = dir + "/" + imageFiles[rand.Intn(len(imageFiles))].Name()

    bs, err = ioutil.ReadFile(fn)
    if err != nil {
        log.Fatal(err)
    }
    //size += len(bs)
    return fn[2:], bs, nil

}