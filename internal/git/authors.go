package git

import (
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type Author struct {
	Name    string
	Commits int
}

func GetTopAuthors() []Author {
	cmd := exec.Command("git", "shortlog", "-s", "-n", "--all", "--no-merges")
	out, err := cmd.Output()
	if err != nil {
		return []Author{}
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var authors []Author

	for _, line := range lines {
		// örnek satır: "   42\tJohn Doe"
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		commitCount, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		name := strings.Join(parts[1:], " ")
		authors = append(authors, Author{Name: name, Commits: commitCount})
	}

	// yine de garanti sıralı olsun
	sort.Slice(authors, func(i, j int) bool {
		return authors[i].Commits > authors[j].Commits
	})

	return authors
}
