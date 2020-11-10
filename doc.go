/*
Package Box CLI Maker is a Highly Customized Terminal Box Creator written in Go.

It provides many styles and options to make Boxes. There are 8 inbuilt styles and has Color support via RGB uint, RGB Array of [3]uint and string (given).

Inbuilt Box Styles:
Single,
Double,
Bold,
Single Double,
Double Single,
Round,
Hidden and
Classic

Inbuilt Colors:
Black,
Blue,
Red,
Yellow,
Green,
Cyan,
Magenta,
HiBlack,
HiBlue,
HiRed,
HiYellow,
HiGreen,
HiCyan and
HiMagenta

It also has Unicode and Emoji support which works across all terminals though there might be some terminals which do not support
Unicode and Emoji like Windows CMD and Powershell. Unlike other CLI Makers, Box CLI Maker supports tab and multi line string.

Note: As different terminals have different font by default so the right vertical alignment may not be aligned well.
You will have to change your font accordingly to make it work.

Basic Example:

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")

You can specify and change the options by changing the above Config struct.

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", TitlePos: "Top", ContentAlign: "Left"})

You can also customize and change the TitlePos to Inside, Top, Bottom and ContentAlign to Left, Right and Center.
By default TitlePos is Inside, ContentAlign is Left and Style is Single.

You can also use the String() method for the string representation of the Box.

If you want the Box to be printed correctly irrespective of if it will form the correct Color or not on Windows (as Windows Console only supports 16 Colors) then you will have to add Box.Output
as the passing stream to stream based functions:

fmt.Fprintf(box.Output, boxStr)

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
boxStr := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")

The Custom Color option will not be applicable for terminals which don't have 24 bit support i.e. True Color ANSI Code.
If it is used then the Color effect will not be there.

RGB Uint Example:

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: uint(0x34562f)})
Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")

Note: Uint must be in a range of [0x000000, 0xFFFFFF] else it will panic.

RGB [3]uint Example:

Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: [3]uint{23, 56, 78}})
Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")

Note: [3]uint array elements must be in a range of [0x0, 0xFF] else it will panic.

You can even make your custom Box Style by using struct box.Box:

config := box.Config{Px: 2, Py: 3, Type: "", TitlePos: "Inside"}
boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
*/
package box
