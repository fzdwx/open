package history

import (
	"fmt"
	"github.com/MarvinJWendt/testza"
	"github.com/gookit/goutil/jsonutil"
	"testing"
	"time"
)

func TestHistoryFile(t *testing.T) {
	fmt.Println(FileName())
}

func TestEncode(t *testing.T) {
	encode, _ := jsonutil.Encode(&Model{
		Time: time.Now().UnixMilli(),
		Url:  "11",
	})

	fmt.Println(string(encode))
}

func TestWrite(t *testing.T) {
	testza.AssertNoError(t, Write("aaaaaaaaaaaa"))
}
