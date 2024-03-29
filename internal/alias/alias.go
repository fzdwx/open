package alias

import (
	"bytes"
	"errors"
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/fzdwx/infinite"
	"github.com/fzdwx/infinite/components/input/text"
	"github.com/fzdwx/open/internal/browser"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/util"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"github.com/spf13/cobra"
)

type Model struct {
	Time  int64  // created time
	Url   string // url
	Name  string // alias name
	count int64  // count of open
}

type Map map[string]*Model

func (m Map) refresh() error {
	buffer := bytes.NewBuffer(nil)
	for _, model := range m {
		data, err := jsonutil.Encode(model)
		if err != nil {
			continue
		}
		buffer.Write(data)
		buffer.WriteByte('\n')
	}

	return fsutil.WriteFile(cons.AliasFileName(), buffer.Bytes(), 0644, fsutil.FsCWTFlags)
}

func (m *Model) String() string {
	return fmt.Sprintf("%s -> %s", m.Name, m.Url)
}

func (m *Model) Command(name string) *cobra.Command {
	return &cobra.Command{
		Use:    name,
		Hidden: true,
		Short:  "Open " + m.Url + " in browser",
		Run: func(cmd *cobra.Command, args []string) {
			url := m.Url
			if len(args) > 0 {
				url = fmt.Sprintf(url, toAnySlice(args)...)
			}
			cobra.CheckErr(browser.Open(url))
		},
	}
}

// Add alias
// if alias exists, overwrite url.
func Add(url string, name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	if url == "" {
		return errors.New("url is required")
	}

	val := &Model{
		Time: time.Now().UnixMilli(),
		Url:  url,
		Name: name,
	}

	if alias, err := ReadToMap(); err == nil {
		if pre, ok := alias[name]; ok {
			pre.Url = url
			if err = alias.refresh(); err != nil {
				return err
			}
			return nil
		}
	}

	err := util.AppendJson(val, cons.AliasFileName())
	if err == nil {
		fmt.Printf("%s %s -> %s\n",
			lipgloss.NewStyle().Bold(true).Foreground(util.Highlight).Render("√"),
			lipgloss.NewStyle().Bold(true).Foreground(util.Special).Render(name),
			url,
		)
	}
	return err
}

func AddInteractive() error {
	name, err := infinite.NewText(
		text.WithPrompt("alias name"),
	).Display()

	if err != nil {
		return err
	}
	url, err := infinite.NewText(
		text.WithPrompt("alias url"),
	).Display()
	if err != nil {
		return err
	}

	return Add(url, name)
}

// Remove alias
func Remove(name string) (*Model, error) {
	alias, err := ReadToMap()
	if err != nil {
		return nil, err
	}

	if val, ok := alias[name]; ok {
		delete(alias, name)
		return val, alias.refresh()
	}

	return nil, nil
}

func Read() ([]*Model, error) {
	pairs, err := util.Read(cons.AliasFileName())
	if err != nil {
		return nil, err
	}

	models := make([]*Model, len(pairs))
	for i, val := range pairs {
		var model *Model
		err := jsonutil.DecodeString(val, &model)
		if err != nil {
			return nil, err
		}
		models[i] = model
	}

	return models, nil
}

func ReadToMap() (Map, error) {
	models, err := Read()
	if err != nil {
		return nil, err
	}

	aliasMap := make(Map)
	for _, model := range models {
		aliasMap[model.Name] = model
	}

	return aliasMap, nil
}

// ForeachAlias collection alias and call func
func ForeachAlias(f func(*Model)) error {
	alias, err := Read()
	if err != nil {
		return err
	}

	for _, model := range alias {
		f(model)
	}
	return nil
}

// LsName list alias names
func LsName() error {
	return ForeachAlias(func(model *Model) {
		fmt.Println(model.Name)
	})
}

func toAnySlice[T any](args []T) []interface{} {
	a := make([]interface{}, len(args))
	for i, v := range args {
		a[i] = v
	}
	return a
}
