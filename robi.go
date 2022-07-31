package robi

type Robi struct {
	Base   string
	Engine RuleEngine
}

func NewRobi(base string) (*Robi, error) {

	store, err := NewStoreFromDirectory("testfiles/rules")
	if err != nil {
		return nil, err
	}

	robi := &Robi{
		Base: base,
	}

	var vars map[string]interface{} = map[string]interface{}{
		"robi": robi,
	}

	engine := NewMemoryRuleEngine(store, vars)
	robi.Engine = engine

	return robi, nil
}

func (robi *Robi) Import(def string, srcFile string) interface{} {
	return nil
}

func (robi *Robi) Export(def string, targetFile string) interface{} {
	return nil
}

func (robi *Robi) Execute(task string, args interface{}) (interface{}, error) {
	return robi.Engine.ApplyRule(task, args)
}
