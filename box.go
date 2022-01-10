package box

import (
	"fmt"
	"strings"

	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

const (
	n1     = "\n"
	inside = "Inside"

	// 1 = separator, 2 = spacing, 3 = line; 4 = oddSpace; 5 = space; 6 = sideMargin
	centerAlign = "%[1]s%[2]s%[3]s%[4]s%[2]s%[1]s"
	leftAlign   = "%[1]s%[6]s%[3]s%[4]s%[2]s%[5]s%[1]s"
	rightAlign  = "%[1]s%[2]s%[4]s%[5]s%[3]s%[6]s%[1]s"
)

// Box struct defines the Box to be made
type Box struct {
	TopRight    string // Symbols used for TopRight Corner
	TopLeft     string // Symbols used for TopLeft Corner
	Vertical    string // Symbols used for Vertical Bars
	BottomRight string // Symbols used for BottomRight Corner
	BottomLeft  string // Symbols used for BottomRight Corner
	Horizontal  string // Symbols used for Horizontal Bars
	Config             // Config for the Box struct
}

// Config is the configuration needed for the Box to be designed
type Config struct {
	Py           int         // Horizontal Padding
	Px           int         // Vertical Padding
	ContentAlign string      // Content Alignment inside Box
	Type         string      // Type of Box
	TitlePos     string      // Title Position
	TitleColor   interface{} // Color of Title
	ContentColor interface{} // Color of Content
	Color        interface{} // Color of Box
}

// New takes struct Config and returns the specified Box struct
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

	// Obtain Title and Content color
	title = b.obtainTitleColor(title)
	lines = b.obtainContentColor(lines)

	// Default Position is Inside, no warning for invalid TitlePos as it is done
	// in toString() method
	if b.TitlePos == "" {
		b.TitlePos = inside
	}
	// if Title is empty then TitlePos should be Inside
	if title != "" {
		if b.TitlePos != inside && strings.Contains(title, "\n") {
			panic("Multilines are only supported inside only")
		}
		if b.TitlePos == inside {
			lines2 = append(lines2, strings.Split(title, n1)...)
			lines2 = append(lines2, []string{""}...) // for empty line between title and content
		}
	}
	lines2 = append(lines2, strings.Split(lines, n1)...)
	return b.toString(title, lines2)
}

// toString is same as String except that the main Box generation is done here
func (b Box) toString(title string, lines []string) string {
	titleLen := len(strings.Split(color.ClearCode(title), n1))
	sideMargin := strings.Repeat(" ", b.Px)
	_longestLine, lines2 := longestLine(lines)

	// Get padding on one side
	paddingCount := b.Px

	n := _longestLine + (paddingCount * 2) + 2

	if b.TitlePos != inside && runewidth.StringWidth(title) > n-2 {
		panic("Title must be shorter than the Top & Bottom Bars")
	}

	// Create Top and Bottom Bars
	Bar := strings.Repeat(b.Horizontal, n-2)
	TopBar := b.TopLeft + Bar + b.TopRight
	BottomBar := b.BottomLeft + Bar + b.BottomRight
	TitleBar := repeatWithString(b.Horizontal, n-2, color.ClearCode(title))

	// Check b.TitlePos if it is not Inside
	if b.TitlePos != inside {
		titleLongLineLen, _ := longestLine(strings.Split(TitleBar, n1))
		switch b.TitlePos {
		case "Top":
			TopBar = b.TopLeft + TitleBar + b.TopRight
			// Check if title has tab lines then change BottomBar width accordingly
			if strings.Contains(title, "\t") {
				BottomBar = b.BottomLeft + strings.Repeat(b.Horizontal, titleLongLineLen-1) + b.BottomRight
				// Check if b.TitleColor isn't nil as TopBar and BottomBar won't be equal so they will need to changed
				if b.TitleColor != nil {
					println(strings.Contains(lines[0], "\t"))
					TopBar = b.TopLeft + repeatWithString(b.Horizontal, n+9, color.ClearCode(title)) + b.TopRight
					/*if strings.Contains(lines[0], "\t") {
						//sub := 7 - (3 * b.Px) + 3
						println(titleLongLineLen, len(title), len(repeatWithString(b.Horizontal, titleLongLineLen-14, color.ClearCode(title))))
						TopBar = b.TopLeft + repeatWithString(b.Horizontal, titleLongLineLen-14, color.ClearCode(title)) + b.TopRight
						println(len(TopBar))
					}*/
					// color.ClearCode is used here so that ANSI Color Code also don't get repeated with title
					BottomBar = b.BottomLeft + strings.Repeat(b.Horizontal, titleLongLineLen+10) + b.BottomRight
					println(len(BottomBar))
				}
				// Check if b.TitleColor isn't nil
			} else if b.TitleColor != nil {
				TopBar = b.TopLeft + repeatWithString(b.Horizontal, n+10, color.ClearCode(title)) + b.TopRight
				BottomBar = b.BottomLeft + strings.Repeat(b.Horizontal, n+10) + b.BottomRight
			}
		case "Bottom":
			BottomBar = b.BottomLeft + TitleBar + b.BottomRight
			// Check if title has tab lines then change TopBar width accordingly
			if strings.Contains(title, "\t") {
				TopBar = b.TopLeft + strings.Repeat(b.Horizontal, titleLongLineLen-1) + b.TopRight
				// Check if b.TitleColor isn't nil as TopBar and BottomBar won't be equal so they will need to changed
				if b.TitleColor != nil {
					// color.ClearCode is used here so that ANSI Color Code also don't get repeated with title
					BottomBar = b.BottomLeft + repeatWithString(b.Horizontal, n+9, color.ClearCode(title)) + b.BottomRight
					TopBar = b.TopLeft + strings.Repeat(b.Horizontal, titleLongLineLen+10) + b.TopRight
				}
				// Check if b.TitleColor isn't nil
			} else if b.TitleColor != nil {
				BottomBar = b.BottomLeft + repeatWithString(b.Horizontal, n+10, color.ClearCode(title)) + b.BottomRight
				TopBar = b.TopLeft + strings.Repeat(b.Horizontal, n+10) + b.TopRight
			}
		default:
			// Duplicate warning done here if the String() method is used
			// instead of using Print() and Println() methods
			errorMsg("[warning]: invalid value provided for TitlePos, using default")
			// Using goto here to inorder to exit the current if branch
			goto inside
		}
	}
inside:
	// Check type of b.Color then assign the Colors to TopBar and BottomBar accordingly
	TopBar, BottomBar = b.checkColorType(TopBar, BottomBar, title)
	if b.TitlePos == inside && runewidth.StringWidth(TopBar) != runewidth.StringWidth(BottomBar) {
		panic("cannot create a Box with different sizes of Top and Bottom Bars")
	}

	// Create lines to print
	var texts, vertpadding []string

	// Check if b.TitlePos isn't Inside and title has tab lines it
	// so that vertical padding would be need to be updated as per the need
	if b.TitlePos != "Inside" && strings.Contains(title, "\t") {
		titleLongLineLen, _ := longestLine(strings.Split(TitleBar, n1))
		texts = b.addVertPadding(titleLongLineLen + 1)
		/*if strings.Contains(lines[0], "\t") && b.ContentColor == nil {
			println("here 4")
			texts = b.formatLine(lines2, titleLongLineLen-5, titleLen, sideMargin, color.ClearCode(title), texts)
		}*/
		// Check if b.TitleColor is not nil so that the
		// vertical padding to be needed to update accordingly
		if b.TitleColor != nil {
			texts = b.addVertPadding(titleLongLineLen + 12)
			// Check if Content has tabbed lines so that vertical padding will be need to be updated accordingly
			sub := 7 - (3 * b.Px) + 3
			texts = b.formatLine(lines2, titleLongLineLen+b.Px+sub, titleLen, sideMargin, color.ClearCode(title), texts)
			/*
				if strings.Contains(lines[0], "\t") {
					sub := 7 - (3 * b.Px) + 3
					println("here 3")
					fmt.Println(TopBar)
					texts = b.formatLine(lines2, titleLongLineLen+b.Px+sub, titleLen, sideMargin, color.ClearCode(title), texts)
				} else {
					if len(lines) > 1 {
						println("here 1")

						// Create number of spaces needed for the vertical padding
						sub := 7 - (3 * b.Px) + 3
						//println(sub)
						texts = b.formatLine(lines2, titleLongLineLen+b.Px+sub, titleLen, sideMargin, color.ClearCode(title), texts)
					} else {
						println("here 22", len(lines), titleLongLineLen, _longestLine)
						sub := 7 - (3 * b.Px) + 3
						texts = b.formatLine(lines2, titleLongLineLen+b.Px+sub, titleLen, sideMargin, color.ClearCode(title), texts)
					}
				}*/
			vertpadding = b.addVertPadding(titleLongLineLen + 12)
			texts = append(texts, vertpadding...)
		} else {
			sub := -2 - (1 * b.Px) + 1
			texts = b.formatLine(lines2, titleLongLineLen-b.Px+sub, titleLen, sideMargin, color.ClearCode(title), texts)
			/*println("here 5")
			if strings.Contains(lines[0], "\t") && b.ContentColor != nil {
				println("here 6")
				sub := -2 - (1 * b.Px) + 1
				texts = b.formatLine(lines2, titleLongLineLen-b.Px+sub, titleLen, sideMargin, color.ClearCode(title), texts)
			} else if b.ContentColor == nil {
				sub := -2 - (1 * b.Px) + 1
				//println(sub)
				texts = b.formatLine(lines2, titleLongLineLen-b.Px+sub, titleLen, sideMargin, color.ClearCode(title), texts)
			}*/
			vertpadding = b.addVertPadding(titleLongLineLen + 1)
			texts = append(texts, vertpadding...)
		}

	} else if b.TitleColor != nil && b.TitlePos != "Inside" && !strings.Contains(title, "\t") {
		titleLongLineLen, _ := longestLine(strings.Split(TitleBar, n1))
		texts = b.addVertPadding(titleLongLineLen + 14)
		texts = b.formatLine(lines2, _longestLine+12, titleLen, sideMargin, color.ClearCode(title), texts)
		vertpadding = b.addVertPadding(titleLongLineLen + 14)
		texts = append(texts, vertpadding...)
	} else {
		texts = b.addVertPadding(n)
		// Check if Content has tabbed lines so that vertical padding will be need to be updated accordingly
		texts = b.formatLine(lines2, _longestLine, titleLen, sideMargin, title, texts)
		vertpadding := b.addVertPadding(n)
		texts = append(texts, vertpadding...)
	}

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

// obtainTitleColor obtains the TitleColor from string, uint and [3]uint respectively
func (b Box) obtainTitleColor(title string) string {
	if b.TitleColor == nil { // if nil then just return the string
		return title
	}
	// Check if type of b.TitleColor is string
	if str, ok := b.TitleColor.(string); ok {
		// Hi Intensity Color
		if strings.HasPrefix(str, "Hi") {
			if _, ok := fgHiColors[str]; ok {
				// If title has newlines in it then spliting would be needed
				// as color won't be applied on all
				if strings.Contains(title, "\n") {
					return b.applyColorToAll(title, str, color.RGBColor{}, false)
				}
				return fgHiColors[str].Sprintf(title)
			}
		} else if _, ok := fgColors[str]; ok {
			// If title has newlines in it then spliting would be needed
			// as color won't be applied on all
			if strings.Contains(title, "\n") {
				return b.applyColorToAll(title, str, color.RGBColor{}, false)
			}
			return fgColors[str].Sprintf(title)
		}
		// Return a warning as TitleColor provided as a string is unknown and
		// return without the color effect
		errorMsg("[warning]: invalid value provided to Color, using default")
		return title

		// Check if type of b.TitleColor is uint
	} else if hex, ok := b.TitleColor.(uint); ok {
		// Break down the hex into R, G and B respectively
		hexArray := [3]uint{hex >> 16, hex >> 8 & 0xff, hex & 0xff}
		col := color.RGB(uint8(hexArray[0]), uint8(hexArray[1]), uint8(hexArray[2]))

		// If title has newlines in it then spliting would be needed
		// as color won't be applied on all
		if strings.Contains(title, "\n") {
			return b.applyColorToAll(title, "", col, true)
		}
		return b.roundOffTitleColor(col, title)

		// Check if type of b.TitleColor is [3]uint
	} else if rgb, ok := b.TitleColor.([3]uint); ok {
		col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))

		// If title has newlines in it then spliting would be needed
		// as color won't be applied on all
		if strings.Contains(title, "\n") {
			return b.applyColorToAll(title, "", col, true)
		}
		return b.roundOffTitleColor(col, title)
	}
	// Panic if b.TitleColor is an unexpected type
	panic(fmt.Sprintf("expected string, [3]uint or uint not %T", b.TitleColor))
}

// obtainContentColor obtains the ContentColor from string, uint and [3]uint respectively
func (b Box) obtainContentColor(content string) string {
	if b.ContentColor == nil { // if nil then just return the string
		return content
	}
	// Check if type of b.ContentColor is string
	if str, ok := b.ContentColor.(string); ok {
		// Hi Intensity Color
		if strings.HasPrefix(str, "Hi") {
			if _, ok := fgHiColors[str]; ok {
				// If content has newlines in it then spliting would be needed
				// as color won't be applied on all
				if strings.Contains(content, "\n") {
					return b.applyColorToAll(content, str, color.RGBColor{}, false)
				}
				return fgHiColors[str].Sprintf(content)
			}
		} else if _, ok := fgColors[str]; ok {
			// If content has newlines in it then spliting would be needed
			// as color won't be applied on all
			if strings.Contains(content, "\n") {
				return b.applyColorToAll(content, str, color.RGBColor{}, false)
			}
			return fgColors[str].Sprintf(content)
		}
		// Return a warning as ContentColor provided as a string is unknown and
		// return without the color effect
		errorMsg("[warning]: invalid value provided to Color, using default")
		return content

		// Check if type of b.ContentColor is uint
	} else if hex, ok := b.ContentColor.(uint); ok {
		// Break down the hex into R, G and B respectively
		hexArray := [3]uint{hex >> 16, hex >> 8 & 0xff, hex & 0xff}
		col := color.RGB(uint8(hexArray[0]), uint8(hexArray[1]), uint8(hexArray[2]))

		// If content has newlines in it then spliting would be needed
		// as color won't be applied on all
		if strings.Contains(content, "\n") {
			return b.applyColorToAll(content, "", col, true)
		}
		return b.roundOffTitleColor(col, content)

		// Check if type of b.ContentColor is [3]uint
	} else if rgb, ok := b.ContentColor.([3]uint); ok {
		col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))

		// If content has newlines in it then spliting would be needed
		// as color won't be applied on all
		if strings.Contains(content, "\n") {
			return b.applyColorToAll(content, "", col, true)
		}
		return b.roundOffTitleColor(col, content)
	}
	// Panic if b.ContentColor is an unexpected type
	panic(fmt.Sprintf("expected string, [3]uint or uint not %T", b.ContentColor))
}

// obtainColor obtains the BoxColor from string, uint and [3]uint respectively
func (b Box) obtainBoxColor() string {
	if b.Color == nil { // if nil then just return the string
		return b.Vertical
	}
	// Check if type of b.Color is string
	if str, ok := b.Color.(string); ok {
		// Hi Intensity Color
		if strings.HasPrefix(str, "Hi") {
			if _, ok := fgHiColors[str]; ok {
				return fgHiColors[str].Sprintf(b.Vertical)
			}
		} else if _, ok := fgColors[str]; ok {
			return fgColors[str].Sprintf(b.Vertical)
		}
		errorMsg("[warning]: invalid value provided to Color, using default")
		// Return a warning as Color provided as a string is unknown and
		// return without the color effect
		return b.Vertical
		// Check if type of b.Color is uint
	} else if hex, ok := b.Color.(uint); ok {
		// Break down the hex into r, g and b respectively
		hexArray := [3]uint{hex >> 16, hex >> 8 & 0xff, hex & 0xff}
		col := color.RGB(uint8(hexArray[0]), uint8(hexArray[1]), uint8(hexArray[2]))
		return b.roundOffColorVertical(col)
		// Check if type of b.Color is [3]uint
	} else if rgb, ok := b.Color.([3]uint); ok {
		col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))
		return b.roundOffColorVertical(col)
	}
	// Panic if b.Color is an unexpected type
	panic(fmt.Sprintf("expected string, [3]uint or uint not %T", b.Color))
}

// Print prints the Box
func (b Box) Print(title, lines string) {
	var lines2 []string

	// Obtain Title and Content color
	title = b.obtainTitleColor(title)
	lines = b.obtainContentColor(lines)

	// Default Position is Inside, if invalid position is given then just raise a warning
	// then use Default Position which is Inside
	if b.TitlePos == "" {
		b.TitlePos = inside
	} else if b.TitlePos != inside && b.TitlePos != "Bottom" && b.TitlePos != "Top" {
		errorMsg("[warning]: invalid value provided for TitlePos, using default")
		b.TitlePos = inside
	}
	// if Title is empty then TitlePos should be Inside
	if title != "" {
		if b.TitlePos != inside && strings.Contains(title, "\n") {
			panic("Multilines are only supported inside only")
		}
		if b.TitlePos == inside {
			lines2 = append(lines2, strings.Split(title, n1)...)
			lines2 = append(lines2, []string{""}...) // for empty line between title and content
		}
	}
	lines2 = append(lines2, strings.Split(lines, n1)...)
	color.Print(b.toString(title, lines2))
}

// Println adds a newline before and after the Box
func (b Box) Println(title, lines string) {
	var lines2 []string

	// Obtain Title and Content color
	title = b.obtainTitleColor(title)
	lines = b.obtainContentColor(lines)

	// Default Position is Inside, if invalid position is given then just raise a warning
	// then use Default Position which is Inside
	if b.TitlePos == "" {
		b.TitlePos = inside
	} else if b.TitlePos != inside && b.TitlePos != "Bottom" && b.TitlePos != "Top" {
		errorMsg("[warning]: invalid value provided for TitlePos, using default")
		b.TitlePos = inside
	}
	// if Title is empty then TitlePos should be Inside
	if title != "" {
		if b.TitlePos != inside && strings.Contains(title, "\n") {
			panic("Multilines are only supported inside only")
		}
		if b.TitlePos == inside {
			lines2 = append(lines2, strings.Split(title, n1)...)
			lines2 = append(lines2, []string{""}...) // for empty line between title and content
		}
	}
	lines2 = append(lines2, strings.Split(lines, n1)...)
	color.Printf("\n%s\n", b.toString(title, lines2))
}
