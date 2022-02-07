package main

import (
	"math"
	"strings"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/gookit/color"
)

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
	Box.Print(lolcat("Box CLI Maker"), lolcat("Make Highly Customized Terminal Boxes"))
}

func lolcat(str string) string {
	var Output string
	Freq := float64(0.1)
	for _, s := range strings.Split(str, "") {
		Output += normalStyle(Freq, s)
		Freq += 0.1

	}
	return Output
}

func normalStyle(num float64, s string) string {
	Red := uint8(math.Sin(num+0)*127 + 128)
	Green := uint8(math.Sin(num+2)*127 + 128)
	Blue := uint8(math.Sin(num+4)*127 + 128)

	return color.Rgb(Red, Blue, Green).Sprint(s)
}
