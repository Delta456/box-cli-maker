package main

import (
	"github.com/Delta456/box-cli-maker"
)

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Double"})
	Box.Println(`Box CLI Maker`, `
	Make
	 Highly
			Customized
						Terminal
								 Boxes
`)
	/*box := Box.String(`Box CLI Maker`, `Make
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
