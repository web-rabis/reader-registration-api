package minio

import (
	"bytes"
	"context"
	"io"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/web-rabis/reader-registration-api/internal/storage"
)

type Client struct {
	cl     *minio.Client
	bucket string
}

func New(endpoint, accessKey, secretKey, bucket string, useSSL bool) (*Client, error) {
	cl, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return &Client{cl: cl, bucket: bucket}, nil
}

// EnsureBucket создаёт бакет, если он ещё не существует - бакет
// reader-photos обычно уже создан вручную в консоли MinIO, но это
// защищает от расхождения окружений (например, свежий staging).
func (c *Client) EnsureBucket(ctx context.Context) error {
	exists, err := c.cl.BucketExists(ctx, c.bucket)
	if err != nil {
		return err
	}
	if !exists {
		return c.cl.MakeBucket(ctx, c.bucket, minio.MakeBucketOptions{})
	}
	return nil
}

func (c *Client) Upload(ctx context.Context, key string, data []byte, contentType string) error {
	_, err := c.cl.PutObject(ctx, c.bucket, key, bytes.NewReader(data), int64(len(data)), minio.PutObjectOptions{
		ContentType: contentType,
	})
	return err
}

func (c *Client) Download(ctx context.Context, key string) ([]byte, error) {
	obj, err := c.cl.GetObject(ctx, c.bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer func() { _ = obj.Close() }()
	return io.ReadAll(obj)
}

var _ storage.Store = (*Client)(nil)
