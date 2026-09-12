package main

import (
	"fmt"
	"image/color"
	"math"
	"strings"

	box "github.com/box-cli-maker/box-cli-maker/v3"
	"github.com/charmbracelet/x/ansi"
)

func main() {
	b := box.NewBox().Padding(2, 1).Style(box.Single).Color("#8B75FF").ContentAlign(box.Center)
	fmt.Println(b.MustRender(lolcat("Box CLI Maker"), lolcat("Render highly customizable boxes\nin the terminal")))
}

func lolcat(str string) string {
	var output strings.Builder
	freq := float64(0.1)
	for s := range strings.SplitSeq(str, "") {
		output.WriteString(normalStyle(freq, s))
		freq += 0.1
	}
	return output.String()
}

func normalStyle(num float64, s string) string {
	red := uint8(math.Sin(num+0)*127 + 128)
	green := uint8(math.Sin(num+2)*127 + 128)
	blue := uint8(math.Sin(num+4)*127 + 128)

	c := &color.RGBA{R: red, G: green, B: blue, A: 255}
	style := ansi.Style{}.ForegroundColor(c)
	return style.Styled(s)
}
