/*
Package Box CLI Maker is a Highly Customized Terminal Box Creator written in Go.

It provides many styles and options to make Boxes. There are 8 inbuilt styles and has Color support via RGB uint, RGB Array of [3]uint and string (given).
It can also detect the terminal's maximum color supported accordingly.

True Color is supported for even Windows Console though you need ^v1.2.4 for using it.

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
White,
HiBlack,
HiBlue,
HiRed,
HiYellow,
HiGreen,
HiCyan,
HiMagenta and
HiWhite

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

	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
	boxStr := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")

Some terminals may not support True Color so it will detect the terminal's max color capacity
then will round off True Color to either 4 bit or 8 bit respectively.

RGB Uint Example:

	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: uint(0x34562f)})
	Box.Println("Box CLI Maker", "Highly Customized Terminal Box Maker")

Note: Uint must be in a range of [0x000000, 0xFFFFFF].

RGB [3]uint Example:

	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: [3]uint{23, 56, 78}})
	Box.Println("Box CLI Maker", "Highly Customized Terminal Box Maker")

Note: [3]uint array elements must be in a range of [0x0, 0xFF].

You can even make your custom Box Style by using struct box.Box:

	config := box.Config{Px: 2, Py: 3, Type: "", TitlePos: "Inside"}
	boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
*/
package box
