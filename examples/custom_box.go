package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	config := box.Config{Px: 2, Py: 3, Type: "", TitlePos: "Inside", Color: "Green"}
	boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
	boxNew.Println("Box CLI Maker", "Make Highly Customized Terminal Boxes")
}
