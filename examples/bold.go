package main

import (
	"fmt"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	Box := box.New(box.Config{Px: 2, Py: 1, Type: "Single", TitlePos: "Bottom", TitleColor: uint(0x1d9bf0), ContentColor: uint(0x1d9bf0)})
	Box.Println("\033[1mBroken\033[0m, sorry", "Btw \033[1mNot Broken\033[0m, very nice \n hello")
	//fmt.Println("\x1b[36m┌ \x1b[1mBroken\x1b[0m\x1b[36m, sorry ──────────────┐\x1b[0m")
	fmt.Println(color.New(color.OpBold, color.Red).Sprint("Broken"))

}
