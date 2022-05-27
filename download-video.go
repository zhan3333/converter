package converter

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/iawia002/lux/app"
	"io"
	"net/url"
)

// DownloadVideo 下载视频
// todo 下载进度需要写到 writer 中
func DownloadVideo(videoURL string, writer io.Writer) error {
	if _, err := url.ParseRequestURI(videoURL); err != nil {
		return fmt.Errorf("%s 不是一个有效的地址: %w", videoURL, err)
	}
	args := []string{"", videoURL}

	a := app.New()
	if err := a.Run(args); err != nil {
		return fmt.Errorf("下载 %s 失败: %w", videoURL, err)
	}
	return nil
}

func ActionDownload() {
	args := []string{"", ""}
	var videoURL string
	_ = survey.AskOne(&survey.Input{
		Message: "输入视频链接",
	}, &videoURL, survey.WithValidator(survey.Required))
	args[1] = videoURL

	if _, err := url.ParseRequestURI(videoURL); err != nil {
		color.Red("%s 不是一个有效的地址: %s", videoURL, err.Error())
		return
	}

	color.Green("开始下载")
	if err := app.New().Run(args); err != nil {
		fmt.Fprintf(
			color.Output,
			"Run %s failed: %s\n",
			color.CyanString("%s", app.Name), color.RedString("%v", err),
		)
	} else {
		color.Green("下载完成")
	}
}
