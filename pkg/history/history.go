package history

import (
	"fmt"
	"github.com/fzdwx/open/pkg/cons"
	"github.com/fzdwx/open/pkg/user"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"os"
	"time"
)

// FileName get open history file name
func FileName() string {
	return fmt.Sprintf("%s%s", cons.UserDir(), cons.HistoryFile)
}

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
	return fsutil.WriteFile(FileName(), data, os.ModePerm, fsutil.FsCWAFlags)
}
