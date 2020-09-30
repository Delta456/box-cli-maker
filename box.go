package box

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/mattn/go-runewidth"
)

const (
	n1 = "\n"
	// sep = separator, sp = spacing, ln = line; os = oddSpace; s = space
	centerAlign = "{sep}{sp}{ln}{os}{sp}{sep}"
	leftAlign   = "{sep}{px}{ln}{os}{sp}{s}{sep}"
	rightAlign  = "{sep}{sp}{os}{s}{ln}{px}{sep}"
)

// Box struct defines the Box to be made.
type Box struct {
	TopRight    string // TopRight corner used for Symbols
	TopLeft     string // TopLeft corner used for Symbols
	Vertical    string // Symbols used for Vertical Bars
	BottomRight string // BottomRight corner used for Symbols
	BottomLeft  string // BotromLeft corner used for Symbols
	Horizontal  string // Symbols used for Horizontal Bars
	Con         Config // Config for the Box struct
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
	if _, ok := boxs[config.Type]; ok {
		BoxNew := boxs[config.Type]
		BoxNew.Con = config
		return BoxNew
	}
	panic("Invalid Box Type provided")

}

// String returns the string representation of Box.
func (b Box) String(title, lines string) string {
	var lines2 []string

	// Default Position is Inside
	if b.Con.TitlePos == "" {
		b.Con.TitlePos = "Inside"
	}
	// if Title is empty then TitlePos should be Inside
	if title != "" {
		if b.Con.TitlePos != "Inside" && strings.Contains(title, "\n") {
			panic("Multilines are only supported inside only")
		}
		if b.Con.TitlePos == "Inside" {
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
	sideMargin := strings.Repeat(" ", b.Con.Px)
	longestLine := longestLine(lines)

	// get padding on one side
	paddingCount := b.Con.Px

	n := longestLine + (paddingCount * 2) + 2

	if b.Con.TitlePos != "Inside" && runewidth.StringWidth(title) > n-2 {
		panic("Title must be lower in length than the Top & Bottom Bars")
	}

	// create Top and Bottom Bars
	Bar := strings.Repeat(b.Horizontal, n-2)
	TopBar := b.TopLeft + Bar + b.TopRight
	BottomBar := b.BottomLeft + Bar + b.BottomRight

	if b.Con.TitlePos != "Inside" {
		TitleBar := repeatWithString(b.Horizontal, n-2, title)
		if b.Con.TitlePos == "Top" {
			TopBar = b.TopLeft + TitleBar + b.TopRight
		} else if b.Con.TitlePos == "Bottom" {
			BottomBar = b.BottomLeft + TitleBar + b.BottomRight
		} else {
			fmt.Fprintln(os.Stderr, color.RedString("[error]: invalid value provided for TitlePos, using default"))
		}
	}
	if str, ok := b.Con.Color.(string); ok {
		if b.Con.Color != "" {
			if strings.HasPrefix(str, "Hi") {
				if _, ok := fgHiColors[str]; ok {
					Style := color.New(fgHiColors[str]).SprintfFunc()
					TopBar = Style(TopBar)
					BottomBar = Style(BottomBar)
				}
			} else if _, ok := fgColors[str]; ok {
				Style := color.New(fgColors[str]).SprintfFunc()
				TopBar = Style(TopBar)
				BottomBar = Style(BottomBar)
			} else {
				fmt.Fprintln(os.Stderr, color.RedString("[error]: invalid value provided to Color, using default"))
			}
		}
	} else if hex, ok := b.Con.Color.(int); ok {
		TopBar = rbg_hex(hex, TopBar)
		BottomBar = rbg_hex(hex, BottomBar)
	} else if rgb, ok := b.Con.Color.(Rgb); ok {
		TopBar = rbg_struct(rgb, TopBar)
		BottomBar = rbg_struct(rgb, BottomBar)
	} else {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("Expected string/Rgb/int not %T using default", b.Con.Color))
		b.Con.Color = ""
	}

	if b.Con.TitlePos == "Inside" && runewidth.StringWidth(TopBar) != runewidth.StringWidth(BottomBar) {
		panic("cannot create a Box with different sizes of Top and Bottom Bars")
	}

	// create lines to print
	var texts []string
	texts = b.addVertPadding(n)

	for i, line := range lines {
		length := runewidth.StringWidth(line)

		// use later
		var space, oddSpace string

		// if current text is shorter than the longest one
		// center the text, so it looks better
		if length < longestLine {
			// difference between longest and current one
			diff := longestLine - length

			// the spaces to add on each side
			toAdd := diff / 2
			space = strings.Repeat(" ", toAdd)

			// if the difference between the longest and current one
			// is odd, we have to add one additional space before the last vertical separator
			if diff%2 != 0 {
				oddSpace = " "
			}
		}

		spacing := space + sideMargin
		format := b.findAlign()

		if i < titleLen && title != "" {
			format = centerAlign
		}
		// obtain color
		sep := b.obtainColor()

		// TODO: find a better way
		formatted := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(format, "{sep}", sep), "{sp}", spacing), "{ln}", line), "{os}", oddSpace), "{s}", space), "{px}", sideMargin)
		texts = append(texts, formatted)
	}
	vertpadding := b.addVertPadding(n)
	texts = append(texts, vertpadding...)

	return TopBar + n1 + strings.Join(texts, n1) + n1 + BottomBar + n1
}

func (b Box) obtainColor() string {
	if str, ok := b.Con.Color.(string); ok {
		if b.Con.Color != "" {
			if strings.HasPrefix(str, "Hi") {
				if _, ok := fgHiColors[str]; ok {
					Style := color.New(fgHiColors[str]).SprintfFunc()
					return Style(b.Vertical)
				}
			} else if _, ok := fgColors[str]; ok {
				Style := color.New(fgColors[str]).SprintfFunc()
				return Style(b.Vertical)
			}
			fmt.Fprintln(os.Stderr, color.RedString("[error]: invalid value provided to Color, using default"))
			return b.Vertical
		}
	} else if hex, ok := b.Con.Color.(int); ok {
		return rbg_hex(hex, b.Vertical)
	} else if rgb, ok := b.Con.Color.(Rgb); ok {
		return rbg_struct(rgb, b.Vertical)
	}
	panic(fmt.Sprintf("Expected string/Rgb/int not %T", b.Con.Color))
}

// Print prints the box
func (b Box) Print(title, lines string) {
	var lines2 []string

	// Default Position is Inside
	if b.Con.TitlePos == "" {
		b.Con.TitlePos = "Inside"
	}
	// if Title is empty then TitlePos should be Inside
	if title != "" {
		if b.Con.TitlePos != "Inside" && strings.Contains(title, "\n") {
			panic("Multilines are only supported inside only")
		}
		if b.Con.TitlePos == "Inside" {
			lines2 = append(lines2, strings.Split(title, n1)...)
			lines2 = append(lines2, []string{""}...) // for empty line between title and content
		}
	}
	lines2 = append(lines2, strings.Split(lines, n1)...)
	fmt.Print(b.toString(title, lines2))
}

// Println adds a newline before and after the box
func (b Box) Println(title, lines string) {
	var lines2 []string

	// Default Position is Inside
	if b.Con.TitlePos == "" {
		b.Con.TitlePos = "Inside"
	}
	// if Title is empty then TitlePos should be Inside
	if title != "" {
		if b.Con.TitlePos != "Inside" && strings.Contains(title, "\n") {
			panic("Multilines are only supported inside only")
		}
		if b.Con.TitlePos == "Inside" {
			lines2 = append(lines2, strings.Split(title, n1)...)
			lines2 = append(lines2, []string{""}...) // for empty line between title and content
		}
	}
	lines2 = append(lines2, strings.Split(lines, n1)...)
	fmt.Printf("\n%s\n", b.toString(title, lines2))
}
