package history

import (
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/user"
	"github.com/fzdwx/open/internal/util"
	"time"
)

type Model struct {
	Time     int64
	Url      string
	Username string
}

func Write(url string) error {
	return util.AppendJson(&Model{
		Time:     time.Now().UnixMilli(),
		Url:      url,
		Username: user.Name(),
	}, cons.HistoryFileName())
}
