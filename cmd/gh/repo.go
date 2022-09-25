package gh

import (
	"github.com/fzdwx/open/pkg/browser"
	"github.com/gookit/goutil/strutil"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"regexp"
)

var repo = &cobra.Command{
	Use:     "repo",
	Short:   "open github repository in browser. eg: https://github.com/fzdwx/open",
	Example: `$ open gh repo # must in git project`,
	Run: func(cmd *cobra.Command, args []string) {
		buffer := &strutil.Buffer{}
		command := exec.Command("git", "remote", "-v")
		command.Stdout = buffer
		command.Stderr = os.Stderr
		command.Stdin = os.Stdin

		cobra.CheckErr(command.Run())
		remotePairs := strutil.Split(buffer.String(), "\n")

		pattern, err := regexp.Compile("(http|ftp|https):\\/\\/[\\w\\-_]+(\\.[\\w\\-_]+)+([\\w\\-\\.,@?^=%&:/~\\+#]*[\\w\\-\\@?^=%&/~\\+#])?")
		if err != nil {
			cobra.CheckErr(err)
		}

		url := pattern.FindString(remotePairs[0])
		cobra.CheckErr(browser.Open(url))
	},
}
