package box

import (
	"os"

	"github.com/charmbracelet/colorprofile"
)

// BoxStyle defines a built‑in border style for a Box.
type BoxStyle string

const (
	// Single is a box style with single-line borders.
	Single BoxStyle = "Single"
	// Double is a box style with double-line borders.
	Double BoxStyle = "Double"
	// Round is a box style with rounded corners.
	Round BoxStyle = "Round"
	// Bold is a box style with heavy lines.
	Bold BoxStyle = "Bold"
	// SingleDouble is a box style with single horizontal and double vertical lines.
	SingleDouble BoxStyle = "SingleDouble"
	// DoubleSingle is a box style with double horizontal and single vertical lines.
	DoubleSingle BoxStyle = "DoubleSingle"
	// Classic is a box style using plus and minus characters.
	Classic BoxStyle = "Classic"
	// Hidden is a box style with invisible (space) edges and visible "+"
	// corner markers indicating the box's extent.
	//
	// For a fully invisible box, override the corners with spaces:
	//
	//	box.NewBox().Style(box.Hidden).TopLeft(" ").TopRight(" ").BottomLeft(" ").BottomRight(" ")
	Hidden BoxStyle = "Hidden"
	// Block is a box style with solid block characters.
	Block BoxStyle = "Block"
)

// AlignType represents the horizontal alignment of content inside the box.
type AlignType string

const (
	// Center represents centered content alignment.
	Center AlignType = "Center"
	// Left represents left-aligned content.
	Left AlignType = "Left"
	// Right represents right-aligned content.
	Right AlignType = "Right"
)

// TitlePosition represents the position of the title relative to the box.
type TitlePosition string

const (
	// Inside places the title as the first lines inside the box.
	Inside TitlePosition = "Inside"
	// Top places the title on the top border of the box.
	Top TitlePosition = "Top"
	// Bottom places the title on the bottom border of the box.
	Bottom TitlePosition = "Bottom"
)

// Standard and bright ANSI color name constants usable with Color,
// TitleColor, and ContentColor.
const (
	Black   = "Black"
	Red     = "Red"
	Green   = "Green"
	Yellow  = "Yellow"
	Blue    = "Blue"
	Magenta = "Magenta"
	Cyan    = "Cyan"
	White   = "White"

	BrightBlack   = "BrightBlack"
	HiBlack       = "HiBlack"
	BrightRed     = "BrightRed"
	HiRed         = "HiRed"
	BrightGreen   = "BrightGreen"
	HiGreen       = "HiGreen"
	BrightYellow  = "BrightYellow"
	HiYellow      = "HiYellow"
	BrightBlue    = "BrightBlue"
	HiBlue        = "HiBlue"
	BrightMagenta = "BrightMagenta"
	HiMagenta     = "HiMagenta"
	BrightCyan    = "BrightCyan"
	HiCyan        = "HiCyan"
	BrightWhite   = "BrightWhite"
	HiWhite       = "HiWhite"
)

var (
	// boxes are inbuilt Box styles provided by the module
	boxes = map[BoxStyle]Box{
		Single: {
			topRight:    "┐",
			topLeft:     "┌",
			bottomRight: "┘",
			bottomLeft:  "└",
			horizontal:  "─",
			vertical:    "│",
		},
		Double: {
			topRight:    "╗",
			topLeft:     "╔",
			bottomRight: "╝",
			bottomLeft:  "╚",
			horizontal:  "═",
			vertical:    "║",
		},
		Round: {
			topRight:    "╮",
			topLeft:     "╭",
			bottomRight: "╯",
			bottomLeft:  "╰",
			horizontal:  "─",
			vertical:    "│",
		},
		Bold: {
			topRight:    "┓",
			topLeft:     "┏",
			bottomRight: "┛",
			bottomLeft:  "┗",
			horizontal:  "━",
			vertical:    "┃",
		},
		SingleDouble: {
			topRight:    "╖",
			topLeft:     "╓",
			bottomRight: "╜",
			bottomLeft:  "╙",
			horizontal:  "─",
			vertical:    "║",
		},
		DoubleSingle: {
			topRight:    "╕",
			topLeft:     "╒",
			bottomRight: "╛",
			bottomLeft:  "╘",
			horizontal:  "═",
			vertical:    "│",
		},
		Classic: {
			topRight:    "+",
			topLeft:     "+",
			bottomRight: "+",
			bottomLeft:  "+",
			horizontal:  "-",
			vertical:    "|",
		},
		Hidden: {
			topRight:    "+",
			topLeft:     "+",
			bottomRight: "+",
			bottomLeft:  "+",
			horizontal:  " ",
			vertical:    " ",
		},
		Block: {
			topRight:    "█",
			topLeft:     "█",
			bottomRight: "█",
			bottomLeft:  "█",
			horizontal:  "█",
			vertical:    "█",
		},
	}

	// colorToHex maps color names to their hexadecimal codes.
	// This includes both standard and bright ANSI colors.
	colorToHex = map[string]string{
		// Basic ANSI colors (0-7)
		Black:   "#000000",
		Red:     "#800000",
		Green:   "#008000",
		Yellow:  "#808000",
		Blue:    "#000080",
		Magenta: "#800080",
		Cyan:    "#008080",
		White:   "#C0C0C0",
		// Bright ANSI colors (8-15)
		BrightBlack:   "#808080",
		HiBlack:       "#808080",
		BrightRed:     "#FF0000",
		HiRed:         "#FF0000",
		BrightGreen:   "#00FF00",
		HiGreen:       "#00FF00",
		BrightYellow:  "#FFFF00",
		HiYellow:      "#FFFF00",
		BrightBlue:    "#0000FF",
		HiBlue:        "#0000FF",
		BrightMagenta: "#FF00FF",
		HiMagenta:     "#FF00FF",
		BrightCyan:    "#00FFFF",
		HiCyan:        "#00FFFF",
		BrightWhite:   "#FFFFFF",
		HiWhite:       "#FFFFFF",
	}

	profile = colorprofile.Detect(os.Stdout, os.Environ())
)
