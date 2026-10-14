package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/logutils"
	"golang.org/x/sync/errgroup"

	readerclient "github.com/web-rabis/reader-client"
	readerclientmodel "github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-registration-api/internal/config"
	"github.com/web-rabis/reader-registration-api/internal/manager/dictionary"
	"github.com/web-rabis/reader-registration-api/internal/manager/reader"
	"github.com/web-rabis/reader-registration-api/internal/server/http"
)

var (
	version = "unknown"
)

// @title Reader Registration API
// @version 1.0
// @description Pay System Admin API
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fmt.Printf("reader-registration-api %s\n", version)

	opts := config.Parse(new(config.APIServer)).(*config.APIServer)

	setupLogUtils(opts.Dbg)

	appCtx, appCtxCancel := context.WithCancel(context.Background())
	defer appCtxCancel()

	go catchForTermination(appCtxCancel, os.Interrupt, syscall.SIGTERM)

	readerCli, err := readerclient.NewReaderClient(&readerclientmodel.ConnectionConfig{
		Address:  opts.ReaderConfig.GrpcAddress,
		Protocol: "grpc",
		Insecure: true,
	})
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := readerCli.Close(); err != nil {
			log.Printf("[ERROR] failed to close reader client: %v", err)
		}
	}()
	log.Printf("[INFO] connected to reader grpc-server(%s)", opts.ReaderConfig.GrpcAddress)

	dictMan := dictionary.NewCachingManager(dictionary.NewManager(readerCli.DictionarySvc()))
	readerMan := reader.NewManager(readerCli.ReaderSvc())
	servers, serversCtx := errgroup.WithContext(appCtx)

	servers.Go(func() error {
		return http.Run(serversCtx, opts, dictMan, readerMan, version)
	})

	if err := servers.Wait(); err != nil {
		log.Printf("[INFO] process terminated, %s", err)
		return
	}

}

func setupLogUtils(inDebugMode bool) {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "INFO", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("INFO"),
		Writer:   os.Stdout,
	}

	log.SetFlags(log.Ldate | log.Ltime)

	if inDebugMode {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
		filter.MinLevel = "DEBUG"
	}

	log.SetOutput(filter)
}

func catchForTermination(cancel context.CancelFunc, signals ...os.Signal) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, signals...)
	<-stop
	log.Print("[WARN] interrupt signal")
	cancel()
}
