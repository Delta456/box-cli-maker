package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
	Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")
}
