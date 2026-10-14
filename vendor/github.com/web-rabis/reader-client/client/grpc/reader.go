package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/web-rabis/reader-client/client"
	"github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-client/protobuf"
)

type ReaderService struct {
	client protobuf.ReaderSvcClient
}

var _ client.ReaderService = &ReaderService{}

func NewReaderServiceClient(client protobuf.ReaderSvcClient) client.ReaderService {
	return &ReaderService{
		client: client,
	}
}

func (s *ReaderService) ReaderByTicketNumber(ctx context.Context, ticketNumber int64) (*model.Reader, error) {
	response, err := s.client.ReaderByTicketNumber(ctx, &protobuf.ReaderByTicketNumberRequest{TicketNumber: ticketNumber})
	switch status.Code(err) {
	case codes.OK:
		return model.ReaderFromProto(response), nil
	default:
		return nil, err
	}
}
func (s *ReaderService) ReaderList(ctx context.Context, filters *model.ReaderFilter, paging *model.Paging) (int64, []*model.Reader, error) {
	request := &protobuf.ReaderListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	response, err := s.client.ReaderList(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.ReaderListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}
func (s *ReaderService) ReaderRegistration(ctx context.Context, readerRegistrationRequest *model.ReaderRegistrationRequest) (int64, string, error) {
	request := readerRegistrationRequest.ToProto()
	response, err := s.client.ReaderRegistration(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.TicketNumber, response.Status, nil
	default:
		return 0, "", err
	}

}
