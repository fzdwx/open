package cons

import (
	"github.com/fzdwx/open/pkg/env"
	"github.com/gookit/goutil/testutil/assert"
	"testing"
)

func TestGetLogFileName(t *testing.T) {
	assert.Eq(t, "/tmp/fzdwx_open.log", GetLogFileName())
}

func TestGetLogFileNameCustom(t *testing.T) {
	env.Set(EnvLogFile, "qwe.log")
	assert.Eq(t, "/tmp/qwe.log", GetLogFileName())
}
