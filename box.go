package box

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/term"
	"github.com/huandu/xstrings"
	"github.com/mattn/go-runewidth"
)

const (
	// 1 = separator, 2 = spacing, 3 = line; 4 = oddSpace; 5 = space; 6 = sideMargin
	centerAlign = "%[1]s%[2]s%[3]s%[4]s%[2]s%[1]s"
	leftAlign   = "%[1]s%[6]s%[3]s%[4]s%[2]s%[5]s%[1]s"
	rightAlign  = "%[1]s%[2]s%[4]s%[5]s%[3]s%[6]s%[1]s"

	defaultWrapDivisor = 3  // 2/3 of terminal width
	minWrapWidth       = 20 // Minimum width to wrap content
)

// Box renders styled borders around text content.
type Box struct {
	// topRight renders the glyph used in the upper-right corner.
	topRight string
	// topLeft renders the glyph used in the upper-left corner.
	topLeft string
	// vertical renders the glyph used for the left and right walls.
	vertical string
	// bottomRight renders the glyph used in the lower-right corner.
	bottomRight string
	// bottomLeft renders the glyph used in the lower-left corner.
	bottomLeft string
	// horizontal renders the glyph used for the top and bottom edges.
	horizontal string
	config
}

// config contains configuration options for the Box.
type config struct {
	py            int           // Vertical padding.
	px            int           // Horizontal padding.
	contentAlign  AlignType     // Alignment for content inside the box.
	style         BoxStyle      // Active box style preset.
	titlePos      TitlePosition // Where the title, if any, is rendered.
	titleAlign    AlignType     // Alignment for the title based on TitlePosition.
	titleColor    string        // ANSI color (or hex code) for the title.
	contentColor  string        // ANSI color (or hex code) for the content.
	color         string        // ANSI color (or hex code) for the box chrome.
	allowWrapping bool          // Whether long content may wrap.
	wrappingLimit int           // Custom wrap width when wrapping is enabled.
	styleSet      bool          // Tracks if a style preset has already been applied.
}

// NewBox creates a new Box with the box.Single style preset applied.
func NewBox() *Box {
	b := &Box{}
	b.Style(Single)
	return b
}

// Copy returns a shallow copy of the Box so further mutations do not affect the original.
//
// Useful for creating base styles and deriving multiple boxes from them.
func (b *Box) Copy() *Box {
	if b == nil {
		return nil
	}
	clone := *b
	return &clone
}

// Padding sets horizontal (px) and vertical (py) inner padding.
func (b *Box) Padding(px, py int) *Box {
	b.px = px
	b.py = py
	return b
}

// HPadding sets horizontal padding (left and right).
func (b *Box) HPadding(px int) *Box {
	b.px = px
	return b
}

// VPadding sets vertical padding (top and bottom).
func (b *Box) VPadding(py int) *Box {
	b.py = py
	return b
}

// Style selects one of the built-in BoxStyle presets.
//
// Common styles include box.Single, box.Double, box.Round, box.Bold,
// box.SingleDouble, box.DoubleSingle, box.Classic, box.Hidden, and box.Block.
//
// To make custom styles, call TopRight, TopLeft, BottomRight, BottomLeft,
// Horizontal, and Vertical after Style to override individual glyphs.
//
// Example:
//
//	b := box.NewBox()
//	b.TopRight("+").TopLeft("+").BottomRight("+").BottomLeft("_").Horizontal("-").Vertical("|")
func (b *Box) Style(box BoxStyle) *Box {
	b.style = box
	b.styleSet = true
	// Set the box style characters from predefined styles
	// This also allows manual overrides after setting style
	// and have a standard base.
	if styleDef, ok := boxes[box]; ok {
		b.BottomLeft(styleDef.bottomLeft).
			BottomRight(styleDef.bottomRight).
			TopLeft(styleDef.topLeft).
			TopRight(styleDef.topRight).
			Horizontal(styleDef.horizontal).
			Vertical(styleDef.vertical)
	}
	return b
}

// TopRight sets the glyph used in the upper-right corner.
func (b *Box) TopRight(glyph string) *Box {
	b.topRight = glyph
	return b
}

// TopLeft sets the glyph used in the upper-left corner.
func (b *Box) TopLeft(glyph string) *Box {
	b.topLeft = glyph
	return b
}

// BottomRight sets the glyph used in the lower-right corner.
func (b *Box) BottomRight(glyph string) *Box {
	b.bottomRight = glyph
	return b
}

// BottomLeft sets the glyph used in the lower-left corner.
func (b *Box) BottomLeft(glyph string) *Box {
	b.bottomLeft = glyph
	return b
}

// Horizontal sets the glyph used for the horizontal edges.
func (b *Box) Horizontal(glyph string) *Box {
	b.horizontal = glyph
	return b
}

// Vertical sets the glyph used for the vertical edges.
func (b *Box) Vertical(glyph string) *Box {
	b.vertical = glyph
	return b
}

// TitleColor sets the color used for the title text.
//
// Accepts one of the first 16 ANSI color name constants (e.g. box.Green,
// box.BrightRed) or a #RGB / #RRGGBB / rgb:RRRR/GGGG/BBBB /
// rgba:RRRR/GGGG/BBBB/AAAA value.
//
// Invalid colors cause Render to return an error.
func (b *Box) TitleColor(color string) *Box {
	b.titleColor = color
	return b
}

// ContentColor sets the color used for the content text.
//
// Accepts one of the first 16 ANSI color name constants (e.g. box.Green,
// box.BrightRed) or a #RGB / #RRGGBB / rgb:RRRR/GGGG/BBBB /
// rgba:RRRR/GGGG/BBBB/AAAA value.
//
// Invalid colors cause Render to return an error.
func (b *Box) ContentColor(color string) *Box {
	b.contentColor = color
	return b
}

// Color sets the color used for the box border (chrome).
//
// Accepts one of the first 16 ANSI color name constants (e.g. box.Green,
// box.BrightRed) or a #RGB / #RRGGBB / rgb:RRRR/GGGG/BBBB /
// rgba:RRRR/GGGG/BBBB/AAAA value.
//
// Invalid colors cause Render to return an error.
func (b *Box) Color(color string) *Box {
	b.color = color
	return b
}

// TitlePosition sets where the title is rendered relative to the box.
//
// Valid positions are box.Inside, box.Top, and box.Bottom.
func (b *Box) TitlePosition(pos TitlePosition) *Box {
	b.titlePos = pos
	return b
}

// WrapContent enables or disables automatic wrapping of content.
//
// When enabled, content is wrapped to fit roughly two-thirds of the terminal
// width by default. For custom limits or non-TTY outputs, use WrapLimit
// instead.
func (b *Box) WrapContent(allow bool) *Box {
	b.allowWrapping = allow
	return b
}

// WrapLimit enables wrapping and sets an explicit maximum width for content.
func (b *Box) WrapLimit(limit int) *Box {
	b.allowWrapping = true
	b.wrappingLimit = limit
	return b
}

// TitleAlign sets the horizontal alignment of the title.
//
// When TitlePosition is Inside, it aligns the title lines inside the box.
// When TitlePosition is Top or Bottom, it aligns the title on the border.
//
// Supported values are box.Left, box.Center, and box.Right.
func (b *Box) TitleAlign(align AlignType) *Box {
	b.titleAlign = align
	return b
}

// ContentAlign sets the horizontal alignment of content inside the box.
//
// Supported values are box.Left, box.Center, and box.Right.
func (b *Box) ContentAlign(align AlignType) *Box {
	b.contentAlign = align
	return b
}

// MustRender is like Render but panics if an error occurs.
//
// Use MustRender in examples or CLIs where failures should abort execution
// instead of being handled explicitly.
func (b *Box) MustRender(title, content string) string {
	s, err := b.Render(title, content)
	if err != nil {
		panic(err)
	}
	return s
}

// wrapContent applies wrapping to the content string based on the Box configuration.
func (b *Box) wrapContent(content string) (string, error) {
	if !b.allowWrapping {
		return content, nil
	}
	if b.wrappingLimit < 0 {
		return "", fmt.Errorf("wrapping limit cannot be negative")
	}
	// If limit not provided then use 2*TermWidth/3 as limit else
	// use the one provided
	if b.wrappingLimit != 0 {
		return ansi.Wrap(content, b.wrappingLimit, ""), nil
	}
	if !isTTY(os.Stdout.Fd()) {
		return "", fmt.Errorf("cannot determine terminal width; use WrapLimit to set an explicit wrap limit when wrapping on non-TTY outputs")
	}
	width, _, err := term.GetSize(os.Stdout.Fd())
	if err != nil {
		return "", fmt.Errorf("cannot determine terminal width: %v", err)
	}
	// Use 2/3 of terminal width as default wrapping limit
	wrapWidth := max(2*width/defaultWrapDivisor, minWrapWidth)
	return ansi.Wrap(content, wrapWidth, ""), nil
}

// boxLayout holds the computed dimensions needed to render a box.
type boxLayout struct {
	innerWidth      int
	longestLine     int
	lineWidth       int
	horizontalWidth int
	lines           []expandedLine
	sideMargin      string
}

// prepareContentLines validates the title position and padding, splits the title
// and content into display lines, and returns those lines along with the number
// of title lines.
func (b *Box) prepareContentLines(title, content string) ([]string, int, error) {
	if b.titlePos == "" {
		b.titlePos = Inside
	} else if b.titlePos != Inside && b.titlePos != Top && b.titlePos != Bottom {
		return nil, 0, fmt.Errorf("invalid TitlePosition %s", b.titlePos)
	}

	var contentLines []string
	if title != "" {
		if b.titlePos != Inside && strings.Contains(title, "\n") {
			return nil, 0, fmt.Errorf("multiline titles are only supported Inside title position only")
		}
		if b.titlePos == Inside {
			contentLines = append(contentLines, strings.Split(title, "\n")...)
			contentLines = append(contentLines, "") // empty line between title and content
		}
	}
	contentLines = append(contentLines, strings.Split(content, "\n")...)

	titleLen := 0
	if title != "" {
		titleLen = len(strings.Split(ansi.Strip(title), "\n"))
	}

	if b.px < 0 {
		return nil, 0, fmt.Errorf("horizontal padding cannot be negative")
	}
	if b.py < 0 {
		return nil, 0, fmt.Errorf("vertical padding cannot be negative")
	}

	return contentLines, titleLen, nil
}

// computeLayout determines all the box dimensions from the prepared content
// lines and the (possibly empty) title string.
func (b *Box) computeLayout(contentLines []string, title string) boxLayout {
	sideMargin := strings.Repeat(" ", b.px)
	longest, lines := longestLine(contentLines)

	contentInnerWidth := longest + 2*b.px
	innerWidth := contentInnerWidth

	// Make sure the box is wide enough to fit the title when it's on Top/Bottom.
	if b.titlePos != Inside && title != "" {
		titleWidth := runewidth.StringWidth(ansi.Strip(title))
		if minW := titleWidth + 2; minW > innerWidth {
			innerWidth = minW
		}
	}

	// If we enlarged the inner width to fit the title, reflect that in longestLine.
	if innerWidth > contentInnerWidth {
		longest = max(innerWidth-2*b.px, 0)
	}

	verticalWidth := charWidth(b.vertical)
	horizontalWidth := charWidth(b.horizontal)

	// Ensure the inner width is a multiple of the horizontal glyph width so
	// the bar is visually uniform.
	if horizontalWidth > 1 && innerWidth%horizontalWidth != 0 {
		innerWidth += horizontalWidth - (innerWidth % horizontalWidth)
		longest = max(innerWidth-2*b.px, 0)
	}

	return boxLayout{
		innerWidth:      innerWidth,
		longestLine:     longest,
		lineWidth:       innerWidth + 2*verticalWidth,
		horizontalWidth: horizontalWidth,
		lines:           lines,
		sideMargin:      sideMargin,
	}
}

// buildAndColorBars constructs the top and bottom bars (optionally embedding a title)
// and applies both box-chrome and title coloring.
func (b *Box) buildAndColorBars(title string, lay boxLayout) (string, string, error) {
	tlw := charWidth(b.topLeft)
	trw := charWidth(b.topRight)
	blw := charWidth(b.bottomLeft)
	brw := charWidth(b.bottomRight)

	topBar := buildPlainBar(b.topLeft, b.horizontal, b.topRight, tlw, trw, lay.lineWidth, lay.horizontalWidth)
	bottomBar := buildPlainBar(b.bottomLeft, b.horizontal, b.bottomRight, blw, brw, lay.lineWidth, lay.horizontalWidth)

	if b.titlePos != Inside {
		switch b.titlePos {
		case Top:
			if title != "" {
				align, err := b.findTitleAlign(Left)
				if err != nil {
					return "", "", err
				}
				topBar = buildTitledBar(b.topLeft, b.horizontal, b.topRight, tlw, trw, lay.lineWidth, lay.horizontalWidth, title, align)
			}
		case Bottom:
			if title != "" {
				align, err := b.findTitleAlign(Left)
				if err != nil {
					return "", "", err
				}
				bottomBar = buildTitledBar(b.bottomLeft, b.horizontal, b.bottomRight, blw, brw, lay.lineWidth, lay.horizontalWidth, title, align)
			}
		}
	}

	var err error
	if topBar, err = applyColor(topBar, b.color); err != nil {
		return "", "", err
	}
	if bottomBar, err = applyColor(bottomBar, b.color); err != nil {
		return "", "", err
	}

	// Apply title coloring to the bars, expanding tabs in the title if needed.
	titleForBar := title
	if strings.Contains(titleForBar, "\t") {
		titleForBar = xstrings.ExpandTabs(titleForBar, 4)
	}
	if topBar, bottomBar, err = b.applyColorBar(topBar, bottomBar, titleForBar); err != nil {
		return "", "", err
	}

	return topBar, bottomBar, nil
}

// Render generates the box with the given title and content.
//
// It returns an error if:
//   - the BoxStyle is invalid,
//   - the TitlePosition is invalid,
//   - the wrapping limit is negative,
//   - padding is negative,
//   - a multiline title is used with a non-Inside TitlePosition, or
//   - any configured colors are invalid.
func (b *Box) Render(title, content string) (string, error) {
	if b.styleSet {
		if _, ok := boxes[b.style]; !ok {
			return "", fmt.Errorf("invalid Box style %s", b.style)
		}
	}

	content, err := b.wrapContent(content)
	if err != nil {
		return "", err
	}

	title, err = applyColor(title, b.titleColor)
	if err != nil {
		return "", err
	}
	content, err = applyColor(content, b.contentColor)
	if err != nil {
		return "", err
	}

	contentLines, titleLen, err := b.prepareContentLines(title, content)
	if err != nil {
		return "", err
	}

	lay := b.computeLayout(contentLines, title)

	topBar, bottomBar, err := b.buildAndColorBars(title, lay)
	if err != nil {
		return "", err
	}

	texts, err := b.addVertPadding(lay.innerWidth)
	if err != nil {
		return "", err
	}
	texts, err = b.formatLine(lay.lines, lay.longestLine, titleLen, lay.sideMargin, title, texts)
	if err != nil {
		return "", err
	}
	vertPadding, err := b.addVertPadding(lay.innerWidth)
	if err != nil {
		return "", err
	}
	texts = append(texts, vertPadding...)

	var sb strings.Builder
	sb.WriteString(topBar)
	sb.WriteString("\n")
	sb.WriteString(strings.Join(texts, "\n"))
	sb.WriteString("\n")
	sb.WriteString(bottomBar)
	sb.WriteString("\n")

	return sb.String(), nil
}
