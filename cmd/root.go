package cmd

import (
	"github.com/fzdwx/open/alias"
	"github.com/fzdwx/open/api"
	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "open [subcommand]",
	Short:   "Open url in browser",
	Version: "v0.1.0",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		if val, ok := alias.Get(args[0]); ok {
			api.BrowseWithCheck(val.Path)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		pcli.CheckForUpdates()
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		pcli.CheckForUpdates()
		os.Exit(1)
	}

	//pcli.CheckForUpdates()
}

func init() {
	api.InitBrowser(os.Stdout)
	rootCmd.AddCommand(initGh())
	rootCmd.AddCommand(initSet())
	rootCmd.AddCommand(initDel())
	rootCmd.AddCommand(alias.LoadCmds()...)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-open.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// Use https://github.com/pterm/pcli to style the output of cobra.
	pcli.SetRepo("fzdwx/open")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()
}
