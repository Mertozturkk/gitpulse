package git

import (
	"strings"
)

func Bar(n int) string {
	unit := n / 10
	if unit < 1 {
		unit = 1
	}
	return strings.Repeat("█", unit)
}
