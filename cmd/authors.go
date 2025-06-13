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
		authors := git.GetTopAuthors(topN)

		fmt.Println("👤 Contributors:")
		fmt.Println("────────────────────")
		for i, a := range authors {
			fmt.Printf("%d. %-15s - %d commits\n", i+1, a.Name, a.Commits)
		}
	},
}

var topN int

func init() {
	authorsCmd.Flags().IntVar(&topN, "top", 10, "Limit the number of authors to show")
	rootCmd.AddCommand(authorsCmd)
}
