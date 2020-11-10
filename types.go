package box

import "github.com/fatih/color"

var (
	// boxes are inbuilt styles provided by the module
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
	// fgColors are the inbuilt Foreground Colors provided by the module
	fgColors map[string]color.Attribute = map[string]color.Attribute{
		"Black":   color.FgBlack,
		"Blue":    color.FgBlue,
		"Red":     color.FgRed,
		"Green":   color.FgGreen,
		"Yellow":  color.FgYellow,
		"Cyan":    color.FgCyan,
		"Magenta": color.FgMagenta,
	}
	// fgHiColors are the inbuilt Hi-intensity Foreground Colors provided by the module
	fgHiColors map[string]color.Attribute = map[string]color.Attribute{
		"HiBlack":   color.FgHiBlack,
		"HiBlue":    color.FgHiBlue,
		"HiRed":     color.FgHiRed,
		"HiGreen":   color.FgHiGreen,
		"HiYellow":  color.FgHiYellow,
		"HiCyan":    color.FgHiCyan,
		"HiMagenta": color.FgHiMagenta,
	}
	// Output is an io.Writer and instance of
	// color.Output which is needed
	// for outputting on Windows Console
	Output = color.Output
)
