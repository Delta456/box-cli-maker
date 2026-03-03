package box

import (
	"image/color"
	"strings"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-runewidth"
)

func TestAddVertPadding(t *testing.T) {
	b := &Box{vertical: "|"}
	b.py = 2

	// innerWidth is the visible width between the vertical borders.
	got, err := b.addVertPadding(4)
	if err != nil {
		t.Fatalf("addVertPadding unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("expected 2 padding lines, got %d", len(got))
	}

	want := "|    |" // len-2 = 4 spaces
	for i, line := range got {
		if line != want {
			t.Errorf("line %d: expected %q, got %q", i, want, line)
		}
	}
}

func TestLongestLineBasicAndTabs(t *testing.T) {
	lines := []string{"short", "longer"}
	longest, expanded := longestLine(lines)

	if longest != len("longer") {
		t.Errorf("expected longest %d, got %d", len("longer"), longest)
	}
	if len(expanded) != len(lines) {
		t.Fatalf("expected %d expanded lines, got %d", len(lines), len(expanded))
	}
	if expanded[0].line != "short" || expanded[0].len != len("short") {
		t.Errorf("unexpected expansion for 'short': %#v", expanded[0])
	}

	// Tab expansion: tab stops every 8 columns; "a\tb" -> "a" + 7 spaces + "b" (visible width 9).
	lines = []string{"a\tb"}
	longest, expanded = longestLine(lines)
	wantLine := "a" + strings.Repeat(" ", 7) + "b"
	if expanded[0].line != wantLine {
		t.Errorf("tab-expanded line mismatch: want %q, got %q", wantLine, expanded[0].line)
	}
	if longest != 9 {
		t.Errorf("expected longest 9 for tab-expanded line, got %d", longest)
	}

	// ANSI-colored line should be measured by visible width
	plain := "abc"
	colored := "\x1b[31mabc\x1b[0m" // same visible width as plain
	longest, _ = longestLine([]string{plain, colored})
	if longest != len(plain) {
		t.Errorf("expected longest visible width %d, got %d", len(plain), longest)
	}
}

func TestFormatLine_LeftAlignAndError(t *testing.T) {
	// Happy path: left-aligned content, no title
	b := &Box{vertical: "|"}
	b.contentAlign = Left

	lines := []expandedLine{{line: "hi", len: 2}}
	sideMargin := " "
	texts, err := b.formatLine(lines, 2, 0, sideMargin, "", nil)
	if err != nil {
		t.Fatalf("formatLine unexpected error: %v", err)
	}
	if len(texts) != 1 {
		t.Fatalf("expected 1 formatted line, got %d", len(texts))
	}

	// With left alignment and no additional padding the layout is:
	// sep + sideMargin + line + spacing + sideMargin + sep
	// where spacing == sideMargin and sep == "|".
	want := "| hi |"
	if texts[0] != want {
		t.Errorf("formatted line mismatch: want %q, got %q", want, texts[0])
	}

	// Title line with Inside position should not call findAlign (even if contentAlign is invalid).
	b = &Box{vertical: "|"}
	b.titlePos = Inside
	b.contentAlign = AlignType("InvalidAlign")
	lines = []expandedLine{{line: "Title", len: len("Title")}}
	texts, err = b.formatLine(lines, len("Title"), 1, sideMargin, "Title", nil)
	if err != nil {
		t.Fatalf("formatLine for title line should not error, got: %v", err)
	}
	if len(texts) != 1 || !strings.Contains(texts[0], "Title") {
		t.Errorf("expected formatted title line containing 'Title', got %q", texts[0])
	}

	// Error path: invalid alignment when not formatting a title line.
	b = &Box{vertical: "|"}
	b.contentAlign = AlignType("InvalidAlign")
	lines = []expandedLine{{line: "hi", len: 2}}
	_, err = b.formatLine(lines, 2, 0, sideMargin, "", nil)
	if err == nil {
		t.Fatalf("expected error for invalid content alignment, got nil")
	}
}

func TestFindAlign(t *testing.T) {
	cases := []struct {
		name   string
		align  AlignType
		want   string
		wantOK bool
	}{
		{"default-left", "", leftAlign, true},
		{"left", Left, leftAlign, true},
		{"center", Center, centerAlign, true},
		{"right", Right, rightAlign, true},
		{"invalid", AlignType("Invalid"), "", false},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			b := &Box{}
			b.contentAlign = tt.align
			got, err := b.findAlign()
			if tt.wantOK && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tt.wantOK && err == nil {
				t.Fatalf("expected error for invalid align, got nil")
			}
			if tt.wantOK && got != tt.want {
				t.Errorf("expected %q, got %q", tt.want, got)
			}
		})
	}
}

func TestGetConvertedColorAndApplyColor(t *testing.T) {
	c, err := getConvertedColor(Green)
	if err != nil {
		t.Fatalf("expected non-error from getConvertedColor: %v", err)
	}
	if c == nil {
		t.Fatalf("expected non-nil color from getConvertedColor")
	}

	text := "hello"
	got, err := applyColor(text, "")
	if err != nil {
		t.Fatalf("expected no error when color is empty: %v", err)
	}
	if got != text {
		t.Errorf("expected text unchanged when color is empty, got %q", got)
	}

	colored, err := applyColor(text, Green)
	if err != nil {
		t.Fatalf("unexpected error applying valid color: %v", err)
	}
	if ansi.Strip(colored) != text {
		t.Errorf("expected stripped colored text to equal %q, got %q", text, ansi.Strip(colored))
	}

	if _, err := applyColor(text, "NotAColor"); err == nil {
		t.Fatalf("expected error when applying unknown color name")
	}
}

func TestStringColorToHex(t *testing.T) {
	if got := stringColorToHex(Green); got != "#008000" {
		t.Errorf("expected Green -> #008000, got %q", got)
	}
	if got := stringColorToHex("UnknownColor"); got != "" {
		t.Errorf("expected unknown color to map to empty string, got %q", got)
	}
}

func TestAddStylePreservingOriginalFormat(t *testing.T) {
	// Without reset sequence, the whole string is passed once to f.
	got := addStylePreservingOriginalFormat("foo", strings.ToUpper)
	if got != "FOO" {
		t.Errorf("expected FOO, got %q", got)
	}

	// With reset sequences, they are removed and styling is applied around segments.
	input := "foo\033[0mbar"
	got = addStylePreservingOriginalFormat(input, func(a string) string { return "[" + a + "]" })
	if strings.Contains(got, "\033[0m") {
		t.Errorf("expected reset sequences to be removed, got %q", got)
	}
	if got != "[foo][bar]" {
		t.Errorf("expected styled segments [foo][bar], got %q", got)
	}
}

func TestParseColorString(t *testing.T) {
	// Known color name should result in a non-nil color.
	c, err := parseColorString(Green)
	if err != nil {
		t.Fatalf("expected non-error for Green: %v", err)
	}
	if c == nil {
		t.Fatalf("expected non-nil color for Green")
	}

	// Hex and rgb/rgba forms supported by ansi.XParseColor should also parse.
	for _, tc := range []string{
		"#0F0",                     // #RGB short hex
		"#00FF00",                  // #RRGGBB full hex
		"rgb:0000/ffff/0000",       // rgb:RRRR/GGGG/BBBB
		"rgba:ffff/0000/0000/ffff", // rgba:RRRR/GGGG/BBBB/AAAA
	} {
		c, err := parseColorString(tc)
		if err != nil {
			t.Fatalf("expected non-error for %q: %v", tc, err)
		}
		if c == nil {
			t.Fatalf("expected non-nil color for %q", tc)
		}
	}

	// Invalid color string should return an error.
	if _, err := parseColorString("not-a-real-color"); err == nil {
		t.Fatalf("expected error for invalid color string")
	}
}

func TestApplyConvertedColor(t *testing.T) {
	// Nil color should return the original string.
	orig := "plain"
	if got := applyConvertedColor(orig, nil); got != orig {
		t.Errorf("expected original string when color is nil, got %q", got)
	}

	c := color.RGBA{R: 1, G: 2, B: 3, A: 255}

	// Single line: fast path.
	single := "one line"
	coloredSingle := applyConvertedColor(single, c)
	if ansi.Strip(coloredSingle) != single {
		t.Errorf("expected stripped colored single-line text to equal %q, got %q", single, ansi.Strip(coloredSingle))
	}

	// Multi-line: ensure each line is styled and newlines preserved.
	multi := "line1\nline2"
	coloredMulti := applyConvertedColor(multi, c)
	if strings.Count(coloredMulti, "\n") != 1 {
		t.Errorf("expected one newline in colored multi-line text, got %q", coloredMulti)
	}
	if ansi.Strip(coloredMulti) != multi {
		t.Errorf("expected stripped colored multi-line text to equal %q, got %q", multi, ansi.Strip(coloredMulti))
	}
}

func TestApplyColorBar(t *testing.T) {
	top := "+------TITLE------+"
	bottom := "+-----------------+"
	title := "TITLE"

	// Early return when titleColor is empty or title is empty.
	b := &Box{}
	gotTop, gotBottom, err := b.applyColorBar(top, bottom, title)
	if err != nil {
		t.Fatalf("applyColorBar unexpected error: %v", err)
	}
	if gotTop != top || gotBottom != bottom {
		t.Errorf("expected bars unchanged when titleColor is empty")
	}

	// Title at top: top bar should be recolored but visually unchanged when stripped.
	b = &Box{}
	b.titleColor = BrightRed
	b.color = BrightBlue
	b.titlePos = Top
	gotTop, gotBottom, err = b.applyColorBar(top, bottom, title)
	if err != nil {
		t.Fatalf("applyColorBar unexpected error for top title: %v", err)
	}
	if ansi.Strip(gotTop) != top {
		t.Errorf("expected stripped top bar to remain %q, got %q", top, ansi.Strip(gotTop))
	}
	if ansi.Strip(gotBottom) != bottom {
		t.Errorf("expected stripped bottom bar to remain %q, got %q", bottom, ansi.Strip(gotBottom))
	}
	if !strings.Contains(ansi.Strip(gotTop), title) {
		t.Errorf("expected title to remain present in top bar, got %q", ansi.Strip(gotTop))
	}

	// Title at bottom: bottom bar should be recolored but visually unchanged when stripped.
	topPlain := "+-----------------+"
	bottomWithTitle := "+------TITLE------+"
	b = &Box{}
	b.titleColor = BrightRed
	b.color = BrightBlue
	b.titlePos = Bottom
	gotTop, gotBottom, err = b.applyColorBar(topPlain, bottomWithTitle, title)
	if err != nil {
		t.Fatalf("applyColorBar unexpected error for bottom title: %v", err)
	}
	if ansi.Strip(gotTop) != topPlain {
		t.Errorf("expected stripped top bar to remain %q, got %q", topPlain, ansi.Strip(gotTop))
	}
	if ansi.Strip(gotBottom) != bottomWithTitle {
		t.Errorf("expected stripped bottom bar to remain %q, got %q", bottomWithTitle, ansi.Strip(gotBottom))
	}
	if !strings.Contains(ansi.Strip(gotBottom), title) {
		t.Errorf("expected title to remain present in bottom bar, got %q", ansi.Strip(gotBottom))
	}

	// No box color set: bars should remain unchanged so existing styling is preserved.
	coloredTitle, err := applyColor(title, BrightRed)
	if err != nil {
		t.Fatalf("unexpected error coloring title: %v", err)
	}
	topWithColoredTitle := strings.Replace(top, title, coloredTitle, 1)
	b = &Box{}
	b.titleColor = BrightRed
	b.color = ""
	b.titlePos = Top
	gotTop, gotBottom, err = b.applyColorBar(topWithColoredTitle, bottom, coloredTitle)
	if err != nil {
		t.Fatalf("applyColorBar unexpected error when Color is empty: %v", err)
	}
	if gotTop != topWithColoredTitle {
		t.Errorf("expected top bar unchanged when Color is empty; got %q", gotTop)
	}
	if gotBottom != bottom {
		t.Errorf("expected bottom bar unchanged when Color is empty; got %q", gotBottom)
	}
}

func TestCharWidth(t *testing.T) {
	if w := charWidth("abc"); w != 3 {
		t.Errorf("expected width 3 for 'abc', got %d", w)
	}

	colored := "\x1b[31mabc\x1b[0m" // red "abc"
	if w := charWidth(colored); w != 3 {
		t.Errorf("expected visible width 3 for colored 'abc', got %d", w)
	}

	if w := charWidth(""); w != 1 {
		t.Errorf("expected fallback width 1 for empty string, got %d", w)
	}
}

func TestBuildSegment(t *testing.T) {
	fill := "📦"
	hw := runewidth.StringWidth(fill)

	seg := buildSegment(fill, hw*3, hw)
	if runewidth.StringWidth(seg) != hw*3 {
		t.Fatalf("expected segment visual width %d, got %d", hw*3, runewidth.StringWidth(seg))
	}

	seg = buildSegment(fill, hw*3+1, hw)
	if runewidth.StringWidth(seg) != hw*3+1 {
		t.Fatalf("expected segment visual width %d, got %d", hw*3+1, runewidth.StringWidth(seg))
	}
}

func TestBuildPlainBar(t *testing.T) {
	fill := "📦"
	hw := runewidth.StringWidth(fill)
	lineWidth := hw*10 + 2 // 10 emojis + 2 corners

	left := fill
	right := fill
	bar := buildPlainBar(left, fill, right, hw, hw, lineWidth, hw)
	if w := runewidth.StringWidth(bar); w != lineWidth {
		t.Fatalf("expected bar visual width %d, got %d", lineWidth, w)
	}
	if !strings.HasPrefix(bar, fill) || !strings.HasSuffix(bar, fill) {
		t.Errorf("expected bar to start and end with fill, got %q", bar)
	}
}

func TestBuildTitledBar_LeftAlignedWithEmojiFill(t *testing.T) {
	fill := "📦"
	hw := runewidth.StringWidth(fill)
	title := "Box CLI Maker"

	left := fill
	right := fill
	leftW := hw
	rightW := hw
	lineWidth := hw*20 + leftW + rightW

	bar := buildTitledBar(left, fill, right, leftW, rightW, lineWidth, hw, title)
	if w := runewidth.StringWidth(ansi.Strip(bar)); w != lineWidth {
		t.Fatalf("expected bar visual width %d, got %d", lineWidth, w)
	}
	plain := ansi.Strip(bar)
	if !strings.Contains(plain, " "+title+" ") {
		t.Fatalf("expected bar to contain title with spaces, got %q", plain)
	}
	if !strings.HasPrefix(plain, fill+" ") {
		t.Errorf("expected left corner followed by space then title, got %q", plain)
	}
	if !strings.HasSuffix(plain, fill) {
		t.Errorf("expected bar to end with fill glyph, got %q", plain)
	}
}
