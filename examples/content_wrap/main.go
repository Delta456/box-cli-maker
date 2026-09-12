package main

import (
	"fmt"
	"strings"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	content := strings.TrimSpace(strings.Repeat("Box CLI Maker 盒子製造商 📦 ", 12))

	// WrapLimit sets an explicit wrap width (and enables wrapping), so the
	// output is identical on a TTY, in a pipe, and in CI.
	b := box.NewBox().
		Padding(2, 0).
		Margin(2, 1).
		Style(box.Single).
		Color(box.Green).
		TitlePosition(box.Top).
		TitleAlign(box.Center).
		WrapLimit(48)

	fmt.Println(b.MustRender("Content Wrap", content))

	// For automatic wrapping to two-thirds of the terminal width, use
	// WrapContent(true) instead — it measures the terminal, so it needs
	// a TTY on stdout.
}
