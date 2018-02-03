package main

import (
    "sort"
    "fmt"
    "ms/sun2/shared/golib/sorter"
    "math/rand"
)

func main() {
    arr := []sorter.IntWeight{
        sorter.IntWeight{1, 8},
        sorter.IntWeight{2, 11},
        sorter.IntWeight{19, 42},
        sorter.IntWeight{18, 22},
        sorter.IntWeight{14, 2},
        sorter.IntWeight{10, 0},
        sorter.IntWeight{100, 12},
        sorter.IntWeight{15, 100},
    }
    for i := 0; i < 1000; i++ {
        arr = append(arr, sorter.IntWeight{rand.Intn(5000), rand.Intn(50000)})
    }
    a := sorter.IntWeightSlice(arr)
    sort.Sort(a)
    fmt.Println(a)

}

func f1()  {

}
