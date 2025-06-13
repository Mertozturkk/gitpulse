package git

import (
	"os/exec"
	"strconv"
	"strings"
)

type Author struct {
	Name    string
	Commits int
}

func GetTopAuthors(topN int) []Author {
	cmd := exec.Command("git", "shortlog", "-s", "-n", "--all", "--no-merges")
	out, err := cmd.Output()
	if err != nil {
		return []Author{}
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var authors []Author

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		commitCount, _ := strconv.Atoi(parts[0])
		name := strings.Join(parts[1:], " ")
		authors = append(authors, Author{Name: name, Commits: commitCount})
	}

	if topN > 0 && topN < len(authors) {
		return authors[:topN]
	}

	return authors
}
