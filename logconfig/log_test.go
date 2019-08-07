package logconfig

import (
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
		false,
		//"logstash:5000",
		"localhost:5000",
		true,
		"/Users/shikuanxu/go/src/github.com/sillyhatxu/convenient-utils/logs",
	)

	logConfig.InitialLogConfig()
}

func TestInputLogstash(t *testing.T) {
	var i int64
	for {
		logrus.Infof("test info[%d] %v", i, gotime.FormatLocation(time.Now()))
		logrus.Errorf("test info[%d] %v", i, gotime.FormatLocation(time.Now()))
		i++
		time.Sleep(5 * time.Second)
	}
}
