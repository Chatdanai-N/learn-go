package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("My Os is\t", runtime.GOOS)
	fmt.Println("My Arch is\t", runtime.GOARCH)
}
