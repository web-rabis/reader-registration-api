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
	VisitSvc() VisitService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name ReaderService
type ReaderService interface {
	ReaderByTicketNumber(ctx context.Context, ticketNumber int64) (*model.Reader, error)
	// ReaderByGlobalTicketNumber - адресация по внутреннему суррогатному PK
	// (reader.GlobalTicketNumber), а не по локальному номеру в рамках
	// библиотеки; нужен, когда идентификатор читателя известен только из
	// visit (Visit.ReaderId - это как раз global_ticket_num).
	ReaderByGlobalTicketNumber(ctx context.Context, globalTicketNumber int64) (*model.Reader, error)
	ReaderByBarcode(ctx context.Context, barcode string) (*model.Reader, error)
	ReaderList(ctx context.Context, filters *model.ReaderFilter, paging *model.Paging) (int64, []*model.Reader, error)
	ReaderRegistration(ctx context.Context, readerRegistrationRequest *model.ReaderRegistrationRequest) (int64, string, error)
	// userId - id сотрудника, выполняющего действие (для reader_audit в
	// reader-service); 0, если действие без участия сотрудника.
	ReaderCreate(ctx context.Context, reader *model.Reader, userId int64) (int64, error)
	ReaderUpdate(ctx context.Context, reader *model.Reader, userId int64) error
	ReaderDelete(ctx context.Context, ticketNumber, userId int64) error

	ReaderApprove(ctx context.Context, ticketNumber, userId int64) (*model.ReaderApproveResult, error)
	ReaderReject(ctx context.Context, ticketNumber, userId int64) error
	ReaderRenewalRegistration(ctx context.Context, ticketNumber int64, years, months, days int, userId int64) error
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
	TypeCardById(ctx context.Context, id int64) (*model.TypeCard, error)
}

// VisitCheckIn/VisitCheckOut больше не принимают department строкой -
// branch_id/department_id передаются через tenant-контекст grpc-metadata
// (см. WithTenant в корне модуля), а не в теле запроса.
type VisitService interface {
	VisitCheckIn(ctx context.Context, ticketNumber int64, user *model.User) error
	VisitCheckOut(ctx context.Context, ticketNumber int64, user *model.User) error
	VisitList(ctx context.Context, filters *model.VisitFilter, paging *model.Paging) (int64, []*model.Visit, error)
}
