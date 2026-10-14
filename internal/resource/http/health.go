package http

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/web-rabis/httperrors"
	rcmodel "github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-registration-api/internal/manager/reader"
)

const healthCheckTimeout = 3 * time.Second

// HealthResponse - ответ на запрос проверки работоспособности.
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthResource - проверяет, что сервис и его зависимости (reader grpc-бэкенд) работоспособны.
type HealthResource struct {
	path      string
	readerMan reader.IManager
}

func NewHealthResource(path string, readerMan reader.IManager) *HealthResource {
	return &HealthResource{path: path, readerMan: readerMan}
}

func (hr *HealthResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", hr.get)

	return r
}

func (hr *HealthResource) Path() string {
	return hr.path
}

// @Summary Health check
// @Description Check that the service and its reader grpc-backend dependency are operational
// @Accept json
// @Produce json
// @Tags Health
// @Success 200 {object} http.HealthResponse
// @Failure 503 {object} httperrors.Response
// @Router /health [get]
func (hr *HealthResource) get(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), healthCheckTimeout)
	defer cancel()

	if _, _, err := hr.readerMan.List(ctx, nil, &rcmodel.Paging{Limit: 1}); err != nil {
		_ = render.Render(w, r, httperrors.Custom(err, http.StatusServiceUnavailable))
		return
	}

	render.JSON(w, r, HealthResponse{Status: "ok"})
}
