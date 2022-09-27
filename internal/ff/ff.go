package ff

import (
	"converter/internal/util"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"strings"
)

type FF struct {
	log io.Writer
}

func NewFF(log io.Writer) *FF {
	return &FF{log: log}
}

func (f FF) Convert(in string, out string, overwrite bool) error {
	// 本地文件才检查是否存在
	if !strings.HasPrefix(in, "http") {
		exist, _ := util.IsFileExists(in)
		if !exist {
			return fmt.Errorf("file %s not exist", in)
		}
	}

	s := ffmpeg.Input(in).
		Output(out).
		WithOutput(f.log, f.log)
	if overwrite {
		s.OverWriteOutput()
	}
	if err := s.Run(); err != nil {
		return fmt.Errorf("run ff convert: %w", err)
	}
	return nil
}
