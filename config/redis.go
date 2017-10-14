package config

import (
	"crud/utils"
)

type redisConfig struct {
	redisURI    string
	maxConn     int
	idleConn    int
	idleTimeout int
}

func (conf *redisConfig) URI() string {
	return conf.redisURI
}

func (conf *redisConfig) MaxConn() int {
	return conf.maxConn
}

func (conf *redisConfig) IdleConn() int {
	return conf.idleConn
}

func (conf *redisConfig) IdleTimeout() int {
	return conf.idleTimeout
}

func newRedisConfig() *redisConfig {
	return &redisConfig{
		redisURI:    utils.FatalGetString("redis_uri"),
		maxConn:     utils.GetIntOrPanic("redis_max_conn"),
		idleConn:    utils.GetIntOrPanic("redis_idle_conn"),
		idleTimeout: utils.GetIntOrPanic("redis_idle_timeout"),
	}
}
