package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	b := box.NewBox().
		Style(box.Round).
		Padding(2, 1).
		Color(box.White).
		TopBorderColor(box.Red).
		RightBorderColor(box.Green).
		BottomBorderColor(box.Blue).
		LeftBorderColor(box.Yellow).
		TitleColor(box.BrightMagenta)

	fmt.Print(b.MustRender("Custom border colors", "Each side has its own color."))
}
