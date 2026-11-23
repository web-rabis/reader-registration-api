package model

import (
	"time"

	"github.com/web-rabis/reader-client/protobuf"
)

// Visit - поля совпадают с internal/model/visit.Visit в reader-сервисе
// (см. mapper.VisitToProto) - без легаси-полей старой control (zakazy,
// f51, count).
type Visit struct {
	Id           int64      `json:"id,omitempty"`
	ReaderId     int64      `json:"readerId,omitempty"`
	LibraryId    int64      `json:"libraryId,omitempty"`
	BranchId     int64      `json:"branchId,omitempty"`
	DepartmentId *int64     `json:"departmentId,omitempty"`
	CheckedInAt  *time.Time `json:"checkedInAt,omitempty"`
	CheckedOutAt *time.Time `json:"checkedOutAt,omitempty"`
	CheckedInBy  int64      `json:"checkedInBy,omitempty"`
	CheckedOutBy int64      `json:"checkedOutBy,omitempty"`
}

func VisitFromProto(v *protobuf.Visit) *Visit {
	if v == nil {
		return nil
	}
	var checkedInAt *time.Time
	if v.CheckedInAt != nil {
		t := v.CheckedInAt.AsTime()
		checkedInAt = &t
	}
	var checkedOutAt *time.Time
	if v.CheckedOutAt != nil {
		t := v.CheckedOutAt.AsTime()
		checkedOutAt = &t
	}
	return &Visit{
		Id:           v.Id,
		ReaderId:     v.ReaderId,
		LibraryId:    v.LibraryId,
		BranchId:     v.BranchId,
		DepartmentId: v.DepartmentId,
		CheckedInAt:  checkedInAt,
		CheckedOutAt: checkedOutAt,
		CheckedInBy:  v.CheckedInBy,
		CheckedOutBy: v.CheckedOutBy,
	}
}

func VisitListFromProto(l []*protobuf.Visit) []*Visit {
	list := make([]*Visit, len(l))
	for i, d := range l {
		list[i] = VisitFromProto(d)
	}
	return list
}
