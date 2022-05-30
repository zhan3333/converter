package converter

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// ReadDirFiles 读取目录下所有文件
func ReadDirFiles(dir string) ([]string, error) {
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

// FilterVideoFiles 过滤出所有视频文件
func FilterVideoFiles(files []string) []string {
	var res []string
	for _, f := range files {
		ext := filepath.Ext(f)
		if SupportVideoExtensions[ext] {
			res = append(res, f)
		}
	}
	if len(res) == 0 {
		return []string{}
	}
	return res
}

var SupportVideoExtensions = map[string]bool{
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

func GetSupportVideoExtensions() []string {
	var ret []string
	for k := range SupportVideoExtensions {
		ret = append(ret, k)
	}
	return ret
}

func OpenSystemDir(dir string) {
	var err error
	switch GetOS() {
	case "windows":
		err = exec.Command("start", dir).Run()
	case "darwin":
		err = exec.Command("open", dir).Run()
	case "linux":
	}
	if err != nil {
		fmt.Printf("open dir=%s error: %s\n", dir, err.Error())
	}
}

func GetOS() string {
	return runtime.GOOS
}

func GetCurrentDir() (string, error) {
	d, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("获取当前文件夹路径失败: %w", err)
	}
	return d, nil
}

func PrintCard() {
	fmt.Println()
	fmt.Println("to: chen")
	fmt.Println()
	fmt.Println("    昨夜有繁星满天，")
	fmt.Println("    今早有朝霞渐起。")
	fmt.Println("    你看见也好，")
	fmt.Println("    看不见也没关系，")
	fmt.Println("    我找到你，")
	fmt.Println("    它们才有意义。")
	fmt.Println()
	fmt.Println("              from: zhan")
	fmt.Println()
}
