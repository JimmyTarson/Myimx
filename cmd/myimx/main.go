// cmd/myimx/main.go
package main

import (
	"fmt"
	"os"

	"github.com/JimmyTarson12/Myimx/internal/cli"
)

func main() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		return
	}

	err := cli.ProcessCommand(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		cli.PrintUsage()
	}
}
