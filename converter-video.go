package converter

import (
	"converter/internal/ff"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type VideoConvert struct {
	FF       *ff.FF
	Progress io.Writer
}

func NewVideoConvert(FF *ff.FF, progress io.Writer) *VideoConvert {
	return &VideoConvert{FF: FF, Progress: progress}
}

func (c VideoConvert) Convert(in string, outExt string, saveDir string) (string, error) {
	outFile := c.getNoExistFilename(filepath.Join(saveDir, filepath.Base(in)), outExt)
	if err := c.FF.Convert(in, outFile, true); err != nil {
		return "", fmt.Errorf("convert file=%s: %w", in, err)
	}
	return outFile, nil
}

func isFileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("stat %s: %w", file, err)
	}
	return true, nil
}

// 获取一个不存在的 mp4 文件名
func (c VideoConvert) getNoExistFilename(file string, outputExt string) string {
	var index int
	for {
		newFile := c.getMP4FileName(file, index, outputExt)
		if exist, _ := isFileExists(newFile); exist {
			index++
		} else {
			return newFile
		}
	}
}

func (c VideoConvert) getMP4FileName(file string, index int, outputExt string) string {
	if index == 0 {
		return strings.TrimSuffix(file, filepath.Ext(file)) + outputExt
	}
	return fmt.Sprintf("%s-%d%s", strings.TrimSuffix(file, filepath.Ext(file)), index, outputExt)
}
