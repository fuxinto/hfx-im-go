package fxlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var fxlog *zap.Logger

func Infof(format string, v ...interface{}) {
	Info(fmt.Sprintf(format, v...))
}
func Info(msg string, fields ...zap.Field) {
	fxlog.Info(msg, fields...)

}

func Warn(msg string, fields ...zap.Field) {
	fxlog.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	fxlog.Error(msg, fields...)

}

func Panic(msg string, fields ...zap.Field) {
	fxlog.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	fxlog.Fatal(msg, fields...)
}

/*log路径文件名*/
func Initlog(pathName, serviceName string) {

	hook := lumberjack.Logger{
		Filename:   pathName,
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     customTimeEncoder,              // 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", serviceName))
	// 构造日志
	fxlog = zap.New(core, caller, development, filed)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}
