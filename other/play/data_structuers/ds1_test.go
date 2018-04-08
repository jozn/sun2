package main_test

import (
	"ms/sun_old/ds"
	"testing"
    "math/rand"
    "fmt"
)

func BenchmarkDs1(b *testing.B) {
	collection := ds.New()
	for i := 0; i < 1000; i++ {
		collection.Add(rand.Intn(1000000))
	}
    collection.SortAsc()
    //c := collection.
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        collection.SortAsc()
    }
    fmt.Println(collection)
}
