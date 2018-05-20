package file_disk_cache_test

import (
	"ms/sun/servises/file_service/file_common"
	"net/url"
	"testing"
    "ms/sun/servises/file_service_kb_without_refrence/file_disk_cache"
    "fmt"
)
var config = &file_common.FileServingConfig{
    DiskDirs:     []string{"D:/sun/a/", "D:/sun/b/"},
    FileServerId: 1,
}
func BenchmarkIsFileExists(b *testing.B) {
	url, _ := url.Parse("http://localhost:5151/post_file/1518506476136010007_180.jpg")
    req:= file_common.NewRowReq(file_common.HttpCategories[1], url)
    req.SetNewFileDataStoreIdAndExtntion(154165465465,".exe")
	for i := 0; i < b.N; i++ {
		file_disk_cache.IsLocalCacheAvailable(req,config)
	}
}

func BenchmarkIsFileExistsMAn(b *testing.B) {
    for i := 0; i < b.N; i++ {
        url, _ := url.Parse(fmt.Sprintf("http://localhost:5151/post_file/%d_180.jpg",i))
        req:= file_common.NewRowReq(file_common.HttpCategories[1], url)
        file_disk_cache.IsLocalCacheAvailable(req,config)
    }
}
