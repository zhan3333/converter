package converter

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var FfmpegWin = "dependencies/ffmpeg-win"
var FfmpegMac = "dependencies/ffmpeg-mac"
var Font = "dependencies/fonts/微软雅黑.ttf"

func init() {
	// 编译前在 build 目录
	if e, _ := isDirExists(FfmpegWin); !e {
		FfmpegWin = "build/" + FfmpegWin
	}
	if e, _ := isDirExists(FfmpegMac); !e {
		FfmpegMac = "build/" + FfmpegMac
	}
	if e, _ := isFileExists(Font); !e {
		Font = "build/" + Font
	}
}

func LoadFont() error {
	d, _ := os.Getwd()
	err := os.Setenv("FYNE_FONT", filepath.Join(d, Font))
	if err != nil {
		return fmt.Errorf("set windows PATH: %w", err)
	}
	return nil
}

// Windows 在 windows 下运行时，设置 ffmpeg 可执行文件路径
// W:\ffmpeg\bin\
func Windows() error {
	d, _ := os.Getwd()
	p := strings.Join([]string{os.Getenv("PATH"), filepath.Join(d, FfmpegWin, "bin")}, ";")
	err := os.Setenv("PATH", p)
	if err != nil {
		return fmt.Errorf("set windows PATH: %w", err)
	}
	v, err := getFfmpegVersion()
	if err != nil {
		return err
	}
	color.White(v)
	return nil
}

var rg = regexp.MustCompile("ffmpeg version (.*) Copyright")

// Mac mac 下的可执行文件
func Mac() error {
	d, _ := os.Getwd()
	p := fmt.Sprintf("%s:%s", os.Getenv("PATH"), filepath.Join(d, FfmpegMac))
	err := os.Setenv("PATH", p)

	if err != nil {
		return fmt.Errorf("set windows PATH: %w", err)
	}

	v, err := getFfmpegVersion()
	if err != nil {
		return err
	}
	color.White(v)
	return nil
}

func getFfmpegVersion() (string, error) {
	c := exec.Command("ffmpeg", "-version")
	if out, err := c.Output(); err != nil {
		color.White("PATH: %s", os.Getenv("PATH"))
		return "", fmt.Errorf("exec ffmpeg -version: %w", err)
	} else {
		if rg.Match(out) {
			return strings.TrimSuffix(rg.FindString(string(out)), " Copyright"), nil
		} else {
			return "", fmt.Errorf("未查找到 ffmpeg 可执行文件: \n%s", string(out))
		}
	}
}

func isDirExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Setup() error {
	var err error
	color.White("run in " + runtime.GOOS + " " + runtime.GOARCH)
	if runtime.GOOS == "windows" {
		err = Windows()
	}
	if runtime.GOOS == "darwin" {
		err = Mac()
	}
	if runtime.GOOS == "linux" {
		err = fmt.Errorf("暂时不支持 %s %s", runtime.GOOS, runtime.GOARCH)
	}
	return err
}
