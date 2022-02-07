<div align="center">
<h1>Box CLI Maker 📦 </h1>
</div>

<p align="center">
Box CLI Maker is a Highly Customized Terminal Box Creator.
</p>

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/Delta456/box-cli-maker/v2.svg)](https://pkg.go.dev/github.com/Delta456/box-cli-maker/v2)
[![godocs.io](http://godocs.io/github.com/Delta456/box-cli-maker?status.svg)](https://godocs.io/github.com/Delta456/box-cli-maker/v2)
[![CI](https://github.com/Delta456/box-cli-maker/workflows/Box%20CLI%20Maker/badge.svg)](https://github.com/Delta456/box-cli-maker/actions?query=workflow%3A"Box+CLI+Maker")
[![Go Report Card](https://goreportcard.com/badge/github.com/Delta456/box-cli-maker)](https://goreportcard.com/report/github.com/Delta456/box-cli-maker)
[![GolangCI](https://golangci.com/badges/github.com/moul/golang-repo-template.svg)](https://golangci.com/r/github.com/Delta456/box-cli-maker)
[![GitHub release](https://img.shields.io/github/release/Delta456/box-cli-maker.svg)](https://github.com/Delta456/box-cli-maker/releases)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

</div>

## Features

- Make Terminal Box in 8️⃣ inbuilt different styles
- 16 Inbuilt Colors and True Color Support 🎨
- Custom Title Positions 📏
- Make your own Terminal Box style 📦
- Support for Tabbed, Multi Lines and Line Wrapping 📑
- Align the text according to the need 📐
- Unicode, Emoji and [Windows Console](https://en.wikipedia.org/wiki/Windows_Console) Support 😋
- Written in 🇬 🇴

## Installation

```terminal
 go get github.com/Delta456/box-cli-maker/v2
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

`box.New(config Config)` takes Box `Config` and returns a `Box` from the given `Config`.

- Parameters
  - `Px` : Horizontal Padding
  - `Py` : Vertical Padding
  - `ContentAlign` : Content Alignment inside Box i.e. `Center`, `Left` and `Right`
  - `Type`: Box Type
  - `TitlePos` : Title Position of Box i.e. `Inside`, `Top` and `Bottom`
  - `Color` : Box Color
  - `TitleColor` : Title Color
  - `ContentColor` : Content Color
  - `AllowWrapping`: Flag to allow custom `Content` wrapping
  - `WrappingLimit`: Wrap the `Content` upto the Limit

### `Box struct` Methods

`Box.Print(title, lines string)` prints Box from the specified arguments.

- Parameters
  - `title` : Box Title
  - `lines` : Box Content

`Box.Println(title, lines string)` prints Box in a newline from the specified arguments.

- Parameters
  - `title` : Box Title
  - `lines` : Box Content

`Box.String(title, lines string) string` return `string` representation of Box.

- Parameters
  - `title` : Box Title
  - `lines` : Box Content

### Box Types

<style>
  .center {
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 50%;
}
</style>

<details>

<summary>Single</summary>


<img src="img/carbon(1).svg" alt="single" width="300" class="center"/>

</details>

<details>


<summary>Double Single</summary>

<img src="img/double_single.png" alt="single" width="300"/>

</details>
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

### Custom Box

You can also make your custom Box by using the inbuilt Box struct provided by the module.

```go
type Box struct {
	TopRight    string // TopRight Corner Symbols
	TopLeft     string // TopLeft Corner Symbols
	Vertical    string // Vertical Bars Symbols
	BottomRight string // BottomRight Corner Symbols
	BottomLeft  string // BottomRight Corner Symbols
	Horizontal  string // Horizontal Bars Symbols
	Config             // Box Config
}
```

#### Usage

In `main.go`:

```go
package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
    config := box.Config{Px: 2, Py: 3, Type: "", TitlePos: "Inside"}
    boxNew := box.Box{TopRight: "*", TopLeft: "*", BottomRight: "*", BottomLeft: "*", Horizontal: "-", Vertical: "|", Config: config}
    boxNew.Println("Box CLI Maker", "Make Highly Customized Terminal Boxes")
}
```

`go run main.go`:

<img src="img/custom.png" alt="custom" width="350"/>

More examples can be found in `examples/` folder.

### Color Types

It has color support from [gookit/color](https://github.com/gookit/color) module from which this module uses `FgColor` and `FgHiColor`. `Color` is a key for the following maps:

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

1. True Color is also possible though you need to provide it as `uint` or `[3]uint` and make sure that the terminals which will be targetted must have it supported.

2. `[3]uint`'s element all must be in a range of `[0, 0xFF]` and `uint` in range of `[0x000000, 0xFFFFFF]`.

As convenience, if the terminal's doesn't support True Color then it will round off according to the terminal's max supported colors which makes it easier for the users not to worry about other terminal for most of the cases.

Here's a list of 24 bit [supported terminals](https://gist.github.com/XVilka/8346728) and 8 bit [supported terminals](https://fedoraproject.org/wiki/Features/256_Color_Terminals).

This module also enables **True Color** and **256 Colors** support on Windows Console through [Virtual Terminal Processing](https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences) but you need have at least [Windows 10 Version 1511](https://en.wikipedia.org/wiki/Windows_10_version_history_(version_1511)) for 256 colors or [Windows 10 Version 1607](https://en.wikipedia.org/wiki/Windows_10_version_history_(version_1607)) for True Color Support.

4-bit Colors are now standardized so it should supported by all Terminals now.

If `ConEmu` or `ANSICON` is installed for Windows systems then it will be also detected. It is highly recommended to use the latest versions of both of them to have the best experience.

### Content Wrapping

This library allows the usage of custom wrapping of `Content` so that the Box formed would not be out of bounds or be according to your own need.

To enable this `Config.AllowWrapping` must be set to `true` and you can also provide your own wrapping limit via `Config.WrappingLimit`. By default the wrapping limit is `2*TermWidth/3` where `TermWidth` is terminal's width when the flag above is set to `true`.

### Note

#### 1. Vertical Alignment

As different terminals have different font by default so the right vertical alignment may not be aligned well. You will have to change your font accordingly to make it work.

#### 2. Limitations of Unicode and Emoji

It uses [mattn/go-runewidth](https://github.com/mattn/go-runewidth) for Unicode and Emoji support though there are some limitations:

- `Windows Terminal`, `ConEmu` and `Mintty` are the only know terminal emulators which can render Unicode and Emojis properly on Windows.
- Indic Text only works on very few Terminals as less support it.
- It is recommended not to use this for Online Playgrounds like [`Go Playground`](https://play.golang.org/) and [`Repl.it`](https://repl.it), `CI/CDs` etc. because they use a font that only has ASCII support and other Character Set is used which becomes problematic for finding the length as the font changes during runtime.
- Some changes will be needed to your font which supports Unicode and Emojis else the right vertical alignment will break.

#### 3. Terminal Color Detection

It is possible to round off True Color provided to 8 bit or 4 bit according to your terminal's maximum capacity.

There is no **standardized way** of detecting the terminal's maximum color capacity so the way of detecting your terminal might not work for you. If this can be fixed for your terminal then you can always make a PR.

The following two points are just applicable for **Unix** systems:

- If you think that the module can't detect True Color of the terminal then you must set your environment variable `COLORTERM` to `truecolor` or `24bit` for True Color support.

- If you are targetting 8 bit color based terminals and the module couln't detect it then set your environment variable `TERM` to name of the terminal emulator with `256color` as suffix like `xterm-256color`.

There might be no color effect for very old terminals like [`Windows Console (Legacy Mode)`](https://docs.microsoft.com/en-us/windows/console/legacymode) or `TERM` environment variable  which gives `DUMB` so the module will output some garbage value or a warning if used.

In `Online Playgrounds`, `CI/CDs`, `Browsers` etc, it is recommended **not** to use this module with color effect as few may have it but this is hard to detect in general. If you think that it's possible then open an issue and address the solution!

#### 4. Tabs

This library supports the usage of tabs but it must be noticed that the no. of tabs used should be limited.

### Projects using Box CLI Maker

- <img src="https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/kubernetes/kubernetes.png" alt="kubernetes logo" width="20"> [kubernetes/minikube](https://github.com/kubernetes/minikube): Run Kubernetes locally.
- +[Many More](https://pkg.go.dev/github.com/Delta456/box-cli-maker/v2?tab=importedby)!

### Acknowledgements

I thank the following people and their packages whom I have studied and was able to port to Go accordingly.

- [thecodrr/boxx](https://github.com/thecodrr/boxx)
- [Atrox/box](https://github.com/Atrox/box)
- [sindreorhus-cli-boxes](https://github.com/sindresorhus/cli-boxes)

Special thanks to [@elimsteve](https://github.com/elimisteve) who helped me to optimize the code and told me the best possible ways to fix my problems and [@JalonSolov](https://github.com/JalonSolov) for tab lines support.

Kudos to [moul/golang-repo-template](https://github.com/moul/golang-repo-template) for their Go template.

### License

Licensed under [MIT](LICENSE)
