package logconfig

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sillyhatxu/convenient-utils/logstashhook"
	"github.com/sillyhatxu/convenient-utils/tcpclient"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

type DefaultFieldHook struct {
	project string
	module  string
}

func (h *DefaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *DefaultFieldHook) Fire(e *logrus.Entry) error {
	e.Data["project"] = h.project
	e.Data["module"] = h.module
	return nil
}

type logConfig struct {
	logLevel        logrus.Level
	reportCaller    bool
	project         string
	module          string
	openLogstash    bool
	logstashAddress string
	openLogfile     bool
	filePath        string
}

func NewLogConfig(logLevel logrus.Level, reportCaller bool, project string, module string, openLogstash bool, logstashAddress string, openLogfile bool, filePath string) *logConfig {
	return &logConfig{
		logLevel:        logLevel,
		reportCaller:    reportCaller,
		project:         project,
		module:          module,
		openLogstash:    openLogstash,
		logstashAddress: logstashAddress,
		openLogfile:     openLogfile,
		filePath:        filePath,
	}
}

func (lc logConfig) String() string {
	return fmt.Sprintf(`logConfig{logLevel=%s, reportCaller=%t, project=%s, module=%s, openLogstash=%t, logstashAddress=%s, openLogfile=%t, filePath=%s}`, lc.logLevel, lc.reportCaller, lc.project, lc.module, lc.openLogstash, lc.logstashAddress, lc.openLogfile, lc.filePath)
}

func (lc logConfig) InitialLogConfig() {
	log.Println("InitialLogConfig :", lc)
	logFormatter := &logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		//TimestampFormat:string("2006-01-02 15:04:05"),
		FieldMap: *&logrus.FieldMap{
			logrus.FieldKeyMsg:  "message",
			logrus.FieldKeyTime: "@timestamp",
		},
	}
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(lc.logLevel)
	logrus.SetReportCaller(lc.reportCaller)
	logrus.SetFormatter(logFormatter)
	logrus.AddHook(&DefaultFieldHook{project: lc.project, module: lc.module})
	if lc.openLogstash {
		conn, err := tcpclient.Dial("tcp", lc.logstashAddress)
		if err != nil {
			logrus.Fatal(err)
		}
		if err != nil {
			log.Panicf("net.Dial(tcp, %v); Error : %v", lc.logstashAddress, err)
		}
		hook := logstashhook.New(conn, logstashhook.DefaultFormatter(logrus.Fields{"project": lc.project, "module": lc.module}))
		logrus.AddHook(hook)
	}
	if lc.openLogfile {
		logPath := lc.filePath + lc.module + "/"
		if !exists(logPath) {
			err := createFolder(logPath)
			if err != nil {
				log.Panicf("createFolder error; Error : %v", err)
			}
		}
		path := logPath + lc.module + ".log"
		WithMaxAge := time.Duration(876000) * time.Hour
		WithRotationTime := time.Duration(24) * time.Hour
		infoWriter, err := rotatelogs.New(
			logPath+"info.log.%Y%m%d",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithMaxAge(WithMaxAge),
			rotatelogs.WithRotationTime(WithRotationTime),
		)
		if err != nil {
			log.Panicf("rotatelogs.New [info writer] error; Error : %v", err)
		}
		errorWriter, err := rotatelogs.New(
			logPath+"error.log.%Y%m%d",
			rotatelogs.WithLinkName(path),
			rotatelogs.WithMaxAge(WithMaxAge),
			rotatelogs.WithRotationTime(WithRotationTime),
		)
		if err != nil {
			log.Panicf("rotatelogs.New [error writer] error; Error : %v", err)
		}
		logrus.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  infoWriter,
				logrus.WarnLevel:  infoWriter,
				logrus.ErrorLevel: infoWriter,
			},
			logFormatter,
		))
		logrus.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.WarnLevel:  errorWriter,
				logrus.ErrorLevel: errorWriter,
			},
			logFormatter,
		))
	}
}

func createFolder(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

//判断所给路径文件/文件夹是否存在
func exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
