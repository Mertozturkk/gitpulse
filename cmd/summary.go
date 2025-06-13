package cmd

import (
	"fmt"

	"github.com/Mertozturkk/gitpulse/internal/git"

	"github.com/spf13/cobra"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show overall repository activity summary",
	Run: func(cmd *cobra.Command, args []string) {
		stats := git.GetRepoSummary()
		fmt.Println("📊 Git Summary")
		fmt.Println("━━━━━━━━━━━━━━━")
		fmt.Printf("📁 Repo: %s\n", stats.Name)
		fmt.Printf("🧾 Total commits: %d\n", stats.TotalCommits)
		fmt.Printf("👥 Contributors: %d\n", stats.UniqueAuthors)
		fmt.Printf("🕐 First commit: %s\n", stats.FirstCommit.Format("2006-01-02"))
		fmt.Printf("📆 Last commit: %s\n", stats.LastCommit.Format("2006-01-02"))
		fmt.Printf("📈 Avg commits/day: %.2f\n", stats.AvgCommitsPerDay)
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
