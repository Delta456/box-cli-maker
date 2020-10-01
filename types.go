package box

import "github.com/fatih/color"

var (
	boxs map[string]Box = map[string]Box{
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

	fgColors map[string]color.Attribute = map[string]color.Attribute{
		"Black":   color.FgBlack,
		"Blue":    color.FgBlue,
		"Red":     color.FgRed,
		"Green":   color.FgGreen,
		"Yellow":  color.FgYellow,
		"Cyan":    color.FgCyan,
		"Magenta": color.FgMagenta,
	}

	fgHiColors map[string]color.Attribute = map[string]color.Attribute{
		"HiBlack":   color.FgHiBlack,
		"HiBlue":    color.FgHiBlue,
		"HiRed":     color.FgHiRed,
		"HiGreen":   color.FgHiGreen,
		"HiYellow":  color.FgHiYellow,
		"HiCyan":    color.FgHiCyan,
		"HiMagenta": color.FgHiMagenta,
	}
	// needed for windows cmd and powershell
	Output = color.Output
)
