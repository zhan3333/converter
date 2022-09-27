package util

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"time"
)

type StatusText struct {
	Str       binding.String
	Widget    *widget.Label
	InProcess bool
}

func NewStatusText(str string) *StatusText {
	s := &StatusText{
		Str: binding.NewString(),
	}
	_ = s.Str.Set(str)
	s.Widget = widget.NewLabelWithData(s.Str)
	return s
}

func (t *StatusText) Set(s string) {
	t.InProcess = false
	_ = t.Str.Set("状态: " + s)
}

func (t *StatusText) SetInProcess(s string) {
	t.InProcess = true
	_ = t.Str.Set("状态: " + s)
	go func(t2 *StatusText, s2 string) {
		r := []string{
			"",
			".",
			"..",
			"...",
		}
		for t2.InProcess {
			ps := s2
			for _, v := range r {
				time.Sleep(500 * time.Millisecond)
				if !t2.InProcess {
					return
				}
				_ = t2.Str.Set(ps + v)
			}
		}
	}(t, s)
}
