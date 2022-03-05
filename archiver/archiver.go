package archiver

import (
	"errors"
	"fmt"
	"github.com/hetianyi/easygo/file"
	"github.com/mholt/archiver"
	"io"
	"os"
)

type CompressType int

const (
	CompressTypeZip = iota
	CompressTypeTar
	CompressTypeTarGz
	CompressTypeTarXZ
	CompressTypeRar
	CompressTypeTarBz2
	CompressTypeTarLz4
	CompressTypeTarSz
)

// CompressFile 将若干文件压缩为指定路径的压缩文件，
// 请自行确保输入正确的文件扩展名。
func CompressFile(ct CompressType, outputFilePath string, srcFiles ...string) (err error) {
	if err = isCompressTypeSupport(ct); err != nil {
		return err
	}
	var outputFile *os.File

	defer func() {
		if err != nil && outputFile != nil {
			file.DeleteFile(outputFile)
		}
	}()

	outputFile, err = file.CreateFile(outputFilePath)
	if err != nil {
		return err
	}
	err = CompressToWriter(ct, outputFile, srcFiles...)
	return
}

// CompressToWriter 将若干文件压缩到指定输出
func CompressToWriter(ct CompressType, out io.Writer, srcFiles ...string) error {
	if err := isCompressTypeSupport(ct); err != nil {
		return err
	}
	switch ct {
	case CompressTypeZip:
		return archiver.Zip.Write(out, srcFiles)
	case CompressTypeTar:
		return archiver.Tar.Write(out, srcFiles)
	case CompressTypeTarGz:
		return archiver.TarGz.Write(out, srcFiles)
	case CompressTypeTarXZ:
		return archiver.TarXZ.Write(out, srcFiles)
	case CompressTypeRar:
		return archiver.Rar.Write(out, srcFiles)
	case CompressTypeTarBz2:
		return archiver.TarBz2.Write(out, srcFiles)
	case CompressTypeTarLz4:
		return archiver.TarLz4.Write(out, srcFiles)
	case CompressTypeTarSz:
		return archiver.TarSz.Write(out, srcFiles)
	}
	return nil
}

// DeCompressFile 将压缩文件解压到指定的outputDir文件夹下
func DeCompressFile(ct CompressType, srcFilePath, outputDir string) error {
	if err := isCompressTypeSupport(ct); err != nil {
		return err
	}
	srcFile, err := file.GetFile(srcFilePath)
	if err != nil {
		return err
	}
	return DeCompressFromReader(ct, srcFile, outputDir)
}

// DeCompressFromReader 将压缩文件流解压到指定的outputDir文件夹下
func DeCompressFromReader(ct CompressType, reader io.Reader, outputDir string) error {
	if err := isCompressTypeSupport(ct); err != nil {
		return err
	}

	switch ct {
	case CompressTypeZip:
		return archiver.Zip.Read(reader, outputDir)
	case CompressTypeTar:
		return archiver.Tar.Read(reader, outputDir)
	case CompressTypeTarGz:
		return archiver.TarGz.Read(reader, outputDir)
	case CompressTypeTarXZ:
		return archiver.TarXZ.Read(reader, outputDir)
	case CompressTypeRar:
		return archiver.Rar.Read(reader, outputDir)
	case CompressTypeTarBz2:
		return archiver.TarBz2.Read(reader, outputDir)
	case CompressTypeTarLz4:
		return archiver.TarLz4.Read(reader, outputDir)
	case CompressTypeTarSz:
		return archiver.TarSz.Read(reader, outputDir)
	}
	return nil
}

func isCompressTypeSupport(ct CompressType) error {
	if ct < CompressTypeZip || ct > CompressTypeTarSz {
		return errors.New(fmt.Sprintf("not support compress type: %d", ct))
	}
	return nil
}
