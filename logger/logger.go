package logger

import (
	"github.com/Kemosabe2911/employee-app-go/constant"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugaredLogger *zap.SugaredLogger

func InitLogger(env string) (*zap.SugaredLogger, error) {
	logger, err := getLoggerbyEnv(env)
	if err != nil {
		return nil, err
	}

	sugaredLogger = logger.Sugar()
	return sugaredLogger, nil
}

func getLoggerbyEnv(env string) (logger *zap.Logger, err error) {
	option := zap.AddCallerSkip(1)

	if env == constant.PRD {
		return zap.NewProduction(option)
	}

	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return config.Build(option)
}

func Errorw(message string, args ...interface{}) {
	sugaredLogger.Errorw(message, args...)
}

func Errorf(message string, args ...interface{}) {
	sugaredLogger.Errorf(message, args...)
}

func Error(args ...interface{}) {
	sugaredLogger.Error(args...)
}

func Infow(message string, args ...interface{}) {
	sugaredLogger.Infow(message, args...)
}

func Infof(message string, args ...interface{}) {
	sugaredLogger.Infof(message, args...)
}

func Info(args ...interface{}) {
	sugaredLogger.Info(args...)
}

func Warnw(message string, args ...interface{}) {
	sugaredLogger.Warnw(message, args...)
}

func Warnf(message string, args ...interface{}) {
	sugaredLogger.Warnf(message, args...)
}

func Warn(args ...interface{}) {
	sugaredLogger.Warn(args...)
}

func Debugw(message string, args ...interface{}) {
	sugaredLogger.Debugw(message, args...)
}

func Debugf(message string, args ...interface{}) {
	sugaredLogger.Debugf(message, args...)
}

func Debug(args ...interface{}) {
	sugaredLogger.Debug(args...)
}

// Fatalf Wrapper.
func Fatalf(message string, args ...interface{}) {
	sugaredLogger.Fatalf(message, args...)
}

// Fatal Wrapper.
func Fatal(args ...interface{}) {
	sugaredLogger.Fatal(args...)
}
