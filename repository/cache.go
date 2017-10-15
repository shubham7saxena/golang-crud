package repository

import (
	"crud/appcontext"
	"fmt"

	redis "github.com/garyburd/redigo/redis"
)

type redisPool struct {
	pool *redis.Pool
}

func GetNewRedisPool() *redisPool {
	return &redisPool{
		pool: appcontext.GetRedis(),
	}
}

func (r *redisPool) Ping() error {

	conn := r.pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		return fmt.Errorf("cannot 'PING' db: %v", err)
	}
	return nil
}

func (r *redisPool) Get(key string) ([]byte, bool) {

	conn := r.pool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		fmt.Errorf("error getting key from redis")
		return data, false
	}
	return data, true
}

func (r *redisPool) Set(key string, value []byte) {

	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		fmt.Errorf("Error setting key in redis")
	}
}

func (r *redisPool) Exists(key string) bool {

	conn := r.pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		fmt.Errorf("cannot check key existence in redis")
		return false
	}
	return ok
}

func (r *redisPool) Delete(key string) error {

	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}
