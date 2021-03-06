package model_service

import (
	"ms/sun_old/base"
	"ms/sun/shared/helper"
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
	user, err := x.UserByUserId(base.DB, pp.UserId)
	if err != nil {
		return 0, err
	}
	user.PostSeq += 1
	post := &x.Post{
		PostId:           helper.NextRowsSeqId(),
		UserId:           pp.UserId,
		PostType:     int(x.PostTypeEnum_POST_TEXT),
		MediaId:          0,
		PostKey:          <-nextPostKey,
		Text:             pp.Text,
		RichText:         "",
		MediaCount:       0,
		SharedTo:         0,
		DisableComment:   0,
		Seq:              user.PostSeq,
		CommentsCount:    0,
		LikesCount:       0,
		ViewsCount:       1,
		EditedTime:       0,
		CreatedTime:      helper.TimeNow(),
	}
	/*if len(pp.MediaBytes) > 0 {
		reader := bytes.NewBuffer(pp.MediaBytes)
		img, imgType, err := image.Decode(reader)
		if err != nil {
			return 0, err
		}

		media := &x.PostMedia{
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
		rowCassndra := file_service_old.Row{
			Id:        media.MediaId,
			Data:      pp.MediaBytes,
			FileType:  int(file_service_old.IMAGE),
			Extension: "." + imageTypeToExt(imgType),
		}
		file_store.SavaFile()
		file_service_old.SavePostFile(rowCassndra)
		err = media.Save(base.DB)
		if err != nil {
			return 0, err
		}
		post.PostTypeEnum = int(x.PostTypeEnum_POST_PHOTO)
		post.MediaId = media.MediaId
	}*/
	postCount := &x.PostCount{
		PostId:     post.PostId,
		ViewsCount: 0,
	}
	err = post.Save(base.DB)
	_ = postCount.Save(base.DB)
	if err != nil {
		return 0, err
	}
	Counter.IncermentUserPostSeq(pp.UserId)

    postTextsRelationshipsHandle(*post)

	return post.PostId, nil
}

//todo update post counter
func DeletePostFully(userId, postId int) {
    Counter.UpdateUserPostsCounts(userId,-1)
	DeletePostsFully([]int{postId})
}



//todo: remove from cassandra too
func DeletePostsFully(postIds []int) {
	if len(postIds) == 0 {
		return
	}

	x.NewPost_Deleter().PostId_In(postIds).Delete(base.DB)

	//related posts types: comments,likes,...
	x.NewAction_Deleter().PostId_In(postIds).Delete(base.DB)
	x.NewComment_Deleter().PostId_In(postIds).Delete(base.DB)
	x.NewLikes_Deleter().PostId_In(postIds).Delete(base.DB)
	x.NewNotify_Deleter().PostId_In(postIds).Delete(base.DB)

	//post
	x.NewPostCount_Deleter().PostId_In(postIds).Delete(base.DB)
	x.NewPostMedia_Deleter().PostId_In(postIds).Delete(base.DB)
	x.NewPostReshared_Deleter().PostId_In(postIds).Delete(base.DB)
	x.NewTagPost_Deleter().PostId_In(postIds).Delete(base.DB)

	//profile
    x.NewProfileAll_Deleter().PostId_In(postIds).Delete(base.DB)
    x.NewProfileMedia_Deleter().PostId_In(postIds).Delete(base.DB)
    x.NewProfileMentioned_Deleter().PostId_In(postIds).Delete(base.DB)

	//metas
	x.NewHomeFanout_Deleter().PostId_In(postIds).Delete(base.DB)
	x.NewSuggestedTopPosts_Deleter().PostId_In(postIds).Delete(base.DB)

	var arr []x.PostDeleted
	for _, id := range postIds {
		pd := x.PostDeleted{
			PostId: id,
		}
		arr = append(arr, pd)
	}
	x.MassReplace_PostDeleted(arr, base.DB)
}

/*func imageTypeToExt(typ string) string {
    if typ == "jpeg" {
        return "jpg"
    }
    return typ
}
*/