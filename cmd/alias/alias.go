package alias

import (
	"io"
	"os"
	"os/exec"

	as "github.com/fzdwx/open/internal/alias"
	"github.com/gookit/goutil/byteutil"
	"github.com/spf13/cobra"
)

var (
	alias = &cobra.Command{
		Use:     "alias [subcommand]",
		Short:   "Manage custom aliases",
		Aliases: []string{"ls"},
		Example: `$ open alias add https://fzdwx.github.io/ --name blog
$ open alias list
$ open alias ls | fzf --preview 'open alias info {}' --bind 'enter:execute(open alias run {})'
$ open alias remove blog`,
		Run: func(cmd *cobra.Command, args []string) {
			fzfCommand := exec.Command("fzf", "--preview", "open alias info {}", "--bind", "enter:execute(open alias run {})")
			buffer := byteutil.NewBuffer()
			cobra.CheckErr(as.ForeachAlias(func(model *as.Model) {
				io.WriteString(buffer, model.Name+"\n")
			}))

			fzfCommand.Stdin = buffer
			fzfCommand.Stdout = os.Stdout
			fzfCommand.Stderr = os.Stderr

			cobra.CheckErr(fzfCommand.Run())
		},
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
