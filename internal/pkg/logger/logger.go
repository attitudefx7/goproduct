package logger

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func Init(c *Config) {
	// 指定日志将写到哪里去。我们使用zapcore.AddSync()函数并且将打开的文件句柄传进去
	writeSyncer := getLogWriter(c)
	// 编码器(如何写入日志)。我们将使用开箱即用的NewJSONEncoder()，并使用预先设置的ProductionEncoderConfig()。
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, writeSyncer, getLevel(c.LogLevel))

	zapLogger := zap.New(core, zap.AddCaller())
	logger = zapLogger.Sugar()

	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	// 修改时间编码器
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 日志格式大写
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
	//return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(c *Config) zapcore.WriteSyncer {
	// 日志切割
	lumberJackLogger := &lumberjack.Logger{
		Filename:   c.Filename,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func GetLogger() *zap.SugaredLogger {
	return logger
}