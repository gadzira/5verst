package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log" //nolint: gci
)

//nolint: gci

const (
	envHost    = "APP_HOST"
	envPort    = "APP_PORT"
	envDBHost  = "APP_DB_HOST"
	envDBPort  = "APP_DB_PORT"
	envDBUser  = "APP_DB_USER"
	envDBPass  = "APP_DB_PASS"
	envDBName  = "APP_DB_NAME"
	envLogFile = "APP_LOG_FILE"
)

type Config struct {
	Host    string
	Port    string
	DBDsn   string
	LogFile string
}

func MustConfig() Config {
	conf := Config{}

	err := conf.Init()
	if err != nil {
		log.Panic().Msgf("Failed init config: %v", err)
	}

	return conf
}

func (c *Config) Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Panic().Msg(".env file not found")
		return err
	}

	c.Host = os.Getenv(envHost)
	c.Port = os.Getenv(envPort)
	c.DBDsn = fmt.Sprintf(
		`%s:%s@tcp(%s:%s)/%s?parseTime=true`,
		os.Getenv(envDBUser),
		os.Getenv(envDBPass),
		os.Getenv(envDBHost),
		os.Getenv(envDBPort),
		os.Getenv(envDBName),
	)
	c.LogFile = os.Getenv(envLogFile)

	return nil
}
