package history

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

var (
	homeDir string
)

const (
	historyFile = ".open_history"
)

func init() {
	current, err := user.Current()
	cobra.CheckErr(err)
	homeDir = current.HomeDir
}

func Append(url string) error {
	f, err := openHistory()
	if err != nil {
		return err
	}
	defer f.Close()

	data := fmt.Sprintf("%s|;|%d", url, time.Now().UnixNano())
	_, err = f.WriteString(data + "\n")
	if err != nil {
		return err
	}

	return nil
}

func openHistory() (*os.File, error) {
	filename := fmt.Sprintf("%s%s%s", homeDir, string(filepath.Separator), historyFile)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}

	return f, nil
}
