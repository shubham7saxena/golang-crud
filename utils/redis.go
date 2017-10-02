package utils

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

func (r *redisPool) Get(key string) ([]byte, error, bool) {

	conn := r.pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s: %v", key, err), false
	}
	return data, err, true
}

func (r *redisPool) Set(key string, value []byte) error {

	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return fmt.Errorf("error setting key %s to %s: %v", key, value, err)
	}
	return err
}

func (r *redisPool) Exists(key string) (bool, error) {

	conn := r.pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking if key %s exists: %v", key, err)
	}
	return ok, err
}

func (r *redisPool) Delete(key string) error {

	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}
