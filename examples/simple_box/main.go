package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	b := box.NewBox().Padding(2, 1).Style(box.Single).Color(box.Cyan)
	s, err := b.Render("Box CLI Maker", "Render highly customizable boxes in the terminal")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
