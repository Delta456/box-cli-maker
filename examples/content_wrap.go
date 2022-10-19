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
		Color:         "Green",
		TitlePos:      "Top",
		AllowWrapping: true,
		/*WrappingLimit: num,*/ // Uncomment and replace the placeholder with any num to see the change.
	})
	bx.Println("Content Wrappingg works!", strings.Repeat("	Box CLI Maker 盒子製 造商,📦 ", 160))
}
