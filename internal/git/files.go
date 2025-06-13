package git

import (
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type DirStat struct {
	Path    string
	Changes int
}

func GetChangedDirectories() []DirStat {
	cmd := exec.Command("git", "log", "--name-only", "--pretty=format:", "| grep -v '^$'")
	out, err := cmd.Output()
	if err != nil {
		return []DirStat{}
	}

	lines := strings.Split(string(out), "\n")
	dirMap := make(map[string]int)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		dir := filepath.Dir(line)
		dirMap[dir]++
	}

	var stats []DirStat
	for path, count := range dirMap {
		stats = append(stats, DirStat{Path: path, Changes: count})
	}

	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Changes > stats[j].Changes
	})

	return stats
}
