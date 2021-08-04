package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
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
		OutputPaths:      []string{"storage/logs/default.log"},
		ErrorOutputPaths: []string{"storage/logs/default.log"},
	}.Build()
	if err != nil {
		return err
	}

	zap.ReplaceGlobals(logger)
	return nil
}

// CustomLogger 按级别切割文件、按大小切割文件
func CustomLogger() error {
	file, _ := os.Create("storage/logs/zzz.log")

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(file),
		zapcore.DebugLevel,
	)

	lg := zap.New(core)

	zap.ReplaceGlobals(lg)
	return nil
}

func Test() error {
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		CallerKey:      "file",
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})

	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	infoWriter := getWriter("storage/logs/info.log")
	errorWriter := getWriter("storage/logs/error.log")

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(infoWriter), infoLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(errorWriter), errorLevel),
	)

	log := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(log)
	return nil
}

func getWriter(path string) zapcore.WriteSyncer {
	file, _ := os.Create(path)
	return zapcore.AddSync(file)
}
