package converter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var FfmpegWin = "dependencies/ffmpeg-win"

func init() {
	// 编译前在 build 目录
	if e, _ := isDirExists(FfmpegWin); !e {
		FfmpegWin = "build/" + FfmpegWin
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
	fmt.Println("版本: " + v)
	return nil
}

var rg = regexp.MustCompile("ffmpeg version (.*) Copyright")

// Mac mac 下的可执行文件
// 现在不做处理，因为本地安装了 ffmpeg
func Mac() error {
	v, err := getFfmpegVersion()
	if err != nil {
		return err
	}
	fmt.Println("版本: " + v)
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
