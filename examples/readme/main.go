package main

import (
	"fmt"
	"strings"

	box "github.com/box-cli-maker/box-cli-maker/v3"
	"github.com/charmbracelet/x/ansi"
)

// CharmTone‑inspired palette
const (
	// Primary roles
	colorBorderPrimary  = "#8B75FF" // violet box border
	colorTitlePrimary   = "#00FFB2" // mint title
	colorContentPrimary = "#12C78F" // teal content

	// Tints and shades derived from the primaries
	colorBorderSoft  = "#C3B1FF" // lighter violet for accents
	colorBorderDeep  = "#4A35B8" // deeper violet for emphasis
	colorContentSoft = "#6EF0C1" // lighter teal highlight

)

func main() {
	b := box.NewBox().
		Style(box.Round).
		Padding(2, 1).
		TitlePosition(box.Top).
		ContentAlign(box.Left).
		Color(colorBorderPrimary).
		TitleColor(colorTitlePrimary)

	lines := []string{
		"• " + accent("9 styles", colorBorderSoft) + " (Single, Double, Round, Bold, etc.)",
		"• " + accent("Typed API", colorTitlePrimary) + " BoxStyle / TitlePosition / AlignType",
		"• " + accent("Custom styles", colorBorderDeep) + " Corner/edge glyphs + Copy()",
		"• " + accent("Titles", colorContentPrimary) + " Inside • Top • Bottom",
		"• " + accent("Align", colorContentPrimary) + " titles & content Left • Center • Right",
		"• " + accent("Padding & margin", colorBorderSoft) + " inner and outer spacing",
		"• " + accent("Wrapping", colorBorderSoft) + " WrapContent + WrapLimit",
		"• " + accent("Colors", colorContentSoft) + " ANSI names, hex, rgb/rgba",
		"• " + accent("ANSI-safe", colorTitlePrimary) + " styled input never leaks into borders",
		"• " + accent("Unicode & emoji", colorContentPrimary) + " мир",
		"• " + accent("Render / MustRender", colorBorderDeep) + " explicit error handling",
	}

	content := strings.Join(lines, "\n")
	fmt.Println(b.MustRender("Box CLI Maker", content))
}

func accent(text, hex string) string {
	c := ansi.XParseColor(hex)
	if c == nil {
		return text
	}
	style := ansi.Style{}.ForegroundColor(c)
	return style.Styled(text)
}
