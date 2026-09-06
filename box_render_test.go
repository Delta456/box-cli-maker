package box

import (
	"fmt"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-runewidth"
)

func TestRenderBasicBox(t *testing.T) {
	b := NewBox().Padding(2, 1).Style(Single)

	out, err := b.Render("Box CLI Maker", "Highly Customizable Terminal Box Maker")
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}
	if out == "" {
		t.Fatalf("expected non-empty render output")
	}

	if !strings.Contains(out, "Box CLI Maker") || !strings.Contains(out, "Highly Customizable Terminal Box Maker") {
		t.Fatalf("rendered output does not contain title/content: %q", out)
	}

	// Basic structural checks: top and bottom lines should use the Single style corners.
	lines := strings.Split(out, "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines in rendered box, got %d", len(lines))
	}

	// Last element is empty due to trailing newline; bottom bar is at len(lines)-2.
	top := lines[0]
	bottom := lines[len(lines)-2]

	if !strings.HasPrefix(top, "┌") || !strings.HasSuffix(top, "┐") {
		t.Errorf("top bar does not use Single style corners: %q", top)
	}
	if !strings.HasPrefix(bottom, "└") || !strings.HasSuffix(bottom, "┘") {
		t.Errorf("bottom bar does not use Single style corners: %q", bottom)
	}
}
func TestRenderInbuiltStylesPorts(t *testing.T) {
	tests := []BoxStyle{
		Single,
		SingleDouble,
		Double,
		DoubleSingle,
		Bold,
		Round,
		Hidden,
		Classic,
		Block,
	}

	for _, style := range tests {
		preset, ok := boxes[style]
		if !ok {
			t.Fatalf("no preset found for style %q", style)
		}

		b := NewBox().Padding(2, 5).Style(style)
		out, err := b.Render("Box CLI Maker", "Highly Customized Terminal Box Maker")
		if err != nil {
			t.Fatalf("Render returned error for style %q: %v", style, err)
		}

		lines := strings.Split(out, "\n")
		if len(lines) < 3 {
			t.Fatalf("style %q: expected at least 3 lines, got %d", style, len(lines))
		}

		// Last element is empty due to trailing newline; bottom bar is at len-2.
		top := lines[0]
		bottom := lines[len(lines)-2]

		if !strings.HasPrefix(top, preset.topLeft) || !strings.HasSuffix(top, preset.topRight) {
			t.Errorf("style %q: unexpected top corners: %q", style, top)
		}
		if !strings.HasPrefix(bottom, preset.bottomLeft) || !strings.HasSuffix(bottom, preset.bottomRight) {
			t.Errorf("style %q: unexpected bottom corners: %q", style, bottom)
		}

		// Check that interior lines use the expected vertical glyphs (including
		// the Hidden style, where vertical is a space).
		if len(lines) > 3 {
			interior := lines[1 : len(lines)-2]
			mid := interior[len(interior)/2]
			if len(mid) == 0 {
				t.Errorf("style %q: mid interior line unexpectedly empty", style)
			} else {
				if !strings.HasPrefix(mid, preset.vertical) || !strings.HasSuffix(mid, preset.vertical) {
					t.Errorf("style %q: unexpected vertical borders in interior line: %q", style, mid)
				}
			}
		}
	}
}

func TestRenderDefaultStyleWithoutExplicitStyle(t *testing.T) {
	b := NewBox().Padding(1, 1)

	out, err := b.Render("Default", "Content")
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines in rendered box, got %d", len(lines))
	}
	top := lines[0]
	bottom := lines[len(lines)-2]

	if !strings.HasPrefix(top, "┌") || !strings.HasSuffix(top, "┐") {
		t.Errorf("expected Single style corners by default; top=%q", top)
	}
	if !strings.HasPrefix(bottom, "└") || !strings.HasSuffix(bottom, "┘") {
		t.Errorf("expected Single style corners by default; bottom=%q", bottom)
	}
}

func TestManualBorderOverridesAfterStyle(t *testing.T) {
	b := NewBox().
		Style(Double).
		TopLeft("*").
		TopRight("*").
		BottomLeft("*").
		BottomRight("*")

	out, err := b.Render("Title", "Content")
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}
	lines := strings.Split(out, "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines in rendered box, got %d", len(lines))
	}
	top := lines[0]
	bottom := lines[len(lines)-2]

	if !strings.HasPrefix(top, "*") || !strings.HasSuffix(top, "*") {
		t.Errorf("expected custom top corners '*', got %q", top)
	}
	if !strings.HasPrefix(bottom, "*") || !strings.HasSuffix(bottom, "*") {
		t.Errorf("expected custom bottom corners '*', got %q", bottom)
	}

	if !strings.Contains(top, "═") || !strings.Contains(bottom, "═") {
		t.Errorf("expected Double style horizontal borders '═', got top=%q bottom=%q", top, bottom)
	}
}

func TestBoxCopy(t *testing.T) {
	t.Run("independent copies", func(t *testing.T) {
		original := NewBox().
			Padding(1, 2).
			Margin(3, 4).
			Color(Red).
			TitleColor(Blue).
			ContentColor(Yellow).
			TitlePosition(Top).
			Style(Double).
			WrapContent(true).
			WrapLimit(30)
		original.TopLeft("[").TopRight("]").BottomLeft("{").BottomRight("}").Horizontal("-").Vertical("|")

		clone := original.Copy()
		if clone == nil {
			t.Fatalf("expected non-nil copy")
		}
		if clone == original {
			t.Fatalf("Copy should return a distinct pointer")
		}

		if clone.mx != 3 || clone.my != 4 {
			t.Fatalf("expected cloned margin (3,4), got (%d,%d)", clone.mx, clone.my)
		}

		clone.Color(Green).Padding(5, 6).Margin(10, 11).TitlePosition(Bottom).TopLeft("*")
		if original.color != Red {
			t.Fatalf("expected original color to remain Red, got %q", original.color)
		}
		if original.px != 1 || original.py != 2 {
			t.Fatalf("expected original padding (1,2), got (%d,%d)", original.px, original.py)
		}
		if original.mx != 3 || original.my != 4 {
			t.Fatalf("expected original margin (3,4) after mutating clone, got (%d,%d)", original.mx, original.my)
		}
		if original.titlePos != Top {
			t.Fatalf("expected original title position to stay Top, got %v", original.titlePos)
		}
		if original.topLeft != "[" {
			t.Fatalf("expected original topLeft to stay '[', got %q", original.topLeft)
		}

		original.Color(Magenta)
		if clone.color != Green {
			t.Fatalf("expected clone color to remain Green after mutating original, got %q", clone.color)
		}
	})

	t.Run("nil receiver", func(t *testing.T) {
		var b *Box
		if copy := b.Copy(); copy != nil {
			t.Fatalf("expected nil copy from nil receiver, got %#v", copy)
		}
	})
}

func TestHPaddingAndVPadding(t *testing.T) {
	b := NewBox().Padding(1, 2)

	if b.px != 1 || b.py != 2 {
		t.Fatalf("expected initial padding (1,2), got (%d,%d)", b.px, b.py)
	}

	b.HPadding(5)
	if b.px != 5 {
		t.Errorf("expected HPadding to set horizontal padding to 5, got %d", b.px)
	}
	if b.py != 2 {
		t.Errorf("expected HPadding to leave vertical padding unchanged at 2, got %d", b.py)
	}

	b.VPadding(7)
	if b.py != 7 {
		t.Errorf("expected VPadding to set vertical padding to 7, got %d", b.py)
	}
	if b.px != 5 {
		t.Errorf("expected VPadding to leave horizontal padding unchanged at 5, got %d", b.px)
	}
}

func TestRenderTitlePositions(t *testing.T) {
	title := "My Title"
	content := "Some content"

	cases := []struct {
		name string
		pos  TitlePosition
	}{
		{"inside", Inside},
		{"top", Top},
		{"bottom", Bottom},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBox().Padding(2, 1).Style(Single).TitlePosition(tc.pos)
			out, err := b.Render(title, content)
			if err != nil {
				t.Fatalf("Render returned error for position %v: %v", tc.pos, err)
			}

			lines := strings.Split(out, "\n")
			if len(lines) < 3 {
				t.Fatalf("expected at least 3 lines, got %d", len(lines))
			}
			top := lines[0]
			bottom := lines[len(lines)-2]
			interior := lines[1 : len(lines)-2]

			hasTitleInside := false
			for _, l := range interior {
				if strings.Contains(l, title) {
					hasTitleInside = true
					break
				}
			}

			switch tc.pos {
			case Inside:
				if !hasTitleInside {
					t.Errorf("expected title to appear inside box for Inside position; output: %q", out)
				}
			case Top:
				if !strings.Contains(top, title) {
					t.Errorf("expected title to appear in top bar for Top position; top: %q", top)
				}
			case Bottom:
				if !strings.Contains(bottom, title) {
					t.Errorf("expected title to appear in bottom bar for Bottom position; bottom: %q", bottom)
				}
			}
		})
	}
}

func TestRenderTitleAlignInside(t *testing.T) {
	title := "Hi"
	content := "1234567890"
	px := 2

	contentWidth := runewidth.StringWidth(content)
	titleWidth := runewidth.StringWidth(title)
	diff := contentWidth - titleWidth

	cases := []struct {
		name        string
		align       AlignType
		expectedPad int
	}{
		{"left", Left, px},
		{"center", Center, px + diff/2},
		{"right", Right, px + diff},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBox().Padding(px, 0).Style(Single).TitlePosition(Inside).TitleAlign(tc.align)
			out, err := b.Render(title, content)
			if err != nil {
				t.Fatalf("Render returned error: %v", err)
			}

			lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
			if len(lines) < 3 {
				t.Fatalf("expected at least 3 lines, got %d", len(lines))
			}
			interior := lines[1 : len(lines)-1]
			titleLine := findLineContainingTitle(interior, title)
			if titleLine == "" {
				t.Fatalf("expected to find title line for alignment %s", tc.align)
			}

			startCol := titleStartColumn(titleLine, title)
			if startCol < 0 {
				t.Fatalf("could not find title in line: %q", titleLine)
			}

			verticalWidth := runewidth.StringWidth(b.vertical)
			expectedStart := verticalWidth + tc.expectedPad
			if startCol != expectedStart {
				t.Errorf("alignment %s: expected title to start at column %d, got %d; line=%q", tc.align, expectedStart, startCol, titleLine)
			}
		})
	}
}

func TestRenderTitleAlignTopBottom(t *testing.T) {
	title := "Title"
	content := strings.Repeat("x", 20)

	cases := []struct {
		name  string
		pos   TitlePosition
		align AlignType
	}{
		{"top-left", Top, Left},
		{"top-center", Top, Center},
		{"top-right", Top, Right},
		{"bottom-left", Bottom, Left},
		{"bottom-center", Bottom, Center},
		{"bottom-right", Bottom, Right},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBox().Padding(1, 0).Style(Single).TitlePosition(tc.pos).TitleAlign(tc.align)
			out, err := b.Render(title, content)
			if err != nil {
				t.Fatalf("Render returned error: %v", err)
			}

			lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
			if len(lines) < 3 {
				t.Fatalf("expected at least 3 lines, got %d", len(lines))
			}

			bar := lines[0]
			leftCorner := b.topLeft
			rightCorner := b.topRight
			if tc.pos == Bottom {
				bar = lines[len(lines)-1]
				leftCorner = b.bottomLeft
				rightCorner = b.bottomRight
			}

			startCol := titleStartColumn(bar, title)
			if startCol < 0 {
				t.Fatalf("could not find title in bar: %q", bar)
			}

			stripped := ansi.Strip(bar)
			lineWidth := runewidth.StringWidth(stripped)
			leftW := runewidth.StringWidth(leftCorner)
			rightW := runewidth.StringWidth(rightCorner)
			inner := lineWidth - leftW - rightW
			titleWidth := runewidth.StringWidth(title)
			titleSegWidth := titleWidth + 2
			remaining := inner - titleSegWidth

			leftSegWidth := 0
			switch tc.align {
			case Center:
				leftSegWidth = remaining / 2
			case Right:
				leftSegWidth = remaining
			case Left:
				leftSegWidth = 0
			}

			expectedStart := leftW + leftSegWidth + 1
			if startCol != expectedStart {
				t.Errorf("alignment %s/%s: expected title to start at column %d, got %d; bar=%q", tc.pos, tc.align, expectedStart, startCol, stripped)
			}
		})
	}
}

func TestRenderTitleAlignDefaults(t *testing.T) {
	title := "Title"
	content := strings.Repeat("x", 12)
	px := 2

	// Inside defaults to Center.
	b := NewBox().Padding(px, 0).Style(Single).TitlePosition(Inside)
	out, err := b.Render(title, content)
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines, got %d", len(lines))
	}
	interior := lines[1 : len(lines)-1]
	titleLine := findLineContainingTitle(interior, title)
	if titleLine == "" {
		t.Fatalf("expected to find title line for default inside alignment")
	}
	startCol := titleStartColumn(titleLine, title)
	if startCol < 0 {
		t.Fatalf("could not find title in line: %q", titleLine)
	}
	diff := runewidth.StringWidth(content) - runewidth.StringWidth(title)
	expectedInsideStart := runewidth.StringWidth(b.vertical) + px + diff/2
	if startCol != expectedInsideStart {
		t.Errorf("default inside alignment: expected title to start at column %d, got %d; line=%q", expectedInsideStart, startCol, titleLine)
	}

	// Top defaults to Left.
	b = NewBox().Padding(px, 0).Style(Single).TitlePosition(Top)
	out, err = b.Render(title, content)
	if err != nil {
		t.Fatalf("Render returned error: %v", err)
	}
	lines = strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines, got %d", len(lines))
	}
	top := lines[0]
	startCol = titleStartColumn(top, title)
	if startCol < 0 {
		t.Fatalf("could not find title in top bar: %q", top)
	}
	leftW := runewidth.StringWidth(b.topLeft)
	expectedTopStart := leftW + 1
	if startCol != expectedTopStart {
		t.Errorf("default top alignment: expected title to start at column %d, got %d; bar=%q", expectedTopStart, startCol, ansi.Strip(top))
	}
}

func TestRenderInvalidTitleAlign(t *testing.T) {
	b := NewBox().Padding(1, 0).Style(Single).TitlePosition(Top).TitleAlign(AlignType("Weird"))
	if _, err := b.Render("Title", "Content"); err == nil {
		t.Fatalf("expected error for invalid title alignment, got nil")
	} else if !strings.Contains(err.Error(), "invalid Title Alignment") {
		t.Errorf("unexpected error message for invalid title alignment: %v", err)
	}
}

func TestRenderInvalidBoxStyle(t *testing.T) {
	b := NewBox().Padding(2, 1).Style(BoxStyle("InvalidStyle"))
	_, err := b.Render("Title", "Content")
	if err == nil {
		t.Fatalf("expected error for invalid Box style, got nil")
	}
	if !strings.Contains(err.Error(), "invalid Box style") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestRenderInvalidTitlePosition(t *testing.T) {
	b := NewBox().Padding(2, 1).Style(Single).TitlePosition(TitlePosition("Weird"))
	_, err := b.Render("Title", "Content")
	if err == nil {
		t.Fatalf("expected error for invalid TitlePosition, got nil")
	}
	if !strings.Contains(err.Error(), "invalid TitlePosition") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestRenderMultilineTitleNonInside(t *testing.T) {
	b := NewBox().Padding(2, 1).Style(Single).TitlePosition(Top)
	_, err := b.Render("Line1\nLine2", "Content")
	if err == nil {
		t.Fatalf("expected error for multiline title at non-Inside position, got nil")
	}
	if !strings.Contains(err.Error(), "multiline titles are only supported with the Inside title position") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestRenderNonPositiveWrapLimit(t *testing.T) {
	// WrapLimit(0) used to silently fall through to terminal-width detection
	// and, on a non-TTY, fail by advising the caller to use WrapLimit.
	// Zero and negative limits now error explicitly.
	for _, limit := range []int{-1, 0} {
		b := NewBox().Padding(1, 1).Style(Single).WrapLimit(limit)
		if _, err := b.Render("Title", "Content"); err == nil {
			t.Fatalf("expected error for WrapLimit(%d), got nil", limit)
		} else if !strings.Contains(err.Error(), "wrap limit must be positive") {
			t.Errorf("unexpected error message for WrapLimit(%d): %v", limit, err)
		}
	}

	// The automatic path (WrapContent without WrapLimit) must be unaffected.
	oldIsTTY, oldGetTermSize := isTTY, getTermSize
	defer func() {
		isTTY = oldIsTTY
		getTermSize = oldGetTermSize
	}()
	isTTY = func(fd uintptr) bool { return true }
	getTermSize = func(fd uintptr) (int, int, error) { return 80, 24, nil }
	if _, err := NewBox().Style(Single).WrapContent(true).Render("Title", "Content"); err != nil {
		t.Fatalf("WrapContent(true) auto mode should not error: %v", err)
	}
}

func TestRenderNonTTYLWrapContent(t *testing.T) {
	oldIsTTY := isTTY
	defer func() { isTTY = oldIsTTY }() // restore after test

	isTTY = func(fd uintptr) bool { return false } // mock as non-TTY
	b := NewBox().Padding(1, 1).Style(Single).WrapContent(true)

	if _, err := b.Render("Title", "Content"); err == nil {
		t.Fatalf("expected error for non TTY output with wrapping enabled, got nil")
	} else if !strings.Contains(err.Error(), "cannot determine terminal width") {
		t.Errorf("unexpected error message for non TTY wrap content: %v", err)
	}
}

func TestRenderWrapContentWithMarginFitsTerminal(t *testing.T) {
	const termWidth = 80

	oldIsTTY := isTTY
	oldGetTermSize := getTermSize
	defer func() {
		isTTY = oldIsTTY
		getTermSize = oldGetTermSize
	}()
	isTTY = func(fd uintptr) bool { return true }
	getTermSize = func(fd uintptr) (int, int, error) { return termWidth, 24, nil }

	content := strings.Repeat("Box CLI Maker ", 40)
	margins := []int{0, 10, 20, 25, 30, termWidth/3 - 1, termWidth / 3, termWidth/3 + 1}

	for _, mx := range margins {
		t.Run(fmt.Sprintf("HMargin%d", mx), func(t *testing.T) {
			b := NewBox().Style(Single).Padding(2, 0).HMargin(mx).WrapContent(true)
			out, err := b.Render("Demo", content)
			if err != nil {
				t.Fatalf("HMargin(%d): unexpected error: %v", mx, err)
			}
			for line := range strings.SplitSeq(strings.TrimRight(out, "\n"), "\n") {
				if w := runewidth.StringWidth(ansi.Strip(line)); w > termWidth {
					t.Errorf("HMargin(%d): line exceeds terminal width %d (got %d): %q", mx, termWidth, w, line)
				}
			}
		})
	}
}

// TestRenderWrapLimitMarginOverflowRegression verifies the fix for the bug
// where WrapContent used the full terminal width without subtracting HMargin.
// The old workaround (WrapLimit set to 2/3 of terminal) overflows when HMargin
// is large; WrapContent(true) must not.
func TestRenderWrapLimitMarginOverflowRegression(t *testing.T) {
	const termWidth = 80
	largeMx := termWidth/3 + 1        // 27: large enough to cause overflow before the fix
	oldWrapWidth := 2 * termWidth / 3 // 53: what wrapContent used to compute (ignoring margin)

	oldIsTTY := isTTY
	oldGetTermSize := getTermSize
	defer func() {
		isTTY = oldIsTTY
		getTermSize = oldGetTermSize
	}()
	isTTY = func(fd uintptr) bool { return true }
	getTermSize = func(fd uintptr) (int, int, error) { return termWidth, 24, nil }

	content := strings.Repeat("Box CLI Maker ", 40)

	// Old behaviour: WrapLimit(2/3 of terminal) ignores margin → overflows.
	before := NewBox().Style(Single).Padding(2, 0).HMargin(largeMx).WrapLimit(oldWrapWidth)
	outBefore, err := before.Render("Overflow Demo", content)
	if err != nil {
		t.Fatalf("before: unexpected error: %v", err)
	}
	overflows := false
	for line := range strings.SplitSeq(strings.TrimRight(outBefore, "\n"), "\n") {
		if runewidth.StringWidth(ansi.Strip(line)) > termWidth {
			overflows = true
			break
		}
	}
	if !overflows {
		t.Errorf("expected WrapLimit(%d)+HMargin(%d) to overflow a %d-col terminal, but no line exceeded it",
			oldWrapWidth, largeMx, termWidth)
	}

	// Fixed behaviour: WrapContent(true) subtracts margin before computing wrap width → fits.
	after := NewBox().Style(Single).Padding(2, 0).HMargin(largeMx).WrapContent(true)
	outAfter, err := after.Render("Overflow Demo", content)
	if err != nil {
		t.Fatalf("after: unexpected error: %v", err)
	}
	for line := range strings.SplitSeq(strings.TrimRight(outAfter, "\n"), "\n") {
		if w := runewidth.StringWidth(ansi.Strip(line)); w > termWidth {
			t.Errorf("after fix: line exceeds terminal width %d (got %d): %q", termWidth, w, line)
		}
	}
}

func TestRenderNegativePadding(t *testing.T) {
	// Horizontal padding < 0.
	b := NewBox().Padding(-1, 1).Style(Single)
	if _, err := b.Render("Title", "Content"); err == nil {
		t.Fatalf("expected error for negative horizontal padding, got nil")
	} else if !strings.Contains(err.Error(), "horizontal padding cannot be negative") {
		t.Errorf("unexpected error for negative horizontal padding: %v", err)
	}

	// Vertical padding < 0.
	b = NewBox().Padding(1, -1).Style(Single)
	if _, err := b.Render("Title", "Content"); err == nil {
		t.Fatalf("expected error for negative vertical padding, got nil")
	} else if !strings.Contains(err.Error(), "vertical padding cannot be negative") {
		t.Errorf("unexpected error for negative vertical padding: %v", err)
	}
}

func TestRenderInvalidContentAlign(t *testing.T) {
	b := NewBox().Padding(1, 1).Style(Single).ContentAlign(AlignType("Weird"))

	if _, err := b.Render("Title", "Content"); err == nil {
		t.Fatalf("expected error for invalid content alignment, got nil")
	} else if !strings.Contains(err.Error(), "invalid Content Alignment") {
		t.Errorf("unexpected error message for invalid content alignment: %v", err)
	}
}

func TestRenderWithWrapLimit(t *testing.T) {
	longContent := strings.Repeat("word ", 20)
	b := NewBox().Padding(2, 0).Style(Single).Color(Green).WrapContent(true).WrapLimit(10)

	out, err := b.Render("Wrapped", longContent)
	if err != nil {
		t.Fatalf("Render with wrapping returned error: %v", err)
	}
	if !strings.Contains(out, "Wrapped") {
		t.Errorf("expected title to appear in wrapped box output")
	}
	if !strings.Contains(out, "word") {
		t.Errorf("expected content to appear in wrapped box output")
	}
}

// TestRenderWrapLimitWithTabs guards against wrapping running before tab
// expansion: ansi.Wrap counts a tab as one cell, so a tab expanding to up to
// 8 spaces afterwards pushed the content area past the configured limit.
func TestRenderWrapLimitWithTabs(t *testing.T) {
	const limit = 10
	b := NewBox().Style(Single).Padding(0, 0)
	b.WrapLimit(limit)

	out, err := b.Render("", "aa\tbb cc dd ee ff")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	for line := range strings.SplitSeq(strings.TrimRight(out, "\n"), "\n") {
		// Inner content width = visible line width minus the two border cells.
		if inner := runewidth.StringWidth(ansi.Strip(line)) - 2; inner > limit {
			t.Errorf("content width %d exceeds WrapLimit(%d): %q", inner, limit, ansi.Strip(line))
		}
	}
}

func TestRenderWithVariousColorFormats(t *testing.T) {
	title := "Color Formats"
	content := "content"

	tests := []struct {
		name      string
		configure func(*Box)
	}{
		{
			name: "short hex #RGB",
			configure: func(b *Box) {
				b.TitleColor("#0F0")
			},
		},
		{
			name: "full hex #RRGGBB",
			configure: func(b *Box) {
				b.ContentColor("#00FF00")
			},
		},
		{
			name: "rgb:RRRR/GGGG/BBBB",
			configure: func(b *Box) {
				b.Color("rgb:0000/ffff/0000")
			},
		},
		{
			name: "rgba:RRRR/GGGG/BBBB/AAAA",
			configure: func(b *Box) {
				b.Color("rgba:ffff/0000/0000/ffff")
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBox().Padding(1, 1).Style(Single)
			// Apply specific color configuration.
			ct := *b
			box := &ct
			// Configure colors on the copy so one test's colors don't bleed into another.
			tc.configure(box)

			out, err := box.Render(title, content)
			if err != nil {
				t.Fatalf("Render returned error for %s: %v", tc.name, err)
			}
			if out == "" {
				t.Fatalf("expected non-empty output for %s", tc.name)
			}
			if !strings.Contains(out, title) || !strings.Contains(out, content) {
				t.Fatalf("rendered output for %s missing title or content: %q", tc.name, out)
			}
		})
	}
}

func TestRenderInvalidColors(t *testing.T) {
	title := "Title"
	content := "Content"

	tests := []struct {
		name string
		mut  func(*Box)
	}{
		{
			name: "invalid title color",
			mut:  func(b *Box) { b.TitleColor("NotAColor") },
		},
		{
			name: "invalid content color",
			mut:  func(b *Box) { b.ContentColor("NotAColor") },
		},
		{
			name: "invalid border color",
			mut:  func(b *Box) { b.Color("NotAColor") },
		},
	}

	for _, tc := range tests {
		b := NewBox().Padding(1, 1).Style(Single)
		// Apply the specific invalid color configuration.
		tc.mut(b)

		if _, err := b.Render(title, content); err == nil {
			t.Fatalf("%s: expected error for invalid color, got nil", tc.name)
		} else if !strings.Contains(err.Error(), "unable to parse color") {
			t.Errorf("%s: unexpected error message: %v", tc.name, err)
		}
	}
}

func TestMustRenderSuccessAndPanic(t *testing.T) {
	// Success case: MustRender should not panic when Render succeeds.
	t.Run("success", func(t *testing.T) {
		b := NewBox().Padding(1, 1).Style(Single)
		_ = b.MustRender("Title", "Content")
	})

	// Panic case: invalid style causes Render to error, hence MustRender panics.
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("expected MustRender to panic for invalid style, but it did not")
			}
		}()
		b := NewBox().Padding(1, 1).Style(BoxStyle("InvalidStyle"))
		_ = b.MustRender("Title", "Content")
	})
}

func TestRenderEmojiBordersHaveConsistentWidth(t *testing.T) {
	b := NewBox().Padding(2, 1)
	b.TopLeft("📦").TopRight("📦").BottomLeft("📦").BottomRight("📦").Horizontal("📦").Vertical("📦")

	out, err := b.Render("Emoji Box", "With emoji borders")
	if err != nil {
		t.Fatalf("Render with emoji borders returned error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines in rendered box, got %d", len(lines))
	}

	top := ansi.Strip(lines[0])
	interior := ansi.Strip(lines[1])
	bottom := ansi.Strip(lines[len(lines)-1])

	topW := runewidth.StringWidth(top)
	interiorW := runewidth.StringWidth(interior)
	bottomW := runewidth.StringWidth(bottom)

	if topW != interiorW || interiorW != bottomW {
		t.Fatalf("expected equal visual widths for emoji box borders, got top=%d interior=%d bottom=%d", topW, interiorW, bottomW)
	}
}

// TestRenderEmptyTitleWithTitleColor guards against applyColor turning an
// empty title into a non-empty ANSI-only string ("\x1b[38;…m\x1b[m"), which
// made every title != "" check downstream render a phantom title: a gap in
// the Top/Bottom bar, or a spurious title line plus separator Inside.
func TestRenderEmptyTitleWithTitleColor(t *testing.T) {
	for _, pos := range []TitlePosition{Inside, Top, Bottom} {
		t.Run(string(pos), func(t *testing.T) {
			plain, err := NewBox().Style(Single).TitlePosition(pos).Color(Cyan).Render("", "content here")
			if err != nil {
				t.Fatalf("Render error: %v", err)
			}
			colored, err := NewBox().Style(Single).TitlePosition(pos).Color(Cyan).TitleColor(Red).Render("", "content here")
			if err != nil {
				t.Fatalf("Render error: %v", err)
			}
			if plain != colored {
				t.Errorf("TitleColor with empty title must be a no-op:\nwithout: %q\nwith:    %q", plain, colored)
			}
		})
	}

	// Same for empty content with ContentColor.
	plain, err := NewBox().Style(Single).Render("Title", "")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	colored, err := NewBox().Style(Single).ContentColor(Green).Render("Title", "")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	if plain != colored {
		t.Errorf("ContentColor with empty content must be a no-op:\nwithout: %q\nwith:    %q", plain, colored)
	}

	// Invalid colors must still error even when the title is empty.
	if _, err := NewBox().Style(Single).TitleColor("NotAColor").Render("", "content"); err == nil {
		t.Errorf("expected error for invalid TitleColor with empty title")
	}
}

// TestRenderHyperlinkContentWithTabs guards against OSC-8 hyperlinks breaking
// tab handling end to end: a raw \t used to leak through expandTabs into the
// rendered box, where the terminal's own tab expansion breaks the border.
func TestRenderHyperlinkContentWithTabs(t *testing.T) {
	link := "\x1b]8;;https://example.com\x1b\\docs\x1b]8;;\x1b\\"
	out, err := NewBox().Style(Single).Render("", link+"\tafter-tab\nplain line here padding")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	if strings.Contains(out, "\t") {
		t.Errorf("raw tab in rendered output: %q", out)
	}

	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	w0 := runewidth.StringWidth(ansi.Strip(lines[0]))
	for i, line := range lines[1:] {
		if w := runewidth.StringWidth(ansi.Strip(line)); w != w0 {
			t.Errorf("line %d width %d != line 0 width %d:\n%s", i+1, w, w0, out)
		}
	}
	// "docs" is 4 visible columns, so the tab should expand to 4 spaces.
	if !strings.Contains(ansi.Strip(out), "docs    after-tab") {
		t.Errorf("tab did not expand from the hyperlink's visible column:\n%q", ansi.Strip(out))
	}
}

// TestRenderZeroWidthGlyphs guards against border glyphs with zero visible
// width (empty strings, zero-width characters, ANSI-only strings) producing
// ragged boxes: charWidth's width-1 fallback counted a column the glyph never
// rendered. Render normalizes such glyphs to a single space.
func TestRenderZeroWidthGlyphs(t *testing.T) {
	cases := []struct {
		name string
		b    *Box
	}{
		{"empty-vertical", NewBox().Vertical("")},
		{"empty-corners", NewBox().TopLeft("").TopRight("").BottomLeft("").BottomRight("")},
		{"empty-horizontal", NewBox().Horizontal("")},
		{"empty-top-corners-only", NewBox().TopLeft("").TopRight("")},
		{"zero-width-space-vertical", NewBox().Vertical("\u200b")},
		{"ansi-only-horizontal", NewBox().Horizontal("\x1b[31m\x1b[0m")},
		{"zero-value-box", new(Box)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := tc.b.Render("Title", "some content")
			if err != nil {
				t.Fatalf("Render error: %v", err)
			}
			lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
			w0 := runewidth.StringWidth(ansi.Strip(lines[0]))
			for i, line := range lines[1:] {
				if w := runewidth.StringWidth(ansi.Strip(line)); w != w0 {
					t.Errorf("line %d width %d != line 0 width %d:\n%s", i+1, w, w0, out)
				}
			}
		})
	}

	// Semantic pin: an empty glyph renders identically to an explicit space.
	emptyOut, err := NewBox().Vertical("").Render("T", "content")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	spaceOut, err := NewBox().Vertical(" ").Render("T", "content")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	if emptyOut != spaceOut {
		t.Errorf("Vertical(\"\") and Vertical(\" \") render differently:\n%q\n%q", emptyOut, spaceOut)
	}

	// Normalization must not mutate the receiver.
	b := NewBox().Vertical("")
	if _, err := b.Render("T", "content"); err != nil {
		t.Fatalf("Render error: %v", err)
	}
	if b.vertical != "" {
		t.Errorf("Render mutated the box's vertical glyph: %q", b.vertical)
	}
}

// TestRenderCRLFContent guards against raw \r surviving into the output:
// "line one\r" between the borders makes the terminal carriage-return
// mid-line, so the right border overwrites the left edge.
func TestRenderCRLFContent(t *testing.T) {
	crlfOut, err := NewBox().Style(Single).Render("Ti\r\ntle", "line one\r\nline two")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	if strings.Contains(crlfOut, "\r") {
		t.Errorf("output contains raw \\r: %q", crlfOut)
	}

	// CRLF input must render identically to LF input.
	lfOut, err := NewBox().Style(Single).Render("Ti\ntle", "line one\nline two")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	if crlfOut != lfOut {
		t.Errorf("CRLF and LF input render differently:\nCRLF: %q\nLF:   %q", crlfOut, lfOut)
	}
}

// TestRenderWideCornersWithTitleBar guards against computeLayout sizing the box
// from titleWidth+2 alone: the titled bar's span between its corners is
// innerWidth + 2*verticalWidth - leftW - rightW, so corner glyphs wider than
// the vertical glyph made the titled bar overflow the box whenever the title
// dictated the box width.
func TestRenderWideCornersWithTitleBar(t *testing.T) {
	title := "A moderately long title here"
	cases := []struct {
		name   string
		corner string
		pos    TitlePosition
	}{
		{"emoji-corners-top", "🌸", Top},
		{"emoji-corners-bottom", "🌸", Bottom},
		{"ascii-wide-corners-top", "++", Top},
		{"ascii-wide-corners-bottom", "++", Bottom},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := NewBox().Style(Single).TitlePosition(tc.pos).
				TopLeft(tc.corner).TopRight(tc.corner).
				BottomLeft(tc.corner).BottomRight(tc.corner)
			// Short content so the title dictates the box width.
			out, err := b.Render(title, "ab")
			if err != nil {
				t.Fatalf("Render error: %v", err)
			}

			lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
			w0 := runewidth.StringWidth(ansi.Strip(lines[0]))
			for i, line := range lines[1:] {
				if w := runewidth.StringWidth(ansi.Strip(line)); w != w0 {
					t.Errorf("line %d width %d != line 0 width %d:\n%s", i+1, w, w0, out)
				}
			}
			if !strings.Contains(out, title) {
				t.Errorf("title missing from output:\n%s", out)
			}
		})
	}
}

func TestRenderBoxCustomGlyphsWithoutNewBoxMethod(t *testing.T) {
	b := new(Box)
	b = b.TopLeft("+").TopRight("+").BottomLeft("+").BottomRight("+").Horizontal("-").Vertical("|")

	out, err := b.Render("Custom Glyphs", "Using custom border glyphs")
	if err != nil {
		t.Fatalf("Render with custom glyphs returned error: %v", err)
	}

	lines := strings.Split(out, "\n")
	if len(lines) < 3 {
		t.Fatalf("expected at least 3 lines in rendered box, got %d", len(lines))
	}

	top := lines[0]
	bottom := lines[len(lines)-2]
	interior := lines[1 : len(lines)-2]

	if !strings.HasPrefix(top, "+") || !strings.HasSuffix(top, "+") {
		t.Errorf("top border does not use custom corners: %q", top)
	}
	if !strings.HasPrefix(bottom, "+") || !strings.HasSuffix(bottom, "+") {
		t.Errorf("bottom border does not use custom corners: %q", bottom)
	}

	for _, line := range interior {
		if !strings.HasPrefix(line, "|") || !strings.HasSuffix(line, "|") {
			t.Errorf("interior line does not use custom vertical borders: %q", line)
		}
	}
}

func TestRenderMargin(t *testing.T) {
	const title, content = "Title", "Content"

	// Horizontal margin: every line (including blank vertical-margin lines) must be prefixed.
	out, err := NewBox().Style(Single).HMargin(4).Render(title, content)
	if err != nil {
		t.Fatalf("HMargin: unexpected error: %v", err)
	}
	for line := range strings.SplitSeq(strings.TrimRight(out, "\n"), "\n") {
		if !strings.HasPrefix(line, "    ") {
			t.Errorf("HMargin: line missing 4-space prefix: %q", line)
		}
	}

	// Vertical margin: output must start and end with my blank lines.
	out, err = NewBox().Style(Single).VMargin(2).Render(title, content)
	if err != nil {
		t.Fatalf("VMargin: unexpected error: %v", err)
	}
	if !strings.HasPrefix(out, "\n\n") {
		t.Errorf("VMargin: output does not start with 2 blank lines: %q", out[:min(len(out), 10)])
	}
	// 2 blank lines = 2 extra newlines after the final box line's own newline → 3 total.
	nTrailing := len(out) - len(strings.TrimRight(out, "\n"))
	if nTrailing != 3 {
		t.Errorf("VMargin: expected 3 trailing newlines (2 blank lines), got %d", nTrailing)
	}

	// Margin(mx, my): blank vertical-margin lines are bare newlines; only box lines carry the prefix.
	out, err = NewBox().Style(Single).Margin(3, 1).Render(title, content)
	if err != nil {
		t.Fatalf("Margin: unexpected error: %v", err)
	}
	if !strings.HasPrefix(out, "\n") {
		t.Errorf("Margin: output does not start with blank line, got %q", out[:min(len(out), 10)])
	}
	for line := range strings.SplitSeq(strings.TrimRight(out, "\n"), "\n") {
		if line != "" && !strings.HasPrefix(line, "   ") {
			t.Errorf("Margin: non-blank line missing 3-space prefix: %q", line)
		}
	}

	// Zero margin produces the same output as no margin.
	plain, err := NewBox().Style(Single).Render(title, content)
	if err != nil {
		t.Fatalf("plain: unexpected error: %v", err)
	}
	zero, err := NewBox().Style(Single).Margin(0, 0).Render(title, content)
	if err != nil {
		t.Fatalf("Margin(0,0): unexpected error: %v", err)
	}
	if plain != zero {
		t.Errorf("Margin(0,0) should produce identical output to no margin")
	}

	// Negative margin must error.
	if _, err := NewBox().Style(Single).Margin(-1, 0).Render(title, content); err == nil {
		t.Errorf("expected error for negative horizontal margin, got nil")
	}
	if _, err := NewBox().Style(Single).Margin(0, -1).Render(title, content); err == nil {
		t.Errorf("expected error for negative vertical margin, got nil")
	}
}

// expandTabsStop4 simulates the old title tab behaviour (stop-4) so the before
// case can be reproduced without touching the library.
func expandTabsStop4(s string) string {
	var b strings.Builder
	col := 0
	for _, c := range s {
		if c == '\t' {
			spaces := 4 - (col & 3)
			b.WriteString(strings.Repeat(" ", spaces))
			col += spaces
		} else {
			w := runewidth.StringWidth(string(c))
			b.WriteRune(c)
			col += w
		}
	}
	return b.String()
}

// TestRenderTabTitleAndContentAlign verifies that tabs in the title and tabs in
// the content expand using the same tab stop (8), so columns line up visually.
// Previously the title used stop-4 while content used stop-8, causing misalignment.
func TestRenderTabTitleAndContentAlign(t *testing.T) {
	title := "Name\tAge\tCity"
	content := "Alice\t30\tLondon\nBob\t25\tParis\nCharlie\t35\tTokyo"

	tokenCol := func(out, titleToken, contentToken string) (int, int) {
		lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
		topBar := ansi.Strip(lines[0])
		// skip top bar + 2 vertical padding lines to reach first data line
		firstDataLine := ansi.Strip(lines[3])
		return strings.Index(topBar, titleToken), strings.Index(firstDataLine, contentToken)
	}

	// Before: pre-expand title with stop-4, leave content for the library (stop-8).
	// City (title) and London (content) should NOT align.
	oldOut, err := NewBox().Style(Single).Padding(1, 2).TitlePosition(Top).
		Render(expandTabsStop4(title), content)
	if err != nil {
		t.Fatalf("before Render error: %v", err)
	}
	cityColOld, londonColOld := tokenCol(oldOut, "City", "London")
	if cityColOld == -1 || londonColOld == -1 {
		t.Fatalf("before: tokens not found in output:\n%s", oldOut)
	}
	if cityColOld == londonColOld {
		t.Errorf("before: expected misalignment with stop-4 title vs stop-8 content, but columns matched at %d", cityColOld)
	}

	// After: pass raw tab title — library expands both title and content with stop-8.
	// City (title) and London (content) must align.
	newOut, err := NewBox().Style(Single).Padding(1, 2).TitlePosition(Top).
		Render(title, content)
	if err != nil {
		t.Fatalf("after Render error: %v", err)
	}
	cityColNew, londonColNew := tokenCol(newOut, "City", "London")
	if cityColNew == -1 || londonColNew == -1 {
		t.Fatalf("after: tokens not found in output:\n%s", newOut)
	}
	if cityColNew != londonColNew {
		t.Errorf("after: title and content tab columns differ: 'City' at col %d, 'London' at col %d",
			cityColNew, londonColNew)
	}
}

// TestRenderContentColorWithTabs guards against expandTabs being called after
// applyColor for content: same root cause as the title bug, but for content lines.
func TestRenderContentColorWithTabs(t *testing.T) {
	content := "Name\tAge\nAlice\t30"

	out, err := NewBox().Style(Single).Padding(1, 0).ContentColor(Red).Render("", content)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	// lines[1] = header row "Name    Age", lines[2] = data row "Alice   30"
	headerLine := ansi.Strip(lines[1])
	dataLine := ansi.Strip(lines[2])

	ageCol := strings.Index(headerLine, "Age")
	col30 := strings.Index(dataLine, "30")
	if ageCol == -1 || col30 == -1 {
		t.Fatalf("tokens not found:\nheader: %q\ndata:   %q", headerLine, dataLine)
	}
	if ageCol != col30 {
		t.Errorf("tab columns differ: 'Age' at col %d, '30' at col %d\nheader: %q\ndata:   %q",
			ageCol, col30, headerLine, dataLine)
	}
}

// TestRenderInsideTitleNoExtraWidth guards against computeLayout applying the
// Top/Bottom min-width constraint (titleWidth+2) to Inside titles when
// b.titlePos == "" (default). With px=0 this made the box 2 cols wider than needed.
func TestRenderInsideTitleNoExtraWidth(t *testing.T) {
	// Title and content are the same width; box inner width must equal that width, not width+2.
	out, err := NewBox().Style(Single).Padding(0, 0).Render("Hello", "World")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	// "┌─────┐" = 7 runes; "┌───────┐" = 9 runes would indicate the bug.
	wantWidth := len([]rune("┌─────┐"))
	if got := len([]rune(lines[0])); got != wantWidth {
		t.Errorf("top bar width: want %d, got %d\n%s", wantWidth, got, out)
	}
}

// TestRenderDefaultTitlePosMatchesExplicitInside guards against the default titlePos ("")
// being treated differently from explicit Inside in formatLine, causing title lines to
// use content alignment (Left) instead of title alignment (Center).
func TestRenderDefaultTitlePosMatchesExplicitInside(t *testing.T) {
	title := "Title"
	content := "Content"

	defaultOut, err := NewBox().Style(Single).Render(title, content)
	if err != nil {
		t.Fatalf("default Render error: %v", err)
	}
	explicitOut, err := NewBox().Style(Single).TitlePosition(Inside).Render(title, content)
	if err != nil {
		t.Fatalf("explicit Inside Render error: %v", err)
	}
	if defaultOut != explicitOut {
		t.Errorf("default titlePos and explicit Inside produce different output:\ndefault:\n%s\nexplicit:\n%s",
			defaultOut, explicitOut)
	}
}

// TestRenderTitleColorWithTabs guards against expandTabs being called after applyColor:
// when TitleColor is set, the title gains ANSI escape bytes whose ASCII characters
// (e.g. '[','3','8',';'…) were being counted as visible columns, shifting every
// subsequent tab stop. The fix is to expand tabs before applying color in Render.
func TestRenderTitleColorWithTabs(t *testing.T) {
	title := "Name\tAge"
	content := "Alice\t30"

	out, err := NewBox().Style(Single).Padding(1, 0).
		TitlePosition(Top).TitleColor(Red).
		Render(title, content)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}

	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	topBar := ansi.Strip(lines[0])
	// No vertical padding: content is the first interior line.
	contentLine := ansi.Strip(lines[1])

	titleAgeCol := strings.Index(topBar, "Age")
	contentAgeCol := strings.Index(contentLine, "30")
	if titleAgeCol == -1 || contentAgeCol == -1 {
		t.Fatalf("tokens not found in output:\n%s", out)
	}
	if titleAgeCol != contentAgeCol {
		t.Errorf("tab columns differ: 'Age' in title at col %d, '30' in content at col %d\n%s",
			titleAgeCol, contentAgeCol, out)
	}
}

// TestRenderIsIdempotent guards against B1: prepareContentLines previously mutated
// b.titlePos on the receiver, causing repeated Render calls on the same Box to diverge.
func TestRenderIsIdempotent(t *testing.T) {
	b := NewBox().Style(Single).Padding(1, 2)
	out1, err := b.Render("Title", "content")
	if err != nil {
		t.Fatalf("first Render error: %v", err)
	}
	out2, err := b.Render("Title", "content")
	if err != nil {
		t.Fatalf("second Render error: %v", err)
	}
	if out1 != out2 {
		t.Errorf("Render mutated box state between calls:\nfirst:\n%s\nsecond:\n%s", out1, out2)
	}
}

// TestRenderColoredBorderConsistentAcrossLines guards against B2: applyColor for the
// vertical border was previously called once per content line instead of once per Render.
func TestRenderColoredBorderConsistentAcrossLines(t *testing.T) {
	content := "a\nb\nc\nd\ne"
	out, err := NewBox().Style(Single).Padding(0, 1).Color("Cyan").Render("", content)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	contentLines := lines[1 : len(lines)-1]
	if len(contentLines) < 2 {
		t.Fatal("need at least 2 content lines")
	}
	borderEnd := strings.Index(contentLines[0], "│") + len("│")
	wantBorder := contentLines[0][:borderEnd]
	for i, line := range contentLines[1:] {
		if len(line) < borderEnd || line[:borderEnd] != wantBorder {
			t.Errorf("line %d border differs:\ngot  %q\nwant %q", i+1, line[:min(len(line), borderEnd)], wantBorder)
		}
	}
}

// TestTitleColorNotDoubleApplied guards against B3: applyColorBar previously called
// applyColor(title, titleColor) again even though Render had already colored the title.
func TestTitleColorNotDoubleApplied(t *testing.T) {
	out, err := NewBox().
		Style(Single).
		Padding(1, 3).
		TitlePosition(Top).
		Color("Cyan").
		TitleColor("Red").
		Render("Hello", "world")
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	topBar := strings.SplitN(strings.TrimRight(out, "\n"), "\n", 2)[0]
	// Consecutive resets are the visible symptom of double-coloring.
	if strings.Contains(topBar, "\x1b[m\x1b[m") {
		t.Errorf("top bar contains consecutive ANSI resets — title color was applied twice: %q", topBar)
	}
}

// TestRenderANSIColoredContentWidth guards against B5: formatLine previously called
// runewidth.StringWidth(line.line) twice instead of reusing line.len, which could
// miscount visible width for lines containing ANSI escape sequences.
func TestRenderANSIColoredContentWidth(t *testing.T) {
	colored, err := applyColor("hello", "Red")
	if err != nil {
		t.Fatalf("applyColor error: %v", err)
	}
	content := colored + "\nnormal line"
	out, err := NewBox().Style(Single).Padding(1, 2).Render("", content)
	if err != nil {
		t.Fatalf("Render error: %v", err)
	}
	lines := strings.Split(strings.TrimRight(out, "\n"), "\n")
	w0 := runewidth.StringWidth(ansi.Strip(lines[0]))
	for i, line := range lines[1:] {
		if w := runewidth.StringWidth(ansi.Strip(line)); w != w0 {
			t.Errorf("line %d width %d != expected %d: %q", i+1, w, w0, ansi.Strip(line))
		}
	}
}

func findLineContainingTitle(lines []string, title string) string {
	for _, line := range lines {
		if strings.Contains(line, title) {
			return line
		}
	}
	return ""
}

func titleStartColumn(line, title string) int {
	stripped := ansi.Strip(line)
	before, _, ok := strings.Cut(stripped, title)
	if !ok {
		return -1
	}
	return runewidth.StringWidth(before)
}
