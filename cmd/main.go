package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"

	"converter"
)

func main() {
	// 昨夜有繁星满天，今早有朝霞渐起。 你看见也好，看不见也没关系， 我找到你，它们才有意义。
	logrus.Info("程序开始运行")
	printCard()
	defer func() {
		logrus.Info("程序运行结束")
		fmt.Println("请按回车键退出程序，或直接关闭窗口")
		fmt.Scanf("a")
	}()
	var err error
	fmt.Println("run in " + runtime.GOOS + " " + runtime.GOARCH)
	if runtime.GOOS == "windows" {
		err = converter.Windows()
	}
	if runtime.GOOS == "darwin" {
		err = converter.Mac()
	}
	if runtime.GOOS == "linux" {
		err = fmt.Errorf("暂时不支持 %s %s", runtime.GOOS, runtime.GOARCH)
	}
	if err != nil {
		fmt.Println("发生错误: " + err.Error())
		return
	}

	converter.Run()
}

func printCard() {
	fmt.Println()
	fmt.Println("to: 老婆")
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

func init() {
	initLog()
}

func initLog() {
	w1 := os.Stdout
	w2, err := os.OpenFile("logs/log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("init log.txt failed: ", err.Error())
		logrus.SetOutput(io.MultiWriter(w1))
	} else {
		logrus.SetOutput(io.MultiWriter(w1, w2))
	}
}
