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
	Short: "Launch tab-based interactive terminal UI",
	Run: func(cmd *cobra.Command, args []string) {
		tabs := []string{"Summary", "Authors", "Heatmap"}
		currentTab := 0

		app := tview.NewApplication()

		var tabBar *tview.TextView = tview.NewTextView()
		tabBar.SetDynamicColors(true)
		tabBar.SetTextAlign(tview.AlignCenter)
		tabBar.SetBorder(true)
		tabBar.SetTitle(" Tabs ")

		var mainContent *tview.TextView = tview.NewTextView()
		mainContent.SetDynamicColors(true)
		mainContent.SetBorder(true)
		mainContent.SetTitle(" Content ")

		draw := func() {
			tabLine := ""
			for i, tab := range tabs {
				if i == currentTab {
					tabLine += fmt.Sprintf("[yellow][ %s ][-] ", tab)
				} else {
					tabLine += fmt.Sprintf("  %s   ", tab)
				}
			}
			tabBar.SetText(tabLine)

			switch tabs[currentTab] {
			case "Summary":
				summary := git.GetRepoSummary()
				mainContent.SetTitle(" Repo Summary ")
				mainContent.SetText(fmt.Sprintf(`📁 Repo: %s
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
				))
			case "Authors":
				authors := git.GetTopAuthors(10)
				mainContent.SetTitle(" Top Contributors ")
				authorsText := ""
				for i, a := range authors {
					authorsText += fmt.Sprintf("%d. %-15s (%d commits)\n", i+1, a.Name, a.Commits)
				}
				mainContent.SetText(authorsText)
			case "Heatmap":
				heatmap := git.GetCommitHeatmap("")
				mainContent.SetTitle(" Commit Heatmap ")
				heatmapText := ""
				for _, day := range git.Weekdays {
					bar := git.Bar(heatmap[day])
					heatmapText += fmt.Sprintf("%s: %-20s (%d)\n", day, bar, heatmap[day])
				}
				mainContent.SetText(heatmapText)
			}
		}

		draw()

		flex := tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(tabBar, 3, 1, false).
			AddItem(mainContent, 0, 1, true)

		flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			switch event.Key() {
			case tcell.KeyRight:
				currentTab = (currentTab + 1) % len(tabs)
				draw()
				return nil
			case tcell.KeyLeft:
				currentTab = (currentTab - 1 + len(tabs)) % len(tabs)
				draw()
			case tcell.KeyRune:
				if event.Rune() == 'q' || event.Rune() == 'Q' {
					app.Stop()
					return nil
				}
			}
			return event
		})

		if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
