package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/mrityunjaygr8/vlcm-go/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := &logrus.Logger{
		Out:       os.Stdout,
		Hooks:     make(logrus.LevelHooks),
		Formatter: new(logrus.JSONFormatter),
		Level:     logrus.DebugLevel,
	}

	c, err := getConfig(logger)
	if err != nil {
		logger.Fatal(err)
	}

	srvConf := server.ServerConf{
		Addr: c.SrvAddr,
		Port: c.SrvPort,
	}

	if c.DbDsn == "" && c.DbHost == "" {
		logger.WithFields(c.toFields()).Fatal("DB configuration not found. Either specify the DSN or the individual components.")
	}

	if c.DbDsn == "" {
		c.DbDsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.DbUser, c.DbPass, c.DbHost, c.DbPort, c.DbName, c.DbSSL)
	}
	db, err := sql.Open("postgres", c.DbDsn)
	if err != nil {
		logger.WithFields(c.toFields()).Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		logger.WithFields(c.toFields()).Fatal(err)
	}

	a := server.New(logger, db, srvConf)
	a.Serve()
}
