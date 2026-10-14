package reader

import (
	"context"

	readerclient "github.com/web-rabis/reader-client/client"
	rcmodel "github.com/web-rabis/reader-client/model"
)

type IManager interface {
	List(ctx context.Context, filter *rcmodel.ReaderFilter, paging *rcmodel.Paging) (int64, []*rcmodel.Reader, error)
	ReaderRegistration(ctx context.Context, readerRegistrationRequest *rcmodel.ReaderRegistrationRequest) (int64, string, error)
}

type Manager struct {
	readerCli readerclient.ReaderService
}

func NewManager(readerCli readerclient.ReaderService) *Manager {
	return &Manager{
		readerCli: readerCli,
	}
}

func (m *Manager) List(ctx context.Context, filter *rcmodel.ReaderFilter, paging *rcmodel.Paging) (int64, []*rcmodel.Reader, error) {
	return m.readerCli.ReaderList(ctx, filter, paging)
}

func (m *Manager) ReaderRegistration(ctx context.Context, readerRegistrationRequest *rcmodel.ReaderRegistrationRequest) (int64, string, error) {
	return m.readerCli.ReaderRegistration(ctx, readerRegistrationRequest)
}
