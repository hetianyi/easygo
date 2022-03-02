package logger_test

import (
	"github.com/hetianyi/easygo/logger"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	log := logger.New(os.Stdout, "<EASYGO> ", logger.LevelInfo, logger.DefaultFormatter)
	log.Debug("你好")
	log.Info("你好")
	log.Warn("你好")
	log.Error("你好")
	log.Fatal("你好")
}

func TestNewDefaultLogger(t *testing.T) {
	log := logger.NewDefaultLogger()
	log.SetPrefix("[WEB] ")
	log.Debug("你好")
	log.Info("你好")
	log.Warn("你好")
	log.Error("你好")
	log.Fatal("你好")
}
