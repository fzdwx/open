package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

func initCompletion() *cobra.Command {
	var completionCmd = &cobra.Command{
		Use:   "completion [bash|zsh|fish|powershell]",
		Short: "Generate completion script",
		Long: `To load completions:

Bash:

$ source <(open completion bash)

# To load completions for each session, execute once:
Linux:
  $ open completion bash > /etc/bash_completion.d/open
MacOS:
  $ open completion bash > /usr/local/etc/bash_completion.d/open

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ open completion zsh > "${fpath[1]}/_open"

# You will need to start a new shell for this setup to take effect.

Fish:

$ open completion fish | source

# To load completions for each session, execute once:
$ open completion fish > ~/.config/fish/completions/open.fish
`,
		DisableFlagsInUseLine: true,
		ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
		Args:                  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			switch args[0] {
			case "bash":
				cmd.Root().GenBashCompletion(os.Stdout)
			case "zsh":
				cmd.Root().GenZshCompletion(os.Stdout)
			case "fish":
				cmd.Root().GenFishCompletion(os.Stdout, true)
			case "powershell":
				cmd.Root().GenPowerShellCompletion(os.Stdout)
			}
		},
	}

	return completionCmd
}
