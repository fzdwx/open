package history

import (
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/user"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"os"
	"time"
)

type Model struct {
	Time     int64
	Url      string
	Username string
}

func Write(url string) error {
	data, err := jsonutil.Encode(&Model{
		Time:     time.Now().UnixMilli(),
		Url:      url,
		Username: user.Name(),
	})

	if err != nil {
		return err
	}

	if err := cons.MkOpenDir(); err != nil {
		return err
	}

	data = append(data, '\n')
	return fsutil.WriteFile(cons.HistoryFileName(), data, os.ModePerm, fsutil.FsCWAFlags)
}
