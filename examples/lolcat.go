package main

import (
	"math"
	"strings"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
	Box.Print(Lolcat("Box CLI Maker"), Lolcat("Make Highly Customized Terminal Boxes"))
}

func Lolcat(str string) string {
	var Output string
	Freq := float64(0.1)
	for _, s := range strings.Split(str, "") {
		Output += normalStyle(Freq, s)
		Freq += 0.1

	}
	return Output
}

func normalStyle(Num float64, s string) string {
	Red := uint8(math.Sin(Num+0)*127 + 128)
	Green := uint8(math.Sin(Num+2)*127 + 128)
	Blue := uint8(math.Sin(Num+4)*127 + 128)

	return color.Rgb(Red, Blue, Green).Sprint(s)
}
