package box

var boxs map[string]Box = map[string]Box{
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
