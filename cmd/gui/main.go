package main

import (
	"converter"
	"converter/internal/ff"
	"converter/internal/view"
	"fmt"
	"os"
)

func init() {
	if err := converter.LoadFont(); err != nil {
		fontPath, _ := os.LookupEnv("FYNE_FONT")
		fmt.Println("字体路径: " + fontPath)
		fmt.Println("加载字体失败: " + err.Error())
	}
}

func main() {
	converter.PrintCard()
	if err := converter.Setup(); err != nil {
		fmt.Println("程序环境准备失败: %s" + err.Error())
	}
	ffmpegLog, err := os.OpenFile("logs/ffmpeg.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic("init ffmpeg log file failed: " + err.Error())
	}
	ffm := ff.NewFF(ffmpegLog)
	videoConvert := converter.NewVideoConvert(ffm, os.Stderr)
	// 显示窗口
	view.WindowView(videoConvert).ShowAndRun()
}
