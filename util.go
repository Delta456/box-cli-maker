package box

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

// errorMsg prints the msg to os.Stderr in Red ANSI Color according to the system
func errorMsg(msg string) {
	fmt.Fprintln(os.Stderr, color.Red.Sprint(msg))
}

func detectTerminalColor() terminfo.ColorLevel {
	// Detect WSL as it has True Color support
	wsl, err := ioutil.ReadFile("/proc/sys/kernel/osrelease")
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

// roundOffColorVertical rounds off the 24 bit Color to the terminals maximum color capacity for Vertical.
func (b Box) roundOffColorVertical(col color.RGBColor) string {
	if runtime.GOOS != "windows" {
		if detectTerminalColor() == terminfo.ColorLevelHundreds {
			return col.C256().Sprint(b.Vertical)
		} else if detectTerminalColor() == terminfo.ColorLevelBasic {
			return col.C16().Sprint(b.Vertical)
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

// roundOffColorVertical rounds off the 24 bit Color to the terminals maximum color capacity for TopBar and BottomBar.
func roundOffColor(col color.RGBColor, topBar, bottomBar string) (string, string) {
	if runtime.GOOS != "windows" {
		if detectTerminalColor() == terminfo.ColorLevelHundreds {
			TopBar := col.C256().Sprint(topBar)
			BottomBar := col.C256().Sprint(bottomBar)
			return TopBar, BottomBar
		} else if detectTerminalColor() == terminfo.ColorLevelBasic {
			TopBar := col.C16().Sprint(topBar)
			BottomBar := col.C16().Sprint(bottomBar)
			return TopBar, BottomBar
		} else if detectTerminalColor() == terminfo.ColorLevelMillions {
			TopBar := col.Sprint(topBar)
			BottomBar := col.Sprint(bottomBar)
			return TopBar, BottomBar
		} else {
			fmt.Fprintln(os.Stderr, "[warning]: terminal does not support colors, using no effect")
			return topBar, bottomBar
		}
	} else {
		TopBar := col.Sprint(topBar)
		BottomBar := col.Sprint(bottomBar)
		return TopBar, BottomBar
	}
}

func (b Box) checkColorType(TopBar, BottomBar string) (string, string) {
	if b.Color != nil {
		if str, ok := b.Color.(string); ok {
			if strings.HasPrefix(str, "Hi") {
				if _, ok := fgHiColors[str]; ok {
					Style := fgHiColors[str].Sprint
					TopBar = Style(TopBar)
					BottomBar = Style(BottomBar)
				}
			} else if _, ok := fgColors[str]; ok {
				Style := fgColors[str].Sprint
				TopBar = Style(TopBar)
				BottomBar = Style(BottomBar)
			} else {
				errorMsg("[warning]: invalid value provided to Color, using default")
			}
		} else if hex, ok := b.Color.(uint); ok {
			hexArray := [3]uint{hex >> 16, hex >> 8 & 0xff, hex & 0xff}
			col := color.RGB(uint8(hexArray[0]), uint8(hexArray[1]), uint8(hexArray[2]))
			TopBar, BottomBar = roundOffColor(col, TopBar, BottomBar)
		} else if rgb, ok := b.Color.([3]uint); ok {
			col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))
			TopBar, BottomBar = roundOffColor(col, TopBar, BottomBar)
		} else {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("expected string, [3]uint or uint not %T using default", b.Color))
		}
		return TopBar, BottomBar
	}
	return TopBar, BottomBar
}
