package model

import "github.com/web-rabis/reader-client/protobuf"

type TypeCard struct {
	Id         int64  `json:"id" db:"id"`
	Type       string `json:"type" db:"type"`
	Code       string `json:"code" db:"code"`
	Web        string `json:"web" db:"saitbibl"`
	Email      string `json:"email" db:"emailbibl"`
	Instagram  string `json:"instagram" db:"instagram"`
	IsTerm     bool   `json:"isTerm" db:"srokcard"`
	Photo      []byte `json:"photo" db:"photo"`
	PhotoTitle []byte `json:"photoTitle" db:"phototitul"`
}

func (tc *TypeCard) ToProto() *protobuf.TypeCard {
	if tc == nil {
		return nil
	}
	return &protobuf.TypeCard{
		Id:         tc.Id,
		Type:       tc.Type,
		Code:       tc.Code,
		Web:        tc.Web,
		Email:      tc.Email,
		Instagram:  tc.Instagram,
		IsTerm:     tc.IsTerm,
		Photo:      tc.Photo,
		PhotoTitle: tc.PhotoTitle,
	}
}

func TypeCardFromProto(tc *protobuf.TypeCard) *TypeCard {
	if tc == nil {
		return nil
	}
	return &TypeCard{
		Id:         tc.Id,
		Type:       tc.Type,
		Code:       tc.Code,
		Web:        tc.Web,
		Email:      tc.Email,
		Instagram:  tc.Instagram,
		IsTerm:     tc.IsTerm,
		Photo:      tc.Photo,
		PhotoTitle: tc.PhotoTitle,
	}
}
func TypeCardListFromProto(l []*protobuf.TypeCard) []*TypeCard {
	list := make([]*TypeCard, len(l))
	for i, d := range l {
		list[i] = TypeCardFromProto(d)
	}
	return list
}
