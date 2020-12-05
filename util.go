package box

import (
	"fmt"

	"strings"

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

// checkColorType checks the type of b.Color then from the preferences and options
func (b Box) checkColorType(TopBar, BottomBar string) (string, string) {
	if b.Color != nil {
		// Check if type of b.Color is string
		if str, ok := b.Color.(string); ok {
			// Hi Intensity Colors
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
				// Return TopBar and BottomBar with a warning as Color provided as a string is unknown
				errorMsg("[warning]: invalid value provided to Color, using default")
				return TopBar, BottomBar
			}
			// Check if type of b.Color is uint
		} else if hex, ok := b.Color.(uint); ok {
			// Break down the hex into r, g and b respectively
			hexArray := [3]uint{hex >> 16, hex >> 8 & 0xff, hex & 0xff}
			col := color.RGB(uint8(hexArray[0]), uint8(hexArray[1]), uint8(hexArray[2]))
			TopBar, BottomBar = roundOffColor(col, TopBar, BottomBar)
			// Check if type of b.Color is uint
		} else if rgb, ok := b.Color.([3]uint); ok {
			col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))
			TopBar, BottomBar = roundOffColor(col, TopBar, BottomBar)
		} else {
			// Panic if b.Color is an unexpected type
			panic(fmt.Sprintf("expected string, [3]uint or uint not %T", b.Color))
		}
		return TopBar, BottomBar
	}
	// As b.Color is nil then apply no color effect and return
	return TopBar, BottomBar
}
