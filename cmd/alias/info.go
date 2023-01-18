package alias

import (
	"fmt"
	as "github.com/fzdwx/open/internal/alias"
	"github.com/spf13/cobra"
)

var (
	info = &cobra.Command{
		Use:     "info [name]",
		Short:   "Get alias info",
		Hidden:  true,
		Aliases: []string{"i"},
		Example: `$ open alias info blog`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("name is required")
			}

			return nil
		},

		Run: func(cmd *cobra.Command, args []string) {
			if aliasMap, err := as.ReadToMap(); err == nil {
				if val, ok := aliasMap[args[0]]; ok {
					fmt.Println(val)
				}
			}
		},
	}
)
