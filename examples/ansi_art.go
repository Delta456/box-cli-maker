package main

import "github.com/Delta456/box-cli-maker/v2"

func main() {
	Box := box.New(box.Config{Px: 2, Py: 5, Type: "Single", Color: "Cyan", ContentColor: "Green"})
	// All ANSI Art may not fit inside the box as they can exceed the terminal width
	Box.Print("", `
	__________                 _________  .____     .___     _____            __                   
	\______   \  ____ ___  ___ \_   ___ \ |    |    |   |   /     \  _____   |  | __  ____ _______ 
	 |    |  _/ /  _ \\  \/  / /    \  \/ |    |    |   |  /  \ /  \ \__  \  |  |/ /_/ __ \\_  __ \
	 |    |   \(  <_> )>    <  \     \____|    |___ |   | /    Y    \ / __ \_|    < \  ___/ |  | \/
	 |______  / \____//__/\_ \  \______  /|_______ \|___| \____|__  /(____  /|__|_ \ \___  >|__|   
			\/              \/         \/         \/              \/      \/      \/     \/ `)
}
