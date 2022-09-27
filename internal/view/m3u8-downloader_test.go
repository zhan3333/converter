package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSniffM3U8List(t *testing.T) {
	url := "https://www.zxx.edu.cn/syncClassroom/classActivity?activityId=66ce497d-9777-11ec-92ef-246e9675e50c"
	list, err := SniffM3U8List(url)
	assert.NoError(t, err)
	t.Logf("list: %+v", list)
	assert.Equal(t, 1, len(list))

}
