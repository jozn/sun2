package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

func main() {
	for {
		fmt.Println(genUUIDSeq())
		time.Sleep(time.Millisecond * 50)
	}
}

func genUUIDSeq() (string, error) {
	var uuid [16]byte
	var rander = rand.Reader
	_, err := io.ReadFull(rander, uuid[:])
	if err != nil {
		return "", err
	}
	now := time.Now()
    //fmt.Printf("%x\n",now.UnixNano())
    //fmt.Printf("%x\n",now.Unix())
    PutUint64(uuid[:],now.UnixNano())

    b := uuid
	uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuidStr, nil
}

func PutUint64(b []byte, v int64) {
    _ = b[7] // early bounds check to guarantee safety of writes below
    b[0] = byte(v >> 56)
    b[1] = byte(v >> 48)
    b[2] = byte(v >> 40)
    b[3] = byte(v >> 32)
    b[4] = byte(v >> 24)
    b[5] = byte(v >> 16)
    //fmt.Println( byte(v >> 56))
    //b[6] = byte(v >> 8)
    //b[7] = byte(v)
}
func genUUIDSeq_dep() (string, error) {
    var uuid [16]byte
    var rander = rand.Reader
    _, err := io.ReadFull(rander, uuid[:])
    if err != nil {
        return "", err
    }
    now := time.Now()
    timeSecPart := now.Unix()
    timeAllPart := now.UnixNano()
    tiSec32 := uint32(timeSecPart)
    tiMilli32 := uint16((now.UnixNano() - now.Unix()) % 65535) //more tahtn mili
    binary.BigEndian.PutUint32(uuid[0:4], tiSec32)
    binary.BigEndian.PutUint16(uuid[4:6], tiMilli32)
    uuid[4] = byte(timeAllPart >> 32)
    uuid[5] = byte(timeAllPart >> 40)
    fmt.Println(uuid[4:6])
    fmt.Printf("%x\n",timeAllPart)

    b := uuid
    uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
    return uuidStr, nil
}