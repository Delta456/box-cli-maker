package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	baseBox := box.NewBox().Padding(2, 1).Style(box.Single)

	greenBox := baseBox.Copy().Color(box.Green)
	redBox := baseBox.Copy().Color(box.Red)

	fmt.Println(baseBox.MustRender("Base box", "This is the base box."))
	fmt.Println(greenBox.MustRender("Green box", "This is a green box."))
	fmt.Println(redBox.MustRender("Red box", "This is a red box."))
}
