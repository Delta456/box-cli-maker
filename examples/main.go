package main

func main() {
	/*titles := []string{"Box CLI Maker", "ボックスメーカー", "盒子製造商", "박스 메이커", "Créateur de boîte CLI", "Fabricante de cajas", "Qui fecit me arca CLI", "Κουτί CLI Maker"}
	lines := []string{"Make Highly Customized Terminal Boxes", "高度にカスタマイズされた端子ボックスを作成する", "製作高度定制的接線盒", "고도로 맞춤화 된 터미널 박스 만들기", "Créez des boîtes à bornes hautement personnalisées", "Haga cajas de terminales altamente personalizadas", "Fac multum Customized Terminal Pyxidas", "Δημιουργήστε πολύ προσαρμοσμένα τερματικά κουτιά"}
	for i := range titles {
		Box := box.New(box.Config{Px: 2, Py: 5})
		Box.Println(titles[i], lines[i])
	}
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Double"})
	Box.Println(`高度にカスタマイズされた端子ボック
			スを作成	する`, `
		Make
		 Highly
				Customized
							Terminal
									 Boxes
	`)
	box := Box.String(`Box CLI Maker`, `Make
	Highly
		Customized
			Terminal
					 Boxes`)
	str, _ := json.Marshal(box)
	fmt.Println(string(str))*/
	//fmt.Println("╔═════════════════════════════════╗\n║                                 ║\n║                                 ║\n║                                 ║\n║                                 ║\n║     ║\n║                                 ║\n║          Box CLI Maker          ║\n║                                 ║\n║                                 ║\n║ ║\n║  Make                           ║\n║      Highly                     ║\n║          Customized             ║\n║                      Terminal   ║\n║ Boxes  ║\n║                                 ║\n║                                 ║\n║                                 ║\n║                                 ║\n║ ║\n║                                 ║\n╚═════════════════════════════════╝")

	//StyleCases := []string{"Single", "Double", "Single Double", "Double Single", "Bold", "Round", "Hidden", "Classic"}
	//ColorTypes := []string{"Black", "Blue", "Red", "Green", "Yellow", "Cyan", "Magenta"}

	/*for i := 0; i < len(StyleCases); i++ {
		Box := box.New(box.Config{Px: 2, Py: 5, Type: StyleCases[i], ContentAlign: "Center"})
		Box2 := box.New(box.Config{Px: 2, Py: 5, Type: StyleCases[i], ContentAlign: "Right"})
		fmt.Println(Box.String("foooooo", "h") == Box2.String("foooooo", "h"))

		//str, _ := json.Marshal(box)
		//fmt.Println(StyleCases[i], ":", string(str))
		//fmt.Println(box)
	}*/
}
