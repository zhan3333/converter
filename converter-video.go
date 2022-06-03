package converter

import (
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var ffmpegLog io.Writer

func init() {
	var err error
	ffmpegLog, err = os.OpenFile("logs/ffmpeg.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic("init ffmpeg log file failed: " + err.Error())
	}
}

func ConvertVideo(file string, outputExt string, toDir string, write io.Writer) (string, error) {
	outFile := getNoExistFilename(filepath.Join(toDir, filepath.Base(file)), outputExt)
	if err := convertFile(file, outFile); err != nil {
		return "", fmt.Errorf("convert file=%s: %w", file, err)
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

func convertFile(inFile string, outFile string) error {
	exist, _ := isFileExists(inFile)
	if !exist {
		return fmt.Errorf("file %s not exist", inFile)
	}
	err := ffmpeg.Input(inFile).
		Output(outFile).
		OverWriteOutput().
		WithOutput(ffmpegLog, ffmpegLog).
		Run()
	return err
}

// 获取一个不存在的 mp4 文件名
func getNoExistFilename(file string, outputExt string) string {
	var index int
	for {
		newFile := getMP4FileName(file, index, outputExt)
		if exist, _ := isFileExists(newFile); exist {
			index++
		} else {
			return newFile
		}
	}
}

func getMP4FileName(file string, index int, outputExt string) string {
	if index == 0 {
		return strings.TrimSuffix(file, filepath.Ext(file)) + outputExt
	}
	return fmt.Sprintf("%s-%d%s", strings.TrimSuffix(file, filepath.Ext(file)), index, outputExt)
}
