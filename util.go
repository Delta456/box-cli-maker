package box

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
)

// addVertPadding adds Vertical Padding
func (b Box) addVertPadding(len int) []string {
	padding := strings.Repeat(" ", len-2)
	vertical := b.obtainColor()

	var texts []string
	for i := 0; i < b.Con.Py; i++ {
		texts = append(texts, (vertical + padding + vertical))
	}
	return texts
}

// findAlign checks ContentAlign and returns Alignment
func (b Box) findAlign() string {
	if b.Con.ContentAlign == "Center" {
		return centerAlign
	} else if b.Con.ContentAlign == "Right" {
		return rightAlign
	} else if b.Con.ContentAlign == "Left" {
		return leftAlign
	} else {
		fmt.Fprintln(os.Stderr, "Invalid Alignment provided, using default Alignment")
		return leftAlign
	}
}

// longestLine returns longest line
func longestLine(lines []string) int {
	longest := 0
	for _, line := range lines {
		length := runewidth.StringWidth(line)
		if length > longest {
			longest = length
		}
	}
	return longest
}

func repeatWithString(c string, n int, str string) string {
	count := n - runewidth.StringWidth(str) - 2
	bar := strings.Repeat(c, count)
	strNew := fmt.Sprintf(" %s %s", str, bar)
	return strNew
}
