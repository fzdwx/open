package alias

import (
	"bufio"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/open/api"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	homeDir  string
	aliasMap = make(map[string]Alias)
	loadFlag bool
)

type Alias struct {
	Path string
	time string
}

const (
	aliasFile = ".open_alias"
	delimiter = "|-|"
)

func init() {
	current, err := user.Current()
	cobra.CheckErr(err)
	homeDir = current.HomeDir
}

func LoadCmds() []*cobra.Command {
	checkLoad()

	prefix := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("4")).Render("This is alias will open")

	var cmds []*cobra.Command
	for k, a := range aliasMap {
		cmd := &cobra.Command{
			Use:   k,
			Short: prefix + " " + a.Path,
			Run: func(cmd *cobra.Command, args []string) {
				api.BrowseWithCheck(a.Path)
			},
			ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
				return []string{a.Path}, cobra.ShellCompDirectiveNoSpace
			},
		}

		cmds = append(cmds, cmd)
	}

	return cmds
}

func Load() {
	if loadFlag {
		return
	}

	file, err := open()
	defer file.Close()
	api.Check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strArr := strings.Split(scanner.Text(), delimiter)
		if len(strArr) < 3 {
			fmt.Fprint(os.Stderr, "alias load error:", strArr)
			continue
		}
		alias := Alias{
			Path: strArr[1],
			time: strArr[2],
		}
		aliasMap[strArr[0]] = alias
	}

	loadFlag = true
}

func Save() error {
	file, err := open()
	defer file.Close()
	if err != nil {
		return err
	}

	for k, v := range aliasMap {
		_, err := file.WriteString(fmt.Sprintf("%s%s%s%s%s\n", k, delimiter, v.Path, delimiter, v.time))
		if err != nil {
			return err
		}
	}

	return file.Sync()
}

func Set(alias string, path string) {
	checkLoad()
	aliasMap[alias] = Alias{
		Path: path,
		time: strconv.FormatInt(time.Now().UnixMilli(), 10),
	}
}

func Del(alias string) {
	checkLoad()
	delete(aliasMap, alias)
}

func Get(alias string) (Alias, bool) {
	checkLoad()
	v, ok := aliasMap[alias]
	return v, ok
}

func TruncateFile() error {
	file, err := open()
	if err != nil {
		return err
	}

	return file.Truncate(0)
}

func checkLoad() {
	if loadFlag {
		return
	}

	Load()
}

func open() (*os.File, error) {
	filename := fmt.Sprintf("%s%s%s", homeDir, string(filepath.Separator), aliasFile)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return f, nil
}
