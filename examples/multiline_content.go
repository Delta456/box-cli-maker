package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan", ContentColor: [3]uint{215, 58, 74}})
	Box.Print("Lorem ipsum", `Lorem ipsum dolor sit amet, 
	ボックスメーカー
	consectetur adipiscing elit. Integer nec odio. Praesent libero.
	 Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh 
	 elementum imperdiet. Duis sagittis ipsum. Praesent mauris. Fusce nec 
	 tellus sed augue semper porta. Mauris massa. Vestibulum lacinia arcu eget nulla.
	Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos.`)
}
