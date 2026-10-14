package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/web-rabis/reader-client/client"
	"github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-client/protobuf"
)

type DictionaryService struct {
	client protobuf.DictionarySvcClient
}

var _ client.DictionaryService = &DictionaryService{}

func NewDictionaryServiceClient(client protobuf.DictionarySvcClient) client.DictionaryService {
	return &DictionaryService{
		client: client,
	}
}
func (d DictionaryService) AcademicDegreeList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error) {
	response, err := d.client.AcademicDegreeList(ctx, dictionaryListRequest(paging))
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.DictionaryListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

func (d DictionaryService) EducationList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error) {
	response, err := d.client.EducationList(ctx, dictionaryListRequest(paging))
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.DictionaryListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

func (d DictionaryService) FacultyList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error) {
	response, err := d.client.FacultyList(ctx, dictionaryListRequest(paging))
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.DictionaryListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

func (d DictionaryService) InstitutionList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error) {
	response, err := d.client.InstitutionList(ctx, dictionaryListRequest(paging))
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.DictionaryListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

func (d DictionaryService) NationalityList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error) {
	response, err := d.client.NationalityList(ctx, dictionaryListRequest(paging))
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.DictionaryListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

func (d DictionaryService) SexList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error) {
	response, err := d.client.SexList(ctx, dictionaryListRequest(paging))
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.DictionaryListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

func (d DictionaryService) SocialStatusList(ctx context.Context, paging *model.Paging) (int64, []*model.SocialStatus, error) {
	request := &protobuf.SocialStatusListRequest{Paging: paging.ToProto()}
	response, err := d.client.SocialStatusList(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.SocialStatusListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

func (d DictionaryService) SpecialityList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error) {
	response, err := d.client.SexList(ctx, dictionaryListRequest(paging))
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.DictionaryListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

func (d DictionaryService) TypeCardList(ctx context.Context, paging *model.Paging) (int64, []*model.TypeCard, error) {
	request := &protobuf.TypeCardListRequest{Paging: paging.ToProto()}
	response, err := d.client.TypeCardList(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.TypeCardListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}

func dictionaryListRequest(paging *model.Paging) *protobuf.DictionaryListRequest {
	return &protobuf.DictionaryListRequest{Paging: paging.ToProto()}
}
