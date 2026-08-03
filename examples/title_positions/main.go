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
	positions := []box.TitlePosition{
		box.Inside,
		box.Top,
		box.Bottom,
	}

	for _, pos := range positions {
		for _, style := range styles {
			b := box.NewBox().
				Padding(2, 5).
				Style(style).
				TitlePosition(pos)

			out, err := b.Render("Box CLI Maker", "Render highly customizable boxes\nin the terminal")
			if err != nil {
				panic(err)
			}

			fmt.Printf("Style: %s, TitlePosition: %s\n%s\n\n", style, pos, out)
		}
	}
}
