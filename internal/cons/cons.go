package cons

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

const (
	// EnvGhToken env key with github token
	EnvGhToken = "GH_TOKEN"

	GithubUrl = "https://github.com"

	Version = "v0.5"

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

// ~/.config/fzdwx/path
func join(path ...string) string {
	return filepath.Join(getUserDir(), ".config", "fzdwx", filepath.Join(path...))
}

func OpenDir() string {
	return join("open")
}

func HistoryFile() string {
	return join("open", "history")
}

func LogFileName() string {
	return join("open", "log")
}

func MkOpenDir() error {
	err := os.Mkdir(OpenDir(), os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}
