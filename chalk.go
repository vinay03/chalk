package chalk

import (
	"fmt"
	"strings"
)

type PrintableText interface {
	String() string
}

type Color struct {
	color   int
	style   int
	bgcolor int
	text    []interface{}
}

func (c Color) _formatting() string {
	var parts []string
	var ext string

	if c.bgcolor > 0 {
		parts = append(parts, fmt.Sprintf("%d", c.bgcolor))
	}
	if c.color > 0 {
		parts = append(parts, fmt.Sprintf("%d", c.color))
	}
	if c.style > 0 {
		parts = append(parts, fmt.Sprintf("%d", c.style))
	}

	ext = strings.Join(parts[:], ";")
	return "\u001b[" + ext + "m"
}

func (c Color) Print(a ...any) (int, error) {
	items := append([]any{c._formatting()}, a...)
	items = append(items, []any{Reset()}...)
	return fmt.Print(items...)
}
func (c Color) Printf(str string, a ...any) (int, error) {
	text := strings.Join([]string{
		c._formatting(),
		str,
		Reset(),
	}, "")
	return fmt.Printf(text, a...)
}
func (c Color) Println(a ...any) (int, error) {
	items := append([]any{c._formatting()}, a...)
	items = append(items, []any{Reset()}...)
	return fmt.Println(items...)
}

func (c Color) Sprint(a ...any) string {
	items := append([]any{c._formatting()}, a...)
	items = append(items, []any{Reset()}...)
	return fmt.Sprint(items...)
}
func (c Color) Sprintf(str string, a ...any) string {
	text := strings.Join([]string{
		c._formatting(),
		str,
		Reset(),
	}, "")
	return fmt.Sprintf(text, a...)
}
func (c Color) Sprintln(a ...any) string {
	items := append([]any{c._formatting()}, a...)
	items = append(items, []any{Reset()}...)
	return fmt.Sprintln(items...)
}

func (c Color) String() string {
	var ext string
	var parts []string
	fmt.Println()

	if c.bgcolor > 0 {
		parts = append(parts, fmt.Sprintf("%d", c.bgcolor))
	}
	if c.color > 0 {
		parts = append(parts, fmt.Sprintf("%d", c.color))
	}
	if c.style > 0 {
		parts = append(parts, fmt.Sprintf("%d", c.style))
	}

	ext = strings.Join(parts[:], ";")

	returnString := "\u001b[" + ext + "m"

	if len(c.text) > 0 {
		for _, item := range c.text {
			switch f := item.(type) {
			default:
				returnString += fmt.Sprint(f)
			}
		}
		returnString += Reset()
	}

	return returnString
}

func (c Color) Apply(msgs ...interface{}) Color {
	c.SetText(msgs)
	return c
}

func (c *Color) SetText(msgs []interface{}) *Color {
	if len(msgs) > 0 {
		if len(c.text) == 0 {
			c.text = msgs
		} else {
			c.text = append(c.text, msgs...)
		}
	}
	return c
}

// Colors
func (c *Color) Black(msgs ...interface{}) *Color {
	c.color = 30
	c.SetText(msgs)
	return c
}
func (c *Color) Red(msgs ...interface{}) *Color {
	c.color = 31
	c.SetText(msgs)
	return c
}
func (c *Color) Green(msgs ...interface{}) *Color {
	c.color = 32
	c.SetText(msgs)
	return c
}
func (c *Color) Yellow(msgs ...interface{}) *Color {
	c.color = 33
	c.SetText(msgs)
	return c
}
func (c *Color) Blue(msgs ...interface{}) *Color {
	c.color = 34
	c.SetText(msgs)
	return c
}
func (c *Color) Magenta(msgs ...interface{}) *Color {
	c.color = 35
	c.SetText(msgs)
	return c
}
func (c *Color) Cyan(msgs ...interface{}) *Color {
	c.color = 36
	c.SetText(msgs)
	return c
}
func (c *Color) White(msgs ...interface{}) *Color {
	c.color = 37
	c.SetText(msgs)
	return c
}

func (c *Color) BlackLight(msgs ...interface{}) *Color {
	c.color = 90
	c.SetText(msgs)
	return c
}
func (c *Color) RedLight(msgs ...interface{}) *Color {
	c.color = 91
	c.SetText(msgs)
	return c
}
func (c *Color) GreenLight(msgs ...interface{}) *Color {
	c.color = 92
	c.SetText(msgs)
	return c
}
func (c *Color) YellowLight(msgs ...interface{}) *Color {
	c.color = 93
	c.SetText(msgs)
	return c
}
func (c *Color) BlueLight(msgs ...interface{}) *Color {
	c.color = 94
	c.SetText(msgs)
	return c
}
func (c *Color) MagentaLight(msgs ...interface{}) *Color {
	c.color = 95
	c.SetText(msgs)
	return c
}
func (c *Color) CyanLight(msgs ...interface{}) *Color {
	c.color = 96
	c.SetText(msgs)
	return c
}
func (c *Color) WhiteLight(msgs ...interface{}) *Color {
	c.color = 97
	c.SetText(msgs)
	return c
}

// Text Styles
func (c *Color) Bold(msgs ...interface{}) *Color {
	c.style = 1
	c.SetText(msgs)
	return c
}
func (c *Color) Dim(msgs ...interface{}) *Color {
	c.style = 2
	c.SetText(msgs)
	return c
}
func (c *Color) Italic(msgs ...interface{}) *Color {
	c.style = 3
	c.SetText(msgs)
	return c
}
func (c *Color) Underline(msgs ...interface{}) *Color {
	c.style = 4
	c.SetText(msgs)
	return c
}
func (c *Color) Inverse(msgs ...interface{}) *Color {
	c.style = 7
	c.SetText(msgs)
	return c
}
func (c *Color) Hidden(msgs ...interface{}) *Color {
	c.style = 8
	c.SetText(msgs)
	return c
}
func (c *Color) Strikethrough(msgs ...interface{}) *Color {
	c.style = 9
	c.SetText(msgs)
	return c
}

// Background Colors
func (c *Color) BgBlack(msgs ...interface{}) *Color {
	c.bgcolor = 40
	c.SetText(msgs)
	return c
}
func (c *Color) BgRed(msgs ...interface{}) *Color {
	c.bgcolor = 41
	c.SetText(msgs)
	return c
}
func (c *Color) BgGreen(msgs ...interface{}) *Color {
	c.bgcolor = 42
	c.SetText(msgs)
	return c
}
func (c *Color) BgYellow(msgs ...interface{}) *Color {
	c.bgcolor = 43
	c.SetText(msgs)
	return c
}
func (c *Color) BgBlue(msgs ...interface{}) *Color {
	c.bgcolor = 44
	c.SetText(msgs)
	return c
}
func (c *Color) BgMagenta(msgs ...interface{}) *Color {
	c.bgcolor = 45
	c.SetText(msgs)
	return c
}
func (c *Color) BgCyan(msgs ...interface{}) *Color {
	c.bgcolor = 46
	c.SetText(msgs)
	return c
}
func (c *Color) BgWhite(msgs ...interface{}) *Color {
	c.bgcolor = 47
	c.SetText(msgs)
	return c
}

func (c *Color) BgBlackLight(msgs ...interface{}) *Color {
	c.bgcolor = 100
	c.SetText(msgs)
	return c
}
func (c *Color) BgRedLight(msgs ...interface{}) *Color {
	c.bgcolor = 101
	c.SetText(msgs)
	return c
}
func (c *Color) BgGreenLight(msgs ...interface{}) *Color {
	c.bgcolor = 102
	c.SetText(msgs)
	return c
}
func (c *Color) BgYellowLight(msgs ...interface{}) *Color {
	c.bgcolor = 103
	c.SetText(msgs)
	return c
}
func (c *Color) BgBlueLight(msgs ...interface{}) *Color {
	c.bgcolor = 104
	c.SetText(msgs)
	return c
}
func (c *Color) BgMagentaLight(msgs ...interface{}) *Color {
	c.bgcolor = 105
	c.SetText(msgs)
	return c
}
func (c *Color) BgCyanLight(msgs ...interface{}) *Color {
	c.bgcolor = 106
	c.SetText(msgs)
	return c
}
func (c *Color) BgWhiteLight(msgs ...interface{}) *Color {
	c.bgcolor = 107
	c.SetText(msgs)
	return c
}

// For all reset. BackgroundColor and TextColor
func Reset() string {
	return fmt.Sprintf("\033[%dm", 0)
}

func newColor() *Color {
	var c Color
	c.color = 0
	c.style = 0
	c.bgcolor = 0
	return &c
}

var (
	// Text Colors
	Black   = func(msgs ...interface{}) *Color { return newColor().Black(msgs...) }
	Red     = func(msgs ...interface{}) *Color { return newColor().Red(msgs...) }
	Green   = func(msgs ...interface{}) *Color { return newColor().Green(msgs...) }
	Yellow  = func(msgs ...interface{}) *Color { return newColor().Yellow(msgs...) }
	Blue    = func(msgs ...interface{}) *Color { return newColor().Blue(msgs...) }
	Magenta = func(msgs ...interface{}) *Color { return newColor().Magenta(msgs...) }
	Cyan    = func(msgs ...interface{}) *Color { return newColor().Cyan(msgs...) }
	White   = func(msgs ...interface{}) *Color { return newColor().White(msgs...) }

	// Text Light Colors
	BlackLight   = func(msgs ...interface{}) *Color { return newColor().BlackLight(msgs...) }
	RedLight     = func(msgs ...interface{}) *Color { return newColor().RedLight(msgs...) }
	GreenLight   = func(msgs ...interface{}) *Color { return newColor().GreenLight(msgs...) }
	YellowLight  = func(msgs ...interface{}) *Color { return newColor().YellowLight(msgs...) }
	BlueLight    = func(msgs ...interface{}) *Color { return newColor().BlueLight(msgs...) }
	MagentaLight = func(msgs ...interface{}) *Color { return newColor().MagentaLight(msgs...) }
	CyanLight    = func(msgs ...interface{}) *Color { return newColor().CyanLight(msgs...) }
	WhiteLight   = func(msgs ...interface{}) *Color { return newColor().WhiteLight(msgs...) }

	// Text Styles
	Bold          = func(msgs ...interface{}) *Color { return newColor().Bold(msgs...) }
	Dim           = func(msgs ...interface{}) *Color { return newColor().Dim(msgs...) }
	Italic        = func(msgs ...interface{}) *Color { return newColor().Italic(msgs...) }
	Underline     = func(msgs ...interface{}) *Color { return newColor().Underline(msgs...) }
	Inverse       = func(msgs ...interface{}) *Color { return newColor().Inverse(msgs...) }
	Hidden        = func(msgs ...interface{}) *Color { return newColor().Hidden(msgs...) }
	Strikethrough = func(msgs ...interface{}) *Color { return newColor().Strikethrough(msgs...) }

	// Background Colors
	BgBlack   = func(msgs ...interface{}) *Color { return newColor().BgBlack(msgs...) }
	BgRed     = func(msgs ...interface{}) *Color { return newColor().BgRed(msgs...) }
	BgGreen   = func(msgs ...interface{}) *Color { return newColor().BgGreen(msgs...) }
	BgYellow  = func(msgs ...interface{}) *Color { return newColor().BgYellow(msgs...) }
	BgBlue    = func(msgs ...interface{}) *Color { return newColor().BgBlue(msgs...) }
	BgMagenta = func(msgs ...interface{}) *Color { return newColor().BgMagenta(msgs...) }
	BgCyan    = func(msgs ...interface{}) *Color { return newColor().BgCyan(msgs...) }
	BgWhite   = func(msgs ...interface{}) *Color { return newColor().BgWhite(msgs...) }

	BgBlackLight   = func(msgs ...interface{}) *Color { return newColor().BgBlackLight(msgs...) }
	BgRedLight     = func(msgs ...interface{}) *Color { return newColor().BgRedLight(msgs...) }
	BgGreenLight   = func(msgs ...interface{}) *Color { return newColor().BgGreenLight(msgs...) }
	BgYellowLight  = func(msgs ...interface{}) *Color { return newColor().BgYellowLight(msgs...) }
	BgBlueLight    = func(msgs ...interface{}) *Color { return newColor().BgBlueLight(msgs...) }
	BgMagentaLight = func(msgs ...interface{}) *Color { return newColor().BgMagentaLight(msgs...) }
	BgCyanLight    = func(msgs ...interface{}) *Color { return newColor().BgCyanLight(msgs...) }
	BgWhiteLight   = func(msgs ...interface{}) *Color { return newColor().BgWhiteLight(msgs...) }
)
