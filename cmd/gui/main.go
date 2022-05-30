package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"converter"
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
	a := app.NewWithID("convert")
	w := a.NewWindow("Converter")
	tabs := container.NewAppTabs(
		container.NewTabItem("下载网络视频", downloadVideoBox(w)),
		container.NewTabItem("视频转 mp4 格式", convertVideoBox(w)),
	)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(640, 460))
	// 异步检查一下环境
	go func() {
		time.Sleep(500 * time.Millisecond)
		if err := converter.Setup(); err != nil {
			dialog.ShowError(err, w)
		}
	}()
	w.ShowAndRun()
}

func downloadVideoBox(w fyne.Window) (box *fyne.Container) {
	box = container.NewVBox()
	input := widget.NewEntry()
	input.SetPlaceHolder("输入视频链接，例如 https://www.bilibili.com/video/BV1B34y1j73J")
	input.Text = "https://www.bilibili.com/video/BV1B34y1j73J"
	input.MultiLine = true
	input.Validator = func(s string) error {
		if _, err := url.ParseRequestURI(s); err != nil {
			return fmt.Errorf("%s 不是一个有效的地址: %s", s, err.Error())
		}
		return nil
	}
	statusText := NewStatusText("填写视频链接后按开始按钮下载")
	processOutput := widget.NewLabel("")
	processWrite := &strings.Builder{}
	downloadDir, err := converter.GetCurrentDir()
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	openDirButton := widget.NewButton("打开下载目录", func() {
		converter.OpenSystemDir(downloadDir)
	})
	button := widget.NewButton("开始下载", nil)
	button.OnTapped = func() {
		defer box.Refresh()
		videoURL := input.Text
		if _, err := url.ParseRequestURI(videoURL); err != nil {
			dialog.ShowError(fmt.Errorf("%s 不是一个有效的地址: %s", videoURL, err.Error()), w)
			return
		}
		statusText.SetInProcess("下载中")
		button.Disable()
		defer button.Enable()
		// 循环渲染结果到屏幕
		end := false
		defer func() { end = true }()
		go func() {
			for !end {
				processOutput.Text = processWrite.String()
				processOutput.Refresh()
				time.Sleep(100 * time.Millisecond)
			}
		}()
		if err := converter.DownloadVideo(videoURL, processWrite); err != nil {
			statusText.Set("下载失败: " + err.Error())
			return
		}
		statusText.Set("下载完成")
	}

	box.Add(statusText.Widget)
	box.Add(processOutput)

	box.Add(input)
	box.Add(container.NewGridWithColumns(2, button, openDirButton))

	return
}

func convertVideoBox(w fyne.Window) (box *fyne.Container) {
	box = container.NewVBox()
	var (
		selectFile     string
		selectFileName = binding.NewString()
		outputFile     string
		outputFileName = binding.NewString()
		statusText     = NewStatusText("选择视频文件进行格式转换")
	)
	downloadDir, err := converter.GetCurrentDir()
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	openDirButton := widget.NewButton("打开结果目录", func() {
		converter.OpenSystemDir(downloadDir)
	})
	selectFileButton := widget.NewButton("选择待转换格式的视频文件", func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			selectFile = reader.URI().Path()
			_ = selectFileName.Set(filepath.Base(selectFile))
			processWriter := &strings.Builder{}
			statusText.SetInProcess("转换中")
			outputFile, err = converter.ConvertVideo(reader.URI().Path(), downloadDir, processWriter)
			if err != nil {
				statusText.Set("转换失败: " + err.Error())
				dialog.ShowError(err, w)
			} else {
				statusText.Set("转换完成")
				_ = outputFileName.Set(filepath.Base(outputFile))
				dialog.ShowForm("转格式完成", "打开文件夹", "关闭", nil, func(b bool) {
					if b {
						converter.OpenSystemDir(downloadDir)
					}
				}, w)
			}
		}, w)
		fd.SetFilter(storage.NewExtensionFileFilter(converter.GetSupportVideoExtensions()))
		d, err := CurrentDir()
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		fd.SetLocation(d)
		fd.Show()
	})

	box.Add(container.NewGridWithColumns(2,
		widget.NewLabel("输入文件: "),
		widget.NewLabelWithData(selectFileName)),
	)
	box.Add(container.NewGridWithColumns(2,
		widget.NewLabel("输出文件: "),
		widget.NewLabelWithData(outputFileName)),
	)
	box.Add(statusText.Widget)
	box.Add(container.NewGridWithColumns(2, selectFileButton, openDirButton))

	return
}

func CurrentDir() (fyne.ListableURI, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get current dir path: %w", err)
	}
	d, err := storage.ListerForURI(storage.NewFileURI(pwd))
	if err != nil {
		return nil, fmt.Errorf("ListerForURI: %w", err)
	}
	return d, nil
}

type StatusText struct {
	Str       binding.String
	Widget    *widget.Label
	InProcess bool
}

func NewStatusText(str string) *StatusText {
	s := &StatusText{
		Str: binding.NewString(),
	}
	_ = s.Str.Set(str)
	s.Widget = widget.NewLabelWithData(s.Str)
	return s
}

func (t *StatusText) Set(s string) {
	t.InProcess = false
	_ = t.Str.Set(s)
}

func (t *StatusText) SetInProcess(s string) {
	t.InProcess = true
	_ = t.Str.Set(s)
	go func(t2 *StatusText, s2 string) {
		r := []string{
			"",
			".",
			"..",
			"...",
		}
		for t2.InProcess {
			ps := s2
			for _, v := range r {
				time.Sleep(500 * time.Millisecond)
				if !t2.InProcess {
					return
				}
				_ = t2.Str.Set(ps + v)
			}
		}
	}(t, s)
}
