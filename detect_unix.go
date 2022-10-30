//go:build !windows
// +build !windows

package box

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gookit/color"
	"github.com/xo/terminfo"
)

// errorMsg prints msg to os.Stderr and uses Red ANSI Color too if supported
func errorMsg(msg string) {
	// If the terminal doesn't supports the basic 4 bit
	if detectTerminalColor() == terminfo.ColorLevelNone {
		fmt.Fprintln(os.Stderr, msg)
	} else {
		fmt.Fprintln(os.Stderr, color.Red.Sprint(msg))
	}
}

// detectTerminalColor detects Color Level Supported
func detectTerminalColor() terminfo.ColorLevel {
	// Detect WSL as it has True Color support
	wsl, err := ioutil.ReadFile("/proc/sys/kernel/osrelease")
	// Additional check needed for Mac Os as it doesn't have "/proc/" folder
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
	}
	// Lowercasing every content inside "/proc/sys/kernal/osrelease"
	// because it gives "Microsoft" for WSL and "microsoft" for WSL 2
	// so no need of checking twice
	if string(wsl) != "" {
		wslLower := strings.ToLower(string(wsl))
		if strings.Contains(wslLower, "microsoft") {
			return terminfo.ColorLevelMillions
		}
	}
	level, err := terminfo.ColorLevelFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	return level
}

// roundOffColorVertical rounds off 24 bit Color to the terminals maximum color capacity for Vertical.
func (b Box) roundOffColorVertical(col color.RGBColor) string {
	switch detectTerminalColor() {
	// Check if the terminal supports 256 Colors only
	case terminfo.ColorLevelHundreds:
		return col.C256().Sprint(b.Vertical)

	// Check if the terminal supports 16 Colors only
	case terminfo.ColorLevelBasic:
		return col.C16().Sprint(b.Vertical)

	// Check if the terminal supports True Color
	case terminfo.ColorLevelMillions:
		return col.Sprint(b.Vertical)

	default:
		// Return with a warning as the terminal supports no Color
		errorMsg("[warning]: terminal does not support colors, using no effect")
		return b.Vertical
	}
}

// roundOffTitleColor rounds off 24 bit Color to the terminals maximum color capacity for Title.
func (b Box) roundOffTitleColor(col color.RGBColor, title string) string {
	switch detectTerminalColor() {
	// Check if the terminal supports 256 Colors only
	case terminfo.ColorLevelHundreds:
		return col.C256().Sprint(title)

	// Check if the terminal supports 16 Colors only
	case terminfo.ColorLevelBasic:
		return col.C16().Sprint(title)

	// Check if the terminal supports True Color
	case terminfo.ColorLevelMillions:
		return col.Sprint(title)

	default:
		// Return with a warning as the terminal supports no Color
		errorMsg("[warning]: terminal does not support colors, using no effect")
		return title
	}
}

// roundOffColor checks terminlal color level then rounds off 24 bit color to the level supported
// for TopBar and BottomBar
func roundOffColor(col color.RGBColor, topBar, bottomBar string) (string, string) {
	switch detectTerminalColor() {
	// Check if the terminal supports 256 Colors only
	case terminfo.ColorLevelHundreds:
		TopBar := addStylePreservingOriginalFormat(topBar, col.C256().Sprint)
		BottomBar := addStylePreservingOriginalFormat(bottomBar, col.C256().Sprint)
		return TopBar, BottomBar

	// Check if the terminal supports 16 Colors only
	case terminfo.ColorLevelBasic:
		TopBar := addStylePreservingOriginalFormat(bottomBar, col.C16().Sprint)
		BottomBar := addStylePreservingOriginalFormat(bottomBar, col.C16().Sprint)
		return TopBar, BottomBar

	// Check if the terminal supports True Color
	case terminfo.ColorLevelMillions:
		TopBar := addStylePreservingOriginalFormat(topBar, col.Sprint)
		BottomBar := addStylePreservingOriginalFormat(bottomBar, col.Sprint)
		return TopBar, BottomBar

	default:
		// Return with a warning as the terminal supports no Color
		errorMsg("[warning]: terminal does not support colors, using no effect")
		return topBar, bottomBar
	}
}
