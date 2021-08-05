package log

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitLogger 将日志写入本地文件
func InitLogger() error {
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"storage/logs/log.log"},
		ErrorOutputPaths: []string{"storage/logs/log.log"},
	}.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)
	return nil
}

// CustomLogger 按级别切割文件、按大小切割文件
func CustomLogger() error {
	encoder := getEncoder()

	debugWriter := getWriter("storage/logs/debug.log")
	infoWriter := getWriter("storage/logs/info.log")
	errorWriter := getWriter("storage/logs/error.log")

	debugLever := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return zapcore.DebugLevel < lvl && lvl < zapcore.ErrorLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(debugWriter), debugLever),
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)

	log := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(log)
	return nil
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "",
		LevelKey:       "level",
		CallerKey:      "file",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func getWriter(path string) io.Writer {
	return &lumberjack.Logger{
		Filename:   path,
		MaxSize:    10,
		MaxAge:     14,
		MaxBackups: 5,
		LocalTime:  true,
		Compress:   false,
	}
}
