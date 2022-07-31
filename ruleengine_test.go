package robi

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	input = `{
"x":10,
"y":20
}`
)

func TestMemoryRuleEngine_ApplyRulePassMap(t *testing.T) {
	store, err := NewStoreFromDirectory("testfiles/rules")
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}
	jsonErr := json.Unmarshal([]byte(input), &m)
	if jsonErr != nil {
		panic(jsonErr)
	}

	// fmt.Println(m)
	engine := NewMemoryRuleEngine(store, nil)
	result, _ := engine.ApplyRule("testPassMap", m)

	expected := `{
	"z": 30
}`
	fmt.Println(result, expected)
	// fmt.Println(result)
	// r := imhfmt.ToJson(result.(map[string]interface{}))
	// // fmt.Println()
	// if r != expected {
	// 	t.Errorf("Epected %v but got %v", expected, r)
	// }

}

func TestMemoryRuleEngine_ApplyRule(t *testing.T) {
	store, err := NewStoreFromDirectory("testfiles/rules")
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}
	jsonErr := json.Unmarshal([]byte(input), &m)
	if jsonErr != nil {
		panic(jsonErr)
	}

	// fmt.Println(m)
	engine := NewMemoryRuleEngine(store, nil)
	result, _ := engine.ApplyRule("testRule", m)

	var expected int64 = 200
	// fmt.Println(result)
	r := result.(map[string]interface{})["z"].(int64)
	// fmt.Println()
	if r != expected {
		t.Errorf("Epected %v but got %v", expected, r)
	}

	sresult, _ := engine.ApplyRule("testSimple", 11)
	if sresult.(int64) != 121 {
		t.Errorf("Epected %v but got %v", 121, sresult)
	}

}
