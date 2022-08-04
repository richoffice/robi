package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/richoffice/robi"
)

func main() {

	fmt.Println(os.Args)
	if len(os.Args) < 2 {
		fmt.Println("please provide a robi file")
		os.Exit(1)
	}

	src := os.Args[1]
	input := os.Args[2:]

	dir := filepath.Dir(src)
	script := filepath.Base(src)
	if strings.HasSuffix(script, ".js") {
		script = strings.ReplaceAll(script, ".js", "")
	}

	r, err := robi.NewRobi(dir)
	if err != nil {
		fmt.Printf("init robi error: %v \n", err)
		os.Exit(1)
	}

	out, err := r.Execute(script, input)
	if err != nil {
		fmt.Printf("execute robi error: %v \n", err)
		os.Exit(1)
	}

	fmt.Println("Result: ", out)

}
