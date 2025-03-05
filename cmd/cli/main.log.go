package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	//// Suggar loggeer
	//sugar := zap.NewExample().Sugar()
	//sugar.Infof("Hello ===> Name: %s, Age: %d", "Tran Van Dat", 100)
	//
	//// Logger
	//logger := zap.NewExample()
	//logger.Info("Hello", zap.String("Name", "Tran Van Dat"), zap.Int("Age", 100))

	//logger := zap.NewExample()
	//logger.Info("Hello Log Example")
	//
	//// Development
	//logger, _ = zap.NewDevelopment()
	//logger.Info("Hello Log Development")
	//
	//// Product
	//logger, _ = zap.NewProduction()
	//logger.Info("Hello Log Development")

	encoder := getEncoderLog()
	sync := getWriteSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	//1741148263.362946 -->2025-03-05T16:16:07.887+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts --> time
	encodeConfig.TimeKey = "time"
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}
func getWriteSync() zapcore.WriteSyncer {
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncConsole := zapcore.AddSync(os.Stderr)
	syncFile := zapcore.AddSync(file)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
