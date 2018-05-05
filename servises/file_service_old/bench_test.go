package file_service_old

import (
	//"ms/sun/servises/file_service"
	"net/url"
	"testing"
)

func BenchmarkNewRowReq(b *testing.B) {

	r, _ := url.Parse("http://localhost:5151/getfile/1518504671057010006.jpg")
	for i := 0; i < b.N; i++ {
		NewRowReq(r)
	}
}

func BenchmarkDisk(b *testing.B) {

	r, _ := url.Parse("http://localhost:5151/getfile/1518504671057010006.jpg")
	rr, _ := NewRowReq(r)
	for i := 0; i < b.N; i++ {
		rr.isLocalCacheAvailable()
	}
}
