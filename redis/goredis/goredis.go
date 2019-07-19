package goredis

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

var RedisConf RedisConfig

func InitialRedisConfig(address, password string, db int) {
	RedisConf.address = address
	RedisConf.password = password
	RedisConf.db = db
}

type RedisConfig struct {
	address  string
	password string
	db       int
}

func (rc RedisConfig) GetClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     rc.address,
		Password: rc.password,
		DB:       rc.db,
	})
	ping, err := client.Ping().Result()
	log.Debugf("Connect to Redis, ping result -> %v", ping)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (rc RedisConfig) Ping() (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     rc.address,
		Password: rc.password,
		DB:       rc.db,
	})
	return client.Ping().Result()
}

func (rc RedisConfig) Get(key string) (string, error) {
	client, err := rc.GetClient()
	if err != nil {
		return "", err
	}
	value, err := client.Get(key).Result()
	if err != nil {
		return "", err
	}
	return value, nil
}

func (rc RedisConfig) Set(key, value string) error {
	client, err := rc.GetClient()
	if err != nil {
		return err
	}
	err = client.Set(key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc RedisConfig) Exists(key string) (bool, error) {
	client, err := rc.GetClient()
	if err != nil {
		return false, err
	}
	count := client.Exists(key).Val()
	return count > 0, nil
}

func (rc RedisConfig) Expire(key string, expiration time.Duration) error {
	client, err := rc.GetClient()
	if err != nil {
		return err
	}
	err = client.Expire(key, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc RedisConfig) Incr(key string) (int64, error) {
	client, err := rc.GetClient()
	if err != nil {
		return 0, err
	}
	seq := client.Incr(key).Val()
	return seq, nil
}
