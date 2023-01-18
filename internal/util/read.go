package util

import (
	"errors"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/strutil"
)

func Read(filepath string) ([]string, error) {
	bytes := fsutil.ReadExistFile(filepath)
	if bytes == nil {
		return nil, errors.New("alias file not found")
	}
	pairs := strutil.Split(string(bytes), "\n")
	if len(pairs) == 0 {
		return nil, errors.New("alias file is empty")
	}

	return pairs, nil
}
