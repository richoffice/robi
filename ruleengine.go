package robi

import (
	"fmt"

	"github.com/dop251/goja"
)

type Rule struct {
	Name    string
	Source  string
	Program *goja.Program
}

func (rule *Rule) String() string {
	return rule.Name + "->" + rule.Source
}

type RuleEngine interface {
	ApplyRule(name string, input interface{}) (interface{}, error)
}

type MemoryRuleEngine struct {
	Store RuleStore
	VM    *goja.Runtime
}

// type Console struct {
// }

// func (console *Console) Log(values ...string) {
// 	fmt.Println(values)
// }

func Log(v ...interface{}) {
	fmt.Println("[js]", v)
}

func NewMemoryRuleEngine(store RuleStore, vars map[string]interface{}) RuleEngine {
	vm := goja.New()
	vm.Set("log", Log)
	for key, value := range vars {
		vm.Set(key, value)
	}
	return &MemoryRuleEngine{VM: vm, Store: store}
}

func (engine *MemoryRuleEngine) ApplyRule(name string, input interface{}) (interface{}, error) {
	rule := engine.Store.GetRule(name)
	if rule == nil {
		return nil, fmt.Errorf("not a exist rule name '%v'", name)
	}

	engine.VM.Set("input", input)
	// v, runtimeErr := engine.VM.RunProgram(rule.Program)
	v, runtimeErr := engine.VM.RunString(rule.Source)
	if runtimeErr != nil {
		return nil, runtimeErr
	}
	exp, isExp := v.Export().(*goja.Exception)
	if isExp {
		return nil, exp
	}

	return v.Export(), nil
}
