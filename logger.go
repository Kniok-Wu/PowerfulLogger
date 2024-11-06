package PowerfulLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

// NewDefaultLogger 新建一个logger
func NewDefaultLogger(encoder zapcore.Encoder, writer io.Writer) *zap.Logger {
	core := zapcore.NewCore(encoder, zapcore.AddSync(writer), zap.InfoLevel)
	logger := zap.New(core)
	return logger
}

// NewErrorWriter 新建一个Error独立地logger
func NewErrorWriter(encoder zapcore.Encoder, writer io.Writer) *zap.Logger {
	core := zapcore.NewCore(encoder, zapcore.AddSync(writer), zap.ErrorLevel)
	logger := zap.New(core)
	return logger
}
