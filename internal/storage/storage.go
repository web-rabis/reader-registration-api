package storage

import "context"

// Store - хранилище бинарных объектов (MinIO/S3) для фото читателей.
// reader-client/reader хранят только PhotoKey (ссылку на объект), сами
// байты через gRPC не передаются - их загружает caller (registration-api,
// workstation-api) напрямую в Store.
type Store interface {
	Upload(ctx context.Context, key string, data []byte, contentType string) error
	Download(ctx context.Context, key string) ([]byte, error)
}
