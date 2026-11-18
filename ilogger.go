package ysepaygo

import "context"

// 设置外部日志记录器
type ILogger interface {
	Error(ctx context.Context, v ...any)
	Info(ctx context.Context, v ...any)
	Fatal(ctx context.Context, v ...any)
	Debug(ctx context.Context, v ...any)
}

var Logger ILogger = &noopLogger{}

type noopLogger struct{}

func (s *noopLogger) Debug(ctx context.Context, v ...any) {
	if Logger != nil {
		Logger.Debug(ctx, ServiceName, v)
	}
}

func (s *noopLogger) Info(ctx context.Context, v ...any) {
	if Logger != nil {
		Logger.Info(ctx, ServiceName, v)
	}
}

func (s *noopLogger) Error(ctx context.Context, v ...any) {
	if Logger != nil {
		Logger.Error(ctx, ServiceName, v)
	}
}

func (s *noopLogger) Fatal(ctx context.Context, v ...any) {
	if Logger != nil {
		Logger.Fatal(ctx, ServiceName, v)
	}
}

func SetLogger(l ILogger) {
	Logger = l
}
