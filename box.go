package box

import (
	"fmt"
	"os"
	"runtime"

	//"runtime"
	"strings"

	//"syscall"

	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
	"github.com/xo/terminfo"
)

const (
	n1 = "\n"

	// 1 = separator, 2 = spacing, 3 = line; 4 = oddSpace; 5 = space; 6 = sideMargin
	centerAlign = "%[1]s%[2]s%[3]s%[4]s%[2]s%[1]s"
	leftAlign   = "%[1]s%[6]s%[3]s%[4]s%[2]s%[5]s%[1]s"
	rightAlign  = "%[1]s%[2]s%[4]s%[5]s%[3]s%[6]s%[1]s"
)

// Box struct defines the Box to be made.
type Box struct {
	TopRight    string // TopRight corner used for Symbols
	TopLeft     string // TopLeft corner used for Symbols
	Vertical    string // Symbols used for Vertical Bars
	BottomRight string // BottomRight corner used for Symbols
	BottomLeft  string // BotromLeft corner used for Symbols
	Horizontal  string // Symbols used for Horizontal Bars
	Config             // Config for the Box struct
}

// Config is the configuration for the Box struct
type Config struct {
	Py           int         // Horizontal Padding
	Px           int         // Vertical Padding
	ContentAlign string      // Content Alignment inside Box
	Type         string      // Type of Box
	TitlePos     string      // Title Position
	Color        interface{} // Color of Box
}

// New takes struct Config and returns the specified Box struct.
func New(config Config) Box {
	// Default Box Type is Single
	if config.Type == "" {
		boxNew := boxes["Single"]
		boxNew.Config = config
		return boxNew
		// Check if the Box Type provided is valid else panic
	} else if _, ok := boxes[config.Type]; ok {
		boxNew := boxes[config.Type]
		boxNew.Config = config
		return boxNew
	}
	panic("Invalid Box Type provided")
}

// String returns the string representation of Box.
func (b Box) String(title, lines string) string {
	var lines2 []string

	// Default Position is Inside, no warning for invalid TitlePos as it is done
	// in toString() method
	if b.TitlePos == "" {
		b.TitlePos = "Inside"
	}
	// if Title is empty then TitlePos should be Inside
	if title != "" {
		if b.TitlePos != "Inside" && strings.Contains(title, "\n") {
			panic("Multilines are only supported inside only")
		}
		if b.TitlePos == "Inside" {
			lines2 = append(lines2, strings.Split(title, n1)...)
			lines2 = append(lines2, []string{""}...) // for empty line between title and content
		}
	}
	lines2 = append(lines2, strings.Split(lines, n1)...)
	return b.toString(title, lines2)
}

// toString is same as String except that it is used for printing Boxes
func (b Box) toString(title string, lines []string) string {
	titleLen := len(strings.Split(title, n1))
	sideMargin := strings.Repeat(" ", b.Px)
	longestLine, lines2 := longestLine(lines)

	// Get padding on one side
	paddingCount := b.Px

	n := longestLine + (paddingCount * 2) + 2

	if b.TitlePos != "Inside" && runewidth.StringWidth(title) > n-2 {
		panic("Title must be shorter than the Top & Bottom Bars")
	}

	// Create Top and Bottom Bars
	Bar := strings.Repeat(b.Horizontal, n-2)
	TopBar := b.TopLeft + Bar + b.TopRight
	BottomBar := b.BottomLeft + Bar + b.BottomRight

	if b.TitlePos != "Inside" {
		TitleBar := repeatWithString(b.Horizontal, n-2, title)
		if b.TitlePos == "Top" {
			TopBar = b.TopLeft + TitleBar + b.TopRight
		} else if b.TitlePos == "Bottom" {
			BottomBar = b.BottomLeft + TitleBar + b.BottomRight
		} else {
			// Duplicate warning done here if the String() Method is used
			// instead of using Print() and Println() methods
			errorMsg("[warning]: invalid value provided for TitlePos, using default")
			// Using goto here to inorder to exit the current if branch
			goto inside
		}
	}
inside:
	if b.Color != nil {
		if str, ok := b.Color.(string); ok {
			if strings.HasPrefix(str, "Hi") {
				if _, ok := fgHiColors[str]; ok {
					Style := fgHiColors[str].Sprint
					TopBar = Style(TopBar)
					BottomBar = Style(BottomBar)
				}
			} else if _, ok := fgColors[str]; ok {
				Style := fgColors[str].Sprint
				TopBar = Style(TopBar)
				BottomBar = Style(BottomBar)
			} else {
				errorMsg("[warning]: invalid value provided to Color, using default")
			}
		} else if hex, ok := b.Color.(uint); ok {
			hexArray := [3]uint{hex >> 16, hex >> 8 & 0xff, hex & 0xff}
			col := color.RGB(uint8(hexArray[0]), uint8(hexArray[1]), uint8(hexArray[2]))
			if runtime.GOOS != "windows" {
				if detectTerminalColor() == terminfo.ColorLevelHundreds {
					TopBar = col.C256().Sprint(TopBar)
					BottomBar = col.C256().Sprint(BottomBar)
					// TOOD: make a rounding off logic for 24 bit to 8 bit
				} else if detectTerminalColor() == terminfo.ColorLevelBasic {
					TopBar = col.Sprint(TopBar)
					BottomBar = col.Sprint(BottomBar)
				} else if detectTerminalColor() == terminfo.ColorLevelMillions {
					TopBar = col.Sprint(TopBar)
					BottomBar = col.Sprint(BottomBar)
				} else if detectTerminalColor() == terminfo.ColorLevelNone || noColor {
					fmt.Fprintln(os.Stderr, "[warning]: terminal does not support colors, using no effect")
				}
			} else {
				if !noColor {
					TopBar = col.Sprint(TopBar)
					BottomBar = col.Sprint(BottomBar)
				} else {
					fmt.Fprintln(os.Stderr, "[warning]: color not supported for the terminal, using no effect")
				}
			}
		} else if rgb, ok := b.Color.([3]uint); ok {
			col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))
			if runtime.GOOS != "windows" {
				if detectTerminalColor() == terminfo.ColorLevelHundreds {
					TopBar = col.C256().Sprint(TopBar)
					BottomBar = col.C256().Sprint(BottomBar)
					// TOOD: make a rounding off logic for 24 bit to 8 bit
				} else if detectTerminalColor() == terminfo.ColorLevelBasic {
					TopBar = col.Sprint(TopBar)
					BottomBar = col.Sprint(BottomBar)
				} else if detectTerminalColor() == terminfo.ColorLevelMillions {
					TopBar = col.Sprint(TopBar)
					BottomBar = col.Sprint(BottomBar)
				} else {
					fmt.Fprintln(os.Stderr, "[warning]: terminal does not support colors, using no effect")
				}
			} else {
				if !noColor {
					TopBar = col.Sprint(TopBar)
					BottomBar = col.Sprint(BottomBar)
				} else {
					fmt.Fprintln(os.Stderr, "[warning]: terminal does not support colors, using no effect")
				}
			}
		} else {
			fmt.Fprintln(os.Stderr, fmt.Sprintf("expected string, [3]uint or uint not %T using default", b.Color))
		}
	}
	if b.TitlePos == "Inside" && runewidth.StringWidth(TopBar) != runewidth.StringWidth(BottomBar) {
		panic("cannot create a Box with different sizes of Top and Bottom Bars")
	}

	// Create lines to print
	var texts []string
	texts = b.addVertPadding(n)

	for i, line := range lines2 {
		length := line.len

		// Use later
		var space, oddSpace string

		// If current text is shorter than the longest one
		// center the text, so it looks better
		if length < longestLine {
			// Difference between longest and current one
			diff := longestLine - length

			// the spaces to add on each side
			toAdd := diff / 2
			space = strings.Repeat(" ", toAdd)

			// If the difference between the longest and current one
			// is odd, we have to add one additional space before the last vertical separator
			if diff%2 != 0 {
				oddSpace = " "
			}
		}

		spacing := space + sideMargin
		var format string

		if i < titleLen && title != "" && b.TitlePos == "Inside" {
			format = centerAlign
		} else {
			format = b.findAlign()
		}

		// Obtain color
		sep := b.obtainColor()

		formatted := fmt.Sprintf(format, sep, spacing, line.line, oddSpace, space, sideMargin)
		texts = append(texts, formatted)
	}
	vertpadding := b.addVertPadding(n)
	texts = append(texts, vertpadding...)

	// Using strings.Builder is more efficient and faster
	// than concatenating 6 times
	var sb strings.Builder

	sb.WriteString(TopBar)
	sb.WriteString(n1)
	sb.WriteString(strings.Join(texts, n1))
	sb.WriteString(n1)
	sb.WriteString(BottomBar)
	sb.WriteString(n1)

	return sb.String()
}

// obtainColor obtains the Color from string, uint and [3]uint respectively
func (b Box) obtainColor() string {
	if b.Color == nil { // if nil then just return the string
		return b.Vertical
	}
	if str, ok := b.Color.(string); ok {
		if strings.HasPrefix(str, "Hi") {
			if _, ok := fgHiColors[str]; ok {
				return fgHiColors[str].Sprintf(b.Vertical)
			}
		} else if _, ok := fgColors[str]; ok {
			return fgColors[str].Sprintf(b.Vertical)
		}
		errorMsg("[warning]: invalid value provided to Color, using default")
		return b.Vertical
	} else if hex, ok := b.Color.(uint); ok {
		hexArray := [3]uint{hex >> 16, hex >> 8 & 0xff, hex & 0xff}
		col := color.RGB(uint8(hexArray[0]), uint8(hexArray[1]), uint8(hexArray[2]))
		b.roundOffColor(col)
	} else if rgb, ok := b.Color.([3]uint); ok {
		col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))
		return b.roundOffColor(col)
	}
	panic(fmt.Sprintf("expected string, [3]uint or uint not %T", b.Color))
}

// Print prints the Box
func (b Box) Print(title, lines string) {
	var lines2 []string

	// Default Position is Inside, if invalid position is given then just raise a warning
	// then use Default Position which is Inside
	if b.TitlePos == "" {
		b.TitlePos = "Inside"
	} else if b.TitlePos != "Inside" && b.TitlePos != "Bottom" && b.TitlePos != "Top" {
		errorMsg("[warning]: invalid value provided for TitlePos, using default")
		b.TitlePos = "Inside"
	}
	// if Title is empty then TitlePos should be Inside
	if title != "" {
		if b.TitlePos != "Inside" && strings.Contains(title, "\n") {
			panic("Multilines are only supported inside only")
		}
		if b.TitlePos == "Inside" {
			lines2 = append(lines2, strings.Split(title, n1)...)
			lines2 = append(lines2, []string{""}...) // for empty line between title and content
		}
	}
	lines2 = append(lines2, strings.Split(lines, n1)...)
	fmt.Print(b.toString(title, lines2))
}

// Println adds a newline before and after the Box
func (b Box) Println(title, lines string) {
	var lines2 []string

	// Default Position is Inside, if invalid position is given then just raise a warning
	// then use Default Position which is Inside
	if b.TitlePos == "" {
		b.TitlePos = "Inside"
	} else if b.TitlePos != "Inside" && b.TitlePos != "Bottom" && b.TitlePos != "Top" {
		errorMsg("[warning]: invalid value provided for TitlePos, using default")
		b.TitlePos = "Inside"
	}
	// if Title is empty then TitlePos should be Inside
	if title != "" {
		if b.TitlePos != "Inside" && strings.Contains(title, "\n") {
			panic("Multilines are only supported inside only")
		}
		if b.TitlePos == "Inside" {
			lines2 = append(lines2, strings.Split(title, n1)...)
			lines2 = append(lines2, []string{""}...) // for empty line between title and content
		}
	}
	lines2 = append(lines2, strings.Split(lines, n1)...)
	fmt.Printf("\n%s\n", b.toString(title, lines2))
}
