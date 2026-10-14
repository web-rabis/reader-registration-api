package model

import "github.com/web-rabis/reader-client/protobuf"

type Dictionary struct {
	Id   int64  `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func (d *Dictionary) ToProto() *protobuf.Dictionary {
	if d == nil {
		return nil
	}
	return &protobuf.Dictionary{
		Id:   d.Id,
		Name: d.Name,
	}
}
func DictionaryFromProto(d *protobuf.Dictionary) *Dictionary {
	if d == nil {
		return nil
	}
	return &Dictionary{
		Id:   d.Id,
		Name: d.Name,
	}
}
func DictionaryListFromProto(l []*protobuf.Dictionary) []*Dictionary {
	list := make([]*Dictionary, len(l))
	for i, d := range l {
		list[i] = DictionaryFromProto(d)
	}
	return list
}
