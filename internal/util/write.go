package util

import (
	"github.com/fzdwx/open/internal/cons"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"os"
)

// AppendJson write json to file
func AppendJson(val interface{}, filepath string) error {
	data, err := jsonutil.Encode(val)

	if err != nil {
		return err
	}

	if err := cons.MkOpenDir(); err != nil {
		return err
	}

	data = append(data, '\n')
	return fsutil.WriteFile(filepath, data, os.ModePerm, fsutil.FsCWAFlags)
}
