package box

import (
	//	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"strings"

	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
	"github.com/xo/terminfo"
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

// errorMsg prints the msg to os.Stderr in Red ANSI Color according to the system
func errorMsg(msg string) {
	fmt.Fprintln(os.Stderr, color.Red.Sprint(msg))
}

func detectTerminalColor() terminfo.ColorLevel {
	// Detect WSL as it has True Color support
	wsl, err := exec.Command("cat", "/proc/sys/kernel/osrelease").Output()
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(wsl), "microsoft") && strings.Contains(string(wsl), "Microsoft") {
		return terminfo.ColorLevelMillions
	}
	level, err := terminfo.ColorLevelFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	return level
}

func (b Box) roundOffColor(col color.RGBColor) string {
	if runtime.GOOS != "windows" {
		if detectTerminalColor() == terminfo.ColorLevelHundreds {
			return col.C256().Sprint(b.Vertical)
			// TOOD: make a rounding off logic for 24 bit to 8 bit
		} else if detectTerminalColor() == terminfo.ColorLevelBasic {
			return col.Sprint(b.Vertical)
		} else if detectTerminalColor() == terminfo.ColorLevelMillions {
			return col.Sprint(b.Vertical)
		} else {
			fmt.Fprintln(os.Stderr, "[warning]: terminal does not support colors, using no effect")
			return b.Vertical
		}
	} else {
		if !noColor {
			return col.Sprintf(b.Vertical)
		} else {
			fmt.Fprintln(os.Stderr, "[warning]: terminal does not support colors, using no effect")
			return b.Vertical
		}
	}
}
