// Command showcase renders one README showcase box to stdout, named by
// argument (a style, title position/alignment, or content alignment).
// scripts/screenshots/regen.sh pipes each subject through shoot.sh to
// produce the img/*.png files the README references.
package main

import (
	"fmt"
	"os"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

const (
	violet = "#8B75FF"
	mint   = "#00FFB2"
	slate  = "#8A8F98"

	// Content stays short and dim so the border — each image's actual
	// subject — is the dominant visual element.
	tagline = "Render beautiful boxes\nin the terminal"
	// Ragged line lengths make content alignment visible.
	ragged = "Render\nhighly customizable boxes\nin the terminal"
)

func base() *box.Box {
	return box.NewBox().
		Padding(3, 1).
		Style(box.Single).
		Color(violet).
		TitleColor(mint).
		ContentColor(slate).
		ContentAlign(box.Center)
}

func subject(name string) (*box.Box, string) {
	styles := map[string]box.BoxStyle{
		"single": box.Single, "single_double": box.SingleDouble,
		"double": box.Double, "double_single": box.DoubleSingle,
		"bold": box.Bold, "round": box.Round, "hidden": box.Hidden,
		"classic": box.Classic, "block": box.Block,
	}
	if st, ok := styles[name]; ok {
		// No title: nothing competes with the border style itself.
		return base().Style(st), tagline
	}
	switch name {
	case "top":
		return base().TitlePosition(box.Top), tagline
	case "bottom":
		return base().TitlePosition(box.Bottom), tagline
	case "top_center":
		return base().TitlePosition(box.Top).TitleAlign(box.Center), tagline
	case "top_right":
		return base().TitlePosition(box.Top).TitleAlign(box.Right), tagline
	case "bottom_center":
		return base().TitlePosition(box.Bottom).TitleAlign(box.Center), tagline
	case "bottom_right":
		return base().TitlePosition(box.Bottom).TitleAlign(box.Right), tagline
	case "inside_left":
		return base().TitleAlign(box.Left), tagline
	case "inside_right":
		return base().TitleAlign(box.Right), tagline
	case "left":
		return base().ContentAlign(box.Left), ragged
	case "right":
		return base().ContentAlign(box.Right), ragged
	}
	return nil, ""
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: showcase <subject>")
		os.Exit(2)
	}
	name := os.Args[1]
	b, content := subject(name)
	if b == nil {
		fmt.Fprintf(os.Stderr, "unknown subject %q\n", name)
		os.Exit(2)
	}
	title := "Box CLI Maker"
	if _, isStyle := map[string]bool{
		"single": true, "single_double": true, "double": true, "double_single": true,
		"bold": true, "round": true, "hidden": true, "classic": true, "block": true,
		"left": true, "right": true,
	}[name]; isStyle {
		title = ""
	}
	fmt.Println(b.MustRender(title, content))
}
