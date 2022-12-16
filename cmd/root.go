package cmd

import (
	"fmt"
	"github.com/fzdwx/open/cmd/gh"
	"github.com/fzdwx/open/cmd/history"
	"github.com/fzdwx/open/cmd/url"
	"github.com/fzdwx/open/internal/browser"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/user"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/pterm/pcli"
	"github.com/pterm/pterm"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "open xxx",
		Short: "Open url in browser",
		Example: `$ open gh
$ open gh p
$ open gh -s fzdwx -> open https://github.com/search?q=fzdwx
$ open https://github.com`,
		Version: cons.Version,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) <= 1 {
				return nil
			}
			return fmt.Errorf("accept only one / zero argument")
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) >= 1 {
				cobra.CheckErr(browser.Open(args[0]))
				return
			}

			fmt.Printf("Your Name: %s\n", user.Name())
			fmt.Printf("Your token: %s\n", user.Token())
		},
	}

	debug bool
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Fetch user interrupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		pterm.Warning.Println("user interrupt")
		_ = pcli.CheckForUpdates()
		os.Exit(0)
	}()

	// Execute cobra
	if err := rootCmd.Execute(); err != nil {
		_ = pcli.CheckForUpdates()
		os.Exit(1)
	}

	//pcli.CheckForUpdates()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(gh.Command())
	rootCmd.AddCommand(history.Command())
	rootCmd.AddCommand(url.Command())

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-open.yaml)")

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "show log in console")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// Use https://github.com/pterm/pcli to style the output of cobra.
	_ = pcli.SetRepo("fzdwx/open")
	pcli.SetRootCmd(rootCmd)
	pcli.Setup()
}

func initConfig() {
	if !debug {
		slog.SetLogLevel(slog.PanicLevel)
	}
	slog.PushHandler(handler.MustFileHandler(cons.LogFileName(), handler.WithLogLevels(slog.AllLevels)))
}
