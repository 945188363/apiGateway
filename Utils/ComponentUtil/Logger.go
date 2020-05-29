package ComponentUtil

import (
	"apiGateway/Config"
	"apiGateway/DBModels"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/vladoatanasov/logrus_amqp"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/sohlich/elogrus.v2"
	"log"
	"os"
	"path"
	"time"
)

var accessLog *logrus.Logger
var runtimeLog *logrus.Logger
var runtimeLogInfo *DBModels.LogInfo

func Init() {
	runtimeLogInfo.LogType = Config.RuntimeLog
	err := runtimeLogInfo.GetLogInfoByType()
	if err != nil {
		return
	}
}

func AccessLog(logInfo *DBModels.LogInfo) *logrus.Logger {

	// 日志文件
	fileName := path.Join(logInfo.LogAddress, logInfo.LogName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	// 实例化
	logger := logrus.New()
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置输出
	logger.Out = src

	// 记录周期
	var period int
	switch logInfo.LogPeriod {
	case Config.Hour:
		period = 1
	case Config.Day:
		period = 24
	}
	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(默认天)
		rotatelogs.WithMaxAge(time.Duration(logInfo.LogSavedTime)*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(time.Duration(period)*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return logger
}

func RuntimeLog() *logrus.Logger {
	// 日志文件
	fileName := path.Join(runtimeLogInfo.LogAddress, runtimeLogInfo.LogName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	// 实例化
	logger := logrus.New()
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置输出
	logger.Out = src

	// 记录周期
	var period int
	switch runtimeLogInfo.LogPeriod {
	case Config.Hour:
		period = 1
	case Config.Day:
		period = 24
	}

	// 设置 rotatelogs
	logWriter, err := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(默认天)
		rotatelogs.WithMaxAge(time.Duration(runtimeLogInfo.LogSavedTime)*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(time.Duration(period)*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	return logger
}

// config logrus log to amqp
func ConfigAmqpLogger(logger *logrus.Logger) {
	hook := logrus_amqp.NewAMQPHookWithType(Config.MqUrl, Config.MqUsername, Config.MqPassword, Config.MqExchange,
		Config.MqExchangeType, Config.MqVirtualHost, Config.MqRoutingKey)
	logger.AddHook(hook)
}

// config logrus log to es
func ConfigESLogger(logger *logrus.Logger) {
	client, err := elastic.NewClient(elastic.SetURL(Config.EsUrl))
	if err != nil {
		logger.Errorf("config es logger error. %+v", errors.WithStack(err))
	}
	esHook, err := elogrus.NewElasticHook(client, Config.EsHost, logrus.DebugLevel, Config.EsIndex)
	if err != nil {
		logger.Errorf("config es logger error. %+v", errors.WithStack(err))
	}
	logger.AddHook(esHook)
}
