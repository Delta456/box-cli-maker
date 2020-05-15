package box

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// addVertPadding adds Vertical Padding
func (b Box) addVertPadding(len int) []string {
	padding := strings.Repeat(" ", len-2)

	var texts []string
	for i := 0; i < b.Con.Py; i++ {
		texts = append(texts, (b.Vertical + padding + b.Vertical))
	}
	return texts
}

// findAlign checks ContentAlign and returns Alignment
func (b Box) findAlign() string {
	if b.Con.ContentAlign == "Center" {
		return centerAlign
	} else if b.Con.ContentAlign == "Right" {
		return rightAlign
	} else {
		return leftAlign
	}
}

// longestLine returns longest line
func longestLine(lines []string) int {
	longest := 0
	for _, line := range lines {
		length := utf8.RuneCountInString(line)
		if length > longest {
			longest = length
		}
	}
	return longest
}

func repeatWithString(c string, n int, str string) string {
	count := n - len(str) - 2
	bar := strings.Repeat(c, count)
	strNew := fmt.Sprintf(" %s %s", str, bar)
	return strNew
}
