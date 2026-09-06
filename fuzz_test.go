package box

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-runewidth"
)

func FuzzIsolateLineStyles(f *testing.F) {
	f.Add("\x1b[31mred\nline\x1b[0m")
	f.Add("\x1b]8;;https://x.com\x1b\\a\nb\x1b]8;;\x1b\\")
	f.Add("\x1b[1m\x1b[32m\nmix\x1b[m\n\n")
	f.Add("plain\ntext")
	f.Add("\x1b[38;5;10mx\n\x1b[0;31my")
	f.Fuzz(func(t *testing.T, s string) {
		out := isolateLineStyles(s) // must never panic or hang, even on garbage

		// The content oracles below only hold for input that decodes cleanly:
		// on malformed escapes (bare ESC, truncated CSI, unterminated OSC)
		// ansi.Strip and the decoder disagree about what the escape consumed,
		// so "visible text" itself is ill-defined. Those inputs are known
		// pass-through GIGO; the no-panic check above still covers them.
		if !utf8.ValidString(s) || !decodesCleanly(s) {
			t.Skip()
		}

		// Visible text must be untouched.
		if ansi.Strip(out) != ansi.Strip(s) {
			t.Fatalf("visible text changed:\n in %q\nout %q", s, out)
		}
		// Idempotent: a second pass must be a no-op.
		if again := isolateLineStyles(out); again != out {
			t.Fatalf("not idempotent:\n once %q\ntwice %q", out, again)
		}
		// Every line self-contained: last SGR on a line that has any
		// non-reset SGR must be a reset. Decode with the upstream parser so
		// truncated escapes are aborted the way a terminal aborts them.
		for i, line := range strings.Split(out, "\n") {
			last := ""
			var st byte
			for len(line) > 0 {
				seq, _, n, newState := ansi.DecodeSequenceWc(line, st, nil)
				if n == 0 {
					break
				}
				if isSGRSequence(seq) {
					last = seq
				}
				line = line[n:]
				st = newState
			}
			if last != "" && last != "\x1b[0m" && last != "\x1b[m" {
				t.Fatalf("line %d leaves SGR open (%q):\n in %q\nout %q", i, last, s, out)
			}
		}
	})
}

// decodesCleanly reports whether every ESC-initiated token in s is a
// complete, well-formed escape sequence: a CSI with a final byte in
// 0x40-0x7E and no embedded ESC, or an OSC terminated by ST or BEL.
func decodesCleanly(s string) bool {
	var st byte
	for len(s) > 0 {
		seq, _, n, newState := ansi.DecodeSequenceWc(s, st, nil)
		if n == 0 {
			return false
		}
		if seq[0] == 0x1b {
			// A sequence whose payload spans a newline (e.g. an OSC with a
			// literal \n inside) is bisected by the box's line splitting;
			// no line-oriented renderer can process it. GIGO.
			if strings.ContainsRune(seq, '\n') {
				return false
			}
			switch {
			case len(seq) >= 3 && seq[1] == '[':
				final := seq[len(seq)-1]
				if final < 0x40 || final > 0x7e || strings.Contains(seq[1:], "\x1b") {
					return false
				}
			case len(seq) >= 2 && seq[1] == ']':
				if !strings.HasSuffix(seq, "\x1b\\") && !strings.HasSuffix(seq, "\a") {
					return false
				}
			default:
				return false
			}
		}
		s = s[n:]
		st = newState
	}
	return true
}

// FuzzRender hammers the full pipeline with arbitrary title/content and a
// bounded config. Oracles: Render never panics; on success the rendered box
// has uniform visible row widths; rows are ANSI-self-contained for input
// that decodes cleanly.
func FuzzRender(f *testing.F) {
	f.Add("Title", "hello\nworld", uint32(0))
	f.Add("\x1b[31mR", "\x1b[32mspan\nning\x1b[0m", uint32(7))
	f.Add("🔥", "wide 🧱\n\ttabbed", uint32(42))
	f.Add("", "\x1b]8;;https://x\x1b\\link\x1b]8;;\x1b\\ text", uint32(99))
	styles := []BoxStyle{Single, Double, Round, Bold, Classic, Hidden, Block}
	positions := []TitlePosition{Inside, Top, Bottom}
	aligns := []AlignType{Left, Center, Right}
	colors := []string{"", "Red", "Cyan", "#00FF88"}
	f.Fuzz(func(t *testing.T, title, content string, cfg uint32) {
		b := NewBox().
			Style(styles[cfg%7]).
			TitlePosition(positions[(cfg>>3)%3]).
			TitleAlign(aligns[(cfg>>5)%3]).
			ContentAlign(aligns[(cfg>>7)%3]).
			Padding(int((cfg>>9)%4), int((cfg>>11)%3)).
			Color(colors[(cfg>>13)%4]).
			ContentColor(colors[(cfg>>15)%4]).
			TitleColor(colors[(cfg>>17)%4])
		if w := (cfg >> 19) % 4; w > 0 {
			b.WrapLimit(int(w) * 9)
		}
		out, err := b.Render(title, content) // must never panic
		if err != nil {
			return
		}
		if !utf8.ValidString(title) || !utf8.ValidString(content) ||
			strings.ContainsRune(title, '\r') || strings.ContainsRune(content, '\r') ||
			strings.ContainsRune(title, '\x1b') && !decodesCleanly(title) ||
			strings.ContainsRune(content, '\x1b') && !decodesCleanly(content) ||
			hasDanglingCombiningMark(title) || hasDanglingCombiningMark(content) {
			return // GIGO input: no-panic is the only guarantee
		}
		rows := strings.Split(strings.TrimSuffix(out, "\n"), "\n")
		w0 := runewidth.StringWidth(ansi.Strip(rows[0]))
		for i, row := range rows {
			if w := runewidth.StringWidth(ansi.Strip(row)); w != w0 {
				t.Fatalf("row %d width %d != row 0 width %d\ntitle=%q content=%q cfg=%d\nout=%q",
					i, w, w0, title, content, cfg, out)
			}
			if !rowSelfContained(row) {
				t.Fatalf("row %d not self-contained: %q\ntitle=%q content=%q cfg=%d",
					i, row, title, content, cfg)
			}
		}
	})
}

// rowSelfContained reports whether a rendered row closes all SGR/OSC 8 state.
func rowSelfContained(row string) bool {
	last := ""
	var st byte
	for len(row) > 0 {
		seq, _, n, newState := ansi.DecodeSequenceWc(row, st, nil)
		if n == 0 {
			break
		}
		if isSGRSequence(seq) {
			last = seq
		}
		row = row[n:]
		st = newState
	}
	return last == "" || last == "\x1b[0m" || last == "\x1b[m"
}

// hasDanglingCombiningMark reports whether any line's visible width is
// context-dependent: a line starting with a bare combining mark measures
// differently on its own than after the border/padding that the box places
// before it (grapheme clustering attaches the mark to the preceding cell).
// Detected semantically: prepending a space must add exactly one column.
// Known accepted limitation, same family as U+0601.
func hasDanglingCombiningMark(s string) bool {
	for line := range strings.SplitSeq(ansi.Strip(s), "\n") {
		w := runewidth.StringWidth(line)
		if runewidth.StringWidth(" "+line) != 1+w ||
			runewidth.StringWidth(line+" ") != w+1 {
			return true
		}
	}
	return false
}
