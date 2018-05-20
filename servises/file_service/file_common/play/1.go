package main

import (
	"ms/sun/servises/file_service/file_common"
	"net/url"
    "ms/sun/shared/helper"
)

func main() {
	url, _ := url.Parse("http://localhost:5151/post_file/1518506476136010007_180.jpg")
	req := file_common.NewRowReq(file_common.HttpCategories[1], url)

	//helper.PertyPrintNew(req)
	helper.PertyPrint(req)

    url, _ = url.Parse("http://localhost:5151/post_file/1518506476136010007.mov")
    req = file_common.NewRowReq(file_common.HttpCategories[1], url)

    req.SetNewFileDataStoreId(5000000000)

    //helper.PertyPrintNew(req)
    helper.PertyPrint(req)
}
