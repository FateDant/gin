package logs_source

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Log logrus第三方包
var Log = logrus.New()

func init() {
	logConf := LoadLogConfig()
	file, err := os.OpenFile(logConf.LogDir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	Log.Out = file

	//设置日志级别
	levelMapping := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}
	Log.SetLevel(levelMapping[logConf.LogLevel])

	//日志格式化
	Log.SetFormatter(&logrus.TextFormatter{})
}
