package logconfig

import (
	"github.com/sillyhatxu/convenient-utils/gotime"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func init() {
	logConfig := NewLogConfig(
		logrus.InfoLevel, true, "test", "test-backend", true, "localhost:51401", false, "",
	)
	logConfig.InitialLogConfig()
}

func TestInputLogstash(t *testing.T) {
	var i int64
	for {
		logrus.Infof("test info[%d] %v", i, gotime.FormatLocation(time.Now()))
		i++
		time.Sleep(1 * time.Second)
	}
}
