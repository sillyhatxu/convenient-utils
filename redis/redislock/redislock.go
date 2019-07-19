package redislock

import (
	"errors"
	"github.com/bsm/redis-lock"
	redis "github.com/sillyhatxu/go-utils/redis/goredis"
	log "github.com/sirupsen/logrus"
	"time"
)

type LockInterface interface {
	LockKey() string
	Execute() error
}

func RedisLock(lockInterface LockInterface) error {
	defer log.Debugf("... RedisLock End ...")
	client, err := redis.RedisConf.GetClient()
	if err != nil {
		return err
	}
	defer client.Close()
	log.Debug("Connect to Redis")
	var locker *lock.Locker
	for i := 1; i <= 30; i++ {
		locker, err = lock.Obtain(client, lockInterface.LockKey(), nil)
		if err != nil {
			log.Debugf("... Hold [%v] ...", lockInterface.LockKey())
			time.Sleep(time.Duration(i*200) * time.Millisecond)
		} else if locker == nil {
			log.Error("locker is null")
			return err
		} else {
			hasLock, err := locker.Lock()
			if err != nil {
				log.Error(err)
				return err
			}
			if hasLock {
				defer locker.Unlock()
				err = lockInterface.Execute()
				if err != nil {
					return err
				}
				return nil
			}
			return errors.New("[ locker.Lock() ] Unknow error.")
		}
	}
	return err
}
