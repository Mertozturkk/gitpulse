package cmd

import (
	"fmt"

	"github.com/Mertozturkk/gitpulse/internal/git"
	"github.com/spf13/cobra"
)

var dirsCmd = &cobra.Command{
	Use:   "dirs",
	Short: "Show most frequently changed directories",
	Run: func(cmd *cobra.Command, args []string) {
		stats := git.GetChangedDirectories()

		fmt.Println("📂 Most Active Directories:")
		fmt.Println("──────────────────────────────")
		for i, d := range stats {
			fmt.Printf("%d. %-30s - %d changes\n", i+1, d.Path, d.Changes)
		}
	},
}

func init() {
	rootCmd.AddCommand(dirsCmd)
}
