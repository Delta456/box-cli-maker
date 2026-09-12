package main

import (
	"fmt"
	"strings"

	box "github.com/box-cli-maker/box-cli-maker/v3"
	"github.com/charmbracelet/x/ansi"
)

// CharmTone-inspired palette
const (
	violet = "#8B75FF" // border
	mint   = "#00FFB2" // title, temperature
	teal   = "#12C78F" // values
	amber  = "#FFC24B" // sun and sky
	slate  = "#8A8F98" // muted labels
)

// The hero scene: a small weather card — one box, printed the way any CLI
// prints one. The CJK title exercises the titled-bar width handling; the
// glyphs are chosen width-stable so the card renders identically across
// terminals.
func main() {
	fmt.Println(accent("$", mint) + accent(" wthr tokyo", slate))
	fmt.Println()

	card := box.NewBox().
		Style(box.Round).
		Padding(3, 1).
		TitlePosition(box.Top).
		Color(violet).
		TitleColor(mint)

	fmt.Println(card.MustRender("tokyo · 東京", strings.Join([]string{
		accent("🌞  clear sky", amber),
		accent("31°C", mint) + " · feels like 34°C",
		"",
		accent("wind      ", slate) + accent("12 km/h", teal),
		accent("humidity  ", slate) + accent("68%", teal),
		accent("sunset    ", slate) + accent("18:42", teal),
	}, "\n")))
}

// accent colors text via the same x/ansi dependency the library uses.
func accent(text, hex string) string {
	c := ansi.XParseColor(hex)
	if c == nil {
		return text
	}
	return ansi.Style{}.ForegroundColor(c).Styled(text)
}
