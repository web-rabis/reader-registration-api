package config

import (
	"errors"
	"log"
	"os"

	"github.com/jessevdk/go-flags"
)

type ServerConfig struct {
	ListenAddr string `short:"l" long:"listen" env:"LISTEN" description:"Listen Address (format: :8080|127.0.0.1:8080)" required:"false" default:":9191"`
	BasePath   string `long:"base-path" env:"BASE_PATH" description:"base path of the host" required:"false" default:"/reader-registration-api"`
	FilesDir   string `long:"files-directory" env:"FILES_DIR" description:"Directory where all static files are located" required:"false" default:"/usr/share/api"`
	CertFile   string `short:"c" long:"cert" env:"CERT_FILE" description:"Location of the SSL/TLS cert file" required:"false" default:""`
	KeyFile    string `short:"k" long:"key" env:"KEY_FILE" description:"Location of the SSL/TLS key file" required:"false" default:""`

	GrpcListenAddr string `long:"grpc-listen" env:"GRPC_LISTEN" description:"Grpc Listen Address (format: :4000|127.0.0.1:4000)" required:"false" default:":4000"`
}
type AuthConfig struct {
	JWTKey          string `long:"jwt-key" env:"JWT_KEY" description:"JWT secret key" required:"false" default:"nlrk-secret"`
	AccessTokenTTL  int    `long:"access-token-ttl" env:"ACCESS_TOKEN_TTL" required:"false" default:"60"`   // Access token expiration in minutes
	RefreshTokenTTL int    `long:"refresh-token-ttl" env:"REFRESH_TOKEN_TTL" required:"false" default:"24"` // Refresh token expiration in hours
}
type GeneralConfig struct {
	Dbg       bool `long:"dbg" env:"DEBUG" description:"debug mode"`
	IsTesting bool `long:"testing" env:"APP_TESTING" description:"testing mode"`
}
type ReaderConfig struct {
	GrpcAddress string `long:"reader-grpc-address" env:"READER_GRPC_ADDRESS" description:"Reader Grpc Address (format: :4000|127.0.0.1:4000)" required:"false" default:":4000"`
	TypeCard    int64  `long:"reader-type-card" env:"READER_TYPE_CARD" description:"Reader type card" required:"false" default:"130"`
}

func Parse(c interface{}) interface{} {
	p := flags.NewParser(c, flags.Default)

	if _, err := p.Parse(); err != nil {
		log.Println("[ERROR] Error while parsing config options:", err)
		var flagsErr *flags.Error
		if errors.As(err, &flagsErr) && errors.Is(flagsErr.Type, flags.ErrHelp) {
			os.Exit(0)
		}
	}

	return c
}
