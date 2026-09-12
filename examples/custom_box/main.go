package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	emojiBorderBox := box.NewBox().Padding(2, 3).Color(box.HiRed).TitlePosition(box.Top)
	emojiBorderBox.TopRight("📦").TopLeft("📦🚀").BottomRight("📦").BottomLeft("📦").Horizontal("📦").Vertical("📦")

	fmt.Println(emojiBorderBox.MustRender("Box CLI Maker", "Render highly customizable boxes\nin the terminal"))

	unequalBorderBox := box.NewBox().Padding(2, 3).TitlePosition(box.Bottom).TitleColor("#00FFB2").Color("#8B75FF").ContentColor("#12C78F")
	unequalBorderBox.TopRight("+").TopLeft("+").BottomRight("+").BottomLeft("++").Horizontal("-").Vertical("||")

	fmt.Println(unequalBorderBox.MustRender("Box CLI Maker", "Render highly customizable boxes\nin the terminal"))
}
