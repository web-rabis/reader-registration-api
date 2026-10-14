package model

import "github.com/web-rabis/reader-client/protobuf"

type SocialStatus struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Code string `json:"code" db:"code"`
}

func (d *SocialStatus) ToProto() *protobuf.SocialStatus {
	if d == nil {
		return nil
	}
	return &protobuf.SocialStatus{
		Id:   d.Id,
		Code: d.Code,
		Name: d.Name,
	}
}
func SocialStatusFromProto(d *protobuf.SocialStatus) *SocialStatus {
	if d == nil {
		return nil
	}
	return &SocialStatus{
		Id:   d.Id,
		Code: d.Code,
		Name: d.Name,
	}
}
func SocialStatusListFromProto(l []*protobuf.SocialStatus) []*SocialStatus {
	list := make([]*SocialStatus, len(l))
	for i, d := range l {
		list[i] = SocialStatusFromProto(d)
	}
	return list
}
