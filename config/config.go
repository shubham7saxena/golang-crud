package config

import "github.com/spf13/viper"

type config struct {
	dbConfig    *dbConfig
	redisConfig *redisConfig
}

var appConfig *config

func Load() {
	viper.AutomaticEnv()

	viper.SetConfigName("application")
	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

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
