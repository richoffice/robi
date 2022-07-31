package robi

import (
	"fmt"
	"testing"
)

func TestNewStoreFromDirectory(t *testing.T) {
	store, err := NewStoreFromDirectory("testfiles/rules")
	fmt.Println(err)
	fmt.Println(store)
}
