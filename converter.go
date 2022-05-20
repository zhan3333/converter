package converter

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"io/ioutil"
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

func Run() {
	files, err := readDirFiles(".")
	if err != nil {
		color.Red("read dir files failed: %s", err.Error())
		return
	}
	files = filterVideoFiles(files)
	if len(files) == 0 {
		color.Yellow("当前目录未找到待转格式的视频")
		return
	}

	var files2 []string
	_ = survey.AskOne(&survey.MultiSelect{
		Message: "选择需要转码为 mp4 的视频",
		Options: files,
	}, &files2)

	if len(files2) == 0 {
		color.Yellow("未选择任何视频")
		return
	}

	color.White("选择了: %v", files2)
	color.White("转格式开始")
	for _, file := range files2 {
		outFile := getNoExistMP4Filename(file)
		if err = convertFile(file, outFile); err != nil {
			logrus.Errorf("convert file=%s: %s", file, err.Error())
			return
		}
		color.White("- 转换完成: %s -> %s", file, outFile)
	}
	color.Green("所有 %d 个视频转换完成", len(files2))
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
