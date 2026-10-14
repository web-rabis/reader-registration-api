package http

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/web-rabis/reader-registration-api/internal/config"
	"github.com/web-rabis/reader-registration-api/internal/manager/dictionary"
	"github.com/web-rabis/reader-registration-api/internal/manager/reader"
	"github.com/web-rabis/reader-registration-api/internal/resource/http"
	dictionaryV1 "github.com/web-rabis/reader-registration-api/internal/resource/http/dictionary/v1"
	readerV1 "github.com/web-rabis/reader-registration-api/internal/resource/http/reader/v1"
	swaggerV1 "github.com/web-rabis/reader-registration-api/internal/resource/http/swagger/v1"
	cherver "github.com/web-rabis/servers/http"
)

const (
	maxAge        = 300
	compressLevel = 5
)

func Run(serversCtx context.Context, opts *config.APIServer,
	dictMan dictionary.IManager,
	readerMan reader.IManager,
	version string) error {
	resources := []cherver.Resource{
		http.NewVersionResource("/version", version),
		http.NewHealthResource("/health", readerMan),

		http.NewFilesResource("/files", opts.ServerConfig.FilesDir),
		swaggerV1.NewSwaggerResource("/swagger", opts.BasePath, "/files"),

		dictionaryV1.NewDictionaryResource("/api/v1/dictionary", dictMan),
		readerV1.NewReaderResource("/api/v1/reader", opts.ReaderConfig.TypeCard, readerMan),
	}
	httpSrv := cherver.New(
		cherver.WithListenAddress(opts.ServerConfig.ListenAddr),
		cherver.WithCert(opts.ServerConfig.CertFile, opts.ServerConfig.KeyFile),
		cherver.WithResources(resources...),
		cherver.WithMiddlewares(middlewares(opts)...))

	return httpSrv.Run(serversCtx)

}

func middlewaresWithoutLogs(opts *config.APIServer) chi.Middlewares {
	return chi.Middlewares{
		middleware.NoCache,   // no-cache
		middleware.Recoverer, // управляемо обрабатывает паники и выдает stack trace при их возникновении
		middleware.RealIP,    // устанавливает RemoteAddr для каждого запроса с заголовками X-Forwarded-For или X-Real-IP
		middleware.NewCompressor(compressLevel).Handler,

		cors.Handler(cors.Options{
			AllowedOrigins:   allowedOrigins(opts.IsTesting),
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           maxAge, // Maximum value not ignored by any of major browsers
		})}
}

func middlewares(opts *config.APIServer) chi.Middlewares {
	return append(middlewaresWithoutLogs(opts), middleware.Logger)
}
func allowedOrigins(testing bool) []string {
	if testing {
		return []string{"*"}
	}

	return []string{}
}
