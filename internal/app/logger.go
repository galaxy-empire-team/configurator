package app

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/galaxy-empire-team/configurator/internal/config"
)

func newLogger(cfg config.App) (*zap.Logger, error) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	level, err := zap.ParseAtomicLevel(cfg.LogLevel)
	if err != nil {
		return nil, fmt.Errorf("zap.NewAtomicLevelAt(): %w", err)
	}

	config := zap.Config{
		Level:             level,
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          cfg.LogFormat,
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"stdout"},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("config.Build(): %w", err)
	}

	return logger, nil
}
