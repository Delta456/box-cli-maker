# Box CLI Maker 📦


[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/Delta456/box-cli-maker)
[![CI](https://github.com/Delta456/box-cli-maker/workflows/Box%20CLI%20Maker/badge.svg)](https://github.com/Delta456/box-cli-maker/actions?query=workflow%3A"Box+CLI+Maker")
[![Go Report Card](https://goreportcard.com/badge/github.com/Delta456/box-cli-maker)](https://goreportcard.com/report/github.com/Delta456/box-cli-maker)
[![GolangCI](https://golangci.com/badges/github.com/moul/golang-repo-template.svg)](https://golangci.com/r/github.com/Delta456/box-cli-maker)
[![GitHub release](https://img.shields.io/github/release/Delta456/box-cli-maker.svg)](https://github.com/Delta456/box-cli-maker/releases)


Box CLI Maker is a Highly Customized Terminal Box Creator.

## Features

- Make Terminal Box in 8️⃣ inbuilt different style
- Color Support 🎨
- Custom Title Positions
- Make your own Terminal Box style 📦
- Align the text according to the need
- Unicode and Emoji Support 😋
- Written in 🇬 🇴

## Installation

```terminal
 go get github.com/Delta456/box-cli-maker
```

## Usage

In `main.go`

```go
package main

import "github.com/Delta456/box-cli-maker"

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan"})
	Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")
}
```

`box.New(config Config)` accepts a `Config struct` with following parameters and returns a `Box struct`.

- Parameters
  - `Px` : Horizontal Padding
  - `Py` : Vertical Padding
  - `ContentAlign` : Align the content inside the Box i.e. `Center`, `Left` and `Right`
  - `Type`: Type of Box [*click this for more info*](./README.md/#box-types)
  - `TitlePos` : Position of the Title i.e. `Inside`, `Top` and `Bottom` [*click this for more info*](./README.md/#title-positions)
  - `Color` : Color of the Box  [*click this for more info*](./README.md/#color-types)

### `Box struct` Methods

`Box.Print(title, lines string)` prints Box from the specified arguments.

- Parameters
  - `title` : Title of the Box
  - `lines` : Content to be written inside the Box

`Box.Println(title, lines string)` prints Box in a newline from the specified arguments.

- Parameters
  - `title` : Title of the Box
  - `lines` : Content to be written inside the Box
 
`Box.String(title, lines string) string` return `string` representation of Box.

- Parameters
  - `title` : Title of the Box
  - `lines` : Content to be written inside the Box

### Box Types

- `Single`

<img src="img/single.png" alt="single" width="400"/>

- `Single Double`

<img src="img/single_double.png" alt="single_double" width="400"/>

- `Double`

<img src="img/double.png" alt="double" width="400"/>

- `Double Single`

<img src="img/double_single.png" alt="double_single" width="400"/>

- `Bold`

<img src="img/bold.png" alt="bold" width="400"/>

- `Round`

<img src="img/round.png" alt="round" width="400"/>

- `Hidden`

<img src="img/hidden.png" alt="hidden" width="400"/>

- `Classic`

<img src="img/classic.png" alt="classic" width="400"/>

### Title Positions

- `Inside`

<img src="img/single.png" alt="single" width="400"/>

- `Top`

<img src="img/top.png" alt="top" width="400"/>

- `Bottom`

<img src="img/bottom.png" alt="bottom" width="400"/>


### Making custom Box

You can make your custom Box by using the inbuilt Box struct provided by the module.

```go
type Box struct {
	TopRight    string // TopRight corner used for Symbols
	TopLeft     string // TopLeft corner used for Symbols
	Vertical    string // Symbols used for Vertical Bars
	BottomRight string // BottomRight corner used for Symbols
	BottomLeft  string // BotromLeft corner used for Symbols
	Horizontal  string // Symbols used for Horizontal Bars
	Con      Config // Configuration for the Box to be made
}
```

Using it:

```go
package main

import "github.com/Delta456/box-cli-maker"

func main() {
    config := box.Config{Px: 2, Py: 3, Type: "", TitlePos: "Inside"}
    boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
    boxNew.Print("Box CLI Maker", "Make Highly Customized Terminal Boxes")
}
```

Output:

<img src="img/custom.png" alt="custom" width="350"/>

### Color Types

It has color support from [fatih/color](https://github.com/fatih/color) module from which this module uses `FgColor` and `FgHiColor`. `Color` is a key for the following maps:

```go
var fgColors = map[string]color.Attribute{
	"Black":   color.FgBlack,
	"Blue":    color.FgBlue,
	"Red":     color.FgRed,
	"Green":   color.FgGreen,
	"Yellow":  color.FgYellow,
	"Cyan":    color.FgCyan,
	"Magenta": color.FgMagenta,
}

var fgHiColors = map[string]color.Attribute{
	"HiBlack":   color.FgHiBlack,
	"HiBlue":    color.FgHiBlue,
	"HiRed":     color.FgHiRed,
	"HiGreen":   color.FgHiGreen,
	"HiYellow":  color.FgHiYellow,
	"HiCyan":    color.FgHiCyan,
	"HiMagenta": color.FgHiMagenta,
}
```

If you want High Intensity Colors then the Color name should start with `Hi`. If Color option is empty or invalid then Box with default Color is formed.

It can even have custom color which can be provided in `[3]uint` and `uint` (hex notation) though the elements of the array must be greater than `0x0` and lesser than `0xFF` and `uint` must be greater than `0x000000` and lesser than `0xFFFFFF`.

### Note

#### Vertical Alignment

As different terminals have different font by default so the right vertical alignment may not be aligned well. You will have to change your font accordingly to make it work.

#### Limitations of Unicode and Emoji

It uses [mattn/go-runewidth](https://github.com/mattn/go-runewidth) for Unicode and Emoji support though there are some limitations:

- `Windows Terminal` and `Windows System Linux` are the only know terminals which can render Unicode and Emojis properly on Windows.
- Marathi Text only works on very few Terminals as less support it.
- It is recommended not to use Online Playgrounds like [`Go Playground`](https://play.golang.org/) and [`Repl.it`](https://repl.it) because they use a font that only has ASCII support and other Character Set is used which becomes problematic for finding the length as the font changes during runtime.
- Change your font which supports Unicode and Emojis else the right vertical alignment will break.

### Acknowledgements

I thank the following people and their packages whom I have studied and was able to port to Go accordingly.

- [thecodrr/boxx](https://github.com/thecodrr/boxx)
- [Atrox/box](https://github.com/Atrox/box)
- [sindreorhus-cli-boxs](https://github.com/sindresorhus/cli-boxes)

Special thanks to [@elimsteve](https://github.com/elimisteve) who helped me to optimize and tell me the best ways possible to fix my problems.

Kudos to [moul/golang-repo-template](https://github.com/moul/golang-repo-template) for their Go template.

### Related

- [boxcli](https://github.com/NightShade256/boxcli) : Port of this module in Python by [NightShade256](https://github.com/NightShade256).

### License

Licensed under [MIT](LICENSE)
