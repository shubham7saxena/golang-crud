package config

type config struct {
	dbConfig    *dbConfig
	redisConfig *redisConfig
}

var appConfig *config

func Load() {
	appConfig = &config{
		dbConfig:    newDBConfig(),
		redisConfig: newRedisConfig(),
	}
}

func DbConfig() *dbConfig {
	return appConfig.dbConfig
}

func RedisConfig() *redisConfig {
	return appConfig.redisConfig
}
