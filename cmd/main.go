package main

import (
	"fmt"
	"os"
	"quaziiuidownloader/internal/config"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(2)
	}

	cmd := os.Args[1]
	settingsPath := "config.json"

	switch cmd {
	case "setup":
		if err := config.CreateDefaultIfMissing(settingsPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating default config: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Default configuration created at", settingsPath)

	case "run":
		_, err := config.Load(settingsPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

	default:
		printHelp()
		os.Exit(2)
	}
}

func printHelp() {
	fmt.Println("Usage: quaziiuidownloader <command>")
	fmt.Println("Commands:")
	fmt.Println("  setup    Create a default configuration file if it doesn't exist")
	fmt.Println("  run      Run the application with the existing configuration")
}
