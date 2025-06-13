package cmd

import (
	"fmt"

	"github.com/Mertozturkk/gitpulse/internal/git"

	"github.com/spf13/cobra"
)

var heatmapCmd = &cobra.Command{
	Use:   "heatmap",
	Short: "Visualize commit activity per day of the week",
	Run: func(cmd *cobra.Command, args []string) {
		activity := git.GetCommitHeatmap(since)

		fmt.Printf("🗓️ Commit Heatmap")
		if since != "" {
			fmt.Printf(" (since %s)", since)
		}
		fmt.Println(":")
		fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━")

		for _, day := range git.Weekdays {
			count := activity[day]
			bar := git.Bar(count)
			fmt.Printf("%s: %-15s (%d)\n", day, bar, count)
		}
	},
}

var since string

func init() {
	heatmapCmd.Flags().StringVar(&since, "since", "", "Only include commits since this date (e.g. 30d, 2024-01-01)")
	rootCmd.AddCommand(heatmapCmd)
}
