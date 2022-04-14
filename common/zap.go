package common

import (
	// "fmt"
	// "time"
	// "os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *zap.SugaredLogger
)

type ZapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger(serviceName, fileName string) *ZapLogger {
	if len(fileName) == 0 {
		fileName = "micro.log" //日志文件名称
	}
	zapLog := &ZapLogger{
		logger: &zap.SugaredLogger{},
	}
	zapLog.Init(serviceName, fileName)
	return zapLog
}

func (zaplog *ZapLogger) Init(serviceName, fileName string) {
	// now := time.Now()
	hook := &lumberjack.Logger{
		Filename: fileName, //文件名称
		// Filename: fmt.Sprintf("log/%s%04d%02d%02d%02d%02d%02d", fileName, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()), //filePath

		MaxSize: 512, //MB
		//MaxAge:     0, //文件最多保存多少天
		MaxBackups: 0, //最大备份
		LocalTime:  true,
		Compress:   true, //是否启用压缩
	}
	defer hook.Close()
	syncWriter := zapcore.AddSync(hook)
	//编码
	encoder := zap.NewProductionEncoderConfig() //生产环境
	// encoder := zap.NewDevelopmentEncoderConfig() //开发环境
	//时间格式
	encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		// 编码器
		zapcore.NewJSONEncoder(encoder),
		// zapcore.NewConsoleEncoder(encoder), //编码器配置
		syncWriter, //打印到文件
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),syncWriter), // 打印到控制台和文件
		zap.NewAtomicLevelAt(zap.DebugLevel), //日志等级
	)
	log := zap.New(
		core,
		zap.AddCaller(),      //日志添加调用者信息
		zap.AddCallerSkip(1), //设置skip，用户runtime.Caller的参数
		// zap.AddStacktrace(zapcore.DebugLevel), //设置堆栈跟踪
		zap.Fields(zap.String("serviceName", serviceName)), //设置初始化字段
		// zap.Development(),//开发环境 panic开启文件及行号
	)
	zaplog.logger = log.Sugar()
}

func (zaplog *ZapLogger) Debug(args ...interface{}) {
	zaplog.logger.Debug(args)
}

func (zaplog *ZapLogger) Debugf(template string, args ...interface{}) {
	zaplog.logger.Debugf(template, args...)
}

func (zaplog *ZapLogger) Info(args ...interface{}) {
	zaplog.logger.Info(args...)
}

func (zaplog *ZapLogger) Infof(template string, args ...interface{}) {
	zaplog.logger.Infof(template, args...)
}

func (zaplog *ZapLogger) Warn(args ...interface{}) {
	zaplog.logger.Warn(args...)
}

func (zaplog *ZapLogger) Warnf(template string, args ...interface{}) {
	zaplog.logger.Warnf(template, args...)
}

func (zaplog *ZapLogger) Error(args ...interface{}) {
	zaplog.logger.Error(args...)
}

func (zaplog *ZapLogger) Errorf(template string, args ...interface{}) {
	zaplog.logger.Errorf(template, args...)
}

func (zaplog *ZapLogger) DPanic(args ...interface{}) {
	zaplog.logger.DPanic(args...)
}

func (zaplog *ZapLogger) DPanicf(template string, args ...interface{}) {
	zaplog.logger.DPanicf(template, args...)
}

func (zaplog *ZapLogger) Panic(args ...interface{}) {
	zaplog.logger.Panic(args...)
}

func (zaplog *ZapLogger) Panicf(template string, args ...interface{}) {
	zaplog.logger.Panicf(template, args...)
}

func (zaplog *ZapLogger) Fatal(args ...interface{}) {
	zaplog.logger.Fatal(args...)
}

func (zaplog *ZapLogger) Fatalf(template string, args ...interface{}) {
	zaplog.logger.Fatalf(template, args...)
}
