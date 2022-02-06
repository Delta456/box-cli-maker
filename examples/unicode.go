package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	// English, Japanese, Chinese(Traditional), Korean, French, Spanish, Latin, Greek
	titles := []string{"Box CLI Maker", "ボックスメーカー", "盒子製造商", "박스 메이커", "Créateur de boîte CLI", "Fabricante de cajas", "Qui fecit me arca CLI", "Κουτί CLI Maker"}
	lines := []string{"Make Highly Customized Terminal Boxes", "高度にカスタマイズされた端子ボックスを作成する", "製作高度定制的接線盒", "고도로 맞춤화 된 터미널 박스 만들기", "Créez des boîtes à bornes hautement personnalisées", "Haga cajas de terminales altamente personalizadas", "Fac multum Customized Terminal Pyxidas", "Δημιουργήστε πολύ προσαρμοσμένα τερματικά κουτιά"}
	StyleCases := []string{"Single", "Double", "Single Double", "Double Single", "Bold", "Round", "Hidden", "Classic"}
	for i := 0; i < len(titles); i++ {
		for j := 0; j < len(StyleCases); j++ {
			Box := box.New(box.Config{Px: 2, Py: 5, Type: StyleCases[j]})
			Box.Println(titles[i], lines[i])
		}
	}
}
