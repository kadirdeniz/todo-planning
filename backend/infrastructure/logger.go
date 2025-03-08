package infrastructure

import "go.uber.org/zap"

type ILogger interface {
    Info(msg string, fields ...zap.Field)
    Error(msg string, fields ...zap.Field)
    Debug(msg string, fields ...zap.Field)
}

type logger struct {
    log *zap.Logger
}

func NewLogger() (ILogger, error) {
    zapLogger, err := zap.NewDevelopment()
    if err != nil {
        return nil, err
    }
    return &logger{log: zapLogger}, nil
}	

func (l *logger) Info(msg string, fields ...zap.Field)  { l.log.Info(msg, fields...) }
func (l *logger) Error(msg string, fields ...zap.Field) { l.log.Error(msg, fields...) }
func (l *logger) Debug(msg string, fields ...zap.Field) { l.log.Debug(msg, fields...) }