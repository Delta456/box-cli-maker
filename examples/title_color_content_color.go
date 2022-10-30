package main

import (
	"github.com/Delta456/box-cli-maker/v2"
)

func main() {
	StyleCases := []string{"Single", "Double", "Single Double", "Double Single", "Bold", "Round", "Hidden", "Classic"}
	ColorTypes := []string{"Black", "Blue", "Red", "Green", "Yellow", "Cyan", "Magenta", "White", "HiBlack", "HiBlue", "HiRed", "HiGreen", "HiYellow", "HiCyan", "HiMagenta", "HiWhite"}

	for i := 0; i < len(StyleCases); i++ {
		for j := 0; j < len(ColorTypes); j++ {
			Box := box.New(box.Config{Px: 2, Py: 6, Type: StyleCases[i], Color: ColorTypes[j], ContentColor: "Cyan", TitleColor: [3]uint{215, 58, 74}})
			Box.Println("Box CLI Maker 📦", "Highly Customized Terminal	Box Maker")
		}
	}
}
