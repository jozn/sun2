package file_store

import (
    "ms/sun/shared/xc"
    "ms/sun/shared/helper"
    "ms/sun/shared/dbs"
)

const(
    SOTRE_TYPE_DIRECT_FILE = "direct"
    SOTRE_TYPE_GROUP_FILE = "group"
    SOTRE_TYPE_POST_FLLE = "post"
    SOTRE_TYPE_AVATAR_FILE = "avatar"
    SOTRE_TYPE_OTHERS = "other"
)

func SavaFile(file_ref *xc.FileRef, data []byte) (error){
    if file_ref.Md5Key == "" {
        file_ref.Md5Key = helper.MD5BytesToString(data)
    }

    fileStorage,err := xc.NewFileStorage_Selector().
        Select_Md5Key().Select_UserId().Select_FirstFileRefId().
        Md5Key_Eq(file_ref.Md5Key).GetRow(dbs.CassndraSession)
    if err!= nil {
        //not exists save it
        //todo parse file instead of reling on user inputs
        fileStorage = &xc.FileStorage{
            FirstFileRefId: file_ref.FileRefId,
            Extension: file_ref.Extension,
            Height: file_ref.Height,
            Length: file_ref.Length,
            Md5Key: file_ref.Md5Key,
            MimeType: file_ref.MimeType,
            UserId: file_ref.UserId,
            Width: file_ref.Width,
            Zdata: data,
            ZdataThumb: []byte{},
        }
        err = fileStorage.Save(dbs.CassndraSession)
        if err!= nil {
            return err
        }
    }

    su := &xc.StorageUsage{
        FileRefId: file_ref.FileRefId,
        Md5Key: file_ref.Md5Key,
        StoreType: file_ref.StoreType,
    }
    su.Save(dbs.CassndraSession)

    file_ref.StorageUserId = fileStorage.UserId
    file_ref.FirstFileRefId = fileStorage.FirstFileRefId
    file_ref.Save(dbs.CassndraSession)

    return nil
}
