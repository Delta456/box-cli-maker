package main

import (
	"fmt"

	box "github.com/box-cli-maker/box-cli-maker/v3"
)

func main() {
	b := box.NewBox().
		Padding(2, 1).
		Style(box.Single).
		Color(box.Cyan).
		TitlePosition(box.Top)

	title := "\033[1mBold\033[0m · \033[4mUnderline\033[0m · \033[5mBlink\033[0m · \033[9mStrike\033[0m · Hyperlink"

	content := "" +
		"• Normal text\n" +
		"• \033[1mBold text\033[0m\n" +
		"• \033[4mUnderlined text\033[0m\n" +
		"• \033[5mBlinking text\033[0m (if supported by your terminal)\n" +
		"• \033[9mStrikethrough text\033[0m\n" +
		"• Mixed: \033[1;4mBold + Underline\033[0m\n" +
		"• Hyperlink (OSC 8): \x1b]8;;https://github.com/box-cli-maker/box-cli-maker\x07box-cli-maker repo\x1b]8;;\x07"

	out, err := b.Render(title, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
