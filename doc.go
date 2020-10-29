/*
Package Box CLI Maker is a Highly Customized Terminal Box Creator written in Go.

It provides many styles and options to make Boxes. There are 8 inbuilt styles and Color support via RGB uint, RGB Array of [3]uint and string (given).

Inbuilt Box Styles:
Single
Double
Bold
Single Double
Double Single
Round
Hidden
Classic

Inbuilt Colors:
Black
Blue
Red
Yellow
Green
Cyan
Magenta
HiBlack
HiBlue
HiRed
HiYellow
HiGreen
HiCyan
HiMagenta

It also has Unicode and Emoji support which works across all terminals though there might be some terminals which do not support
Unicode and Emoji like Windows CMD and Powershell.

Basic Example:

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")

You can specify and change the options by changing the above Config struct.

You can customize and change the TitlePos to Inside, Top, Bottom and ContentAlign to Left, Right and Center.
By default TitlePos is Inside and ContentAlign is Left.

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", TitlePos: "Top", ContentAlign: "Left"})

If you want the string representation of the Box then you can just use String() method of the Box:

If you want the Box to be printed correctly irrespective of it will form the correct color or not on Windows then you will have to add box.Output
as the passing stream to Fprintf(), Fprintln() and etc to the passing stream functions:

if runtime.GOOS == "windows" {
	fmt.Fprintf(box.Output, box_str)
}

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
b := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")
... // use it afterwards

If you do not want the title and lines to be there then you can just leave i.e. put an empty in the both places.const

The following two will not be applicable for terminals which don't have 24 bit support i.e. True Color ANSI Code.
If used then the color effect will not be there.

RBG Hex Example:

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: uint(0x34562f)})
Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")

Note: Uint must be in a range of [0x000000, 0xFFFFFF] else it will panic.


RBG [3]uint Example:

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: [3]uint{23, 56, 78}})
Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")

Note: [3]uint elements must be in a range of [0x0, 0xFF] else it will panic.


You can even make your custom Box Style by using box.Box struct:

config := box.Config{Px: 2, Py: 3, Type: "", TitlePos: "Inside"}
boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}

*/

package box
