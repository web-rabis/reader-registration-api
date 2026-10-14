package model

import (
	"net/http"
	"strconv"

	"github.com/web-rabis/reader-client/protobuf"
)

type Paging struct {
	Offset  int64  `json:"offset"`
	Limit   int64  `json:"limit"`
	SortKey string `json:"sort_key"`
	SortVal int32  `json:"sort_val"`
}

func PagingParseFromHttp(r *http.Request) (*Paging, error) {
	paging := Paging{}

	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		limit, err := strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			return nil, err
		}
		paging.Limit = limit
	}

	if pageParam := r.URL.Query().Get("offset"); pageParam != "" {
		offset, err := strconv.ParseInt(pageParam, 10, 64)
		if err != nil {
			return nil, err
		}
		paging.Offset = offset
	}

	if orderParam := r.URL.Query().Get("order"); orderParam != "" {
		paging.SortVal = 1
		order, err := strconv.Atoi(orderParam)
		if err == nil {
			paging.SortVal = int32(order)
		}
	}

	if sortKeyFromQuery := r.URL.Query().Get("orderBy"); sortKeyFromQuery != "" {
		paging.SortKey = sortKeyFromQuery
	}

	return &paging, nil
}

func (p *Paging) ToProto() *protobuf.Paging {
	if p == nil {
		return nil
	}
	return &protobuf.Paging{
		Offset:  p.Offset,
		Limit:   p.Limit,
		SortKey: p.SortKey,
		SortVal: p.SortVal,
	}
}
