package api

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func Check(msg interface{}) {
	cobra.CheckErr(msg)
}

func Eprintln(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
}
