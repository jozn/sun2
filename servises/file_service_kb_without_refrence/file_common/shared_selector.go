package file_common

import (
	"fmt"
    "ms/sun/shared/xc"
)

//const cdnAddrecc  = "http://file%s.msapp.ir/"
const cdnAddrecc = "http://file%d.localhost:1100/"

func SelectPostSharedPostMedia(post xc.FileRef) string {
    url := fmt.Sprintf(`%s%s/%d$%s`, getCdnId(post.FileRefId), PostFileCategory.UrlPath, post.FileRefId, post.Extension)
    return url
}
/*
func SelectPostSharedPostMedia(post x.PostMedia) string {
	url := fmt.Sprintf(`%s%s/%d$%s`, getCdnId(post.MediaId), PostFileCategory.UrlPath, post.MediaId, post.Extension)
	return url
}

func SelectSharedMsgFile(m x.PB_FileMsg) string {
	url := fmt.Sprintf(`%s%s/%d$%s`, getCdnId(int(m.Id)), MsgFileCategory.UrlPath, m.Id, m.Extension)
	return url
}

func SelectSharedAvatar(m x.User) string {
	url := fmt.Sprintf(`%s%s/%d$%s`, getCdnId(0), AvatarFileCategory.UrlPath, m.UserId, ".jpg")
	return url
}
*/
func getCdnId(objectId int) string {
	rem := (objectId % 10) + 1
	return fmt.Sprintf(cdnAddrecc, rem)
}
