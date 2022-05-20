package converter

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var FfmpegWin = "dependencies/ffmpeg-win"
var FfmpegMac = "dependencies/ffmpeg-mac"

func init() {
	// 编译前在 build 目录
	if e, _ := isDirExists(FfmpegWin); !e {
		FfmpegWin = "build/" + FfmpegWin
	}
	if e, _ := isDirExists(FfmpegMac); !e {
		FfmpegMac = "build/" + FfmpegMac
	}
}

// Windows 在 windows 下运行时，设置 ffmpeg 可执行文件路径
// W:\ffmpeg\bin\
func Windows() error {
	d, _ := os.Getwd()
	p := os.Getenv("PATH")
	p = p + filepath.Join(d, FfmpegWin, "bin") + `\;`
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
