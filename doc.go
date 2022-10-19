// Package box a.k.a Box CLI Maker is a Highly Customized Terminal Box Creator written in Go.
//
// It provides many styles and options to make Boxes. There are 8 inbuilt styles and has Color support via RGB uint, RGB Array of [3]uint and string (Given).
//
// Inbuilt Box Styles:
// Single,
// Double,
// Bold,
// Single Double,
// Double Single,
// Round,
// Hidden and
// Classic
//
// Inbuilt Colors:
// Black,
// Blue,
// Red,
// Yellow,
// Green,
// Cyan,
// Magenta,
// White,
// HiBlack,
// HiBlue,
// HiRed,
// HiYellow,
// HiGreen,
// HiCyan,
// HiMagenta and
// HiWhite
//
// It also has Unicode and Emoji support which works across almost all terminals. Unlike other CLI Makers, Box CLI Maker also supports tab, multi-line strings and string wrapping.
//
// Disclaimer: As different terminals have different fonts by default so the right vertical alignment may not be aligned well.
// You will have to change your font accordingly to make it work.
//
// Basic Example:
//
//	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
//	Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")
//
// You can specify and change the options by changing the below Config struct.
//
//	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", TitlePos: "Top", ContentAlign: "Left"})
//
// TitlePos can be changed to Inside, Top, Bottom and ContentAlign to be Left, Right and Center.
// By default TitlePos is Inside, ContentAlign is Left and Style is Single.
//
// Manual string wrapping is also allowed via a flag Config.AllowWrapping, by the default padding, is
// 2*TermWidth/3.
//
// String() method can be for the string representation of the Box.
//
//	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
//	boxStr := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")
//
// True Color is also supported and it can be by providing an array of [3]uint or uint.
//
// There might be some terminals not supporting True Color so in this case, it will detect the terminal's max color capacity
// then will round off True Color to either 4-bit or 8-bit respectively.
//
// Title and Content can also be colored by passing the colors needed to the fields box.TitleColor and box.ContentColor respectively.
//
// This module also enables True Color and 256 Colors support on Windows Console but you need have at least Windows 10 Version 1511
// for 256 colors or Windows 10 Version 1607 for True Color Support.
//
// Example of True Color via uint:
//
//	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: uint(0x34562f)})
//	Box.Println("Box CLI Maker", "Highly Customized Terminal Box Maker")
//
// Note: uint must be in a range of [0x000000, 0xFFFFFF].
//
// Example of True Color via [3]uint:
//
//	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: [3]uint{23, 56, 78}})
//	Box.Println("Box CLI Maker", "Highly Customized Terminal Box Maker")
//
// Note: [3]uint array elements must be in a range of [0x0, 0xFF].
//
// Custom Box Style can also be by using box.Box:
//
//	config := box.Config{Px: 2, Py: 3, Type: "", TitlePos: "Inside"}
//	boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
//
// More info and examples can be found in README.md and examples/ folder
package box
