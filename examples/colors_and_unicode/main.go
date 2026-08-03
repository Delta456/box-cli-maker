package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	styles := []box.BoxStyle{
		box.Single,
		box.Double,
		box.SingleDouble,
		box.DoubleSingle,
		box.Bold,
		box.Round,
		box.Hidden,
		box.Classic,
		box.Block,
	}

	colors := []string{
		box.Black, box.Blue, box.Red, box.Green, box.Yellow, box.Cyan, box.Magenta, box.White,
		box.BrightBlack, box.BrightBlue, box.BrightRed, box.BrightGreen, box.BrightYellow,
		box.BrightCyan, box.BrightMagenta, box.BrightWhite,
	}

	// Color and style combinations (border color only).
	for _, style := range styles {
		for _, c := range colors {
			b := box.NewBox().
				Padding(2, 5).
				Style(style).
				Color(c)

			out, err := b.Render("Box CLI Maker", "Render highly customizable boxes\nin the terminal")
			if err != nil {
				panic(err)
			}

			fmt.Printf("Style: %s, Color: %s\n%s\n\n", style, c, out)
		}
	}

	// Multiline + tabs with colored title/content.
	b := box.NewBox().
		Padding(2, 5).
		Style(box.Single).
		TitlePosition(box.Top).
		Color(box.Green).
		TitleColor(box.Cyan).
		ContentColor(box.Red)

	multi := "Render\n\thighly\n\t\tcustomizable\n\t\t\tboxes\n\t\t\t\tin the\n\t\t\t\t\tterminal"

	out, err := b.Render("Box CLI Maker", multi)
	if err != nil {
		panic(err)
	}
	fmt.Println("Multiline + tabs:")
	fmt.Println(out)

	// Unicode / emoji demo.
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
		"在终端中渲染高度可定制的盒子\n",
		"터미널에서 고도로 커스터마이즈 가능한 박스를\n렌더링하기",
		"Rendre des boîtes hautement personnalisables\ndans le terminal",
		"Renderiza cajas de cajas altamente personalizables\nen el terminal",
		"Pyxides terminales maxime configurabiles\nin terminali redde",
		"Απόδωσε εξαιρετικά προσαρμόσιμα κουτιά\nστο τερματικό",
	}

	for i := range titles {
		for _, style := range styles {
			b := box.NewBox().
				Padding(2, 5).
				Style(style).
				TitleColor("#00ffb2").
				Color("#8B75FF").
				ContentColor("#12c78f").
				ContentAlign(box.Center)

			out, err := b.Render(titles[i], lines[i])
			if err != nil {
				panic(err)
			}

			fmt.Printf("Box style: %s\n%s\n\n", style, out)
		}
	}
}
