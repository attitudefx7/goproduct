package logger

import "go.uber.org/zap/zapcore"

// Filename: 日志文件的位置
// MaxSize：在进行切割之前，日志文件的最大大小（以MB为单位）
// MaxBackups：保留旧文件的最大个数
// MaxAges：保留旧文件的最大天数
// Compress：是否压缩/归档旧文件
type Config struct {
	Filename   string        `json:"file_name" yaml:"file_name"`
	LogLevel   int8 `json:"log_level" yaml:"log_level"`
	MaxSize    int           `json:"max_size" yaml:"max_size"`
	MaxBackups int           `json:"max_backups" yaml:"max_backups"`
	MaxAge     int           `json:"max_age" yaml:"max_age"`
	Compress   bool          `json:"compress" yaml:"compress"`
}

func getLevelMap() map[int8]zapcore.Level {
	levelMap := make(map[int8]zapcore.Level)
	levelMap[-1] = zapcore.DebugLevel
	levelMap[0] = zapcore.InfoLevel
	levelMap[1] = zapcore.WarnLevel
	levelMap[2] = zapcore.ErrorLevel
	levelMap[3] = zapcore.DPanicLevel
	levelMap[4] = zapcore.PanicLevel
	levelMap[5] = zapcore.FatalLevel

	return levelMap
}

func getLevel(i int8) zapcore.Level {
	if level, ok := getLevelMap()[i]; ok {
		return level
	}

	return getLevelMap()[-1]
}

