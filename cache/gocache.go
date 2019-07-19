package cache

import (
	gocache "github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"time"
)

var CacheConf cacheConfig

const (
	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

type cacheConfig struct {
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	client            *gocache.Cache
}

func InitialGoCache(defaultExpiration, cleanupInterval time.Duration) {
	CacheConf.defaultExpiration = defaultExpiration
	CacheConf.cleanupInterval = cleanupInterval
	CacheConf.client = gocache.New(defaultExpiration, cleanupInterval)
}

func check() {
	if CacheConf.client == nil {
		log.Warning("Uninitialized go cache")
		InitialGoCache(2*time.Hour, 4*time.Hour)
	}
}

func Set(key string, value interface{}, d time.Duration) {
	check()
	CacheConf.client.Set(key, value, d)
}

func Get(key string) (interface{}, bool) {
	check()
	value, found := CacheConf.client.Get(key)
	return value, found
}

func Delete(key string) {
	check()
	CacheConf.client.Delete(key)
}
