package main__test

import (
    "testing"
    "math/rand"
)

func BenchmarkMap1(b *testing.B) {
    mp := make(map[int]bool,10)

    for i := 0; i < b.N; i++ {
        mp[i] = true
    }
}

func BenchmarkManyMap(b *testing.B) {
    mp := make(map[int]map[int]int, 2000)
    for i := 0; i < b.N; i++ {
        mp[i] = make(map[int]int)
        for j := 0; j < 10000; j++ {
            mp[i][j] = j
        }
    }
}

func BenchmarkMapGet(b *testing.B) {
    mp := make(map[int]int, 10)
    for i := 0; i < 100000; i++ {
        mp[rand.Intn(10000000000000)] = rand.Intn(10000000000000)
    }
    b.ResetTimer()

    for i := 0; i < b.N; i++ {
        _ = mp[i]
    }
}
