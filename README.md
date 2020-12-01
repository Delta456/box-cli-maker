<div align="center">
<h1>Box CLI Maker 📦 </h1>
</div>

<p align="center">
Box CLI Maker is a Highly Customized Terminal Box Creator.
</p>

<div align="center">

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/Delta456/box-cli-maker/v2)
[![CI](https://github.com/Delta456/box-cli-maker/workflows/Box%20CLI%20Maker/badge.svg)](https://github.com/Delta456/box-cli-maker/actions?query=workflow%3A"Box+CLI+Maker")
[![Go Report Card](https://goreportcard.com/badge/github.com/Delta456/box-cli-maker)](https://goreportcard.com/report/github.com/Delta456/box-cli-maker)
[![GolangCI](https://golangci.com/badges/github.com/moul/golang-repo-template.svg)](https://golangci.com/r/github.com/Delta456/box-cli-maker)
[![GitHub release](https://img.shields.io/github/release/Delta456/box-cli-maker.svg)](https://github.com/Delta456/box-cli-maker/releases)

</div>

## Features

- Make Terminal Box in 8️⃣ inbuilt different styles
- 16 Inbuilt Colors and Custom (24 bit) Colors Support 🎨
- Custom Title Positions
- Make your own Terminal Box style 📦
- Align the text according to the need
- Unicode, Emoji and [Windows Console](https://en.wikipedia.org/wiki/Windows_Console) Support 😋
- Written in 🇬 🇴

## Installation

```terminal
 go get github.com/Delta456/box-cli-maker
```

## Usage

In `main.go`

```go
package main

import "github.com/Delta456/box-cli-maker/v2"

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
  - `Type`: Type of Box (by default it's Single) [*click this for more info*](./README.md/#box-types)
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
 Config // Configuration for the Box to be made
}
```

Using it:

```go
package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
    config := box.Config{Px: 2, Py: 3, Type: "", TitlePos: "Inside"}
    boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
    boxNew.Print("Box CLI Maker", "Make Highly Customized Terminal Boxes")
}
```

Output:

<img src="img/custom.png" alt="custom" width="350"/>

### Color Types

It has color support from [gookit/color](github.com/gookit/color) module from which this module uses `FgColor` and `FgHiColor`. `Color` is a key for the following maps:

```go
 fgColors map[string]color.Color = {
  "Black":   color.FgBlack,
  "Blue":    color.FgBlue,
  "Red":     color.FgRed,
  "Green":   color.FgGreen,
  "Yellow":  color.FgYellow,
  "Cyan":    color.FgCyan,
  "Magenta": color.FgMagenta,
  "White":   color.FgWhite,
 }
 fgHiColors map[string]color.Color = {
  "HiBlack":   color.FgDarkGray,
  "HiBlue":    color.FgLightBlue,
  "HiRed":     color.FgLightRed,
  "HiGreen":   color.FgLightGreen,
  "HiYellow":  color.FgLightYellow,
  "HiCyan":    color.FgLightCyan,
  "HiMagenta": color.FgLightMagenta,
  "HiWhite":   color.FgLightWhite,
}
```

If you want High Intensity Colors then the Color name must start with `Hi`. If Color option is empty or invalid then Box with default Color is formed.

True Color is possible, you need to provide it as `uint` or `[3]uint` and make sure that the terminals which will be targetted must have it supported.

`[3]uint`'s element all must be in a range of `[0, 0xFF]` and `uint` in range of `[0x000000, 0xFFFFFF]`.

Here's a list of 24 bit [supported terminals](https://gist.github.com/XVilka/8346728) and 16 bit [supported terminals](https://fedoraproject.org/wiki/Features/256_Color_Terminals).

It is possible to have True Color support on Windows Console but you need have `^v1.2.4` else True Color effect will not be there.

### Note

#### Vertical Alignment

As different terminals have different font by default so the right vertical alignment may not be aligned well. You will have to change your font accordingly to make it work.

#### Limitations of Unicode and Emoji

It uses [mattn/go-runewidth](https://github.com/mattn/go-runewidth) for Unicode and Emoji support though there are some limitations:

- `Windows Terminal` and `Windows SubSystem Linux` are the only know terminals which can render Unicode and Emojis properly on Windows.
- Indic Text only works on very few Terminals as less support it.
- It is recommended not to use Online Playgrounds like [`Go Playground`](https://play.golang.org/) and [`Repl.it`](https://repl.it), `CI/CDs` etc. because they use a font that only has ASCII support and other Character Set is used which becomes problematic for finding the length as the font changes during runtime.
- Change your font which supports Unicode and Emojis else the right vertical alignment will break.

#### Terminal Color Detection

It is possible to round off true color provided to 8 bit or 16 bit according to your terminal's maximum capacity.

There is no **standardized way** of detecting the terminal's maximum color capacity so the way of detecting your terminal might not work for you. If this can be fixed for you then you can always make a PR.

If you think that the module can't detect your terminal's color capacity then you must set your environment variable `COLORTERM` to `truecolor` or `24bit` for true color support.

If you are targetting 16 color bit based terminals and if the module couln't detect it then set your environment variable `TERM` to name of the terminal emulator with `256color` as suffix like `xterm-256color`.

There might be no color effect for very old terminals like [`Windows Console (Legacy Mode)`](https://docs.microsoft.com/en-us/windows/console/legacymode) or environment variable `TERM` gives `DUMB` so it will output some garbage value if used.

### Acknowledgements

I thank the following people and their packages whom I have studied and was able to port to Go accordingly.

- [thecodrr/boxx](https://github.com/thecodrr/boxx)
- [Atrox/box](https://github.com/Atrox/box)
- [sindreorhus-cli-boxes](https://github.com/sindresorhus/cli-boxes)

Special thanks to [@elimsteve](https://github.com/elimisteve) who helped me to optimize the code and told me the best possible ways to fix my problems and [@JalonSolov](https://github.com/JalonSolov) for tab lines support.

Kudos to [moul/golang-repo-template](https://github.com/moul/golang-repo-template) for their Go template.

### Related

- [boxcli](https://github.com/NightShade256/boxcli) : Port of this module in Python by [NightShade256](https://github.com/NightShade256).

### License

Licensed under [MIT](LICENSE)
