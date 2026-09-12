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
	teal   = "#12C78F"

	tagline = "Render highly customizable\nboxes in the terminal\nfor your Go CLIs"
	// Ragged line lengths make content alignment visible.
	ragged = "Render\nhighly customizable boxes\nin the terminal"
)

func base() *box.Box {
	return box.NewBox().
		Padding(2, 2).
		Style(box.Single).
		Color(violet).
		TitleColor(mint).
		ContentColor(teal).
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
	b, content := subject(os.Args[1])
	if b == nil {
		fmt.Fprintf(os.Stderr, "unknown subject %q\n", os.Args[1])
		os.Exit(2)
	}
	fmt.Println(b.MustRender("Box CLI Maker", content))
}
