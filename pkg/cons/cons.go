package cons

import (
	"fmt"
	"github.com/fzdwx/open/pkg/env"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const (
	// EnvGhToken env key with github token
	EnvGhToken = "GH_TOKEN"
	// EnvLogFile env key wtih log file
	//
	// default /tmp/fzdwx_open.log
	EnvLogFile    = "OPEN_LOG_FILE"
	logFilePrefix = "fzdwx_open.log"

	GithubUrl = "https://github.com"

	Version = "v0.4"

	HistoryFile = "/.fzdwx_open/history"

	openDir = "/.fzdwx_open"

	HttpPrefix  = "http://"
	HttpsPrefix = "https://"
)

var (
	userDir = getUserDir()
)

func getUserDir() string {
	dir, err := os.UserHomeDir()
	cobra.CheckErr(err)
	return dir
}

// UserDir get user home
func UserDir() string {
	return userDir
}

func OpenDir() string {
	return fmt.Sprintf("%s%s", UserDir(), openDir)
}

func MkOpenDir() error {
	err := os.Mkdir(OpenDir(), os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func GetLogFileName() string {
	return fmt.Sprintf("%s%s%s", os.TempDir(), string(filepath.Separator), env.Or(EnvLogFile, logFilePrefix))
}
