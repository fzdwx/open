package cons

import (
	"github.com/MarvinJWendt/testza"
	"github.com/fzdwx/open/pkg/env"
	"testing"
)

func TestGetLogFileName(t *testing.T) {
	testza.AssertEqual(t, "/tmp/fzdwx_open.log", GetLogFileName())
}

func TestGetLogFileNameCustom(t *testing.T) {
	env.Set(EnvLogFile, "qwe.log")
	testza.AssertEqual(t, "/tmp/qwe.log", GetLogFileName())
}
