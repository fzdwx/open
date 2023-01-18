package alias

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	as "github.com/fzdwx/open/internal/alias"
	"github.com/fzdwx/open/internal/util"
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
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error

			if name == "" {
				err = fmt.Errorf("alias name is required")
			} else if len(args) == 0 {
				err = fmt.Errorf("url is required")
			}
			return err

		},
		Run: func(cmd *cobra.Command, args []string) {
			cobra.CheckErr(as.Add(args[0], name))

			fmt.Printf("%s %s -> %s\n",
				lipgloss.NewStyle().Bold(true).Foreground(util.Highlight).Render("√"),
				lipgloss.NewStyle().Bold(true).Foreground(util.Special).Render(fmt.Sprintf("%s", name)),
				args[0],
			)
		},
	}

	name string
)

func init() {
	add.Flags().StringVarP(&name, "name", "n", "", "alias name")
}
