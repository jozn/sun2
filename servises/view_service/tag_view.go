package view_service

import "ms/sun2/shared/x"

func ToTagView(tags []*x.Tag) (res []*x.PB_TagView) {
	for _, row := range tags {
		v := &x.PB_TagView{
			TagId:         int64(row.TagId),
			Name:          row.Name,
			Count:         int32(row.Count),
			TagStatusEnum: int32(row.TagStatusEnum),
			CreatedTime:   int32(row.CreatedTime),
		}
		res = append(res, v)
	}
	return
}
