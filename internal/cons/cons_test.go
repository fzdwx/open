package cons

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestUserDir(t *testing.T) {
	fmt.Println(getUserDir())
}

func TestOpenDir(t *testing.T) {
	fmt.Println(OpenDir())
}

func TestHistoryFile(t *testing.T) {
	fmt.Println(HistoryFileName())
}
func TestLogFileName(t *testing.T) {
	fmt.Println(LogFileName())
}
func TestLookUp(t *testing.T) {
	path, _ := exec.LookPath("bat")
	fmt.Println(path)
}
