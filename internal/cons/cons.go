package cons

import (
	"github.com/fzdwx/open/internal/env"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	// EnvGhToken env key with github token
	EnvGhToken = "GH_TOKEN"

	GithubUrl = "https://github.com"

	Version = "v1.3.1"

	HttpPrefix  = "http://"
	HttpsPrefix = "https://"
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

func HistoryFileName() string {
	return join("open", "history")
}

func LogFileName() string {
	return join("open", "log")
}

func AliasFileName() string {
	return join("open", "alias")
}

func PreviewCommand() string {
	return env.OrWithFunc("OPEN_PREVIEW", loopUpPreviewExec)
}

func loopUpPreviewExec() string {
	for _, name := range []string{"bat", "cat"} {
		if _, err := exec.LookPath(name); err == nil {
			return name
		}
	}
	return "cat"
}

func MkOpenDir() error {
	err := os.Mkdir(OpenDir(), os.ModePerm)
	if os.IsExist(err) {
		return nil
	}
	return err
}
