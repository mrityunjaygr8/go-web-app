package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/mrityunjaygr8/go-oink/server"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	// logger.Logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	c, err := getConfig(logger)
	if err != nil {
		logger.Fatal().Err(err)
	}

	if c.Env == envDevelopment {
		logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	logger.Info().Any("config", c).Msg("")

	srvConf := server.ServerConf{
		Addr: c.SrvAddr,
		Port: c.SrvPort,
	}

	if c.DbDsn == "" && c.DbHost == "" {
		logger.Fatal().Any("config", c).Msg("DB configuration not found. Either specify the DSN or the individual components.")
	}

	if c.DbDsn == "" {
		c.DbDsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.DbUser, c.DbPass, c.DbHost, c.DbPort, c.DbName, c.DbSSL)
	}
	db, err := sql.Open("postgres", c.DbDsn)
	if err != nil {
		// logger.WithFields(c.toFields()).Fatal(err)
		logger.Fatal().Any("config", c).Err(err)
	}

	err = db.Ping()
	if err != nil {
		// logger.WithFields(c.toFields()).Fatal(err)
		logger.Fatal().Any("config", c).Err(err)
	}

	a := server.New(logger, db, srvConf)
	a.Serve()
}
