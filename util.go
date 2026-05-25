package box

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
	"github.com/huandu/xstrings"
	"github.com/mattn/go-runewidth"
)

// isTTY points to the function used to determine if a file descriptor is a terminal.
// It is defined as a variable to allow mocking in tests.
var isTTY = term.IsTerminal

// expandedLine stores a tab-expanded line, and its visible length.
type expandedLine struct {
	line string // tab-expanded line
	len  int    // line's visible length
}

// addVertPadding adds vertical padding lines using the given inner width.
//
// innerWidth represents the visible width between the vertical borders.
func (b *Box) addVertPadding(innerWidth int) ([]string, error) {
	if innerWidth < 0 {
		innerWidth = 0
	}
	padding := strings.Repeat(" ", innerWidth)
	vertical, err := applyColor(b.vertical, b.color)
	if err != nil {
		return nil, err
	}

	texts := make([]string, b.py)
	for i := range texts {
		texts[i] = vertical + padding + vertical
	}

	return texts, nil
}

// longestLine expands tabs in lines and determines longest visible
// return longest length and array of expanded lines
func longestLine(lines []string) (int, []expandedLine) {
	longest := 0
	expandedLines := make([]expandedLine, 0, len(lines))
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
		if runewidth.StringWidth(ansi.Strip(tmpLine.String())) < runewidth.StringWidth(tmpLine.String()) {
			lineLen = runewidth.StringWidth(ansi.Strip(tmpLine.String()))
		}

		if lineLen > longest {
			longest = lineLen
		}
	}
	return longest, expandedLines
}

// charWidth returns the visible width of a string, treating zero-width
// results as width 1 so that box calculations always make progress.
func charWidth(s string) int {
	w := runewidth.StringWidth(ansi.Strip(s))
	if w == 0 {
		w = 1
	}
	return w
}

// buildSegment builds a horizontal segment with the given visual width using
// the provided fill glyph, padding with spaces if needed to match width.
func buildSegment(fill string, width, horizontalWidth int) string {
	if width <= 0 {
		return ""
	}
	fillCount := width / horizontalWidth
	seg := strings.Repeat(fill, fillCount)
	padWidth := width - fillCount*horizontalWidth
	if padWidth > 0 {
		seg += strings.Repeat(" ", padWidth)
	}
	return seg
}

// buildAlignedSegment builds a horizontal segment with the given visual width
// and aligns any padding so the fill glyph remains adjacent to the chosen edge.
func buildAlignedSegment(fill string, width, horizontalWidth int, attachLeft bool) string {
	if width <= 0 {
		return ""
	}
	if attachLeft {
		return buildSegment(fill, width, horizontalWidth)
	}
	gapWidth := width % horizontalWidth
	fillWidth := width - gapWidth
	seg := buildSegment(fill, fillWidth, horizontalWidth)
	if gapWidth == 0 {
		return seg
	}
	return strings.Repeat(" ", gapWidth) + seg
}

// buildPlainBar builds a horizontal bar (without title) that matches the
// specified visual line width.
func buildPlainBar(left, fill, right string, leftW, rightW, lineWidth, horizontalWidth int) string {
	inner := max(lineWidth-leftW-rightW, 0)
	bar := buildSegment(fill, inner, horizontalWidth)
	return left + bar + right
}

// buildTitledBar builds a top or bottom bar containing a title with the given
// alignment. Any leftover width that is not divisible by the glyph's width is
// emitted as spaces so the fill glyph remains adjacent to the corners.
func buildTitledBar(left, fill, right string, leftW, rightW, lineWidth, horizontalWidth int, title string, align AlignType) string {
	if title == "" {
		return buildPlainBar(left, fill, right, leftW, rightW, lineWidth, horizontalWidth)
	}

	plainTitle := title
	if strings.Contains(plainTitle, "\t") {
		plainTitle = xstrings.ExpandTabs(plainTitle, 4)
	}
	titleWidth := runewidth.StringWidth(ansi.Strip(plainTitle))
	titleSegWidth := titleWidth + 2 // one space padding on each side

	inner := max(lineWidth-leftW-rightW, titleSegWidth)
	remaining := inner - titleSegWidth

	leftWidth := 0
	rightWidth := 0
	switch align {
	case Center:
		leftWidth = remaining / 2
		rightWidth = remaining - leftWidth
	case Right:
		leftWidth = remaining
	case Left, "":
		rightWidth = remaining
	default:
		rightWidth = remaining
	}

	leftSeg := buildAlignedSegment(fill, leftWidth, horizontalWidth, true)
	rightSeg := buildAlignedSegment(fill, rightWidth, horizontalWidth, false)

	return left + leftSeg + " " + plainTitle + " " + rightSeg + right
}

// formatLine formats the line according to the information passed.
func (b *Box) formatLine(lines2 []expandedLine, longestLine, titleLen int, sideMargin, title string, texts []string) ([]string, error) {
	for i, line := range lines2 {
		length := line.len

		// Use later
		var space, oddSpace string

		// compute stripped width once
		strippedWidth := runewidth.StringWidth(ansi.Strip(line.line))
		if strippedWidth < runewidth.StringWidth(line.line) {
			length = strippedWidth
		}

		// If current text is shorter than the longest one
		// center the text, so it looks better
		if length < longestLine {
			// Difference between longest and current one
			diff := longestLine - length

			// the spaces to add on each side
			toAdd := diff / 2
			space = strings.Repeat(" ", toAdd)

			// If difference between the longest and current one
			// is odd, we have to add one additional space before the last vertical separator
			if diff%2 != 0 {
				oddSpace = " "
			}
		}

		spacing := space + sideMargin

		var format string
		var err error

		if i < titleLen && title != "" && b.titlePos == Inside {
			format, err = b.findTitleAlignFormat(Center)
		} else {
			format, err = b.findContentAlign()
		}
		if err != nil {
			return nil, err
		}

		sep, err := applyColor(b.vertical, b.color)
		if err != nil {
			return nil, err
		}

		formatted := fmt.Sprintf(format, sep, spacing, line.line, oddSpace, space, sideMargin)
		texts = append(texts, formatted)
	}
	return texts, nil
}

func (b *Box) findContentAlign() (string, error) {
	switch b.contentAlign {
	case Center:
		return centerAlign, nil
	case Right:
		return rightAlign, nil
	case Left, "":
		// If ContentAlign isn't provided then by default Alignment is Left
		return leftAlign, nil
	default:
		return "", fmt.Errorf("invalid Content Alignment %s", b.contentAlign)
	}
}

func (b *Box) findTitleAlign(defaultAlign AlignType) (AlignType, error) {
	switch b.titleAlign {
	case "":
		return defaultAlign, nil
	case Center, Right, Left:
		return b.titleAlign, nil
	default:
		return "", fmt.Errorf("invalid Title Alignment %s", b.titleAlign)
	}
}

func (b *Box) findTitleAlignFormat(defaultAlign AlignType) (string, error) {
	align, err := b.findTitleAlign(defaultAlign)
	if err != nil {
		return "", err
	}
	switch align {
	case Center:
		return centerAlign, nil
	case Left:
		return leftAlign, nil
	default:
		return rightAlign, nil
	}
}

func getConvertedColor(colorStr string) (color.Color, error) {
	cv, err := parseColorString(colorStr)
	if err != nil {
		return nil, err
	}
	// If profile conversion results in nil, fall back to the
	// parsed color so we always emit color.
	converted := profile.Convert(cv)
	if converted == nil {
		return cv, nil
	}
	return converted, nil
}

func applyColor(str string, colorStr string) (string, error) {
	// Empty color string means: do not apply any styling.
	if colorStr == "" {
		return str, nil
	}
	convertedColor, err := getConvertedColor(colorStr)
	if err != nil {
		return str, err
	}
	return applyConvertedColor(str, convertedColor), nil
}

func stringColorToHex(color string) string {
	if hex, exists := colorToHex[color]; exists {
		return hex
	}
	// Return empty string for unknown colors to let ansi.XParseColor handle it
	return ""
}

// addStylePreservingOriginalFormat allows to add style around the original formating
func addStylePreservingOriginalFormat(s string, f func(a string) string) string {
	const reset = "\033[0m"
	if !strings.Contains(s, reset) {
		return f(s)
	}

	var sb strings.Builder
	start := 0
	for {
		idx := strings.Index(s[start:], reset)
		if idx == -1 {
			sb.WriteString(f(s[start:]))
			break
		}
		sb.WriteString(f(s[start : start+idx]))
		// skip the reset sequence (preserve original behavior of removing it)
		start += idx + len(reset)
	}
	return sb.String()
}

// parseColorString converts a color string to color.Color using stringColorToHex and ansi.XParseColor
func parseColorString(colorStr string) (color.Color, error) {
	hexColor := stringColorToHex(colorStr)

	if hexColor == "" {
		hexColor = colorStr
	}

	colorValue := ansi.XParseColor(hexColor)
	if colorValue == nil {
		return nil, fmt.Errorf("unable to parse color: %s", colorStr)
	}
	return colorValue, nil
}

func applyConvertedColor(str string, c color.Color) string {
	if c == nil {
		return str
	}

	style := ansi.Style{}.ForegroundColor(c)
	styled := style.Styled

	// Fast path: no newlines
	if !strings.Contains(str, "\n") {
		return addStylePreservingOriginalFormat(str, styled)
	}

	var sb strings.Builder
	start := 0
	for {
		idx := strings.IndexByte(str[start:], '\n')
		if idx == -1 {
			sb.WriteString(addStylePreservingOriginalFormat(str[start:], styled))
			break
		}
		sb.WriteString(addStylePreservingOriginalFormat(str[start:start+idx], styled))
		sb.WriteByte('\n')
		start += idx + 1
	}
	return sb.String()
}

func (b *Box) applyColorBar(topBar, bottomBar, title string) (string, string, error) {
	if b.titleColor == "" || title == "" {
		return topBar, bottomBar, nil
	}

	if strings.TrimSpace(b.color) == "" {
		return topBar, bottomBar, nil
	}

	converted, err := getConvertedColor(b.color)
	if err != nil {
		return "", "", err
	}

	if b.titlePos == Top {
		strippedBar := ansi.Strip(topBar)
		strippedTitle := ansi.Strip(title)
		if idx := strings.Index(strippedBar, strippedTitle); idx != -1 {
			// split around first occurrence to preserve any other repeats
			b0 := applyConvertedColor(strippedBar[:idx], converted)
			b1 := applyConvertedColor(strippedBar[idx+len(strippedTitle):], converted)
			coloredTitle, err := applyColor(title, b.titleColor)
			if err != nil {
				return "", "", err
			}
			topBar = b0 + coloredTitle + b1
		}
	}

	if b.titlePos == Bottom {
		strippedBar := ansi.Strip(bottomBar)
		strippedTitle := ansi.Strip(title)
		if idx := strings.Index(strippedBar, strippedTitle); idx != -1 {
			// split around first occurrence to preserve any other repeats
			b0 := applyConvertedColor(strippedBar[:idx], converted)
			b1 := applyConvertedColor(strippedBar[idx+len(strippedTitle):], converted)
			coloredTitle, err := applyColor(title, b.titleColor)
			if err != nil {
				return "", "", err
			}
			bottomBar = b0 + coloredTitle + b1
		}
	}

	return topBar, bottomBar, nil
}
