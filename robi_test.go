package robi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRobi_Execute(t *testing.T) {
	robi, err := NewRobi("./testfiles/robidemo")

	expected := map[string]interface{}{
		"z": int64(6),
	}
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	result, err := robi.Execute("testRule", map[string]interface{}{
		"x": 2,
		"y": 3,
	})
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v but got %v", expected, result)
	}
	fmt.Println(result.(map[string]interface{}))
}

func TestRobi_ExecuteSimple(t *testing.T) {
	robi, err := NewRobi("./testfiles/robidemo")

	expected := map[string]interface{}{
		"z": int64(6),
	}
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	result, err := robi.Execute("testRule", []string{
		"",
	})
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %v but got %v", expected, result)
	}
	fmt.Println(result.(map[string]interface{}))
}
