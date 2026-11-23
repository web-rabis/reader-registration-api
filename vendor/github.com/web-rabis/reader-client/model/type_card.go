package model

import "github.com/web-rabis/reader-client/protobuf"

// TermYears - срок действия билета в годах, nil = бессрочный. PhotoKey/
// TitlePhotoKey - ключи в MinIO (раньше фото лежало в БД как bytea).
type TypeCard struct {
	Id            int64  `json:"id" db:"id"`
	LibraryId     int64  `json:"libraryId" db:"library_id"`
	BranchId      *int64 `json:"branchId" db:"branch_id"`
	Type          string `json:"type" db:"type"`
	Code          string `json:"code" db:"code"`
	Website       string `json:"website" db:"website"`
	Email         string `json:"email" db:"email"`
	Instagram     string `json:"instagramUrl" db:"instagram_url"`
	TermYears     *int32 `json:"termYears" db:"term_years"`
	PhotoKey      string `json:"photoKey" db:"photo_key"`
	TitlePhotoKey string `json:"titlePhotoKey" db:"title_photo_key"`
}

func (tc *TypeCard) ToProto() *protobuf.TypeCard {
	if tc == nil {
		return nil
	}
	return &protobuf.TypeCard{
		Id:            tc.Id,
		LibraryId:     tc.LibraryId,
		BranchId:      tc.BranchId,
		Type:          tc.Type,
		Code:          tc.Code,
		Website:       tc.Website,
		Email:         tc.Email,
		InstagramUrl:  tc.Instagram,
		TermYears:     tc.TermYears,
		PhotoKey:      tc.PhotoKey,
		TitlePhotoKey: tc.TitlePhotoKey,
	}
}

func TypeCardFromProto(tc *protobuf.TypeCard) *TypeCard {
	if tc == nil {
		return nil
	}
	return &TypeCard{
		Id:            tc.Id,
		LibraryId:     tc.LibraryId,
		BranchId:      tc.BranchId,
		Type:          tc.Type,
		Code:          tc.Code,
		Website:       tc.Website,
		Email:         tc.Email,
		Instagram:     tc.InstagramUrl,
		TermYears:     tc.TermYears,
		PhotoKey:      tc.PhotoKey,
		TitlePhotoKey: tc.TitlePhotoKey,
	}
}
func TypeCardListFromProto(l []*protobuf.TypeCard) []*TypeCard {
	list := make([]*TypeCard, len(l))
	for i, d := range l {
		list[i] = TypeCardFromProto(d)
	}
	return list
}
