package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// type config struct {
// 	DB     server.DbConf
// 	Server server.ServerConf
// }

type config struct {
	DbDsn   string `mapstructure:"PSQL_DSN" json:"PSQL_DSN"`
	DbPort  int    `mapstructure:"PSQL_PORT" json:"PSQL_PORT"`
	DbHost  string `mapstructure:"PSQL_HOST" json:"PSQL_HOST"`
	DbName  string `mapstructure:"PSQL_DBNAME" json:"PSQL_DBNAME"`
	DbPass  string `mapstructure:"PSQL_PASS" json:"PSQL_PASS"`
	DbUser  string `mapstructure:"PSQL_USER" json:"PSQL_USER"`
	DbSSL   string `mapstructure:"PSQL_SSLMODE" json:"PSQL_SSLMODE"`
	SrvAddr string `mapstructure:"SERVER_ADDR" json:"SERVER_ADDR"`
	SrvPort int    `mapstructure:"SERVER_PORT" json:"SERVER_PORT"`
}

func (c config) toFields() logrus.Fields {
	fields := make(logrus.Fields)

	fields["DbDsn"] = c.DbDsn
	fields["DbHost"] = c.DbHost
	fields["DbPort"] = c.DbPort
	fields["DbPass"] = c.DbPass
	fields["DbUser"] = c.DbUser
	fields["DbSSL"] = c.DbSSL
	fields["DbName"] = c.DbName
	fields["SrvAddr"] = c.SrvAddr
	fields["SrvPort"] = c.SrvPort

	return fields
}

func getConfig(logger *logrus.Logger) (config, error) {
	var c config

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	// viper.AddConfigPath(".")

	viper.SetEnvPrefix("vlcm")

	viper.BindEnv("SERVER_ADDR", "SERVER_ADDR")
	viper.BindEnv("SERVER_PORT", "SERVER_PORT")
	viper.BindEnv("PSQL_HOST", "PSQL_HOST")
	viper.BindEnv("PSQL_PORT", "PSQL_PORT")
	viper.BindEnv("PSQL_DBNAME", "PSQL_DBNAME")
	viper.BindEnv("PSQL_PASS", "PSQL_PASS")
	viper.BindEnv("PSQL_USER", "PSQL_USER")
	viper.BindEnv("PSQL_DSN", "PSQL_DSN")
	viper.BindEnv("PSQL_SSLMODE", "PSQL_SSLMODE")

	err := viper.ReadInConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			return config{}, err
		} else {
			logger.Errorln(err.Error())
		}
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return config{}, err
	}
	return c, nil
}
