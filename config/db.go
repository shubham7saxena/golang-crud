package config

import "fmt"

type dbConfig struct {
	host         string
	port         int
	username     string
	databaseName string
	password     string
}

func newDBConfig() *dbConfig {
	return &dbConfig{
		host:         "localhost",
		port:         5432,
		username:     "postgres",
		databaseName: "test",
		password:     "s7saxena",
	}
}

func (dc *dbConfig) GetConnectionString() string {
	return fmt.Sprintf("dbname=%s user=%s password='%s' sslmode=disable", dc.databaseName, dc.username, dc.password)
}
