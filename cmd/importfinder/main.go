package main

import (
	"fmt"

	i "github.com/ldelossa/importfinder"
)

func main() {
	f := i.NewFinder()

	err := f.FindImports()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f)
}
