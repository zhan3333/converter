package converter

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/iawia002/lux/app"
)

func ActionDownload() {
	args := []string{"", ""}
	var videoURL string
	_ = survey.AskOne(&survey.Input{
		Message: "输入视频链接",
	}, &videoURL, survey.WithValidator(survey.Required))
	args[1] = videoURL

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
