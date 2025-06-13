package git

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type RepoSummary struct {
	Name             string
	TotalCommits     int
	FirstCommit      time.Time
	LastCommit       time.Time
	UniqueAuthors    int
	AvgCommitsPerDay float64
}

func GetRepoSummary() RepoSummary {
	total := runGit("rev-list --count HEAD")
	first := runGit("log --reverse --format=%cs | head -n 1")
	last := runGit("log -1 --format=%cs")
	authors := runGit("shortlog -sne | wc -l")

	firstTime, _ := time.Parse("2006-01-02", first)
	lastTime, _ := time.Parse("2006-01-02", last)

	days := int(lastTime.Sub(firstTime).Hours()/24) + 1
	commitCount := toInt(total)
	authorCount := toInt(authors)

	return RepoSummary{
		Name:             getRepoName(),
		TotalCommits:     commitCount,
		FirstCommit:      firstTime,
		LastCommit:       lastTime,
		UniqueAuthors:    authorCount,
		AvgCommitsPerDay: float64(commitCount) / float64(days),
	}
}

func runGit(args string) string {
	out, _ := exec.Command("bash", "-c", "git "+args).Output()
	return strings.TrimSpace(string(out))
}

func toInt(s string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	return n
}

func getRepoName() string {
	out, _ := exec.Command("bash", "-c", "basename `git rev-parse --show-toplevel`").Output()
	return strings.TrimSpace(string(out))
}
