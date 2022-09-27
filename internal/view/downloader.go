package view

import (
	"converter"
	"converter/internal/downloader"
	"converter/internal/util"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"net/url"
	"strings"
	"time"
)

func DownloaderView(w fyne.Window) (box *fyne.Container) {
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
	statusText := util.NewStatusText("填写视频链接后按开始按钮下载")
	var (
		title       = binding.NewString()
		site        = binding.NewString()
		quality     = binding.NewString()
		size        = binding.NewString()
		type1       = binding.NewString()
		processLine = binding.NewString()
	)

	processWrite := &strings.Builder{}
	process := &downloader.Process{
		BarWriter: processWrite,
	}
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
		inProcess := true
		defer func() {
			time.Sleep(200 * time.Millisecond)
			inProcess = false
		}()
		go func() {
			for inProcess {
				_ = title.Set("标题: " + process.Title)
				_ = site.Set("站点: " + process.Site)
				_ = type1.Set("类型: " + process.Type)
				_ = quality.Set("质量: " + process.Quality)
				_ = size.Set("体积: " + process.Size)
				_ = processLine.Set(process.LastProcessLine())
				time.Sleep(100 * time.Millisecond)
			}
		}()
		if err := converter.DownloadVideo(videoURL, process); err != nil {
			statusText.Set("下载失败: " + err.Error())
			return
		}
		statusText.Set("下载完成")
	}

	box.Add(input)
	box.Add(container.NewGridWithColumns(2, button, openDirButton))
	box.Add(statusText.Widget)
	box.Add(widget.NewLabelWithData(title))
	box.Add(widget.NewLabelWithData(type1))
	box.Add(widget.NewLabelWithData(site))
	box.Add(widget.NewLabelWithData(size))
	box.Add(widget.NewLabelWithData(quality))
	box.Add(widget.NewLabelWithData(processLine))

	return
}
