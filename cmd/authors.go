package cmd

import (
	"fmt"

	"github.com/Mertozturkk/gitpulse/internal/git"
	"github.com/spf13/cobra"
)

var authorsCmd = &cobra.Command{
	Use:   "authors",
	Short: "Show top contributors by commit count",
	Run: func(cmd *cobra.Command, args []string) {
		authors := git.GetTopAuthors()

		fmt.Println("👤 Contributors:")
		fmt.Println("────────────────────")
		for i, a := range authors {
			fmt.Printf("%d. %-15s - %d commits\n", i+1, a.Name, a.Commits)
		}
	},
}

func init() {
	rootCmd.AddCommand(authorsCmd)
}

// if contributors count is too high, we can limit the output like --top 10
