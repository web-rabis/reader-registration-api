package v1

import (
	"github.com/go-chi/chi/v5"

	"github.com/web-rabis/reader-registration-api/internal/manager/reader"
)

type Resource struct {
	path      string
	readerMan reader.IManager
}

func NewReaderResource(basePath string, readerMan reader.IManager) *Resource {
	return &Resource{
		path:      basePath,
		readerMan: readerMan,
	}
}

func (res *Resource) Path() string {
	return res.path
}

func (res *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/registration", res.registrationReader)
	})

	return r
}
