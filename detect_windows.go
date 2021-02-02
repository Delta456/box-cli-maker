// +build windows

package box

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/xo/terminfo"
	"golang.org/x/sys/windows"
)

// Get the Windows Version and Build Number
var (
	winVersion, _, buildNumber = windows.RtlGetNtVersionNumbers()
)

// errorMsg prints the msg to os.Stderr and uses Red ANSI Color too if supported
func errorMsg(msg string) {
	// If the terminal doesn't supports the basic 4 bit
	if detectTerminalColor() == terminfo.ColorLevelNone {
		fmt.Fprintln(os.Stderr, msg)
	} else {
		fmt.Fprintln(os.Stderr, color.Red.Sprint(msg))
	}
}

// detectTerminalColor detects the Color Level Supported
func detectTerminalColor() terminfo.ColorLevel {
	if os.Getenv("ConEmuANSI") == "ON" {
		// ConEmuANSI is "ON" for generic ANSI support
		// but True Color option is enabled by default
		// I am just assuming that people wouldn't have disabled it
		// Even if it is not enabled then ConEmu will auto round off
		// accordingly
		return terminfo.ColorLevelMillions
	}
	// Before Windows 10 Build Number 10586, console never supported ANSI Colors
	if buildNumber < 10586 || winVersion < 10 {
		// Detect if using ANSICON on older systems
		if os.Getenv("ANSICON") != "" {
			conVersion := os.Getenv("ANSICON_VER")
			// 8 bit Colors were only supported after v1.81 release
			if conVersion >= "181" {
				return terminfo.ColorLevelHundreds
			}
			return terminfo.ColorLevelBasic
		}
		return terminfo.ColorLevelNone
	}
	if buildNumber < 14931 {
		// True Color is not available before build 14931 so fallback to 8 bit color.
		return terminfo.ColorLevelHundreds
	}
	return terminfo.ColorLevelMillions
}

// roundOffColorVertical rounds off the 24 bit Color to the terminals maximum color capacity for Vertical.
func (b Box) roundOffColorVertical(col color.RGBColor) string {
	switch detectTerminalColor() {
	case terminfo.ColorLevelNone:
		errorMsg("[warning]: terminal does not support colors, using no effects")
		return b.Vertical
	case terminfo.ColorLevelMillions:
		return col.Sprint(b.Vertical)
	default:
		return col.C256().Sprint(b.Vertical)
	}
}

// roundOffColor checks the terminlal color level then rounds off the 24 bit color to the level supported
// for TopBar and BottomBar
func roundOffColor(col color.RGBColor, topBar, bottomBar string) (string, string) {
	switch detectTerminalColor() {
	case terminfo.ColorLevelNone:
		errorMsg("[warning]: terminal does not support colors, using no effects")
		return topBar, bottomBar
	case terminfo.ColorLevelMillions:
		return col.Sprint(topBar), col.Sprint(bottomBar)
	default:
		return col.C256().Sprint(topBar), col.C256().Sprint(bottomBar)
	}
}
