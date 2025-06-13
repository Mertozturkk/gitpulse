package git

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var Weekdays = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

func GetCommitHeatmap(since string) map[string]int {
	args := []string{"log", "--pretty=format:%cd", "--date=iso"}
	if since != "" {
		if strings.HasSuffix(since, "d") {
			days, _ := strconv.Atoi(strings.TrimSuffix(since, "d"))
			t := time.Now().AddDate(0, 0, -days)
			args = append(args, fmt.Sprintf("--since=%s", t.Format("2006-01-02")))
		} else {
			args = append(args, fmt.Sprintf("--since=%s", since))
		}
	}

	cmd := exec.Command("git", args...)
	out, err := cmd.Output()
	if err != nil {
		return map[string]int{}
	}

	lines := strings.Split(string(out), "\n")
	counts := make(map[string]int)

	for _, line := range lines {
		t, err := time.Parse("2006-01-02 15:04:05 -0700", line)
		if err != nil {
			continue
		}
		day := t.Weekday().String()[:3]
		counts[day]++
	}

	return counts
}
