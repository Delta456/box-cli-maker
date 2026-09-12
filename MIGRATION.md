# Migration Guide: v2 → v3

This guide will help you migrate your projects from v2 to v3 of `box-cli-maker`, and highlights the benefits, new features, and breaking changes.

## Why Migrate?

- **v3 is the latest, actively maintained version** under the new organization:  
	`github.com/box-cli-maker/box-cli-maker/v3`
- **Modern, fluent API:** v3 introduces a chainable, builder-style API for easier and more readable configuration.
- **Type safety:** v3 uses constants for style, alignment, and colors instead of strings, reducing runtime errors.
- **Improved error handling:** More robust and explicit error reporting.
- **Better color support:** Consistent handling of ANSI, hex, and XParseColor formats.
- **Cleaner codebase:** Deprecated and confusing features from v2 have been removed for maintainability.
- **Active development:** All new features, bug fixes, and improvements will be in v3.

## New Features in v3 (Not in v2)

- **Fluent, chainable API:** Configure boxes with method chaining for clarity and brevity.
- **Strongly typed options:** Use constants for box styles, alignments, and title positions (e.g., `box.Single`, `box.Center`, `box.Top`).
- **Improved color support:** Accepts named ANSI colors, hex codes (e.g., `#FF00FF`), and XParseColor strings.
- **Explicit content wrapping:** Easily enable/disable wrapping and set wrap limits with `.WrapContent()` and `.WrapLimit()`.
- **Box copying:** Use `.Copy()` to duplicate box configurations.
- **Better error handling:** `Render()` returns errors for invalid configurations; `MustRender()` panics on error for convenience.
- **Simplified custom glyphs:** Set custom border characters with dedicated methods (e.g., `.TopLeft("+")`).
- **Cleaner API surface:** Deprecated methods and fields removed for a more focused experience.
- **Accurate rendering of emoji and custom borders:** v3 ensures correct box layout even when using emojis or custom borders (where border symbols may be different for each side or corner), especially when the title is longer than the content and the title position is set to `Top` or `Bottom`.
- **New Block Box style**: Render boxes with the new built-in `box.Block` style.

## Breaking Changes (v2 → v3)

| Area                | v2 Example / Behavior                                                                 | v3 Example / Behavior                                                                 |
|---------------------|--------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------|
| **Module Path**     | `module github.com/Delta456/box-cli-maker/v2`                                        | `module github.com/box-cli-maker/box-cli-maker/v3`                                   |
| **Import Path**     | `import box "github.com/Delta456/box-cli-maker/v2"`                                 | `import box "github.com/box-cli-maker/box-cli-maker/v3"`                            |
| **Box Construction**| `box.New(box.Config{...})`                                                           | `box.NewBox().Style(...).Padding(...)...`                                            |
| **Config Options**  | Strings for style, alignment, etc. (`Type: "Single"`)                               | Strongly typed constants (`.Style(box.Single)`)                                      |
| **Field Names**     | `Type`, `TitlePos`, `ContentAlign`                                                   | `.Style()`, `.TitlePosition()`, `.ContentAlign()`                                    |
| **Color Handling**  | Various types and string names                                                       | Only named ANSI colors, hex (`#RGB`, `#RRGGBB`), or XParseColor formats              |
| **Rendering**       | `b.String("Title", "Content")`                                                    | `b.MustRender("Title", "Content")` or `b.Render(...)`                             |
| **Print Methods**   | `b.Print()`, `b.Println()`                                                           | Removed; use `fmt.Println(b.MustRender(...))`                                        |
| **Wrapping**        | Implicit/less explicit                                                               | `.WrapContent(true)`, `.WrapLimit(width)`                                            |
| **Custom Glyphs**   | Config struct fields (e.g., `TopLeft: "+"`)                                        | Methods (e.g., `.TopLeft("+")`)                                                    |
| **Cloning/Copying** | Assignment or manual copy                                                            | `.Copy()` method                                                                     |
| **Deprecated/Removed** | `Config` struct, string-based config, undocumented color formats, Print methods, etc. | All replaced by new API; only documented color formats and methods remain            |

## Example Migration

**v2:**
```go
import box "github.com/Delta456/box-cli-maker/v2"

cfg := box.Config{
	Type:         "Single",
	Px:           2,
	Py:           1,
	ContentAlign: "Center",
	TitlePos:     "Top",
	Color:        "Cyan",
	TitleColor:   "BrightYellow",
}
b := box.New(cfg)
fmt.Println(b.String("Box CLI Maker", "Render highly customizable boxes\n in the terminal"))
```

**v3:**
```go
import box "github.com/box-cli-maker/box-cli-maker/v3"

b := box.NewBox().
	Style(box.Single).
	Padding(2, 1).  // horizontal, vertical
	TitlePosition(box.Top).
	ContentAlign(box.Center).
	Color(box.Cyan).
	TitleColor(box.BrightYellow)

fmt.Println(b.MustRender("Box CLI Maker", "Render highly customizable boxes\n in the terminal"))
```

## How to Upgrade

1. **Change your import path:**
	 ```go
	 import box "github.com/box-cli-maker/box-cli-maker/v3"
	 ```
2. **Update your `go.mod`:**
	 ```sh
	 go get github.com/box-cli-maker/box-cli-maker/v3
	 ```
3. **Update your code for the new API:**
	 - Use method chaining for configuration.
	 - Use strongly typed constants for style, alignment, and colors.
	 - See the [v3 README](https://github.com/box-cli-maker/box-cli-maker) for examples.
4. **Remove the old dependency:**
	 ```sh
	 go mod tidy
	 ```

## Troubleshooting

- Check for any remaining v2 import paths.
- Review the [v3 documentation](https://pkg.go.dev/github.com/box-cli-maker/box-cli-maker/v3) for new features and API changes.
- If you encounter issues, open an [issue](https://github.com/box-cli-maker/box-cli-maker/issues).

For more details and advanced usage, see the [v3 README](https://github.com/box-cli-maker/box-cli-maker) and [examples](https://github.com/box-cli-maker/box-cli-maker/tree/master/examples).
