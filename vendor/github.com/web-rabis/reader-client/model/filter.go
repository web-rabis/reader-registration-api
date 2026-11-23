package model

import (
	"time"

	"github.com/web-rabis/reader-client/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ReaderFilter struct {
	Query              string
	Iin                string
	ExpirationDateLow  *time.Time
	ExpirationDateHigh *time.Time
	IsEmployee         *bool
	// "" - все, кроме "deleted"; конкретное значение ("pending" и т.д.) -
	// сузить до него одного.
	Status string
}

// VisitFilter - названия полей совпадают с internal/model/visit.Filter
// в reader-сервисе (CheckedInAt.../CheckedOutAt.../OpenOnly); TicketNumber -
// локальный номер читателя (как и везде на wire), резолвится в reader_id
// на стороне reader-resource (см. VisitList в reader).
type VisitFilter struct {
	CheckedInAtLow   *time.Time
	CheckedInAtHigh  *time.Time
	CheckedOutAtLow  *time.Time
	CheckedOutAtHigh *time.Time
	OpenOnly         *bool
	TicketNumber     *int64
}

func (f *ReaderFilter) ToProto() *protobuf.ReaderFilter {
	if f == nil {
		return nil
	}
	var expirationDateLow *timestamppb.Timestamp
	var expirationDateHigh *timestamppb.Timestamp
	if f.ExpirationDateLow != nil {
		expirationDateLow = timestamppb.New(*f.ExpirationDateLow)
	}
	if f.ExpirationDateHigh != nil {
		expirationDateHigh = timestamppb.New(*f.ExpirationDateHigh)
	}
	return &protobuf.ReaderFilter{
		Query:              f.Query,
		Iin:                f.Iin,
		ExpirationDateLow:  expirationDateLow,
		ExpirationDateHigh: expirationDateHigh,
		IsEmployee:         f.IsEmployee,
		Status:             f.Status,
	}
}

func (f *VisitFilter) ToProto() *protobuf.VisitFilter {
	if f == nil {
		return nil
	}
	var checkedInAtLow *timestamppb.Timestamp
	if f.CheckedInAtLow != nil {
		checkedInAtLow = timestamppb.New(*f.CheckedInAtLow)
	}
	var checkedInAtHigh *timestamppb.Timestamp
	if f.CheckedInAtHigh != nil {
		checkedInAtHigh = timestamppb.New(*f.CheckedInAtHigh)
	}
	var checkedOutAtLow *timestamppb.Timestamp
	if f.CheckedOutAtLow != nil {
		checkedOutAtLow = timestamppb.New(*f.CheckedOutAtLow)
	}
	var checkedOutAtHigh *timestamppb.Timestamp
	if f.CheckedOutAtHigh != nil {
		checkedOutAtHigh = timestamppb.New(*f.CheckedOutAtHigh)
	}
	return &protobuf.VisitFilter{
		CheckedInAtLow:   checkedInAtLow,
		CheckedInAtHigh:  checkedInAtHigh,
		CheckedOutAtLow:  checkedOutAtLow,
		CheckedOutAtHigh: checkedOutAtHigh,
		OpenOnly:         f.OpenOnly,
		TicketNumber:     f.TicketNumber,
	}
}
