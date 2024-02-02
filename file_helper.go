package raselper

import (
	"bufio"
	"embed"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//go:embed files/*
var localFile embed.FS

type Matcher func(file os.FileInfo) bool
type FileNameHandler func(targetPath string) string
type TextReplacer func(bytes []byte) []byte

// CopyFiles 复制一个目录下的所有文件
func CopyFiles(sourcePath string, targetPath string, matcher Matcher, handler FileNameHandler, replaceText TextReplacer) error {
	sourcePath, _ = filepath.Abs(sourcePath)

	err := filepath.Walk(sourcePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			println(err.Error())
		}

		if matcher(info) {
			file, _ := os.Stat(path)
			if file.IsDir() {
				return nil
			}
			_targetPath := filepath.Join(targetPath, strings.Replace(path, sourcePath, "", 1))
			err := CopyFile(path, handler(_targetPath), replaceText)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// CopyFile 复制单个文件
func CopyFile(sourcePath string, targetPath string, replaceText TextReplacer) error {
	if sourcePath == targetPath {
		return errors.New("复制路径相同")
	}

	input, err := os.Open(sourcePath)
	reader := bufio.NewReader(input)
	_ = os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
	output, err := os.Create(targetPath)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	for {
		line, _ := reader.ReadBytes('\n')
		if replaceText != nil {
			line = replaceText(line)
		}
		_, err := output.Write(line)
		if err != nil {
			return err
		}
		if len(line) == 0 {
			break
		}
	}

	if err != nil {
		return err
	}

	err = input.Close()
	err = output.Close()
	if err != nil {
		return err
	}

	println("复制:", sourcePath, " 到:", targetPath)

	return nil
}

// InsertText 在文件中插入字符串
//
// from: 在from字符串之后的位置开始插入
//
// insert: 插入的字符串内容
func InsertText(filePath string, from string, insert string) error {
	tmpPath := filePath + ".tmp"
	err := CopyFile(filePath, tmpPath, func(bytes []byte) []byte {
		line := string(bytes)
		if strings.Contains(line, from) {
			line = strings.Replace(line, from, from+insert, -1)
		}
		return []byte(line)
	})
	if err != nil {
		return err
	}
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	err = os.Rename(tmpPath, filePath)
	if err != nil {
		return err
	}
	return nil
}

func CopyLocalFile(localFileName string, targetPath string) error {
	input, err := localFile.ReadFile(localFileName)
	_ = os.MkdirAll(filepath.Dir(targetPath), os.ModePerm)
	output, err := os.Create(targetPath)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	_, err = output.Write(input)
	if err != nil {
		return err
	}
	if input == nil {
		return errors.New(localFileName + "复制失败")
	}

	err = output.Close()
	if err != nil {
		return err
	}

	println("复制:", localFileName, " 到:", targetPath)

	return nil
}
