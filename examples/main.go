package main

import (
	"fmt"

	"github.com/Delta456/box-cli-maker"
)

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Double"})
	fmt.Println(Box.String(`Box CLI Maker 
	foo bar ehidiuhfidlf
	fdhfkrhf`, `foo 
	bv fd
	feruiegye`))

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
