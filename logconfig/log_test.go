package logconfig

import (
	"fmt"
	"github.com/sillyhatxu/convenient-utils/gotime"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func init() {
	logConfig := NewLogConfig(
		logrus.InfoLevel,
		true,
		"test-backend",
		"test-backend",
		true,
		//"logstash:5000",
		"localhost:51401",
		false,
		"/Users/shikuanxu/go/src/github.com/sillyhatxu/convenient-utils/logs",
	)

	logConfig.InitialLogConfig()
}

func TestLogstash(t *testing.T) {
	logrus.Infof("test : %v", "haha")
	logrus.Errorf("test err : %v", fmt.Errorf("create error"))
}

func TestInputLogstash(t *testing.T) {
	var i int64
	for {
		logrus.Infof("test info[%d] %v", i, gotime.FormatLocation(time.Now()))
		i++
		time.Sleep(5 * time.Second)
	}
}
