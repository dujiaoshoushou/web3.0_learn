package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"strings"
	"time"
)

func SetLogger() {
	SetLoggerWriter()
	initLogger()
}

var logger *slog.Logger

func Logger() *slog.Logger {
	return logger
}

var logWriter io.Writer

func LogWriter() io.Writer {
	return logWriter
}

func SetLoggerWriter() {
	month := time.Now().Format("2006-01-02")
	logfile := viper.GetString("app.log.path")
	logfile += fmt.Sprintf("/app-%s.log", month)
	logWriter = &lumberjack.Logger{
		Filename:   logfile,
		MaxSize:    viper.GetInt("app.log.max_size"),
		MaxBackups: viper.GetInt("app.log.max_backups"),
		MaxAge:     viper.GetInt("app.log.max_age"),
		Compress:   viper.GetBool("app.log.compress"),
	}
}

func initLogger() {
	switch strings.ToUpper(gin.Mode()) {
	case APP_RELEASE:
		logger = slog.New(slog.NewTextHandler(logWriter, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	case APP_TEST, APP_DEBUG:
		fallthrough
	default:
		logger = slog.New(slog.NewJSONHandler(logWriter, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

	}
	slog.SetDefault(logger)
}
