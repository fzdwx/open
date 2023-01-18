package alias

import (
	"bytes"
	"fmt"
	"github.com/fzdwx/open/internal/cons"
	"github.com/fzdwx/open/internal/util"
	"github.com/gookit/goutil/fsutil"
	"github.com/gookit/goutil/jsonutil"
	"time"
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

// Add alias
// if alias exists, overwrite url.
func Add(url string, name string) error {
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

	return util.AppendJson(val, cons.AliasFileName())
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
