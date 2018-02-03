package ds__test


import (
    "ms/sun/ds"
    "testing"
    "math/rand"
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
    //fmt.Println(collection)
}

//740ns : use map it is faster and more reliable
func BenchmarkContains(b *testing.B) {
    collection := ds.New()
    for i := 0; i < 1000; i++ {
        collection.Add(rand.Intn(10000))
    }
    collection.SortAsc()
    //c := collection.
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        collection.Contains(i)
    }
    //fmt.Println(collection)
}


