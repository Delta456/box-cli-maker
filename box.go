package box

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/huandu/xstrings"
	"github.com/mattn/go-runewidth"
	"github.com/muesli/reflow/wrap"
	"golang.org/x/term"
)

const (
	n1     = "\n"
	inside = "Inside"

	// 1 = separator, 2 = spacing, 3 = line; 4 = oddSpace; 5 = space; 6 = sideMargin
	centerAlign = "%[1]s%[2]s%[3]s%[4]s%[2]s%[1]s"
	leftAlign   = "%[1]s%[6]s%[3]s%[4]s%[2]s%[5]s%[1]s"
	rightAlign  = "%[1]s%[2]s%[4]s%[5]s%[3]s%[6]s%[1]s"
)

// Box defines the design
type Box struct {
	TopRight    string // TopRight Corner Symbols
	TopLeft     string // TopLeft Corner Symbols
	Vertical    string // Vertical Bar Symbols
	BottomRight string // BottomRight Corner Symbols
	BottomLeft  string // BottomLeft Corner Symbols
	Horizontal  string // Horizontal Bars Symbols
	Config             // Box Config
}

// Config is the configuration needed for the Box to be designed
type Config struct {
	Py            int         // Horizontal Padding
	Px            int         // Vertical Padding
	ContentAlign  string      // Content Alignment inside Box
	Type          string      // Box Type
	TitlePos      string      // Title Position
	TitleColor    interface{} // Title Color
	ContentColor  interface{} // Content Color
	Color         interface{} // Box Color
	AllowWrapping bool        // Flag to allow custom Content Wrapping
	WrappingLimit int         // Wrap the Content upto the Limit
}

// New takes Box Config and returns a Box from the given Config
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

// String returns the string representation of Box
func (b Box) String(title, lines string) string {
	var lines2 []string

	// Allow Wrapping according to the user
	if b.AllowWrapping {
		// If limit not provided then use 2*TermWidth/3 as limit else
		// use the one provided
		if b.WrappingLimit != 0 {
			lines = wrap.String(lines, b.WrappingLimit)
		} else {
			width, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				log.Fatal(err)
			}
			lines = wrap.String(lines, 2*width/3)
		}
	}

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

// toString is an internal method and same as String method except that the main Box generation is done here
func (b Box) toString(title string, lines []string) string {
	titleLen := len(strings.Split(color.ClearCode(title), n1))
	sideMargin := strings.Repeat(" ", b.Px)
	_longestLine, lines2 := longestLine(lines)

	// Get padding on one side
	paddingCount := b.Px

	n := _longestLine + (paddingCount * 2) + 2

	if b.TitlePos != inside && runewidth.StringWidth(color.ClearCode(title)) > n-2 {
		panic("Title must be shorter than the Top & Bottom Bars")
	}

	// Create Top and Bottom Bars
	Bar := strings.Repeat(b.Horizontal, n-2)
	TopBar := b.TopLeft + Bar + b.TopRight
	BottomBar := b.BottomLeft + Bar + b.BottomRight

	var TitleBar string
	// If title has tabs then expand them accordingly.
	if strings.Contains(title, "\t") {
		TitleBar = repeatWithString(b.Horizontal, n-2, xstrings.ExpandTabs(title, 4))
	} else {
		TitleBar = repeatWithString(b.Horizontal, n-2, title)
	}

	// Check b.TitlePos if it is not Inside
	if b.TitlePos != inside {
		switch b.TitlePos {
		case "Top":
			TopBar = b.TopLeft + TitleBar + b.TopRight
			//fmt.Println(TopBar)
		case "Bottom":
			BottomBar = b.BottomLeft + TitleBar + b.BottomRight
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
	// If title has tabs then expand them accordingly.
	if strings.Contains(title, "\t") {
		TopBar, BottomBar = b.checkColorType(TopBar, BottomBar, xstrings.ExpandTabs(title, 4))
	} else {
		TopBar, BottomBar = b.checkColorType(TopBar, BottomBar, title)
	}

	if b.TitlePos == inside && runewidth.StringWidth(TopBar) != runewidth.StringWidth(BottomBar) {
		panic("cannot create a Box with different sizes of Top and Bottom Bars")
	}

	// Create lines to print
	texts := b.addVertPadding(n)
	texts = b.formatLine(lines2, _longestLine, titleLen, sideMargin, title, texts)
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

// obtainTitleColor obtains TitleColor from types string, uint and [3]uint respectively
func (b Box) obtainTitleColor(title string) string {
	if b.TitleColor == nil { // if nil then just return the string
		return title
	}
	// Check if type of b.TitleColor is string
	if str, ok := b.TitleColor.(string); ok {
		// Hi Intensity Color
		if strings.HasPrefix(str, "Hi") {
			if _, ok := fgHiColors[str]; ok {
				// If title has newlines in it then splitting would be needed
				// as color won't be applied on all
				if strings.Contains(title, "\n") {
					return b.applyColorToAll(title, str, color.RGBColor{}, false)
				}
				return addStylePreservingOriginalFormat(title, fgHiColors[str].Sprint)
			}
		} else if _, ok := fgColors[str]; ok {
			// If title has newlines in it then splitting would be needed
			// as color won't be applied on all
			if strings.Contains(title, "\n") {
				return b.applyColorToAll(title, str, color.RGBColor{}, false)
			}
			return addStylePreservingOriginalFormat(title, fgColors[str].Sprint)
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

		// If title has newlines in it then splitting would be needed
		// as color won't be applied on all
		if strings.Contains(title, "\n") {
			return b.applyColorToAll(title, "", col, true)
		}
		return addStylePreservingOriginalFormat(title, col.Sprint)

		// Check if type of b.TitleColor is [3]uint
	} else if rgb, ok := b.TitleColor.([3]uint); ok {
		col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))

		// If title has newlines in it then splitting would be needed
		// as color won't be applied on all
		if strings.Contains(title, "\n") {
			return b.applyColorToAll(title, "", col, true)
		}
		return b.roundOffTitleColor(col, addStylePreservingOriginalFormat(title, col.Sprint))
	}
	// Panic if b.TitleColor is an unexpected type
	panic(fmt.Sprintf("expected string, [3]uint or uint not %T", b.TitleColor))
}

// obtainContentColor obtains ContentColor from types string, uint and [3]uint respectively
func (b Box) obtainContentColor(content string) string {
	if b.ContentColor == nil { // if nil then just return the string
		return content
	}
	// Check if type of b.ContentColor is string
	if str, ok := b.ContentColor.(string); ok {
		// Hi Intensity Color
		if strings.HasPrefix(str, "Hi") {
			if _, ok := fgHiColors[str]; ok {
				// If Content has newlines in it then splitting would be needed
				// as color won't be applied on all
				if strings.Contains(content, "\n") {
					return b.applyColorToAll(content, str, color.RGBColor{}, false)
				}
				return addStylePreservingOriginalFormat(content, fgHiColors[str].Sprint)
			}
		} else if _, ok := fgColors[str]; ok {
			// If Content has newlines in it then splitting would be needed
			// as color won't be applied on all
			if strings.Contains(content, "\n") {
				return b.applyColorToAll(content, str, color.RGBColor{}, false)
			}
			return addStylePreservingOriginalFormat(content, fgColors[str].Sprint)
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

		// If content has newlines in it then splitting would be needed
		// as color won't be applied on all
		if strings.Contains(content, "\n") {
			return b.applyColorToAll(content, "", col, true)
		}
		return b.roundOffTitleColor(col, content)

		// Check if type of b.ContentColor is [3]uint
	} else if rgb, ok := b.ContentColor.([3]uint); ok {
		col := color.RGB(uint8(rgb[0]), uint8(rgb[1]), uint8(rgb[2]))

		// If content has newlines in it then splitting would be needed
		// as color won't be applied on all
		if strings.Contains(content, "\n") {
			return b.applyColorToAll(content, "", col, true)
		}
		return b.roundOffTitleColor(col, content)
	}
	// Panic if b.ContentColor is an unexpected type
	panic(fmt.Sprintf("expected string, [3]uint or uint not %T", b.ContentColor))
}

// obtainColor obtains BoxColor from types string, uint and [3]uint respectively
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

	// Allow Wrapping according to the user
	if b.AllowWrapping {
		// If limit not provided then use 2*TermWidth/3 as limit else
		// use the one provided
		if b.WrappingLimit != 0 {
			lines = wrap.String(lines, b.WrappingLimit)
		} else {
			width, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				log.Fatal(err)
			}
			lines = wrap.String(lines, 2*width/3)
		}
	}

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

// Println adds a newline before and after printing the Box
func (b Box) Println(title, lines string) {
	var lines2 []string

	// Allow Wrapping according to the user
	if b.AllowWrapping {
		// If limit not provided then use 2*TermWidth/3 as limit else
		// use the one provided
		if b.WrappingLimit != 0 {
			lines = wrap.String(lines, b.WrappingLimit)
		} else {
			width, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				log.Fatal(err)
			}
			lines = wrap.String(lines, 2*width/3)
		}
	}

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
