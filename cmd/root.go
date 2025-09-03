package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "nexus-bulk-downloader",
	Short: "Download mods from NexusMods",
}
