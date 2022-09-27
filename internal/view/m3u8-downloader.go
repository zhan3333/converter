package view

import (
	"context"
	"converter"
	"converter/internal/util"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

func M3u8DownloaderView(w fyne.Window, convert *converter.VideoConvert) (box *fyne.Container) {
	var selectExt = ".mp4"

	box = container.NewVBox()
	input := widget.NewEntry()
	input.SetPlaceHolder("输入视频链接，例如 https://www.zxx.edu.cn/syncClassroom/classActivity?activityId=66ce497d-9777-11ec-92ef-246e9675e50c")
	input.Text = "https://www.zxx.edu.cn/syncClassroom/classActivity?activityId=66ce497d-9777-11ec-92ef-246e9675e50c"
	input.MultiLine = true
	input.Validator = func(s string) error {
		if _, err := url.ParseRequestURI(s); err != nil {
			return fmt.Errorf("%s 不是一个有效的地址: %s", s, err.Error())
		}
		return nil
	}
	statusText := util.NewStatusText("填写链接后按开始按钮下载")

	downloadDir, err := converter.GetCurrentDir()
	if err != nil {
		dialog.ShowError(err, w)
		return
	}
	openDirButton := widget.NewButton("打开下载目录", func() {
		converter.OpenSystemDir(downloadDir)
	})
	var selectM3u8Url = ""
	var downloadButton *widget.Button
	downloadButton = widget.NewButton("开始下载", func() {
		defer box.Refresh()
		downloadButton.Disable()
		defer downloadButton.Enable()
		if selectM3u8Url == "" {
			dialog.ShowError(errors.New("请选择 m3u8 链接"), w)
			return
		}
		fmt.Println("开始下载: " + selectM3u8Url)
		statusText.SetInProcess("下载中")
		if out, err := convert.Convert(selectM3u8Url, selectExt, downloadDir); err != nil {
			fmt.Println("下载失败: ", err.Error())
			statusText.Set("下载失败: " + err.Error())
			return
		} else {
			fmt.Println("下载完成")
			statusText.Set("下载完成: " + out)
		}
	})
	downloadButton.Disable() // 需要选择了 m3u8 才能下载
	extList := widget.NewSelect([]string{".mp4", ".mp3"}, func(s string) {
		selectExt = s
	})
	parseList := widget.NewSelect([]string{}, func(s string) {
		selectM3u8Url = s
		downloadButton.Enable()
	})
	parseList.Disable()
	var parseM3u8Button *widget.Button
	parseM3u8Button = widget.NewButton("嗅探", func() {
		defer box.Refresh()
		videoURL := input.Text
		if _, err := url.ParseRequestURI(videoURL); err != nil {
			dialog.ShowError(fmt.Errorf("%s 不是一个有效的地址: %s", videoURL, err.Error()), w)
			return
		}
		fmt.Println("开始嗅探: " + input.Text)
		list, err := SniffM3U8List(videoURL)
		if err != nil {
			dialog.ShowError(fmt.Errorf("%s m3u8 嗅探失败: %w", videoURL, err), w)
			return
		}
		fmt.Println("开始结束")

		fmt.Printf("嗅探到的地址: %+v", list)
		parseList.Options = list
		if len(list) == 0 {
			parseList.Disable()
			dialog.ShowError(fmt.Errorf("%s 未嗅探到 m3u8 链接", videoURL), w)
			return
		} else {
			parseList.Enable()
		}
	})

	box.Add(input)
	box.Add(container.NewGridWithColumns(2, widget.NewLabel("选择 m3u8"), parseList))
	box.Add(container.NewGridWithColumns(2, widget.NewLabel("选择导出格式"), extList))
	box.Add(container.NewGridWithColumns(3, parseM3u8Button, downloadButton, openDirButton))
	box.Add(statusText.Widget)

	return
}

func SniffM3U8List(url string) ([]string, error) {
	dir, err := ioutil.TempDir("", "chromedp-example")
	if err != nil {
		panic(err)
	}
	defer func() { _ = os.RemoveAll(dir) }()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("headless", true),
		chromedp.Flag("ignore-certificate-errors", true),
		chromedp.Flag("window-size", "50,400"),
		chromedp.UserDataDir(dir),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	taskCtx, cancel = context.WithTimeout(taskCtx, 10*time.Second)
	defer cancel()

	// ensure that the browser process is started
	if err := chromedp.Run(taskCtx); err != nil {
		return []string{}, fmt.Errorf("check run chromedp error: %w", err)
	}

	// listen network event
	var m3u8List []string
	listenForNetworkEvent(taskCtx, &m3u8List)

	err = chromedp.Run(taskCtx,
		network.Enable(),
		chromedp.Navigate(url),
		//chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.WaitVisible(`video`, chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Printf("body is visible")
			return nil
		}),
	)
	if err != nil {
		return []string{}, fmt.Errorf("run chromedp error: %w", err)
	}
	return util.Unique(m3u8List), nil
}

//监听
func listenForNetworkEvent(ctx context.Context, list *[]string) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		switch ev := ev.(type) {
		case *network.EventRequestWillBeSent:
			req := ev.Request
			if strings.Contains(req.URL, ".m3u8") {
				*list = append(*list, req.URL)
			}
		}
	})
}
