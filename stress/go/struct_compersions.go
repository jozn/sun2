package main

import "fmt"

type st1 struct {
    Id int
    Name string
}

type st2 struct {
    Id int
    Name string
}

func main(){
    s1:= st1{12,"abc2"}
    s2 := st2(s1)
    fmt.Println("", s1, s2)
}
