package config

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
		redisURI:    "localhost:6379",
		maxConn:     5,
		idleConn:    5,
		idleTimeout: 5,
	}
}
