// cmd/myimx/main.go
package main

import (
	"fmt"
	"os"

	"github.com/JimmyTarson/Myimx/tree/main/internal/art"
	"github.com/JimmyTarson/Myimx/tree/main/internal/cli"
)

func main() {
	err := art.EnsureArtDirectoryExists()
	if err != nil {
		fmt.Printf("Warning: %s\n", err)
	}

	if len(os.Args) < 2 {
		cli.PrintUsage()
		return
	}

	err = cli.ProcessCommand(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		cli.PrintUsage()
	}
}
