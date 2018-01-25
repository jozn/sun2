package main__test

import "testing"

func BenchmarkArry1(b *testing.B) {
    arr := []int{}

    for i := 0; i < b.N; i++ {
        arr = append(arr,i)
    }
}

func BenchmarkManyMap23(b *testing.B) {
    mp := make(map[int]map[int]int, 2000)
    for i := 0; i < b.N; i++ {
        mp[i] = make(map[int]int)
        for j := 0; j < 1000000; j++ {
            mp[i][j] = j
        }
    }
}