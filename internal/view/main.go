package view

import (
	"converter"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"time"
)

func WindowView(convert *converter.VideoConvert) fyne.Window {
	a := app.NewWithID("convert")
	w := a.NewWindow("Converter")
	tabs := container.NewAppTabs(
		container.NewTabItem("下载网络视频", DownloaderView(w)),
		container.NewTabItem("视频转换格式", ConverterView(w, convert)),
		container.NewTabItem("m3u8 视频下载", M3u8DownloaderView(w, convert)),
	)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(1080, 460))
	// 异步检查一下环境
	go func() {
		time.Sleep(500 * time.Millisecond)
		if err := converter.Setup(); err != nil {
			dialog.ShowError(err, w)
		}
	}()
	return w
}
