package box

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInbuiltStyles(t *testing.T) {
	cases := map[string]string{
		"Single":        `"┌────────────────────────────────────────┐\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│             Box CLI Maker              │\n│                                        │\n│  Highly Customized Terminal Box Maker  │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n│                                        │\n└────────────────────────────────────────┘\n"`,
		"Single Double": `"╓────────────────────────────────────────╖\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║             Box CLI Maker              ║\n║                                        ║\n║  Highly Customized Terminal Box Maker  ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n║                                        ║\n╙────────────────────────────────────────╜\n"`,
		/*"Double":        ``,
		"Double Single": ``,
		"Bold":          ``,
		"Round":   "",
		"Hidden":  "",
		"Classic": "",*/
	}
	for key := range cases {
		Box := New(Config{Px: 2, Py: 5, Type: key})
		box := Box.String("Box CLI Maker", "Highly Customized Terminal Box Maker")
		str, _ := json.Marshal(box)
		if cases[key] != string(str) {
			t.Fatal(fmt.Sprintf(key, "Style", cases[key], "and", string(str), "are not same"))
		}
	}

}
