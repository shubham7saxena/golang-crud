package config

type redisConfig struct {
	redisURI    string
	maxConn     int
	idleConn    int
	idleTimeout int
}

func (self *redisConfig) URI() string {
	return self.redisURI
}

func (self *redisConfig) MaxConn() int {
	return self.maxConn
}

func (self *redisConfig) IdleConn() int {
	return self.idleConn
}

func (self *redisConfig) IdleTimeout() int {
	return self.idleTimeout
}

func NewRedisConfig() *redisConfig {
	return &redisConfig{
		redisURI:    "localhost:6379",
		maxConn:     5,
		idleConn:    5,
		idleTimeout: 5,
	}
}
