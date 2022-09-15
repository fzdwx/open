package user

import (
	"github.com/MarvinJWendt/testza"
	"testing"
)

func Test_Username(t *testing.T) {
	testza.AssertEqual(t, "fzdwx", Name())
}
