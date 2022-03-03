package logger_test

import (
	"github.com/hetianyi/easygo/logger"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	log := logger.New(os.Stdout, "<EASYGO> ", logger.LevelDebug, logger.DefaultFormatter)
	log.Debug("你好")
	log.Info("你好")
	log.Warn("你好")
	log.Error("你好")
	log.Fatal("你好")
}

func TestInitDefaultLogger(t *testing.T) {
	logger.SetPrefix("[WEB] ")
	logger.Debug("你好")
	logger.Info("你好")
	logger.Warn("你好")
	logger.Error("你好")
	logger.Fatal("你好")
}

func TestSimpleFormatter(t *testing.T) {
	sl := logger.New(os.Stdout, "<EASYGO> ", logger.LevelDebug, logger.SimpleFormatter)
	sl.Debug("你好")
	sl.Info("你好")
	sl.Warn("你好")
	sl.Error("你好")
	sl.Fatal("你好")
}
