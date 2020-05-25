# Box CLI Maker 📦

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c)](https://pkg.go.dev/github.com/Delta456/box-cli-maker)
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2FDelta456%2Fbox-cli-maker%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/Delta456/box-cli-maker/goto?ref=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/Delta456/box-cli-maker)](https://goreportcard.com/report/github.com/Delta456/box-cli-maker)


Box CLI Maker is a Highly Customized Terminal Box Creator.

## Features

- Make terminal box in 8️⃣ inbuilt different style
- Color Support 🎨
- Custom Title Positions
- Make your own Box style 📦
- Align the text according to the need
- *Limited* Unicode and Emoji Support 😋
- Written in  🇬 🇴

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

`box.New(config Config)` accepts a `Config struct` with following parameters amd returns a `Box struct`.

- Parameters
  - `Px` : Horizontal Padding
  - `Py` : Vertical Padding
  - `ContentAlign` : Align content in the Box i.e. `Center`, `Left` and `Right`
  - `Type`: Type of Box (listed down below)
  - `TitlePos` : Position of the Title i.e. `Inside`, `Top` and `Bottom`
  - `Color` : Color of the Box

### `Box struct` Methods

`Box.Print(title, lines string)` prints Box from the specified arguements.

- Parameters
  - `title` : Title of the Box
  - `lines` : Content written inside the Box

`Box.Println(title, lines string)` prints Box in a newline from the specified arguements.

- Parameters
  - `title` : Title of the Box
  - `lines` : Content written inside the Box
 
`Box.String(title, lines string) string` return `string` representation of Box.

- Parameters
  - `title` : Title of the Box
  - `lines` : Content written inside the Box

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
    boxNew.Print("Box CLI Maker", "Make highly customized Terminal Boxes")
}
```

Output:

<img src="img/custom.png" alt="custom" width="350"/>

### Color Types

It has color support from [fatih/color](https://github.com/fatih/color) module from which this module uses `FgColor` and `FgHiColor`. `Color` is a key for the following maps:

```go
var fgColors map[string]color.Attribute = map[string]color.Attribute{
	"Black":  color.FgBlack,
	"Blue":   color.FgBlue,
	"Red":    color.FgRed,
	"Green":  color.FgGreen,
	"Yellow": color.FgYellow,
	"Cyan":   color.FgCyan,
	"Magenta": color.FgMagenta,
}

var fgHiColors map[string]color.Attribute = map[string]color.Attribute{
	"HiBlack":  color.FgHiBlack,
	"HiBlue":   color.FgHiBlue,
	"HiRed":    color.FgHiRed,
	"HiGreen":  color.FgHiGreen,
	"HiYellow": color.FgHiYellow,
	"HiCyan":   color.FgHiCyan,
	"HiMagenta": color.FgHiMagenta,
}
```

If you want High Intensity Colors then the Color name should start with `Hi`. If Color option is empty then normal Box is formed.

### Note

#### Vertical Alignment

As different terminals have different font by default so the right vertical alignment may not be aligned well. You will have to change your font accordingly to make it work.

#### Limited Unicode Suppport

It uses [rivo/uniseg](https://github.com/rivo/uniseg) for Unicode support though there are some limitations:

- Different terminals render Unicode differently.

- No known terminal support characters like Japanese and Chinese ones because their width is `1.5` not `1`.

- Windows CMD and Powershell doesn't render emojis well at all so the right vertical alignment will break so it prefered to use Git Bash in IDEs like VSCode. On Linux some terminal can render and some cannot.

- Online Go Compilers like [Go Playground](https://play.golang.org/) don't support Unicode for this package and [repl.it](https://repl.it) supports Unicode but not all characters as stated above.

If you want to use Unicode then do it at your own risk. I will try to find a solution for this soon or if you have a solution for this then please do a PR or so!

### Acknowledgements

I thank the following people and their packages whom I have studied and was able to port to Go accordingly.

- [thecodrr/boxx](https://github.com/thecodrr/boxx)
- [Atrox/box](https://github.com/Atrox/box)
- [sindreorhus-cli-boxs](https://github.com/sindresorhus/cli-boxes)

### Related

- [boxcli](https://github.com/NightShade256/boxcli) : Port of this module in Python by [NightShade256](https://github.com/NightShade256).

### License

Licensed under [MIT](LICENSE)
