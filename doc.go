// Package box renders styled boxes around text for terminal applications.
//
// The core type is Box, constructed with NewBox and configured via a fluent API.
// Boxes support multiple built‑in styles, title positions, alignment, wrapping,
// and ANSI/truecolor output.
//
// Basic example:
//
//	b := box.NewBox().
//		Style(box.Single).
//		Padding(2, 1).
//		TitlePosition(box.Top).
//		ContentAlign(box.Center).
//		Color(box.Cyan).
//		TitleColor(box.BrightYellow).
//
//	out, err := b.Render("Box CLI Maker", "Render highly customizable boxes\n in the terminal")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(out)
//
// # Construction
//
// It is recommended to create boxes using NewBox, which returns a Box with
// sensible defaults. The returned Box can then be configured using methods
// such as Style, Padding, ContentAlign, Color, TitleColor, and TitlePosition.
//
// The zero value of Box is not intended to be used directly with Render.
// If you construct a Box manually (e.g. with &box.Box{} or new(box.Box)),
// you must fully initialize all required fields (style, padding, glyphs,
// colors, etc.) yourself before calling Render.
//
// # Styles
//
// Box styles are selected with Style and the BoxStyle constants:
//
//	box.Single
//	box.Double
//	box.Round
//	box.Bold
//	box.SingleDouble
//	box.DoubleSingle
//	box.Classic
//	box.Hidden
//	box.Block
//
// You can further customize any style by overriding the corner and edge glyphs
// using TopRight, TopLeft, BottomRight, BottomLeft, Horizontal, and Vertical.
//
// # Titles and alignment
//
// Titles can be placed inside the box, on the top border, or on the bottom
// border using TitlePosition with the TitlePosition constants:
//
//	box.Inside
//	box.Top
//	box.Bottom
//
// Title alignment is controlled with TitleAlign and the AlignType constants:
// Inside defaults to Center (within box), Top/Bottom default to Left (on border).
//
//	box.Left
//	box.Center
//	box.Right
//
// Content alignment is controlled with ContentAlign and the AlignType
// constants:
//
//	box.Left
//	box.Center
//	box.Right
//
// # Wrapping
//
// WrapContent enables or disables automatic wrapping of the content. By
// default, when wrapping is enabled, the box width is based on two‑thirds of
// the terminal width. WrapLimit can be used to set an explicit maximum width.
//
// # Colors
//
// TitleColor, ContentColor, and Color accept either one of the first 16 ANSI
// color name constants (e.g. box.Green, box.BrightRed) or a
// #RGB / #RRGGBB / rgb:RRRR/GGGG/BBBB / rgba:RRRR/GGGG/BBBB/AAAA value.
// Invalid colors cause Render to return an error.
//
// # Errors
//
// Render returns an error if the style or title position is invalid, the wrap
// limit or padding is negative, a multiline title is used with a non‑Inside
// title position, any configured colors are invalid, or the terminal width
// cannot be determined. MustRender is a convenience wrapper that panics on
// error.
//
// # Copying
//
// Copy returns a shallow copy of a Box so you can define a base style and
// derive variants without mutating the original:
//
//	base := box.NewBox().Style(box.Single).Padding(2, 1)
//	info := base.Copy().Color(box.Green)
//	warn := base.Copy().Color(box.Yellow)
//
// Each Copy can then be customized and rendered independently.
package box
