package alias

import (
	as "github.com/fzdwx/open/internal/alias"
	"github.com/spf13/cobra"
)

var (
	alias = &cobra.Command{
		Use:   "alias [subcommand]",
		Short: "Manage custom aliases",
		Example: `$ open alias add https://fzdwx.github.io/ --name blog
$ open alias list
$ open alias ls | fzf --preview 'open alias info {}' --bind 'enter:execute(open alias run {})'
$ open alias remove blog`,
	}
)

func Command(root *cobra.Command) *cobra.Command {
	alias.AddCommand(add)
	alias.AddCommand(list)
	alias.AddCommand(remove)

	// hide alias subcommand
	alias.AddCommand(info)
	alias.AddCommand(run)

	loadAlias(root)

	return alias
}

// loadAlias load user custom alias
func loadAlias(root *cobra.Command) {
	if aliasMap, err := as.ReadToMap(); err == nil {
		for name, alias := range aliasMap {
			alias := alias
			root.AddCommand(alias.Command(name))
		}
	}
}
