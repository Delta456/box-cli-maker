package box

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

// addVertPadding adds Vertical Padding
func (b Box) addVertPadding(len int) []string {
	padding := strings.Repeat(" ", len-2)
	vertical := b.obtainColor()

	var texts []string
	for i := 0; i < b.Py; i++ {
		texts = append(texts, (vertical + padding + vertical))
	}
	return texts
}

// findAlign checks ContentAlign and returns Alignment
func (b Box) findAlign() string {
	if b.ContentAlign == "Center" {
		return centerAlign
	} else if b.ContentAlign == "Right" {
		return rightAlign
	} else if b.ContentAlign == "Left" || b.ContentAlign == "" {
		return leftAlign
	} else {
		fmt.Fprintln(os.Stderr, color.RedString("[error]: invalid value provided to Alignment, using default"))
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

// rgb returns the custom rgb formed string
// Taken from https://github.com/vlang/v/blob/master/vlib/term/colors.v#L10-L12
func rgb(r, g, b uint, msg, open, close string) string {
	return fmt.Sprintf("\x1b[%s;2;%s;%s;%sm%s\x1b[%sm", open, fmt.Sprint(r), fmt.Sprint(g), fmt.Sprint(b), msg, close)
}

func rbg_struct(r [3]uint, msg string) string {
	for _, ele := range r {
		if ele > 0xff || ele < 0x0 {
			panic("rgb array elements must be less more than 0x0 and less than 0xFF")
		}
	}
	return rgb(r[0], r[1], r[2], msg, "38", "39")
}

// rgb_hex returns the custom rgb formed string from hexadecimal
// Taken from https://github.com/vlang/v/blob/master/vlib/term/colors.v#L22-L24
func rbg_hex(hex uint, msg string) string {
	if hex < 0x000000 || hex > 0xFFFFFF {
		panic(fmt.Sprint(hex, "must be more than 0x000000 and less than 0xFFFFFF"))
	}
	return rgb(hex>>16, hex>>8&0xFF, hex&0xFF, msg, "38", "39")
}
