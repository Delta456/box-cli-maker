package box

import (
	"fmt"
	"testing"
)

func TestBoxCheckColorType(t *testing.T) {
	tests := []struct {
		name string
		b    Box
	}{
		{
			"string Color Check", Box{Config: Config{Color: "Green"}},
		},
		{
			"string HiColor Check", Box{Config: Config{Color: "HiGreen"}},
		},
		{
			"string Invalid Color Check", Box{Config: Config{Color: "Any"}},
		},
		{
			"uint Color Check", Box{Config: Config{Color: uint(0x345ED3)}},
		},
		{
			"Simple [3]uint Color Check", Box{Config: Config{Color: [3]uint{230, 45, 233}}},
		},
	}
	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			switch tests[i].b.Config.Color.(type) {
			case uint:
			case string:
			case [3]uint:
			default:
				t.Errorf("Expected string, uint and [3]uint not %v", tests[i].b.Config.Color)
			}
			top, bottom := tests[i].b.checkColorType("top", "bottom", "")
			fmt.Println(top, bottom)

			bar := tests[i].b.obtainBoxColor()
			fmt.Println(bar)

		})
	}

	panicTest := struct {
		name string
		b    Box
	}{
		"Unknown type panic", Box{Config: Config{Color: byte(23)}},
	}
	t.Run(panicTest.name, func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Not expected string, uint, [3]uint, but got %v", panicTest.b.Config.Color)
			}
		}()
		_, _ = panicTest.b.checkColorType("top", "bottom", "")
		_ = panicTest.b.obtainBoxColor()
	})
}

func TestFindAlignment(t *testing.T) {
	tests := []struct {
		name string
		b    Box
		want string
	}{
		{
			"b.findAlignment() Default Check", Box{Config: Config{ContentAlign: ""}}, leftAlign,
		},
		{
			"b.findAlignment() Centre Check", Box{Config: Config{ContentAlign: "Center"}}, centerAlign,
		},
		{
			"b.findAlignment() Left Check", Box{Config: Config{ContentAlign: "Left"}}, leftAlign,
		},
		{
			"b.findAlignment() Right Check", Box{Config: Config{ContentAlign: "Right"}}, rightAlign,
		},
		{
			"b.findAlignment() Invalid Check", Box{Config: Config{ContentAlign: "Invalid"}}, leftAlign,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.b.findAlign()
			if want != tt.want {
				t.Errorf("expected %v but got %v", tt.want, want)
			}
		})
	}
}
