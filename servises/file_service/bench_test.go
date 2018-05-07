package file_service_test

import (
	"ms/sun/servises/file_service/file_common"
	"ms/sun/servises/file_service/file_disk_cache"
	"testing"
    url2 "net/url"
    "fmt"
)

func BenchmarkRun(b *testing.B) {
    url,_:= url2.Parse("http://localhost:5153/post_file/1519150360431010060_1080.jpg")
    var config = &file_common.FileServingConfig{
        DiskDirs:     []string{"D:/sun/a/", "D:/sun/b/"},
        FileServerId: 1,
    }
    row := file_common.NewRowReq(file_common.HttpCategories[0], url)
	for i := 0; i < b.N; i++ {
		row = file_common.NewRowReq(file_common.HttpCategories[0], url)
	}
    file_disk_cache.SelectDisk(row, config)

}

func BenchmarkRunIsFileExists(b *testing.B) {
    for i := 0; i < b.N; i++ {
        p := fmt.Sprintf("D:/sun/a/v1/post_file/jpg/15/19/15/03/1519150360431010060%d.jpg",i)
        file_disk_cache.IsFileExists(p)
    }
}
