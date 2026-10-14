package client

import (
	"context"

	"github.com/web-rabis/reader-client/model"
)

type Base interface {
	Connect() error
	Close() error
	ReaderSvc() ReaderService
	DictionarySvc() DictionaryService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name ReaderService
type ReaderService interface {
	ReaderByTicketNumber(ctx context.Context, ticketNumber int64) (*model.Reader, error)
	ReaderList(ctx context.Context, filters *model.ReaderFilter, paging *model.Paging) (int64, []*model.Reader, error)
	ReaderRegistration(ctx context.Context, readerRegistrationRequest *model.ReaderRegistrationRequest) (int64, string, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name DictionaryService
type DictionaryService interface {
	AcademicDegreeList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error)
	EducationList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error)
	FacultyList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error)
	InstitutionList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error)
	NationalityList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error)
	SexList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error)
	SocialStatusList(ctx context.Context, paging *model.Paging) (int64, []*model.SocialStatus, error)
	SpecialityList(ctx context.Context, paging *model.Paging) (int64, []*model.Dictionary, error)
	TypeCardList(ctx context.Context, paging *model.Paging) (int64, []*model.TypeCard, error)
}
