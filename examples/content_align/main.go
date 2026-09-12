package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	content := "Render highly customizable boxes\nin the terminal\nwith aligned content"

	for _, align := range []box.AlignType{box.Left, box.Center, box.Right} {
		b := box.NewBox().
			Padding(2, 1).
			Style(box.Single).
			Color(box.Green).
			ContentAlign(align)

		fmt.Printf("ContentAlign: %s\n%s\n", align, b.MustRender("", content))
	}
}
