package main

import (
	"encoding/json"
	"fmt"

	"github.com/Delta456/box-cli-maker"
)

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Bold"})
	box := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")
	str, _ := json.Marshal(box)
	fmt.Println(string(str))
}
