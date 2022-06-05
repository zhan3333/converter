package internal

import (
	"github.com/iawia002/lux/extractors"

	"converter/internal/downloader"
)

func Download(videoURL string, process *downloader.Process) error {
	data, err := extractors.Extract(videoURL, extractors.Options{})
	if err != nil {
		// if this error occurs, it means that an error occurred before actually starting to extract data
		// (there is an error in the preparation step), and the data list is empty.
		return err
	}

	defaultDownloader := downloader.New(downloader.Options{})
	errors := make([]error, 0)
	for _, item := range data {
		if item.Err != nil {
			// if this error occurs, the preparation step is normal, but the data extraction is wrong.
			// the data is an empty struct.
			errors = append(errors, item.Err)
			continue
		}
		if err = defaultDownloader.Download(item, process); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) != 0 {
		return errors[0]
	}
	return nil
}
