package main

import (
	"github.com/Delta456/box-cli-maker/v2"
)

func main() {
	Box := box.New(box.Config{Px: 2, Py: 1, Type: "Single", Color: "Cyan", TitlePos: "Top"})
	Box.Print("\033[1mBold\033[0m, works", "Btw \033[1mit works here too\033[0m, very nice")

}
