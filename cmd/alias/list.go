package alias

import (
	"fmt"
	as "github.com/fzdwx/open/internal/alias"
	"github.com/spf13/cobra"
)

var (
	list = &cobra.Command{
		Use:     "list",
		Short:   "Show custom aliases",
		Aliases: []string{"ls"},
		Example: `$ open alias list
$ open alias ls | fzf --preview 'open alias info {}'`,
		Run: func(cmd *cobra.Command, args []string) {
			if aliasMap, err := as.ReadToMap(); err == nil {
				for _, alias := range aliasMap {
					fmt.Println(alias.Name)
				}
			}
		},
	}
)
