package util

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	"os"
)

func CurrentDir() (fyne.ListableURI, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("get current dir path: %w", err)
	}
	d, err := storage.ListerForURI(storage.NewFileURI(pwd))
	if err != nil {
		return nil, fmt.Errorf("ListerForURI: %w", err)
	}
	return d, nil
}

func IsFileExists(file string) (bool, error) {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("stat %s: %w", file, err)
	}
	return true, nil
}

func Unique[T comparable](list []T) []T {
	var res []T
	var tmp = map[T]bool{}
	for _, v := range list {
		if !tmp[v] {
			res = append(res, v)
		}
		tmp[v] = true
	}
	if len(res) == 0 {
		return []T{}
	}
	return res
}
