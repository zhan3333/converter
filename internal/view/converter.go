package view

import (
	"converter"
	"converter/internal/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pkg/errors"
	"log"
	"path/filepath"
)

func ConverterView(w fyne.Window, convert *converter.VideoConvert) (box *fyne.Container) {
	box = container.NewVBox()
	var (
		selectFile     string
		selectFileName = binding.NewString()
		outputExt      = binding.NewString()
		outputFile     string
		outputFileName = binding.NewString()
		statusText     = util.NewStatusText("未开始")
	)
	downloadDir, err := converter.GetCurrentDir()
	if err != nil {
		dialog.ShowError(err, w)
		return
	}

	box.Add(container.NewGridWithColumns(2,
		// 选择文件按钮
		widget.NewButton("选择文件", func() {
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
			}, w)
			fd.SetFilter(storage.NewExtensionFileFilter(converter.GetSupportVideoExtensions()))
			d, err := util.CurrentDir()
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			fd.SetLocation(d)
			fd.Show()
		}),
		// 选择输出格式
		func() *widget.Select {
			s := widget.NewSelect(converter.GetSupportVideoExtensions(), func(s string) {
				_ = outputExt.Set(s)
			})
			s.PlaceHolder = "选择输出格式"
			return s
		}(),
	))

	box.Add(container.NewGridWithColumns(2,
		widget.NewLabel("输入文件: "),
		widget.NewLabelWithData(selectFileName)),
	)
	box.Add(container.NewGridWithColumns(2,
		widget.NewLabel("输出格式: "),
		widget.NewLabelWithData(outputExt)),
	)
	box.Add(container.NewGridWithColumns(2,
		// 开始转换按钮
		widget.NewButton("开始转换", func() {
			if selectFile == "" {
				dialog.ShowError(errors.New("请选择输入文件"), w)
				return
			}
			if ext, _ := outputExt.Get(); ext == "" {
				dialog.ShowError(errors.New("请选择输出格式"), w)
				return
			}
			statusText.SetInProcess("转换中")
			ext, _ := outputExt.Get()
			outputFile, err = convert.Convert(selectFile, ext, downloadDir)
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
		}),
		// 打开文件夹
		widget.NewButton("打开结果目录", func() {
			converter.OpenSystemDir(downloadDir)
		}),
	))
	box.Add(container.NewGridWithColumns(2,
		widget.NewLabel("输出文件: "),
		widget.NewLabelWithData(outputFileName)),
	)
	box.Add(container.NewGridWrap(fyne.NewSize(50, 20),
		widget.NewLabel("状态："),
		statusText.Widget,
	))

	return
}
