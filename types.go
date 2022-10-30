package box

import (
	"github.com/gookit/color"
)

var (
	// boxes are inbuilt Box styles provided by the module
	boxes map[string]Box = map[string]Box{
		"Single": {
			TopRight:    "┐",
			TopLeft:     "┌",
			BottomRight: "┘",
			BottomLeft:  "└",
			Horizontal:  "─",
			Vertical:    "│",
		},
		"Double": {
			TopRight:    "╗",
			TopLeft:     "╔",
			BottomRight: "╝",
			BottomLeft:  "╚",
			Horizontal:  "═",
			Vertical:    "║",
		},
		"Round": {
			TopRight:    "╮",
			TopLeft:     "╭",
			BottomRight: "╯",
			BottomLeft:  "╰",
			Horizontal:  "─",
			Vertical:    "│",
		},
		"Bold": {
			TopRight:    "┓",
			TopLeft:     "┏",
			BottomRight: "┛",
			BottomLeft:  "┗",
			Horizontal:  "━",
			Vertical:    "┃",
		},
		"Single Double": {
			TopRight:    "╖",
			TopLeft:     "╓",
			BottomRight: "╜",
			BottomLeft:  "╙",
			Horizontal:  "─",
			Vertical:    "║",
		},
		"Double Single": {
			TopRight:    "╕",
			TopLeft:     "╒",
			BottomRight: "╛",
			BottomLeft:  "╘",
			Horizontal:  "═",
			Vertical:    "│",
		},
		"Classic": {
			TopRight:    "+",
			TopLeft:     "+",
			BottomRight: "+",
			BottomLeft:  "+",
			Horizontal:  "-",
			Vertical:    "|",
		},
		"Hidden": {
			TopRight:    "+",
			TopLeft:     "+",
			BottomRight: "+",
			BottomLeft:  "+",
			Horizontal:  " ",
			Vertical:    " ",
		},
	}
	// fgColors are inbuilt Foreground Colors provided by the module
	fgColors map[string]color.Color = map[string]color.Color{
		"Black":   color.FgBlack,
		"Blue":    color.FgBlue,
		"Red":     color.FgRed,
		"Green":   color.FgGreen,
		"Yellow":  color.FgYellow,
		"Cyan":    color.FgCyan,
		"Magenta": color.FgMagenta,
		"White":   color.FgWhite,
	}
	// fgHiColors are inbuilt Hi-intensity Foreground Colors provided by the module
	fgHiColors map[string]color.Color = map[string]color.Color{
		"HiBlack":   color.FgDarkGray,
		"HiBlue":    color.FgLightBlue,
		"HiRed":     color.FgLightRed,
		"HiGreen":   color.FgLightGreen,
		"HiYellow":  color.FgLightYellow,
		"HiCyan":    color.FgLightCyan,
		"HiMagenta": color.FgLightMagenta,
		"HiWhite":   color.FgLightWhite,
	}
)
