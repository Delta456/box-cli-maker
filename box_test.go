package box

import (
	"fmt"
	"testing"
)

func TestInbuiltStyles(t *testing.T) {
	cases := map[string]string{
		"Single":        "┌────────────────────────────────────────┐\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│             Box CLI Maker              │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n└────────────────────────────────────────┘\n",
		"Single Double": "╓────────────────────────────────────────╖\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║             Box CLI Maker              ║\n║                                        ║\n║  Highly Customized Terminal Box Maker  ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n╙────────────────────────────────────────╜\n",
		"Double":        "╔════════════════════════════════════════╗\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║             Box CLI Maker              ║\n║                                        ║\n║  Highly Customized Terminal Box Maker  ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n╚════════════════════════════════════════╝\n",
		"Double Single": "╒════════════════════════════════════════╕\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│             Box CLI Maker              │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n╘════════════════════════════════════════╛\n",
		"Bold":          "┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃             Box CLI Maker              ┃\n┃                                        ┃\n┃  Highly Customized Terminal Box Maker  ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛\n",
		"Round":         "╭────────────────────────────────────────╮\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│             Box CLI Maker              │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n╰────────────────────────────────────────╯\n",
		"Hidden":        "+                                        +\n                                          \n                                          \n                                          \n                                          \n                                          \n              Box CLI Maker               \n                                          \n   Highly Customized Terminal Box Maker   \n                                          \n                                          \n                                          \n                                          \n                                          \n+                                        +\n",
		"Classic":       "+----------------------------------------+\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|             Box CLI Maker              |\n|                                        |\n|  Highly Customized Terminal Box Maker  |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n+----------------------------------------+\n",
	}
	for key := range cases {
		Box := New(Config{Px: 2, Py: 5, Type: key})
		box := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")
		if cases[key] != box {
			t.Fatalf(key, "Style", cases[key], "and", box, "are not same")
		}
	}
}

func TestPrintColorBox(t *testing.T) {
	StyleCases := []string{"Single", "Double", "Single Double", "Double Single", "Bold", "Round", "Hidden", "Classic"}
	ColorTypes := []string{"Black", "Blue", "Red", "Green", "Yellow", "Cyan", "Magenta", "White", "HiBlack", "HiBlue", "HiRed", "HiGreen", "HiYellow", "HiCyan", "HiMagenta", "HiWhite"}

	for i := 0; i < len(StyleCases); i++ {
		for j := 0; j < len(ColorTypes); j++ {
			Box := New(Config{Px: 2, Py: 5, Type: StyleCases[i], Color: ColorTypes[j]})
			fmt.Print(fmt.Sprint("Using ", StyleCases[i], " as Style and ", ColorTypes[j], " as Color:  "))
			Box.Println("Box CLI Maker", "Highly Customized Terminal Box Maker")
		}
	}
}

func TestTitlePos(t *testing.T) {
	cases := map[string]map[string]string{
		"Inside": {
			"Single":        "┌────────────────────────────────────────┐\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│             Box CLI Maker              │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n└────────────────────────────────────────┘\n",
			"Single Double": "╓────────────────────────────────────────╖\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║             Box CLI Maker              ║\n║                                        ║\n║  Highly Customized Terminal Box Maker  ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n╙────────────────────────────────────────╜\n",
			"Double":        "╔════════════════════════════════════════╗\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║             Box CLI Maker              ║\n║                                        ║\n║  Highly Customized Terminal Box Maker  ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n╚════════════════════════════════════════╝\n",
			"Double Single": "╒════════════════════════════════════════╕\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│             Box CLI Maker              │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n╘════════════════════════════════════════╛\n",
			"Bold":          "┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃             Box CLI Maker              ┃\n┃                                        ┃\n┃  Highly Customized Terminal Box Maker  ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛\n",
			"Round":         "╭────────────────────────────────────────╮\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│             Box CLI Maker              │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n╰────────────────────────────────────────╯\n",
			"Hidden":        "+                                        +\n                                          \n                                          \n                                          \n                                          \n                                          \n              Box CLI Maker               \n                                          \n   Highly Customized Terminal Box Maker   \n                                          \n                                          \n                                          \n                                          \n                                          \n+                                        +\n",
			"Classic":       "+----------------------------------------+\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|             Box CLI Maker              |\n|                                        |\n|  Highly Customized Terminal Box Maker  |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n+----------------------------------------+\n",
		},
		"Bottom": {
			"Single":        "┌────────────────────────────────────────┐\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n└ Box CLI Maker ─────────────────────────┘\n",
			"Single Double": "╓────────────────────────────────────────╖\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║  Highly Customized Terminal Box Maker  ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n╙ Box CLI Maker ─────────────────────────╜\n",
			"Double":        "╔════════════════════════════════════════╗\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║  Highly Customized Terminal Box Maker  ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n╚ Box CLI Maker ═════════════════════════╝\n",
			"Double Single": "╒════════════════════════════════════════╕\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n╘ Box CLI Maker ═════════════════════════╛\n",
			"Bold":          "┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃  Highly Customized Terminal Box Maker  ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┗ Box CLI Maker ━━━━━━━━━━━━━━━━━━━━━━━━━┛\n",
			"Round":         "╭────────────────────────────────────────╮\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n╰ Box CLI Maker ─────────────────────────╯\n",
			"Hidden":        "+                                        +\n                                          \n                                          \n                                          \n                                          \n                                          \n   Highly Customized Terminal Box Maker   \n                                          \n                                          \n                                          \n                                          \n                                          \n+ Box CLI Maker                          +\n",
			"Classic":       "+----------------------------------------+\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|  Highly Customized Terminal Box Maker  |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n+ Box CLI Maker -------------------------+\n",
		},
		"Top": {
			"Single":        "┌ Box CLI Maker ─────────────────────────┐\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n└────────────────────────────────────────┘\n",
			"Single Double": "╓ Box CLI Maker ─────────────────────────╖\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║  Highly Customized Terminal Box Maker  ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n╙────────────────────────────────────────╜\n",
			"Double":        "╔ Box CLI Maker ═════════════════════════╗\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║  Highly Customized Terminal Box Maker  ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n╚════════════════════════════════════════╝\n",
			"Double Single": "╒ Box CLI Maker ═════════════════════════╕\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n╘════════════════════════════════════════╛\n",
			"Bold":          "┏ Box CLI Maker ━━━━━━━━━━━━━━━━━━━━━━━━━┓\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃  Highly Customized Terminal Box Maker  ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┃                                        ┃\n┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛\n",
			"Round":         "╭ Box CLI Maker ─────────────────────────╮\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n╰────────────────────────────────────────╯\n",
			"Hidden":        "+ Box CLI Maker                          +\n                                          \n                                          \n                                          \n                                          \n                                          \n   Highly Customized Terminal Box Maker   \n                                          \n                                          \n                                          \n                                          \n                                          \n+                                        +\n",
			"Classic":       "+ Box CLI Maker -------------------------+\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|  Highly Customized Terminal Box Maker  |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n|                                        |\n+----------------------------------------+\n",
		},
	}
	for titlePos, val := range cases {
		for style := range val {
			Box := New(Config{Px: 2, Py: 5, Type: style, TitlePos: titlePos})
			box := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")
			if box != val[style] {
				t.Errorf("Using %s as Style and %s as TitlePos but %s and %s are not same", style, titlePos, box, val[style])
			}
		}
	}
}

func TestPrintMultiandTabLineString(t *testing.T) {
	StyleCases := []string{"Single", "Double", "Single Double", "Double Single", "Bold", "Round", "Hidden", "Classic"}
	ColorTypes := []string{"Black", "Blue", "Red", "Green", "Yellow", "Cyan", "Magenta", "HiBlack", "HiBlue", "HiRed", "HiGreen", "HiYellow", "HiCyan", "HiMagenta"}

	for i := 0; i < len(StyleCases); i++ {
		for j := 0; j < len(ColorTypes); j++ {
			Box := New(Config{Px: 2, Py: 5, Type: StyleCases[i], Color: ColorTypes[j], TitlePos: "Top", TitleColor: "Cyan", ContentColor: "Red"})
			fmt.Print(fmt.Sprint("Using ", StyleCases[i], " as Style and ", ColorTypes[j], " as Color: "))
			Box.Println("Box CLI Maker", `Make
			Highly
				Customized
					Terminal
							 Boxes`)
		}
	}
}

func TestUnicodeString(t *testing.T) {
	// English, Japanese, Chinese(Traditional), Korean, French, Spanish, Latin, Greek
	titles := []string{"Box CLI Maker", "ボックスメーカー", "盒子製造商", "박스 메이커", "Créateur de boîte CLI", "Fabricante de cajas", "Qui fecit me arca CLI", "Κουτί CLI Maker"}
	lines := []string{"Make Highly Customized Terminal Boxes", "高度にカスタマイズされた端子ボックスを作成する", "製作高度定制的接線盒", "고도로 맞춤화 된 터미널 박스 만들기", "Créez des boîtes à bornes hautement personnalisées", "Haga cajas de terminales altamente personalizadas", "Fac multum Customized Terminal Pyxidas", "Δημιουργήστε πολύ προσαρμοσμένα τερματικά κουτιά"}
	StyleCases := []string{"Single", "Double", "Single Double", "Double Single", "Bold", "Round", "Hidden", "Classic"}
	for i := 0; i < len(titles); i++ {
		for j := 0; j < len(StyleCases); j++ {
			Box := New(Config{Px: 2, Py: 5, Type: StyleCases[j]})
			Box.Println(titles[i], lines[i])
		}
	}
}

func TestBoxPrint(t *testing.T) {
	StyleCases := []string{"Single", "Double", "Single Double", "Double Single", "Bold", "Round", "Hidden", "Classic"}
	for i := 0; i < len(StyleCases); i++ {
		Box := New(Config{Px: 2, Py: 5, Type: StyleCases[i]})
		Box.Print("Box CLI Maker", "Highly Customized Terminal Box Maker")
	}
}

func TestTabWithColorBox(t *testing.T) {
	StyleCases := []string{"Single", "Double", "Single Double", "Double Single", "Bold", "Round", "Hidden", "Classic"}
	ColorTypes := []string{"Black", "Blue", "Red", "Green", "Yellow", "Cyan", "Magenta", "White", "HiBlack", "HiBlue", "HiRed", "HiGreen", "HiYellow", "HiCyan", "HiMagenta", "HiWhite"}

	for i := 0; i < len(StyleCases); i++ {
		for j := 0; j < len(ColorTypes); j++ {
			Box := New(Config{Px: 2, Py: 6, Type: StyleCases[i], Color: ColorTypes[j], ContentColor: "Cyan", TitleColor: [3]uint{215, 58, 74}, TitlePos: "Top"})
			fmt.Print(fmt.Sprint("Using ", StyleCases[i], " as Style and ", ColorTypes[j], " as Color:  "))
			Box.Println("Box 	CLI 	Maker 	📦", "Highly 		Customized 			Terminal	 Box	 Maker")
		}
	}
}

func TestBoxAlign(t *testing.T) {
	bx := New(Config{
		Px:           2,
		Py:           0,
		Type:         "Single",
		ContentAlign: "Left",
		Color:        "Green",
		TitlePos:     "Top",
		ContentColor: uint(0xa77032),
		TitleColor:   "Cyan",
	})
	bx.Print("System		Info", "LoremIpsum\nfoo bar hello world\n123456 abcdefghijk")

}

/* TODO: find a way to make term.GetSize() work in Testing functions.
func TestBoxWrapText(t *testing.T) {
	bx := New(Config{
		Px:            2,
		Py:            0,
		Type:          "Single",
		ContentAlign:  "Left",
		Color:         "Green",
		TitlePos:      "Top",
		ContentColor:  uint(0xa77032),
		TitleColor:    "Cyan",
		AllowWrapping: true,
	})
	//width, h, err := term.GetSize(int(2))
	//fmt.Println(strings.Repeat("a", width), width, h, err)
	bx.Println("title", strings.Repeat("	盒子製 造商,📦 ", 160))

}
*/
