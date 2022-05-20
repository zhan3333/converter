package converter

import (
	"fmt"
	"github.com/sirupsen/logrus"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Run() {
	files, err := readDirFiles(".")
	if err != nil {
		logrus.Errorf("read dir files faield: %s", err.Error())
		return
	}
	files = filterVideoFiles(files)
	if len(files) == 0 {
		logrus.Warn("当前目录未找到待转格式的视频")
		return
	}
	logrus.Infof("找到了这些待转格式的视频: %v", files)
	logrus.Info("转格式开始")
	for _, file := range files {
		outFile := getNoExistMP4Filename(file)
		if err = convertFile(file, outFile); err != nil {
			logrus.Errorf("convert file=%s: %s", file, err.Error())
			return
		}
		logrus.Infof("转换一个文件完成: %s -> %s", file, outFile)
	}
	logrus.Info("所有文件转换结束")
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
		ErrorToStdOut().
		Run()
	return err
}

// 获取一个不存在的 mp4 文件名
func getNoExistMP4Filename(file string) string {
	var index int
	for {
		newFile := getMP4FileName(file, index)
		if exist, _ := isFileExists(newFile); exist {
			index++
		} else {
			return newFile
		}
	}
}

func getMP4FileName(file string, index int) string {
	if index == 0 {
		return strings.TrimSuffix(file, filepath.Ext(file)) + ".mp4"
	}
	return fmt.Sprintf("%s-%d.mp4", strings.TrimSuffix(file, filepath.Ext(file)), index)
}

// 读取目录下所有文件
func readDirFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read dir %s: %d", dir, err)
	}
	var fs []string
	for _, v := range files {
		fs = append(fs, v.Name())
	}
	if len(fs) == 0 {
		return []string{}, nil
	}
	return fs, nil
}

var supportVideoExtensions = map[string]bool{
	".m4v":  true,
	".avi":  true,
	".mov":  true,
	".qt":   true,
	".flv":  true,
	".wmv":  true,
	".asf":  true,
	".mpeg": true,
	".mpg":  true,
	".vob":  true,
	".mkv":  true,
	".rm":   true,
	".rmvb": true,
	".dat":  true,
	".ogg":  true,
	".ts":   true,
}

// 过滤出所有视频文件
func filterVideoFiles(files []string) []string {
	var res []string
	for _, f := range files {
		ext := filepath.Ext(f)
		if supportVideoExtensions[ext] {
			res = append(res, f)
		}
	}
	if len(res) == 0 {
		return []string{}
	}
	return res
}

func delOutFiles(files []string) error {
	for _, file := range files {
		if err := os.Remove(file); err != nil {
			if !os.IsNotExist(err) {
				return fmt.Errorf("del %s: %w", file, err)
			}
		}
	}
	return nil
}
