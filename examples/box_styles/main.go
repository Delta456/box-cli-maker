package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	styles := []box.BoxStyle{
		box.Single,
		box.SingleDouble,
		box.Double,
		box.DoubleSingle,
		box.Bold,
		box.Round,
		box.Hidden,
		box.Classic,
		box.Block,
	}

	for _, style := range styles {
		b := box.NewBox().
			Padding(4, 1).
			HMargin(4).
			Style(style).
			TitleColor("#00FFB2").
			Color("#8B75FF").
			ContentColor("#12C78F").
			ContentAlign(box.Center)

		fmt.Printf("Style: %s\n%s\n", style, b.MustRender("Box CLI Maker",
			"Render highly customizable boxes\nin the terminal"))
	}
}
