package main

import (
	"encoding/json"
	"fmt"

	"github.com/Delta456/box-cli-maker"
)

func main() {
	/*Box := box.New(box.Config{Px: 2, Py: 5, Type: "Classic"})
	box := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")
	str, _ := json.Marshal(box)
	fmt.Println(string(str))*/

	StyleCases := []string{"Single", "Double", "Single Double", "Double Single", "Bold", "Round", "Hidden", "Classic"}
	//ColorTypes := []string{"Black", "Blue", "Red", "Green", "Yellow", "Cyan", "Magenta"}

	for i := 0; i < len(StyleCases); i++ {
		Box := box.New(box.Config{Px: 2, Py: 5, Type: StyleCases[i], TitlePos: "Bottom"})
		box := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")
		str, _ := json.Marshal(box)
		fmt.Println(StyleCases[i], ":", string(str))
	}

}
