package main

import (
	"strings"

	"github.com/Delta456/box-cli-maker/v2"
)

func main() {
	bx := box.New(box.Config{
		Px:            2,
		Py:            0,
		Type:          "Single",
		ContentAlign:  "Left",
		Color:         "Green",
		TitlePos:      "Top",
		ContentColor:  uint(0xa77032),
		TitleColor:    "Cyan",
		AllowWrapping: true,
	})
	bx.Println("title", strings.Repeat("	盒子製 造商,📦 ", 160))
}
