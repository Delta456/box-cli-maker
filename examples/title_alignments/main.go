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
				Padding(2, 5).
				Style(box.Single).
				TitlePosition(pos).
				TitleAlign(align)

			out, err := b.Render("Box CLI Maker", "Render highly customizable boxes\n in the terminal")
			if err != nil {
				panic(err)
			}

			fmt.Printf("Style: %s, TitlePosition: %s, TitleAligment: %s\n%s\n\n", box.Single, pos, align, out)
		}
	}
}
