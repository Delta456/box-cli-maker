package box

import (
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

		clone.Color(Green).Padding(5, 6).TitlePosition(Bottom).TopLeft("*")
		if original.color != Red {
			t.Fatalf("expected original color to remain Red, got %q", original.color)
		}
		if original.px != 1 || original.py != 2 {
			t.Fatalf("expected original padding (1,2), got (%d,%d)", original.px, original.py)
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
	if !strings.Contains(err.Error(), "multiline titles are only supported Inside title position only") {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestRenderNegativeWrapLimit(t *testing.T) {
	b := NewBox().Padding(1, 1).Style(Single).WrapContent(true).WrapLimit(-1)

	if _, err := b.Render("Title", "Content"); err == nil {
		t.Fatalf("expected error for negative wrap limit, got nil")
	} else if !strings.Contains(err.Error(), "wrapping limit cannot be negative") {
		t.Errorf("unexpected error message for negative wrap limit: %v", err)
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
	idx := strings.Index(stripped, title)
	if idx == -1 {
		return -1
	}
	return runewidth.StringWidth(stripped[:idx])
}
