package robi

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type RuleStore interface {
	GetRule(name string) *Rule
}

type DirRuleStore struct {
	Rules map[string]*Rule
}

func (store *DirRuleStore) GetRule(name string) *Rule {
	return store.Rules[name]
}

func NewStoreFromFile(filePath string) (*DirRuleStore, error) {

	f, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	rules := make(map[string]*Rule)

	if !f.IsDir() && strings.HasSuffix(f.Name(), "js") {
		srcBytes, readErr := ioutil.ReadFile(filePath)
		if readErr != nil {
			return nil, readErr
		}
		src := string(srcBytes)
		name := strings.TrimSuffix(f.Name(), ".js")
		rule := &Rule{
			Name:   name,
			Source: src,
		}
		rules[name] = rule

	}

	return &DirRuleStore{Rules: rules}, nil
}

func NewStoreFromDirectory(dir string) (*DirRuleStore, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	rules := make(map[string]*Rule)

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), "js") {
			path := filepath.Join(dir, f.Name())
			srcBytes, readErr := ioutil.ReadFile(path)
			if readErr != nil {
				return nil, readErr
			}
			src := string(srcBytes)
			name := strings.TrimSuffix(f.Name(), ".js")
			rule := &Rule{
				Name:   name,
				Source: src,
			}
			rules[name] = rule

		}
	}

	return &DirRuleStore{Rules: rules}, nil

}
