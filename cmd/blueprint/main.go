package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Println("blueprint-monitor v0.1.0")
		return
	}
	fmt.Println("blueprint-monitor: use 'monitor', 'status', or 'kill'")
}
