package log

import (
	"fenv-eats/config"
	nested "github.com/antonfisher/nested-logrus-formatter"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path"
	"sync"
	"time"
)

var (
	logPath     = "./log"
	logFile     = "cfn-fenv.log"
	Instance    *logrus.Logger
	logInitOnce sync.Once
)

/*
	注意：此为业务型日志，项目启动初始化时建议还是使用系统日志会好点，减少日志文件生成
		例如， print和fatal这种涉及系统级别的操作，使用原生会好点，没必要记录
		业务日志一般记录等级都有 trace，debug, info, warning, error级别即可
*/

func Init() (err error) {
	logInitOnce.Do(func() {
		Instance = logrus.New()
		// 打开文件
		logFileName := path.Join(logPath, logFile)
		// 使用滚动压缩方式记录日志
		rolling(logFileName)

		// 设置日志输出JSON格式
		//LogInstance.SetFormatter(&logrus.JSONFormatter{})
		// 设置日志输出文本格式
		//LogInstance.SetFormatter(&logrus.TextFormatter{
		//	// 以下设置只是为了使输出更美观
		//	DisableColors: true,
		//	// 处理文件名
		//	CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
		//		fileName := path.Base(frame.File)
		//		return frame.Function,fileName
		//	},
		//	TimestampFormat: "2006-01-02 15:04:05",
		//
		//})
		Instance.SetReportCaller(false)
		// 设置日志记录级别
		globalApp := config.App{}
		switch globalApp.Mode {
		case "debug":
			Instance.SetLevel(logrus.DebugLevel)
		case "dev":
			Instance.SetLevel(logrus.TraceLevel)
		case "release":
			Instance.SetLevel(logrus.InfoLevel)
		default:
			Instance.SetLevel(logrus.InfoLevel)
		}
		if Instance.GetLevel() == logrus.TraceLevel || Instance.GetLevel() == logrus.DebugLevel {
			Instance.SetReportCaller(true)
		}

		// 使用第三方库 	go get github.com/antonfisher/nested-logrus-formatter 自定义日志格式
		Instance.SetFormatter(&nested.Formatter{
			HideKeys:        true,
			NoColors:        true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	})
	return nil
}

// 日志滚动设置
func rolling(logFile string) {
	//mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
	//	Filename:   logFile, //日志文件位置
	//	MaxSize:    5,       // 单文件最大容量,单位是MB
	//	MaxBackups: 20,      // 最大保留过期文件个数
	//	MaxAge:     1,       // 保留过期文件的最大时间间隔,单位是天
	//	Compress:   true,    // 是否需要压缩滚动日志, 使用的 gzip 压缩
	//})
	//// 设置输出，lumberjack好像不太好用，建议以后切换为file-rotatelogs
	//LogInstance.SetOutput(mw)

	// go get github.com/lestrrat-go/file-rotatelogs
	/* 日志轮转相关函数
	`WithLinkName` 为最新的日志建立软连接
	`WithRotationTime` 设置日志分割的时间，隔多久分割一次
	WithMaxAge 和 WithRotationCount二者只能设置一个
	 `WithMaxAge` 设置文件清理前的最长保存时间
	 `WithRotationCount` 设置文件清理前最多保存的个数
	*/
	// 下面配置日志每隔 1 分钟轮转一个新文件，保留最近 3 分钟的日志文件，多余的自动清理掉。
	writer, _ := rotateLogs.New(
		logFile+".%Y%m%d%H%M",
		rotateLogs.WithLinkName(logFile),
		rotateLogs.WithRotationCount(30),
		rotateLogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	mv := io.MultiWriter(os.Stdout, writer)
	Instance.SetOutput(mv)
}

func Info(msg ...interface{}) {
	Instance.Info(msg)
}

func InfoLn(msg ...interface{}) {
	Instance.Infoln(msg)
}

func InfoF(tmp string, args ...interface{}) {
	Instance.Infof(tmp, args)
}

func Warning(msg ...interface{}) {
	Instance.Warning(msg)
}

func WarnF(tmp string, args ...interface{}) {
	Instance.Warnf(tmp, args)
}

func ErrorF(tmp string, args ...interface{}) {
	Instance.Errorf(tmp, args)
}

func DebugF(tmp string, args ...interface{}) {
	Instance.Debugf(tmp, args)
}

func TraceF(tmp string, args ...interface{}) {
	Instance.Tracef(tmp, args)
}

/*
	------ 以下是涉及系统启动或者崩溃的日志记录，不会存入文件 ------
*/

func Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

func Println(msg ...interface{}) {
	log.Println(msg)
}

func Fatal(msg ...interface{}) {
	log.Fatal(msg)
}
