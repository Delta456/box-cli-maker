package main

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: uint(0xE7BF55)})
	Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")
	color.HEX("#07BAA5").C16().Println("16 colors foo")
	color.HEX("#FF4081").C256().Println("256 colors foo")
	fmt.Println(Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker"))
}
