package cons

import (
	"fmt"
	"testing"
)

func TestUserDir(t *testing.T) {
	fmt.Println(getUserDir())
}

func TestOpenDir(t *testing.T) {
	fmt.Println(OpenDir())
}

func TestHistoryFile(t *testing.T) {
	fmt.Println(HistoryFile())
}
func TestLogFileName(t *testing.T) {
	fmt.Println(LogFileName())
}
