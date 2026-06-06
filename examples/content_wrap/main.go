package main

import (
	"fmt"
	"strings"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	b := box.NewBox().Padding(2, 0).
		Margin(5, 2).
		Style(box.Single).
		Color(box.Green).
		TitlePosition(box.Top).
		TitleAlign(box.Center).
		WrapContent(true)
		// Provide your limit with WrapLimit if needed

	s, err := b.Render("Content Wrap!", strings.Repeat("\tBox CLI Maker 盒子製 造商,📦 ", 160))
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
