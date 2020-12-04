//+build !windows

// errorMsg prints the msg to os.Stderr and uses Red ANSI Color too if supported
func errorMsg(msg string) {
	// If the terminal doesn't supports the basic 4 bit
	if detectTerminalColor() == terminfo.ColorLevelNone {
		fmt.Fprintln(os.Stderr, msg)
	} else {
		fmt.Fprintln(os.Stderr, color.Red.Sprint(msg))
	}
}

func detectTerminalColor() terminfo.ColorLevel {
	// Detect WSL as it has True Color support
	wsl, err := ioutil.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		log.Fatal(err)
	}
	// Microsoft for WSL and microsoft for WSL 2
	if strings.Contains(string(wsl), "microsoft") && strings.Contains(string(wsl), "Microsoft") {
		return terminfo.ColorLevelMillions
	}
	level, err := terminfo.ColorLevelFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	return level
}

// roundOffColorVertical rounds off the 24 bit Color to the terminals maximum color capacity for Vertical.
func (b Box) roundOffColorVertical(col color.RGBColor) string {
	// Check if the terminal supports 256 Colors only
	if detectTerminalColor() == terminfo.ColorLevelHundreds {
		return col.C256().Sprint(b.Vertical)
		// Check if the terminal supports 16 Colors only
	} else if detectTerminalColor() == terminfo.ColorLevelBasic {
		return col.C16().Sprint(b.Vertical)
		// Check if the terminal supports True Color
	} else if detectTerminalColor() == terminfo.ColorLevelMillions {
		return col.Sprint(b.Vertical)
	} else {
		// Return with a warning as the terminal doesn't supports color effect
		errorMsg("[warning]: terminal does not support colors, using no effect")
		return b.Vertical
	}
}

func roundOffColor(col color.RGBColor, topBar, bottomBar string) (string, string) {
	// Check if the terminal supports 256 Colors only
	if detectTerminalColor() == terminfo.ColorLevelHundreds {
		TopBar := col.C256().Sprint(topBar)
		BottomBar := col.C256().Sprint(bottomBar)
		return TopBar, BottomBar
		// Check if the terminal supports 16 Colors only
	} else if detectTerminalColor() == terminfo.ColorLevelBasic {
		TopBar := col.C16().Sprint(topBar)
		BottomBar := col.C16().Sprint(bottomBar)
		return TopBar, BottomBar
		// Check if the terminal supports True Color
	} else if detectTerminalColor() == terminfo.ColorLevelMillions {
		TopBar := col.Sprint(topBar)
		BottomBar := col.Sprint(bottomBar)
		return TopBar, BottomBar
	} else {
		// Return with a warning as the terminal supports no Color
		errorMsg("[warning]: terminal does not support colors, using no effect")
		return topBar, bottomBar
	}
}