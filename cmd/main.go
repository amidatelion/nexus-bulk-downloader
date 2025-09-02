package main

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"

	"nexusmods-cli/internal/config"
	"nexusmods-cli/internal/nexus"
)

func main() {
	// Load config.json
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Interactive mod selection
	modIDs := cfg.ModOptions()
	prompt := promptui.Select{
		Label: "Select Mod",
		Items: modIDs,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
	}

	modID := config.ExtractModID(result)

	// Query API
	files, err := nexus.FetchFiles(cfg.APIKey, modID)
	if err != nil {
		log.Fatalf("Error fetching files: %v", err)
	}

	// Display results
	fmt.Println("\nAvailable Files:")
	for _, f := range files {
		fmt.Printf("- [%d] %s [%s]\n", f.FileID, f.Name, f.Category)
		fmt.Printf("  File: %s\n\n", f.FileName)
	}
}
