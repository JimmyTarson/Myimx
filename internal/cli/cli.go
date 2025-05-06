// internal/cli/cli.go
package cli

import (
	"fmt"
	"sort"
	"strings"

	"github.com/JimmyTarson12/Myimx/tree/main/internal/art"
)

// ProcessCommand handles the user command
func ProcessCommand(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no command provided")
	}

	command := strings.ToLower(args[0])

	switch command {
	case "help":
		PrintUsage()
		return nil
	case "list":
		ListAvailableArt()
		return nil
	default:
		return PrintRequestedArt(command)
	}
}

// PrintUsage displays help information
func PrintUsage() {
	fmt.Println("Usage: myimx <command>")
	fmt.Println("Available commands:")
	fmt.Println("  list - Shows all available ASCII art options")
	fmt.Println("  help - Shows this help message")
	fmt.Println("  <art-name> - Displays the specified ASCII art")
}

// ListAvailableArt displays all available ASCII art options
func ListAvailableArt() {
	fmt.Println("Available ASCII art:")

	artNames := art.GetAvailableArt()
	sort.Strings(artNames)

	for _, name := range artNames {
		fmt.Printf("  - %s\n", name)
	}
}

// PrintRequestedArt displays the requested ASCII art
func PrintRequestedArt(name string) error {
	asciiArt, exists := art.GetArt(name)
	if !exists {
		return fmt.Errorf("unknown art: %s\nUse 'myimx list' to see available options", name)
	}

	fmt.Println(asciiArt)
	return nil
}
