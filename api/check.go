package api

import "github.com/spf13/cobra"

func Check(msg interface{}) {
	cobra.CheckErr(msg)
}
