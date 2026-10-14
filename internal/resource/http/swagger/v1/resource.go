package v1

import (
	"path/filepath"

	"github.com/go-chi/chi/v5"

	httpSwagger "github.com/swaggo/http-swagger"
)

type SwaggerResource struct {
	path      string
	basePath  string
	filesPath string
}

func NewSwaggerResource(path, basePath, filesPath string) SwaggerResource {
	return SwaggerResource{
		path:      path,
		basePath:  basePath,
		filesPath: filesPath,
	}
}

func (sr SwaggerResource) Path() string {
	return sr.path
}

func (sr SwaggerResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/*", httpSwagger.Handler(
		httpSwagger.URL(filepath.Join(sr.basePath, sr.filesPath, "swagger.yaml")),
	))
	return r
}
