package converter

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
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
	err := DownloadVideo(url, writer)
	if assert.NoError(t, err) {
		t.Logf("output: %s", writer.String())
		assert.FileExists(t, testDownloadVideoName)
	}
}
