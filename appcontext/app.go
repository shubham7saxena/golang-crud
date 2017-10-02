package appcontext

import (
	"crud/config"
	sql "database/sql"
	"fmt"
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
	_ "github.com/lib/pq"
)

type appContext struct {
	db    *sql.DB
	redis *redis.Pool
}

var context *appContext

func Initiate() {
	db := initDB()
	redis := initRedis()
	context = &appContext{
		db:    db,
		redis: redis,
	}
}

func GetDB() *sql.DB {
	return context.db
}

func initDB() *sql.DB {
	var err error
	db, err := sql.Open("postgres", config.NewDBConfig().GetConnectionString())

	if err != nil {
		log.Fatalf("Error connecting to the database")
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Ping to database host failed: %s \n", err)
	}
	return db
}

func initRedis() *redis.Pool {
	conf := config.NewRedisConfig()
	redis := &redis.Pool{
		MaxIdle:     conf.IdleConn(),
		MaxActive:   conf.MaxConn(),
		IdleTimeout: time.Duration(conf.IdleTimeout()) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.URI())
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	conn := redis.Get()
	defer conn.Close()

	err := conn.Send("PING")

	if err != nil {
		fmt.Errorf("Unable to connect to redis server %s: %v", conf.URI(), err)
	}

	return redis
}

func GetRedis() *redis.Pool {
	return context.redis
}
