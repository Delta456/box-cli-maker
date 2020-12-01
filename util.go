package box

import (
	//	"bufio"
	"fmt"
	"os"

	//	"runtime"
	"strings"

	//"github.com/fatih/color"
	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// expandedLine stores a tab-expanded line, and its visible length.
type expandedLine struct {
	line string // tab-expanded line
	len  int    // line's visible length
}

// addVertPadding adds Vertical Padding
func (b Box) addVertPadding(len int) []string {
	padding := strings.Repeat(" ", len-2)
	vertical := b.obtainColor()

	var texts = make([]string, 0, b.Py)
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
		// If ContentAlign isn't provided then by default Alignment is Left
	} else if b.ContentAlign == "Left" || b.ContentAlign == "" {
		return leftAlign
	} else {
		// Raise a warning if the ContentAlign isn't invalid
		errorMsg("[warning]: invalid value provided to Alignment, using default")
		return leftAlign
	}
}

// longestLine expands tabs in lines and determines longest visible
// return longest length and array of expanded lines
func longestLine(lines []string) (int, []expandedLine) {
	longest := 0
	var expandedLines []expandedLine
	var tmpLine strings.Builder
	var lineLen int

	for _, line := range lines {
		tmpLine.Reset()

		for _, c := range line {
			lineLen = runewidth.StringWidth(tmpLine.String())

			if c == '\t' {
				tmpLine.WriteString(strings.Repeat(" ", 8-(lineLen&7)))
			} else {
				tmpLine.WriteRune(c)
			}
		}

		lineLen = runewidth.StringWidth(tmpLine.String())
		expandedLines = append(expandedLines, expandedLine{tmpLine.String(), lineLen})

		if lineLen > longest {
			longest = lineLen
		}
	}

	return longest, expandedLines
}

func repeatWithString(c string, n int, str string) string {
	count := n - runewidth.StringWidth(str) - 2
	bar := strings.Repeat(c, count)
	strNew := fmt.Sprintf(" %s %s", str, bar)
	return strNew
}

// rgb returns the custom RGB formed string
// only works with 24 bit color supported terminals
// Taken from https://github.com/vlang/v/blob/master/vlib/term/colors.v#L10-L12
func rgb(r, g, b uint, msg, open, close string) string {
	return fmt.Sprintf("\x1b[%s;2;%s;%s;%sm%s\x1b[%sm", open, fmt.Sprint(r), fmt.Sprint(g), fmt.Sprint(b), msg, close)
}

// rgbArray returns the custom RGB formed string from [3]uint values
// All the elements must be in a range of [0x0, 0xFF]
func rgbArray(r [3]uint, msg string) string {
	for _, ele := range r {
		if ele > 0xFF || ele < 0x0 {
			panic("RGB Array Elements must be in a range of [0x0, 0xFF]")
		}
	}
	return rgb(r[0], r[1], r[2], msg, "38", "39")
}

// rgbHex returns the custom RGB formed string from hexadecimal
// Taken from https://github.com/vlang/v/blob/master/vlib/term/colors.v#L22-L24
// All the elements must be in a range of [0x000000, 0xFFFFFF]
func rgbHex(hex uint, msg string) string {
	if hex < 0x00_0000 || hex > 0xFF_FFFF {
		panic(fmt.Sprint(hex, "must be in a range of [0x000000, 0xFFFFFF]"))
	}
	return rgb(hex>>16, hex>>8&0xFF, hex&0xFF, msg, "38", "39")
}

// errorMsg prints the msg to os.Stderr in Red ANSI Color according to the system
func errorMsg(msg string) {
	fmt.Fprintln(os.Stderr, color.Red.Sprint(msg))
}
