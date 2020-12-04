//+build windows
package box

import (
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/xo/terminfo"
	"golang.org/x/sys/windows"
)

// Get the Windows Versio and Build Number
var (
	winVersion, _, buildNumber = windows.RtlGetNtVersionNumbers()
)

// errorMsg prints the msg to os.Stderr and uses Red ANSI Color too if supported
func errorMsg(msg string) {
	// If the terminal doesn't supports the basic 4 bit
	if buildNumber < 10586 || winVersion < 10 {
		fmt.Fprintln(os.Stderr, msg)
	} else {
		fmt.Fprintln(os.Stderr, color.Red.Sprint(msg))
	}
}

func detectTerminalColor() terminfo.ColorLevel {
	return terminfo.ColorLevelMillions
}

// roundOffColorVertical rounds off the 24 bit Color to the terminals maximum color capacity for Vertical.
func (b Box) roundOffColorVertical(col color.RGBColor) string {
	// Before Windows Build Number 10586, console never supported ANSI Colors
	if buildNumber < 10586 || winVersion < 10 {
		errorMsg("[warning]: terminal does not support colors, using no effect")
		return b.Vertical
	} else {
		if buildNumber >= 14931 {
			// True Color is only possible after Windows 10 Build Number 14931
			// Virtual Terminal Processing is also enabled
			return col.Sprint(b.Vertical)

		} else {
			// After Windows 10 Build Number 10586 and if not upgraded to at least 14931 then round off
			// True Color to 8 bit Color
			return col.C256().Sprint(b.Vertical)
		}
	}
}

func roundOffColor(col color.RGBColor, topBar, bottomBar string) (string, string) {
	if buildNumber < 10586 || winVersion < 10 {
		// Before Build Number 10586, console never supported ANSI Colors
		fmt.Fprintln(os.Stderr, "[warning]: terminal does not support colors, using no effect")
		return topBar, bottomBar
	} else {
		if buildNumber >= 14931 {
			// True Color is only possible after Windows 10 Build 14931
			// Virtual Terminal Processing is enabled by default in the later versions
			return col.Sprint(topBar), col.Sprint(bottomBar)

		} else {
			// After Windows 10 build 10586 and if not upgraded to at least then round off
			// True Color to 8 bit Color
			return col.C256().Sprint(topBar), col.C256().Sprint(bottomBar)
		}
	}
}
