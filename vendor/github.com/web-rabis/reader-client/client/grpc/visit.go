package grpc

import (
	"context"

	"github.com/web-rabis/reader-client/client"
	"github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-client/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type VisitService struct {
	client protobuf.VisitSvcClient
}

func NewVisitServiceClient(client protobuf.VisitSvcClient) client.VisitService {
	return &VisitService{
		client: client,
	}
}
func (c VisitService) VisitCheckIn(ctx context.Context, ticketNumber int64, user *model.User) error {
	request := &protobuf.VisitCheckInRequest{
		TicketNumber: ticketNumber,
		User:         user.ToProto(),
	}
	_, err := c.client.VisitCheckIn(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}

func (c VisitService) VisitCheckOut(ctx context.Context, ticketNumber int64, user *model.User) error {
	request := &protobuf.VisitCheckOutRequest{
		TicketNumber: ticketNumber,
		User:         user.ToProto(),
	}
	_, err := c.client.VisitCheckOut(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}

func (c VisitService) VisitList(ctx context.Context, filters *model.VisitFilter, paging *model.Paging) (int64, []*model.Visit, error) {
	request := &protobuf.VisitListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	response, err := c.client.VisitList(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.VisitListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

var _ client.VisitService = &VisitService{}
