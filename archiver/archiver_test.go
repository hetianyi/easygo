package archiver

import (
	"github.com/hetianyi/easygo/file"
	"github.com/hetianyi/easygo/logger"
	"testing"
)

func TestCompressFile(t *testing.T) {
	err := CompressFile(CompressTypeZip, "C:\\var\\demo.archiver", "D:\\tmp\\poi-test", "D:\\tmp\\1.txt", "D:\\tmp\\2.txt")
	if err != nil {
		logger.Fatal("压缩失败: ", err)
	}
	logger.Info("压缩成功")
}

func TestCompressToWriter(t *testing.T) {
	outFile, err := file.CreateFile("C:\\var\\demo.archiver")
	if err != nil {
		logger.Fatal("压缩失败: ", err)
	}
	defer outFile.Close()
	err = CompressToWriter(CompressTypeZip, outFile, "D:\\tmp\\poi-test", "D:\\tmp\\1.txt", "D:\\tmp\\2.txt")
	if err != nil {
		logger.Fatal("压缩失败: ", err)
	}
	logger.Info("压缩成功")
}

func TestDeCompressFile(t *testing.T) {
	err := DeCompressFile(CompressTypeZip, "C:\\var\\demo.archiver", "C:\\var\\out")
	if err != nil {
		logger.Fatal("解压失败: ", err)
	}
	logger.Info("解压成功")
}
