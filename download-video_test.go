package converter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"

	"converter/internal/downloader"
)

var testDownloadVideoName = "我不想坐地铁上班啊啊啊啊啊啊啊啊啊.mp4"

func TestMain(m *testing.M) {
	_ = os.Remove(testDownloadVideoName)
	defer func() { _ = os.Remove(testDownloadVideoName) }()
	m.Run()
}

func TestDownloadVideo(t *testing.T) {
	writer := &strings.Builder{}
	url := "https://www.bilibili.com/video/BV1H34y1Z7mm"
	process := downloader.Process{
		BarWriter: writer,
		Title:     "",
		Site:      "",
		Type:      "",
	}
	err := DownloadVideo(url, &process)
	if assert.NoError(t, err) {
		t.Logf("title: %s", process.Title)
		t.Logf("site: %s", process.Site)
		t.Logf("type: %s", process.Type)
		t.Logf("quality: %s", process.Quality)
		t.Logf("size: %s", process.Size)
		t.Logf("output: %s", writer.String())
		assert.FileExists(t, testDownloadVideoName)
	}
}
