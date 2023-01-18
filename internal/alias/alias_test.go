package alias

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	err := Write("https://fzdwx.github.io/", "blog")
	if err != nil {
		t.Error(err)
	}
}

func TestRead(t *testing.T) {
	alias, err := Read()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(alias)
}
