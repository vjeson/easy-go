package util

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var log *logrus.Logger

var logToFile *logrus.Logger

var loggerFile string

func Log() *logrus.Logger {
	if log == nil {
		log = logrus.New()
		log.Out = os.Stdout
		log.Formatter = &logrus.JSONFormatter{TimestampFormat:"2006-01-02 15:04:05"}
		log.SetLevel(logrus.DebugLevel)
	}
	return log
}

func LogFile() *logrus.Logger {

	if logToFile == nil {
		logToFile = logrus.New()

		//logger.Out = nil

		logToFile.SetLevel(logrus.DebugLevel)

		// 设置 rotatelogs
		logWriter, _ := rotatelogs.New(
			// 分割后的文件名称
			loggerFile + "_%Y%m%d.log",

			// 生成软链，指向最新日志文件
			//rotatelogs.WithLinkName(logFile),

			// 设置最大保存时间(7天)
			rotatelogs.WithMaxAge(7*24*time.Hour),

			// 设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),

		)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		// 新增 Hook
		logToFile.AddHook(lfHook)
	}
	return logToFile
}

func SetLogFile(file string) {
	loggerFile = file
}

func createDir(filePath string)  error  {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath,os.ModePerm)
		return err
	}
	return nil
}

func createFile(path string, fileName string) error  {
	file := filepath.Join(path, fileName)
	if !isExist(file) {
		_,  err := os.Create(file)
		return err
	}
	return nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
