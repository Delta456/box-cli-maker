package box

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

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
		// if ContentAlign isn't provided then by default Alignment is Left
	} else if b.ContentAlign == "Left" || b.ContentAlign == "" {
		return leftAlign
	} else {
		// raise a warning if the ContentAlign isn't invalid
		errorMsg("[warning]: invalid value provided to Alignment, using default")
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
			panic("rgb array elements must be less more than 0x0 and less than 0xFF")
		}
	}
	return rgb(r[0], r[1], r[2], msg, "38", "39")
}

// rgbHex returns the custom RGB formed string from hexadecimal
// Taken from https://github.com/vlang/v/blob/master/vlib/term/colors.v#L22-L24
// All the elements must be in a range of [0x000000, 0xFFFFFF]
func rgbHex(hex uint, msg string) string {
	if hex < 0x00_0000 || hex > 0xFF_FFFF {
		panic(fmt.Sprint(hex, "must be more than 0x000000 and less than 0xFFFFFF"))
	}
	return rgb(hex>>16, hex>>8&0xFF, hex&0xFF, msg, "38", "39")
}

// errorMsg prints the msg to os.Stderr in Red ANSI Color according to the system
func errorMsg(msg string) {
	if runtime.GOOS == "windows" {
		// Using Output instead of os.Stderr because Output will enable ANSI Color on Winodws Console
		fmt.Fprintln(Output, color.RedString(msg))
		// Using bufio.NewWriter for flushing the message to os.Stderr stream
		bufio.NewWriter(os.Stderr).Flush()
	} else {
		fmt.Fprintln(os.Stderr, color.RedString(msg))
	}
}
