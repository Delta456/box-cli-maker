package main

import "github.com/Delta456/box-cli-maker"

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: uint(0x233a4)})
	Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")
}
