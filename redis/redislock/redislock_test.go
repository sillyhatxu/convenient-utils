package redislock

import (
	redis "github.com/sillyhatxu/go-utils/redis/goredis"
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"
	"time"
)

type LockDemo struct {
	ThreadName string
}

func (ld LockDemo) Execute() error {
	//log.Infof("I'm on this mission;Thread : %v", ld.ThreadName)
	time.Sleep(3 * time.Second)
	//log.Infof("Mission Finish. Thread : %v", ld.ThreadName)
	log.Infof("Execute --- Thread : %v", ld.ThreadName)
	return nil
}

func (ld LockDemo) LockKey() string {
	//log.Infof("ld.ThreadName : %v", ld.ThreadName)
	return "TEST_LOCK"
}

func TestExample(t *testing.T) {
	redis.InitialRedisConfig("127.0.0.1:6379", "", 0)
	//log.SetLevel(log.ErrorLevel)
	//log.SetLevel(log.DebugLevel)
	log.SetLevel(log.InfoLevel)
	for i := 1; i <= 5; i++ {
		go RedisLock(LockDemo{ThreadName: "THREAD_" + strconv.Itoa(i)})
	}
	time.Sleep(35 * time.Second)
	//forever := make(chan bool)
	//<-forever
}
