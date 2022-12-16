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

	Version = "v0.7"

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

func LogPreview() string {
	return env.OrWithFunc("OPEN_LOG_PREVIEW", loopUpLogPreviewExec)
}

func loopUpLogPreviewExec() string {
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
