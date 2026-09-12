package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

// Padding adds space inside the borders, margin adds space outside them.
// Both take horizontal before vertical: Padding(px, py), Margin(mx, my).
func main() {
	content := "spacing demo"

	base := func() *box.Box {
		return box.NewBox().Color("#8B75FF").ContentColor("#12C78F")
	}

	cases := []struct {
		label string
		b     *box.Box
	}{
		{"no spacing", base()},
		{"Padding(4, 1) — 4 columns inside left/right, 1 blank row above/below", base().Padding(4, 1)},
		{"HPadding(6) — horizontal only", base().HPadding(6)},
		{"VPadding(2) — vertical only", base().VPadding(2)},
		{"Margin(6, 1) — 6 columns before every line, 1 blank line above/below", base().Margin(6, 1)},
		{"Padding(2, 1) + Margin(4, 1) — margin outside, padding inside", base().Padding(2, 1).Margin(4, 1)},
	}

	for _, c := range cases {
		fmt.Printf("%s\n%s\n", c.label, c.b.MustRender("", content))
	}
}
