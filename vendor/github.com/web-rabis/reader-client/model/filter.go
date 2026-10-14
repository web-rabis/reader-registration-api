package model

import (
	"time"

	"github.com/web-rabis/reader-client/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ReaderFilter struct {
	Status             string
	Query              string
	Iin                string
	ExpirationDateLow  *time.Time
	ExpirationDateHigh *time.Time
	IsEmployee         *bool
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
		Status:             f.Status,
		Query:              f.Query,
		Iin:                f.Iin,
		ExpirationDateLow:  expirationDateLow,
		ExpirationDateHigh: expirationDateHigh,
		IsEmployee:         f.IsEmployee,
	}
}
