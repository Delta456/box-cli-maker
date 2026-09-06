package box

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
	"github.com/mattn/go-runewidth"
)

// isTTY points to the function used to determine if a file descriptor is a terminal.
// It is defined as a variable to allow mocking in tests.
var isTTY = term.IsTerminal

// getTermSize returns the dimensions of the terminal attached to fd.
// It is defined as a variable to allow mocking in tests.
var getTermSize = term.GetSize

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

// expandTabs expands tab characters in s using tab stops at every 8 columns,
// consistent with POSIX terminal behavior.
// ANSI escape sequences (CSI, OSC hyperlinks and titles, and other escapes)
// are passed through without affecting the column count; parsing is delegated
// to ansi.DecodeSequenceWc, which returns width 0 for every sequence.
func expandTabs(s string) string {
	if !strings.Contains(s, "\t") {
		return s
	}

	var b strings.Builder
	colPos := 0
	state := byte(ansi.NormalState)
	for len(s) > 0 {
		seq, width, n, newState := ansi.DecodeSequenceWc(s, state, nil)
		if n == 0 {
			// Defensive: never loop on undecodable input.
			b.WriteString(s)
			break
		}
		switch seq {
		case "\t":
			spaces := 8 - (colPos & 7)
			b.WriteString(strings.Repeat(" ", spaces))
			colPos += spaces
		case "\n":
			b.WriteByte('\n')
			colPos = 0
		default:
			b.WriteString(seq)
			colPos += width
		}
		s = s[n:]
		state = newState
	}
	return b.String()
}

// longestLine expands tabs in lines and determines longest visible
// return longest length and array of expanded lines
func longestLine(lines []string) (int, []expandedLine) {
	longest := 0
	expandedLines := make([]expandedLine, 0, len(lines))

	for _, line := range lines {
		expanded := expandTabs(line)
		lineLen := runewidth.StringWidth(expanded)

		// Use visible width: strip ANSI codes before measuring.
		if stripped := runewidth.StringWidth(ansi.Strip(expanded)); stripped < lineLen {
			lineLen = stripped
		}

		expandedLines = append(expandedLines, expandedLine{expanded, lineLen})

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
func (b *Box) buildPlainBar(left, right string, lineWidth int) string {
	fill := b.horizontal
	leftW := charWidth(left)
	rightW := charWidth(right)
	horizontalWidth := charWidth(fill)
	inner := max(lineWidth-leftW-rightW, 0)
	bar := buildSegment(fill, inner, horizontalWidth)
	return left + bar + right
}

// buildTitledBar builds a top or bottom bar containing a title with the given
// alignment. Any leftover width that is not divisible by the glyph's width is
// emitted as spaces so the fill glyph remains adjacent to the corners.
// buildTitledBar returns the assembled bar string and the byte offset within
// that string where plainTitle begins. The offset is -1 when title is empty.
func (b *Box) buildTitledBar(left, right string, lineWidth int, title string, align AlignType) (string, int) {
	fill := b.horizontal
	leftW := charWidth(left)
	rightW := charWidth(right)
	horizontalWidth := charWidth(fill)

	if title == "" {
		return b.buildPlainBar(left, right, lineWidth), -1
	}

	plainTitle := expandTabs(title)
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

	// prefix contains no ANSI, so len(prefix) is the title offset in both
	// the raw bar and the ANSI-stripped bar.
	prefix := left + leftSeg + " "
	return prefix + plainTitle + " " + rightSeg + right, len(prefix)
}

// formatLine formats the line according to the information passed.
func (b *Box) formatLine(lines2 []expandedLine, longestLine, titleLen int, sideMargin, title string, texts []string) ([]string, error) {
	sep, err := applyColor(b.vertical, b.color)
	if err != nil {
		return nil, err
	}

	for i, line := range lines2 {
		length := line.len

		var space, oddSpace string

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

		if i < titleLen && title != "" && (b.titlePos == Inside || b.titlePos == "") {
			format, err = b.findTitleAlignFormat(Center)
		} else {
			format, err = b.findContentAlign()
		}
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
	// A nil conversion means the detected color profile disables color
	// output entirely (NO_COLOR, TERM=dumb, or a non-TTY stdout);
	// applyConvertedColor treats nil as "no styling".
	return profile.Convert(cv), nil
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

// nextReset returns the index and byte length of the earliest SGR reset
// sequence in s, recognizing both the long form "\x1b[0m" and the short form
// "\x1b[m" (an omitted parameter defaults to 0). Returns -1, 0 when s
// contains no reset.
func nextReset(s string) (int, int) {
	idxLong := strings.Index(s, "\033[0m")
	idxShort := strings.Index(s, "\033[m")
	switch {
	case idxShort == -1:
		return idxLong, len("\033[0m")
	case idxLong == -1 || idxShort < idxLong:
		return idxShort, len("\033[m")
	default:
		return idxLong, len("\033[0m")
	}
}

// addStylePreservingOriginalFormat allows to add style around the original formating.
// The string is split at every SGR reset (either spelling; see nextReset) and f is
// applied to each segment separately, so the outer style is re-armed after any
// reset embedded in the original string. The resets themselves are removed; f is
// expected to terminate each styled segment with its own reset.
func addStylePreservingOriginalFormat(s string, f func(a string) string) string {
	var sb strings.Builder
	start := 0
	for {
		idx, n := nextReset(s[start:])
		if idx == -1 {
			if start == 0 {
				return f(s)
			}
			sb.WriteString(f(s[start:]))
			break
		}
		sb.WriteString(f(s[start : start+idx]))
		// skip the reset sequence (preserve original behavior of removing it)
		start += idx + n
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
	// Never style an empty string: coloring "" would produce a non-empty,
	// ANSI-only string, and callers use emptiness checks (e.g. title != "")
	// to decide whether an element exists at all.
	if c == nil || str == "" {
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

func (b *Box) applyColorBar(topBar, bottomBar, title string, titleOffset int) (string, string, error) {
	if b.titlePos != Top && b.titlePos != Bottom {
		return topBar, bottomBar, nil
	}
	if b.titleColor == "" || title == "" {
		return topBar, bottomBar, nil
	}
	if b.color == "" {
		return topBar, bottomBar, nil
	}

	converted, err := getConvertedColor(b.color)
	if err != nil {
		return "", "", err
	}

	colorBarTitle := func(bar string) string {
		strippedBar := ansi.Strip(bar)
		strippedTitle := ansi.Strip(title)
		end := titleOffset + len(strippedTitle)
		if titleOffset < 0 || end > len(strippedBar) {
			return bar
		}
		b0 := applyConvertedColor(strippedBar[:titleOffset], converted)
		b1 := applyConvertedColor(strippedBar[end:], converted)
		return b0 + title + b1
	}

	if b.titlePos == Top {
		topBar = colorBarTitle(topBar)
	}
	if b.titlePos == Bottom {
		bottomBar = colorBarTitle(bottomBar)
	}

	return topBar, bottomBar, nil
}
