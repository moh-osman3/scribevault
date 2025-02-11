package diff

import (
	"strings"
)

type Line struct {
	linenumber int
	text string
}

func diffLines(s1, s2 string) ([]Line, []Line) {
	plus := []Line{}
	minus := []Line{}
	sLines1 := strings.Split(s1, "\n")
	sLines2 := strings.Split(s2, "\n")

	l1 := len(sLines1)
	l2 := len(sLines2)

	minLen := min(l1,l2)
	maxLen := max(l1,l2)

	for i := 0; i < maxLen; i++ {
		if i < minLen && sLines1[i] == sLines2[i] {
			continue
		}

		if i < l2 {
			plusLine := Line{
				linenumber: i,
				text: sLines2[i],
			}
			plus = append(plus, plusLine)
		}

		if i < l1 {
			minusLine := Line{
				linenumber: i,
				text: sLines1[i],
			}
			minus = append(minus, minusLine)
		}
	}

	return plus, minus
}