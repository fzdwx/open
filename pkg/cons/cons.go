package cons

import (
	"fmt"
	"github.com/fzdwx/open/pkg/env"
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

	Version = "v0.2.0"
)

func GetLogFileName() string {
	return fmt.Sprintf("%s%s%s", os.TempDir(), string(filepath.Separator), env.Or(EnvLogFile, logFilePrefix))
}
