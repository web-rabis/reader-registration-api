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
	response, err := s.client.ReaderByTicketNumber(ctx, &protobuf.ByTicketNumberRequest{TicketNumber: ticketNumber})
	switch status.Code(err) {
	case codes.OK:
		return model.ReaderFromProto(response), nil
	default:
		return nil, err
	}
}
func (s *ReaderService) ReaderByGlobalTicketNumber(ctx context.Context, globalTicketNumber int64) (*model.Reader, error) {
	response, err := s.client.ReaderByGlobalTicketNumber(ctx, &protobuf.ByGlobalTicketNumberRequest{GlobalTicketNumber: globalTicketNumber})
	switch status.Code(err) {
	case codes.OK:
		return model.ReaderFromProto(response), nil
	default:
		return nil, err
	}
}
func (s *ReaderService) ReaderByBarcode(ctx context.Context, barcode string) (*model.Reader, error) {
	response, err := s.client.ReaderByBarcode(ctx, &protobuf.ByBarcodeRequest{Barcode: barcode})
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
func (s *ReaderService) ReaderCreate(ctx context.Context, reader *model.Reader, userId int64) (int64, error) {
	request := &protobuf.ReaderCreateRequest{
		Reader: reader.ToProto(),
		UserId: userId,
	}
	response, err := s.client.ReaderCreate(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.TicketNumber, nil
	default:
		return 0, err
	}
}
func (s *ReaderService) ReaderUpdate(ctx context.Context, reader *model.Reader, userId int64) error {
	request := &protobuf.ReaderUpdateRequest{
		Reader: reader.ToProto(),
		UserId: userId,
	}
	_, err := s.client.ReaderUpdate(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}
func (s *ReaderService) ReaderDelete(ctx context.Context, ticketNumber, userId int64) error {
	request := &protobuf.ReaderDeleteRequest{
		TicketNumber: ticketNumber,
		UserId:       userId,
	}
	_, err := s.client.ReaderDelete(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}
func (s *ReaderService) ReaderApprove(ctx context.Context, ticketNumber, userId int64) (*model.ReaderApproveResult, error) {
	request := &protobuf.ReaderApproveRequest{
		TicketNumber: ticketNumber,
		UserId:       userId,
	}
	response, err := s.client.ReaderApprove(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return model.ReaderApproveResultFromProto(response), nil
	default:
		return nil, err
	}
}
func (s *ReaderService) ReaderReject(ctx context.Context, ticketNumber, userId int64) error {
	request := &protobuf.ReaderRejectRequest{
		TicketNumber: ticketNumber,
		UserId:       userId,
	}
	_, err := s.client.ReaderReject(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}
func (s *ReaderService) ReaderRenewalRegistration(ctx context.Context, ticketNumber int64, years, months, days int, userId int64) error {
	request := &protobuf.ReaderRenewalRegistrationRequest{
		TicketNumber: ticketNumber,
		Years:        int32(years),
		Months:       int32(months),
		Days:         int32(days),
		UserId:       userId,
	}
	_, err := s.client.ReaderRenewalRegistration(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}
