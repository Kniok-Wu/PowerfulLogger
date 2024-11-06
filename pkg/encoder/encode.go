package encoder

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DefaultJsonEncoder 默认的JSON格式
func DefaultJsonEncoder() zapcore.Encoder {
	// 创建一个新的生产环境编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	// 配置时间的编码方式为 ISO8601 格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// 配置日志级别的编码方式为大写形式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 设置一个caller
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 创建一个新的 JSON 编码器，应用上述配置
	return zapcore.NewJSONEncoder(encoderConfig)
}

//// DefaultTerminalEncoder 默认的命令行格式
//func DefaultTerminalEncoder() zapcore.Encoder {
//	// 创建一个新的生产环境编码器配置
//	encoderConfig := zap.NewProductionEncoderConfig()
//	// 配置时间的编码方式为 ISO8601 格式
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	// 配置日志级别的编码方式为大写形式
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	// 设置一个caller
//	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
//	// 创建一个新的 JSON 编码器，应用上述配置
//	return zapcore.NewEncoder(encoderConfig)
//}
