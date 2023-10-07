<hr/>
<div align="center">
<img src="img/lib_logo.png" alt="logo">
</div>
<hr/>
<br/>
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

- Make a Terminal Box in 8️⃣ inbuilt different styles
- 16 Inbuilt Colors and True Color Support 🎨
- Custom Title Positions 📏
- Make your Terminal Box style 📦
- Support for ANSI, Tabbed, Multi-line and Line Wrapping boxes 📑
- Align the text according to your needs 📐
- Unicode, Emoji and [Windows Console](https://en.wikipedia.org/wiki/Windows_Console) Support 😋
- Written in 🇬 🇴

## Installation

```terminal
 go get github.com/Delta456/box-cli-maker/v2
```

## Usage Tutorial

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
  - `WrappingLimit`: Wrap the `Content` up to the Limit

### `Box` Methods

`Box.Print(title, lines string)` prints Box from the specified arguments.

- Parameters
  - `title` : Box Title
  - `lines` : Box Content

`Box.Println(title, lines string)` prints Box in a newline from the specified arguments.

- Parameters
  - `title` : Box Title
  - `lines` : Box Content

`Box.String(title, lines string) string` returns `string` representation of Box.

- Parameters
  - `title` : Box Title
  - `lines` : Box Content

### Box Types

- `Single`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/single.svg" alt="single" width=500/>
</p>

- `Single Double`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/single_double.svg" alt="single_double" width="500"/>
</p>

- `Double`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/double.svg" alt="double" width="500"/>
</p>

- `Double Single`
<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/double_single.svg" alt="double_single" width="500"/>
</p>

- `Bold`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/bold.svg" alt="bold" width="500"/>
</p>

- `Round`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/round.svg" alt="round" width="500"/>
</p>

- `Hidden`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/hidden.svg" alt="hidden" width="500"/>
</p>

- `Classic`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/classic.svg" alt="classic" width="500"/>
</p>

### Title Positions

- `Inside`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/single.svg" alt="single" width=500/>
</p>

- `Top`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/top.svg" alt="top" width=500/>
</p>

- `Bottom`

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/bottom.svg" alt="bottom" width=500/>
</p>

### Custom Box

A Custom Box can be created by using the built-in Box struct provided by the module.

```go
type Box struct {
  TopRight    string // TopRight Corner Symbols
  TopLeft     string // TopLeft Corner Symbols
  Vertical    string // Vertical Bar Symbols
  BottomRight string // BottomRight Corner Symbols
  BottomLeft  string // BottomLeft Corner Symbols
  Horizontal  string // Horizontal Bar Symbols
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

<p align="center" style="margin-top: 30px; margin-bottom: 20px;">
<img src="img/custom.svg" alt="custom" width=500/>
</p>

More examples can be found in the `examples/` folder which you can explore yourself.
Feel free to add more examples and submit a pr for the same.

### Color Types

 The [gookit/color](https://github.com/gookit/color) module provides support for colors in your program. It uses two main types: `FgColor` and `FgHiColor`. These types correspond to different color settings.`Color` is a key for the following maps:

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

There are two maps, `fgColors` and `fgHiColors`, which map color names to their respective settings. Here are some examples:

- `fgColors` maps basic colors like Black, Blue, Red, Green, Yellow, Cyan, Magenta, and White.
- `fgHiColors` maps high-intensity colors like HiBlack, HiBlue, HiRed, HiGreen, HiYellow, HiCyan, HiMagenta, and HiWhite.

If you want to use high-intensity colors, make sure the color name starts with "Hi." If the color option is empty or invalid, no color will be applied.

You can also use True Colors by providing them as either a `uint` or `[3]uint`. For `[3]uint`, all elements must be in the range of `[0, 0xFF]`, and for `uint`, it should be in the range of `[0x000000, 0xFFFFFF]`.

If your terminal doesn't support True Colors, the module will automatically round the colors to match the terminal's maximum supported colors. This simplifies things for most users.

If you're curious about supported terminals, you can check the list of 24-bit [supported terminals](https://gist.github.com/XVilka/8346728) and 8-bit [supported terminals](https://fedoraproject.org/wiki/Features/256_Color_Terminals).

Additionally, this module enables True Color and 256 Colors support on Windows Console through [Virtual Terminal Processing](https://docs.microsoft.com/en-us/windows/console/console-virtual-terminal-sequences). However, for 256 colors, you need at least [Windows 10 Version 1511](https://en.wikipedia.org/wiki/Windows_10_version_history_(version_1511)), and for True Color support, you need at least [Windows 10 Version 1607](https://en.wikipedia.org/wiki/Windows_10_version_history_(version_1607)).

Finally, if you're using Windows, the module can detect `ConEmu` or `ANSICON` if they're installed. To ensure the best experience, it's recommended to use the latest versions of both programs.

### Content Wrapping

This library allows the usage of custom wrapping of `Content` so that the Box formed will be created according to your own needs.

To use this feature, you need to do two things:

1. Set Config.AllowWrapping to true. This tells the library to allow custom wrapping.

2. Optionally, you can set your own wrapping limit using Config.WrappingLimit. By default, the wrapping limit is set to 2/3 of the current terminal's width (TermWidth).

So, in simple terms, if you want to customize how content is wrapped in boxes, make sure to enable wrapping with Config.AllowWrapping, and if needed, you can adjust the wrapping limit using Config.WrappingLimit, which is initially set to 2/3 of the terminal's width.

### Note

#### 1. Vertical Alignment

As different terminals have different fonts by default, the right vertical alignment may not be aligned well. You will have to change your font accordingly to make it work.

#### 2. Limitations of Unicode and Emoji

Unicode is a character encoding standard that assigns unique codes to characters from various writing systems, allowing universal text representation. 

This library relies on [mattn/go-runewidth](https://github.com/mattn/go-runewidth) to support Unicode and Emoji characters but consider the following:

1. Proper rendering of Unicode and Emojis on Windows is primarily supported by `Windows Terminal`, `ConEmu`, and `Mintty` as terminal emulators. They handle these characters well.

2. Indic text (a script used in some Indian languages) may not display correctly in most terminals. It's important to note that support for Indic text is quite limited.

3. Avoid using this library in certain environments like online code playgrounds such as [`Go Playground`](https://play.golang.org/), [`Repl.it`](https://repl.it), and in Continuous Integration/Continuous Deployment (CI/CD) pipelines. These platforms typically use fonts that only support basic ASCII characters, making it problematic to accurately calculate character lengths when the font changes dynamically.

4. Depending on your font settings, you may need to make adjustments to ensure proper rendering of Unicode and Emoji characters. Failure to do so could result in issues with vertical alignment.

In essence, while this library supports Unicode and Emojis, there are limitations and considerations to keep in mind, especially when dealing with different terminal environments and character sets.

#### 3. Terminal Color Detection

For beginners, this library can adjust True Color to work with your terminal's capabilities, either 8-bit or 4-bit color.

Detecting your terminal's color capacity isn't always straightforward, and there's no universal method. However, if you know a solution for your specific terminal, you can contribute by making a pull request.

For Unix systems:

- If True Color detection fails, you can set the `COLORTERM` environment variable to `truecolor` or `24bit` for True Color support.

- For 8-bit color-based terminals, if detection doesn't work, set the `TERM` environment variable to your terminal emulator's name followed by `256color`, like `xterm-256color`.

Keep in mind that very old terminals like [`Windows Console (Legacy Mode)`](https://docs.microsoft.com/en-us/windows/console/legacymode) or terminals with `TERM` set to `DUMB` might not show color effects, and the module may output unexpected results or warnings.

It's generally not recommended to use this library with color effects in Online Playgrounds, CI/CD environments, Browsers, and similar platforms since color support varies, and detection is challenging. If you encounter issues, you can open an issue and suggest a solution!

#### 4. Tabs

This library supports the usage of tabs but their use should be limited.

### Projects using Box CLI Maker

- <img src="img\k8s_logo.png" alt="kubernetes logo" width="20"> [kubernetes/minikube](https://github.com/kubernetes/minikube): Run Kubernetes locally.
- +[Many More](https://pkg.go.dev/github.com/Delta456/box-cli-maker/v2?tab=importedby)!

### Acknowledgements

I thank the following people and their packages whom I have studied and was able to port to Go accordingly.

- [thecodrr/boxx](https://github.com/thecodrr/boxx)
- [Atrox/box](https://github.com/Atrox/box)
- [sindreorhus-cli-boxes](https://github.com/sindresorhus/cli-boxes)

Also special thanks to [@elimsteve](https://github.com/elimisteve) who helped me to optimize the code and told me the best possible ways to fix my problems, [@JalonSolov](https://github.com/JalonSolov) for tab lines support and [Kunal Raghav](https://github.com/KunalRaghav) for making the library's logo.

Kudos to [moul/golang-repo-template](https://github.com/moul/golang-repo-template) for their Go template.

### License

Licensed under [MIT](LICENSE)
