package alias

import (
	as "github.com/fzdwx/open/internal/alias"
	"github.com/spf13/cobra"
)

var (
	add = &cobra.Command{
		Use:   "add",
		Short: "Add custom aliases",
		Aliases: []string{
			"new",
			"create",
			"a",
		},
		Example: `$ open alias add https://fzdwx.github.io/ --name blog`,
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" && len(args) < 1 {
				cobra.CheckErr(as.AddInteractive())
			} else {
				if name == "" {
					cobra.CheckErr("name is required")
				}
				if len(args) < 1 {
					cobra.CheckErr("url is required")
				}
				cobra.CheckErr(as.Add(args[0], name))
			}
		},
	}

	name string
)

func init() {
	add.Flags().StringVarP(&name, "name", "n", "", "alias name")
}
