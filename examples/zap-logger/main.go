package main

import (
	"time"

	"github.com/vinay03/chalk"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var formattedLevels = map[string]string{
	"INFO":   chalk.BgBlue().Bold().White().Sprint("INFO "),
	"WARN":   chalk.BgYellowLight().Black().Sprint("WARN "),
	"ERROR":  chalk.White().BgRed().Sprint("ERROR"),
	"DEBUG":  chalk.Black().BgWhite().Sprint("DEBUG"),
	"DPANIC": chalk.White().BgRed().Sprint("DPANIC"),
	"PANIC":  chalk.White().BgRed().Sprint("PANIC"),
	"FATAL":  chalk.White().BgRed().Sprint("FATAL"),
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(chalk.Green().Sprint(t.Format("Jan  2 15:04:05")))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(formattedLevels[level.CapitalString()])
}

func init() {
	cfg := zap.NewDevelopmentConfig()
	cfg.Encoding = "console"
	cfg.EncoderConfig.EncodeTime = SyslogTimeEncoder
	cfg.EncoderConfig.EncodeLevel = CustomLevelEncoder

	logger, _ := cfg.Build()
	logger.Info("This is an info level message")
	logger.Debug("This is an debug level message")
	logger.Warn("This is an warning level message")
	logger.Error("This is an error level message")
	logger.Panic("This is an panic level message")
}
