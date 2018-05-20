package file_common_test

import (
	"ms/sun/servises/file_service/file_common"
	"net/url"
	"testing"
)

func BenchmarkRowReqParsingNotExists(b *testing.B) {
	url, _ := url.Parse("http://localhost:5151/post_file/1518506476136010007_180.jpg")
	for i := 0; i < b.N; i++ {
		req:= file_common.NewRowReq(file_common.HttpCategories[1], url)
		req.SetNewFileDataStoreIdAndExtntion(154165465465,".exe")
	}
}
