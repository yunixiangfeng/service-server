package setup

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(_ context.Context, development bool, debug bool) (*zap.Logger, error) {
	level := zapcore.InfoLevel
	sampling := &zap.SamplingConfig{
		Initial:    0,
		Thereafter: 0,
	}
	if debug {
		level = zapcore.DebugLevel
		sampling.Initial, sampling.Thereafter = 100, 100
	}

	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "M",
		LevelKey:       "L",
		TimeKey:        "T",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		//EncodeName:          nil,
		//NewReflectedEncoder: nil,
		//ConsoleSeparator:    "",
	}

	cfg := &zap.Config{
		Level:             zap.NewAtomicLevelAt(level),
		Development:       development,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          sampling,
		Encoding:          "console",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     nil,
	}
	return cfg.Build()
}
