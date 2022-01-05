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
	vertical := b.obtainBoxColor()

	var texts = make([]string, 0, b.Py)
	for i := 0; i < b.Py; i++ {
		texts = append(texts, (vertical + padding + vertical))
	}
	return texts
}

// findAlign checks ContentAlign and returns Alignment
func (b Box) findAlign() string {
	switch b.ContentAlign {
	case "Center":
		return centerAlign
	case "Right":
		return rightAlign
	case "Left":
	case "":
		// If ContentAlign isn't provided then by default Alignment is Left
		return leftAlign
	default:
		// Raise a warning if the ContentAlign isn't invalid
		errorMsg("[warning]: invalid value provided to Alignment, using default")
		return leftAlign
	}
	return leftAlign
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

		// Check if each line has ANSI Color Code then decrease the length accordingly
		if runewidth.StringWidth(color.ClearCode(tmpLine.String())) < runewidth.StringWidth(tmpLine.String()) {
			lineLen = runewidth.StringWidth(color.ClearCode(tmpLine.String()))
		}

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
func (b Box) checkColorType(topBar, bottomBar string) (string, string) {
	if b.Color != nil {
		// Check if type of b.Color is string
		if str, ok := b.Color.(string); ok {
			// Hi Intensity Colors
			if strings.HasPrefix(str, "Hi") {
				if _, ok := fgHiColors[str]; ok {
					Style := fgHiColors[str].Sprint
					topBar = Style(topBar)
					bottomBar = Style(bottomBar)
				}
			} else if _, ok := fgColors[str]; ok {
				Style := fgColors[str].Sprint
				topBar = Style(topBar)
				bottomBar = Style(bottomBar)
			} else {
				// Return TopBar and BottomBar with a warning as Color provided as a string is unknown
				errorMsg("[warning]: invalid value provided to Color, using default")
				return topBar, bottomBar
			}

			// If TitlePos isn't Inside and TitleColor isn't nil then clear Color Code and then split out Title from Top/Bottom Bars
			// then concatenate them again with the colors provided. This is done so that the color of Vertical after Title
			// won't be in effect.
			// TDLR: color.Red("Hello") + color.Yellow("World") + color.Red("!") != color.Red("Hello" + color.Yellow("World") + "!")
			if b.TitleColor != nil {
				if b.TitlePos == "Top" {

				} else if b.TitlePos == "Bottom" {

				}
			}

			// Check if type of b.Color is uint
		} else if hex, ok := b.Color.(uint); ok {
			// Break down the hex into r, g and b respectively
			hexArray := [3]uint{hex >> 16, hex >> 8 & 0xff, hex & 0xff}
			col := color.RGB(uint8(hexArray[0]), uint8(hexArray[1]), uint8(hexArray[2]))
			topBar, bottomBar = roundOffColor(col, topBar, bottomBar)
			// Check if type of b.Color is uint
		} else if rgb, ok := b.Color.([3]uint); ok {
			col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))
			topBar, bottomBar = roundOffColor(col, topBar, bottomBar)
		} else {
			// Panic if b.Color is an unexpected type
			panic(fmt.Sprintf("expected string, [3]uint or uint not %T", b.Color))
		}
		return topBar, bottomBar
	}
	// As b.Color is nil then apply no color effect and return
	return topBar, bottomBar
}

// formatLine formats the line according to the information passed
func (b Box) formatLine(lines2 []expandedLine, longestLine, titleLen int, sideMargin, title string, texts []string) []string {
	for i, line := range lines2 {
		length := line.len

		// Use later
		var space, oddSpace string

		// Check if the line.line has ANSI Color Code then decrease the length accordingly
		if runewidth.StringWidth(color.ClearCode(line.line)) < runewidth.StringWidth(line.line) {
			length = runewidth.StringWidth(color.ClearCode(line.line))
		}

		// If current text is shorter than the longest one
		// center the text, so it looks better
		if length < longestLine {
			// Difference between longest and current one
			diff := longestLine - length

			// the spaces to add on each side
			toAdd := diff / 2
			space = strings.Repeat(" ", toAdd)

			// If the difference between the longest and current one
			// is odd, we have to add one additional space before the last vertical separator
			if diff%2 != 0 {
				oddSpace = " "
			}
		}

		spacing := space + sideMargin
		var format string

		if i < titleLen && title != "" && b.TitlePos == inside {
			format = centerAlign
		} else {
			format = b.findAlign()
		}

		// Obtain color
		sep := b.obtainBoxColor()

		formatted := fmt.Sprintf(format, sep, spacing, line.line, oddSpace, space, sideMargin)
		texts = append(texts, formatted)
	}
	return texts
}
