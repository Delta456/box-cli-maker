package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	positions := []box.TitlePosition{
		box.Inside,
		box.Top,
		box.Bottom,
	}
	alignments := []box.AlignType{
		box.Left,
		box.Center,
		box.Right,
	}

	for _, pos := range positions {
		for _, align := range alignments {
			b := box.NewBox().
				Padding(2, 1).
				Style(box.Single).
				TitlePosition(pos).
				TitleAlign(align)

			out, err := b.Render("Box CLI Maker", "Render highly customizable boxes\nin the terminal")
			if err != nil {
				panic(err)
			}

			fmt.Printf("TitlePosition: %s, TitleAlign: %s\n%s\n", pos, align, out)
		}
	}
}
