package config

import (
	"crud/utils"
	"fmt"
)

type dbConfig struct {
	host         string
	port         int
	username     string
	databaseName string
	password     string
}

func newDBConfig() *dbConfig {
	return &dbConfig{
		host:         utils.FatalGetString("DB_HOST"),
		port:         utils.GetIntOrPanic("DB_PORT"),
		databaseName: utils.FatalGetString("DB_NAME"),
		username:     utils.GetString("DB_USER"),
		password:     utils.GetString("DB_PASSWORD"),
	}
}

func (dc *dbConfig) GetConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' sslmode=disable", dc.databaseName, dc.username, dc.password)
}
