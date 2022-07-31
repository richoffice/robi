package robi

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/richoffice/richframe"
)

type Robi struct {
	Base   string
	Engine RuleEngine
}

func NewRobi(base string) (*Robi, error) {

	store, err := NewStoreFromDirectory(base)
	if err != nil {
		return nil, err
	}

	robi := &Robi{
		Base: base,
	}

	var vars map[string]interface{} = map[string]interface{}{
		"robi": robi,
		"log":  fmt.Println,
	}

	engine := NewMemoryRuleEngine(store, vars)
	robi.Engine = engine

	return robi, nil
}

func (robi *Robi) Import(defPath string, srcFile string) interface{} {
	fullPath := defPath

	if !strings.HasPrefix(defPath, "/") {
		fullPath = filepath.Join(robi.Base, "defs", defPath)
	}
	rf, err := richframe.LoadRichFrames(srcFile, fullPath, nil)
	if err != nil {
		panic(err)
	}
	return rf
}

func (robi *Robi) Export(def string, targetFile string) interface{} {
	return nil
}

func (robi *Robi) Execute(task string, args interface{}) (interface{}, error) {
	return robi.Engine.ApplyRule(task, args)
}
