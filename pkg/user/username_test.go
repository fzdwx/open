package user

import (
	"github.com/gookit/goutil/testutil/assert"
	"testing"
)

func Test_Username(t *testing.T) {
	assert.Eq(t, "fzdwx", Name())
}
