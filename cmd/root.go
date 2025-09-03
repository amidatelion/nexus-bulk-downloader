package cmd

import (
	"fmt"
	"log"

	"nexus-bulk-downloader/internal/config"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"nexus-bulk-downloader/internal/extractor"
	"nexus-bulk-downloader/internal/nexus"
)

var (
	cfgPath         string
	overrideDir     string
	overrideExtract bool
)

var rootCmd = &cobra.Command{
	Use:   "nexusmods-cli",
	Short: "Download mods from NexusMods",
	Run: func(cmd *cobra.Command, args []string) {
		// Load config
		cfg, err := config.LoadConfig("config.json")
		if err != nil {
			log.Fatalf("failed to load config: %v", err)
		}

		fmt.Println("Loaded config. API key set?", cfg.Config.APIKey != "")

		for game, mods := range cfg.Games {
			fmt.Printf("\n=== Game: %s ===\n", game)
			for modID, modName := range mods {
				fmt.Printf("\n--- Mod %s (%s) ---\n", modID, modName)

				files, err := nexus.FetchFiles(cfg.Config.APIKey, game, modID)
				if err != nil {
					log.Printf("error fetching files for mod %s: %v", modID, err)
					continue
				}

				if len(files) == 0 {
					fmt.Println("No MAIN files found.")
					continue
				}

				// Single MAIN file
				if len(files) == 1 {
					f := files[0]
					fmt.Printf("Auto-selected: %s [%s] (id=%d)\n", f.Name, f.FileName, f.FileID)

					uri, err := nexus.GetDownloadLink(cfg.Config.APIKey, game, modID, f.FileID)
					if err != nil {
						log.Printf("error getting download link: %v", err)
						continue
					}

					path, err := nexus.DownloadFile(uri, cfg.Config.DownloadDir)
					if err != nil {
						log.Printf("download failed: %v", err)
						continue
					}
					fmt.Printf("Downloaded to %s\n", path)

					if cfg.Config.AutoExtract {
						fmt.Printf("Extracting %s...\n", path)
						if err := extractor.ExtractZip(path, cfg.Config.DownloadDir); err != nil {
							log.Printf("failed to extract %s: %v", path, err)
						} else {
							fmt.Printf("Extracted to %s\n", cfg.Config.DownloadDir)
						}
					}
					continue

				}

				// Multiple MAIN files, use promptUI to let user pick
				templates := &promptui.SelectTemplates{
					Label:    "{{ . }}",
					Active:   "▶ {{ .Name | cyan }} ({{ .FileName | red }})",
					Inactive: "  {{ .Name | cyan }} ({{ .FileName | red }})",
					Selected: "✔ {{ .Name | green }}",
				}

				prompt := promptui.Select{
					Label:     fmt.Sprintf("Choose MAIN file for mod %s", modName),
					Items:     files,
					Templates: templates,
					Size:      5,
				}

				i, _, err := prompt.Run()
				if err != nil {
					log.Printf("prompt failed: %v", err)
					continue
				}

				selected := files[i]
				fmt.Printf("Selected: %s [%s] (id=%d)\n", selected.Name, selected.FileName, selected.FileID)

				uri, err := nexus.GetDownloadLink(cfg.Config.APIKey, game, modID, selected.FileID)
				if err != nil {
					log.Printf("error getting download link: %v", err)
					continue
				}

				path, err := nexus.DownloadFile(uri, cfg.Config.DownloadDir)
				if err != nil {
					log.Printf("download failed: %v", err)
					continue
				}
				fmt.Printf("Downloaded to %s\n", path)

				if cfg.Config.AutoExtract {
					fmt.Printf("Extracting %s...\n", path)
					if err := extractor.ExtractZip(path, cfg.Config.DownloadDir); err != nil {
						log.Printf("failed to extract %s: %v", path, err)
					} else {
						fmt.Printf("Extracted to %s\n", cfg.Config.DownloadDir)
					}
				}

			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "config.json", "Path to config file")
	rootCmd.PersistentFlags().StringVarP(&overrideDir, "downloaddir", "d", "", "Override download directory")
	rootCmd.PersistentFlags().BoolVarP(&overrideExtract, "autoextract", "x", false, "Override autoextract setting")
}
