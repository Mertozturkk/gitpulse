package cmd

import (
	"fmt"

	"github.com/Mertozturkk/gitpulse/internal/git"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "Launch interactive terminal dashboard",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		grid := tview.NewGrid().
			SetRows(0, 0).
			SetColumns(0, 0).
			SetBorders(true)

		summary := git.GetRepoSummary()
		summaryText := fmt.Sprintf(`📁 Repo: %s
🧾 Commits: %d
👥 Contributors: %d
🕐 First: %s
📆 Last: %s
📈 Avg/day: %.2f`,
			summary.Name,
			summary.TotalCommits,
			summary.UniqueAuthors,
			summary.FirstCommit.Format("2006-01-02"),
			summary.LastCommit.Format("2006-01-02"),
			summary.AvgCommitsPerDay,
		)

		summaryBox := tview.NewTextView()
		summaryBox.SetText(summaryText)
		summaryBox.SetDynamicColors(true)
		summaryBox.SetBorder(true)
		summaryBox.SetTitle(" Repo Summary ")
		summaryBox.SetChangedFunc(func() { app.Draw() })

		authors := git.GetTopAuthors(5)
		authorsText := ""
		for i, a := range authors {
			authorsText += fmt.Sprintf("%d. %-15s (%d)\n", i+1, a.Name, a.Commits)
		}

		authorsBox := tview.NewTextView()
		authorsBox.SetText(authorsText)
		authorsBox.SetDynamicColors(true)
		authorsBox.SetBorder(true)
		authorsBox.SetTitle(" Top Authors ")
		authorsBox.SetChangedFunc(func() { app.Draw() })

		heatmap := git.GetCommitHeatmap("")
		heatmapText := ""
		for _, day := range git.Weekdays {
			bar := git.Bar(heatmap[day])
			heatmapText += fmt.Sprintf("%s: %-20s (%d)\n", day, bar, heatmap[day])
		}

		heatmapBox := tview.NewTextView()
		heatmapBox.SetText(heatmapText)
		heatmapBox.SetDynamicColors(true)
		heatmapBox.SetBorder(true)
		heatmapBox.SetTitle(" Commit Heatmap ")
		heatmapBox.SetChangedFunc(func() { app.Draw() })

		grid.AddItem(summaryBox, 0, 0, 1, 1, 0, 0, false)
		grid.AddItem(authorsBox, 0, 1, 1, 1, 0, 0, false)
		grid.AddItem(heatmapBox, 1, 0, 1, 2, 0, 0, false)

		grid.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			if event.Rune() == 'q' || event.Rune() == 'Q' {
				app.Stop()
				return nil
			}
			return event
		})

		if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
