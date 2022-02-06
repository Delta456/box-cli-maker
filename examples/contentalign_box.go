package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	bx := box.New(box.Config{
		Px:           2,
		Py:           0,
		Type:         "Single",
		ContentAlign: "Left", // Change this to Right or Center to see the difference
		Color:        "Green",
	})
	bx.Print("Lorem Ipsum", "LoremIpsum\nfoo bar hello world\n123456 abcdefghijk")
}
