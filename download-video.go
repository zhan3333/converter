package converter

import (
	"fmt"
	"io"
	"net/url"

	_ "github.com/iawia002/lux/app"

	"converter/internal"
)

// DownloadVideo 下载视频
// todo 下载进度需要写到 writer 中
func DownloadVideo(videoURL string, writer io.Writer) error {
	if _, err := url.ParseRequestURI(videoURL); err != nil {
		return fmt.Errorf("%s 不是一个有效的地址: %w", videoURL, err)
	}

	if err := internal.Download(videoURL); err != nil {
		return fmt.Errorf("下载 %s 失败: %w", videoURL, err)
	}
	return nil
}
