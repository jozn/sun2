package model_service_bk

import (
	"bytes"
	"image"
	"ms/sun_old/base"
	"ms/sun/shared/helper"
	"ms/sun/servises/file_service"
	"ms/sun/shared/x"
)

type PostAddParam struct {
	UserId     int
	Text       string
	MediaBytes []byte
	TimeSecond int
}

func AddPost(pp PostAddParam) (int, error) {
	var err error
	post := &x.Post{
		PostId:         helper.NextRowsSeqId(),
		UserId:         pp.UserId,
		PostTypeEnum:   int(x.PostTypeEnum_POST_TEXT),
		MediaId:        0,
		Text:           pp.Text,
		RichText:       "",
		MediaCount:     0,
		SharedTo:       0,
		DisableComment: 0,
		HasTag:         0,
		CommentsCount:  0,
		LikesCount:     0,
		ViewsCount:     1,
		EditedTime:     0,
		CreatedTime:    helper.TimeNow(),
		ReSharedPostId: 0,
	}
	if len(pp.MediaBytes) > 0 {
		reader := bytes.NewBuffer(pp.MediaBytes)
		img, imgType, err := image.Decode(reader)
		if err != nil {
			return 0, err
		}

		media := &x.Media{
			MediaId:       helper.NextRowsSeqId(),
			UserId:        pp.UserId,
			PostId:        post.PostId,
			AlbumId:       0,
			MediaTypeEnum: int(x.PostTypeEnum_POST_PHOTO),
			Width:         img.Bounds().Max.X,
			Height:        img.Bounds().Max.Y,
			Size:          len(pp.MediaBytes),
			Duration:      0,
			Md5Hash:       helper.MD5BytesToString(pp.MediaBytes),
			Color:         helper.ExtractColoer(img),
			CreatedTime:   helper.TimeNow(),
		}
		rowCassndra := file_service.Row{
			Id:        media.MediaId,
			Data:      pp.MediaBytes,
			FileType:  int(file_service.IMAGE),
			Extension: "." + imageTypeToExt(imgType),
		}
		file_service.SavePostFile(rowCassndra)
		err = media.Save(base.DB)
		if err != nil {
			return 0, err
		}
		post.PostTypeEnum = int(x.PostTypeEnum_POST_PHOTO)
		post.MediaId = media.MediaId
	}
	err = post.Save(base.DB)
	if err != nil {
		return 0, err
	}
	return post.PostId, nil
}

func imageTypeToExt(typ string) string {
	if typ == "jpeg" {
		return "jpg"
	}
	return typ
}
