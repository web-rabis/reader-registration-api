package reader

import (
	"context"

	"github.com/google/uuid"

	readerclient "github.com/web-rabis/reader-client/client"
	rcmodel "github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-registration-api/internal/storage"
)

const photoContentType = "image/jpeg"

type IManager interface {
	List(ctx context.Context, filter *rcmodel.ReaderFilter, paging *rcmodel.Paging) (int64, []*rcmodel.Reader, error)
	ReaderRegistration(ctx context.Context, readerRegistrationRequest *rcmodel.ReaderRegistrationRequest, photo []byte) (int64, string, error)
}

type Manager struct {
	readerCli readerclient.ReaderService
	storage   storage.Store
}

func NewManager(readerCli readerclient.ReaderService, store storage.Store) *Manager {
	return &Manager{
		readerCli: readerCli,
		storage:   store,
	}
}

func (m *Manager) List(ctx context.Context, filter *rcmodel.ReaderFilter, paging *rcmodel.Paging) (int64, []*rcmodel.Reader, error) {
	return m.readerCli.ReaderList(ctx, filter, paging)
}

// photo - уже обработанные (resize) байты фото заявки; reader/reader-client
// хранят только PhotoKey (ссылку на объект в MinIO/S3), поэтому байты
// загружаются сюда напрямую, в обход gRPC.
func (m *Manager) ReaderRegistration(ctx context.Context, readerRegistrationRequest *rcmodel.ReaderRegistrationRequest, photo []byte) (int64, string, error) {
	if len(photo) > 0 {
		key := uuid.NewString() + ".jpg"
		if err := m.storage.Upload(ctx, key, photo, photoContentType); err != nil {
			return 0, "", err
		}
		readerRegistrationRequest.PhotoKey = key
	}
	return m.readerCli.ReaderRegistration(ctx, readerRegistrationRequest)
}
