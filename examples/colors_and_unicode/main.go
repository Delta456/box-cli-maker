package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	// The 16 ANSI color names, applied to the border of a Single-style box.
	colors := []string{
		box.Black, box.Red, box.Green, box.Yellow, box.Blue, box.Magenta, box.Cyan, box.White,
		box.BrightBlack, box.BrightRed, box.BrightGreen, box.BrightYellow,
		box.BrightBlue, box.BrightMagenta, box.BrightCyan, box.BrightWhite,
	}
	for _, c := range colors {
		b := box.NewBox().Padding(2, 0).Style(box.Single).Color(c)
		fmt.Printf("%s\n%s\n", c, b.MustRender("", "Render highly customizable boxes in the terminal"))
	}

	// Multiline content with tabs, expanded at real 8-column stops.
	tabs := box.NewBox().
		Padding(2, 1).
		Style(box.Single).
		TitlePosition(box.Top).
		Color(box.Green).
		TitleColor(box.Cyan).
		ContentColor(box.Red)
	multi := "Render\n\thighly\n\t\tcustomizable\n\t\t\tboxes\n\t\t\t\tin the\n\t\t\t\t\tterminal"
	fmt.Printf("Multiline + tabs:\n%s\n", tabs.MustRender("Box CLI Maker", multi))

	// Wide (CJK), RTL-free non-Latin, and accented titles and content —
	// widths are measured correctly for all of them.
	titles := []string{
		"Box CLI Maker",
		"ボックスメーカー",
		"盒子製造商",
		"박스 메이커",
		"Créateur de boîte CLI",
		"Fabricante de cajas",
		"Qui fecit me arca CLI",
		"Κουτί CLI Maker",
	}
	lines := []string{
		"Render highly customizable boxes\nin the terminal",
		"端末で高度にカスタマイズ可能なボックスを\nターミナルでレンダリングする",
		"在终端中渲染高度可定制的盒子",
		"터미널에서 고도로 커스터마이즈 가능한 박스를\n렌더링하기",
		"Rendre des boîtes hautement personnalisables\ndans le terminal",
		"Renderiza cajas altamente personalizables\nen el terminal",
		"Pyxides terminales maxime configurabiles\nin terminali redde",
		"Απόδωσε εξαιρετικά προσαρμόσιμα κουτιά\nστο τερματικό",
	}
	for i := range titles {
		b := box.NewBox().
			Padding(2, 1).
			Style(box.Round).
			TitleColor("#00FFB2").
			Color("#8B75FF").
			ContentColor("#12C78F").
			ContentAlign(box.Center)

		fmt.Println(b.MustRender(titles[i], lines[i]))
	}
}
