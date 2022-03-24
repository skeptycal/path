package main

import (
	"fmt"

	shpath "github.com/skeptycal/shpath"
)

func main() {
	var sh shpath.Path = nil

	sh.Load()

	fmt.Print("path")
}
