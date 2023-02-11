package alias

import (
	"fmt"
	as "github.com/fzdwx/open/internal/alias"
	"github.com/fzdwx/open/internal/browser"
	"github.com/spf13/cobra"
)

var (
	run = &cobra.Command{
		Use:     "run [name]",
		Short:   "Run alias",
		Hidden:  true,
		Example: `$ open alias rm blog`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("name is required")
			}

			return nil
		},

		Run: func(cmd *cobra.Command, args []string) {
			if aliasMap, err := as.ReadToMap(); err == nil {
				if val, ok := aliasMap[args[0]]; ok {
					cobra.CheckErr(browser.Open(val.Url))
				}
			}
		},
	}
)
