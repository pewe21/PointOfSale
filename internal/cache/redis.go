package cache

import (
	"github.com/gofiber/storage/redis/v3"
	"github.com/pewe21/PointOfSale/internal/config"
	"log"
	"runtime"
	"strconv"
	"time"
)

type RedisCache struct {
	cnf     config.Redis
	storage *redis.Storage
}

func NewRedisCache(cnf config.Redis) *RedisCache {
	port, err := strconv.Atoi(cnf.Port)
	if err != nil {
		log.Fatalln(err)
	}
	store := redis.New(redis.Config{
		Host:      cnf.Host,
		Port:      port,
		Username:  "",
		Password:  "",
		Database:  0,
		Reset:     false,
		TLSConfig: nil,
		PoolSize:  10 * runtime.GOMAXPROCS(0),
	})

	return &RedisCache{cnf: cnf, storage: store}
}

func (r *RedisCache) Set(key string, value string, expiration time.Duration) error {
	err := r.storage.Set(key, []byte(value), expiration)
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisCache) Get(key string) (string, error) {
	value, err := r.storage.Get(key)
	if err != nil {
		return "", err
	}

	return string(value), nil
}

func (r *RedisCache) Delete(key string) error {
	err := r.storage.Delete(key)
	if err != nil {
		return err
	}

	return nil
}
