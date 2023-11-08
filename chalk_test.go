package chalk

import (
	"encoding/hex"
	"testing"
)

func Test_BasicColors(t *testing.T) {

	var colorTests = []struct {
		Expected  string
		Actual    string
		ColorName string
	}{
		{"\u001b[30m", Black().String(), "Black"},
		{"\u001b[31m", Red().String(), "Red"},
		{"\u001b[32m", Green().String(), "Green"},
		{"\u001b[33m", Yellow().String(), "Yellow"},
		{"\u001b[34m", Blue().String(), "Blue"},
		{"\u001b[35m", Magenta().String(), "Magenta"},
		{"\u001b[36m", Cyan().String(), "Cyan"},
		{"\u001b[37m", White().String(), "White"},
		{"\u001b[90m", BlackLight().String(), "BlackLight"},
		{"\u001b[91m", RedLight().String(), "RedLight"},
		{"\u001b[92m", GreenLight().String(), "GreenLight"},
		{"\u001b[93m", YellowLight().String(), "YellowLight"},
		{"\u001b[94m", BlueLight().String(), "BlueLight"},
		{"\u001b[95m", MagentaLight().String(), "MagentaLight"},
		{"\u001b[96m", CyanLight().String(), "CyanLight"},
		{"\u001b[97m", WhiteLight().String(), "WhiteLight"},
	}

	for _, test := range colorTests {
		if test.Actual != test.Expected {
			t.Error("Expected modifier is not same as actual string for Color : " + test.ColorName + ", Actual: " + hex.EncodeToString([]byte(test.Actual)) + ", Expected: " + hex.EncodeToString([]byte(test.Expected)))
		}
	}
}

func Test_TextStyles(t *testing.T) {
	var styleTests = []struct {
		Expected  string
		Actual    string
		StyleName string
	}{
		{"\u001b[1m", Bold().String(), "Bold"},
		{"\u001b[2m", Dim().String(), "Dim"},
		{"\u001b[3m", Italic().String(), "Italic"},
		{"\u001b[4m", Underline().String(), "Underline"},
		{"\u001b[7m", Inverse().String(), "Inverse"},
		{"\u001b[8m", Hidden().String(), "Hidden"},
		{"\u001b[9m", Strikethrough().String(), "Strikethrough"},
	}
	for _, test := range styleTests {
		if test.Actual != test.Expected {
			t.Error("Expected modifier is not same as actual string for Text Style : " + test.StyleName)
		}
	}
}

func Test_BackgroundColors(t *testing.T) {
	var backgroundColorTests = []struct {
		Expected  string
		Actual    string
		ColorName string
	}{
		{"\u001b[40m", BgBlack().String(), "BgBlack"},
		{"\u001b[41m", BgRed().String(), "BgRed"},
		{"\u001b[42m", BgGreen().String(), "BgGreen"},
		{"\u001b[43m", BgYellow().String(), "BgYellow"},
		{"\u001b[44m", BgBlue().String(), "BgBlue"},
		{"\u001b[45m", BgMagenta().String(), "BgMagenta"},
		{"\u001b[46m", BgCyan().String(), "BgCyan"},
		{"\u001b[47m", BgWhite().String(), "BgWhite"},
		{"\u001b[100m", BgBlackLight().String(), "BgBlackLight"},
		{"\u001b[101m", BgRedLight().String(), "BgRedLight"},
		{"\u001b[102m", BgGreenLight().String(), "BgGreenLight"},
		{"\u001b[103m", BgYellowLight().String(), "BgYellowLight"},
		{"\u001b[104m", BgBlueLight().String(), "BgBlueLight"},
		{"\u001b[105m", BgMagentaLight().String(), "BgMagentaLight"},
		{"\u001b[106m", BgCyanLight().String(), "BgCyanLight"},
		{"\u001b[107m", BgWhiteLight().String(), "BgWhiteLight"},
	}

	for _, test := range backgroundColorTests {
		if test.Expected != test.Actual {
			t.Error("Expected modifier is not same as actual string for Background Color : " + test.ColorName)
		}
	}
}

func Test_ResetSequence(t *testing.T) {
	expected := "\u001b[0m"
	actual := Reset()

	if expected != actual {
		t.Error("Expected modifier is n ot same as actual string for Reset style purpose")
	}
}

func Test_MixedScenarios(t *testing.T) {

	var mixedTests = []struct {
		Expected string
		Actual   string
		ID       string
	}{
		{"\u001b[107;31;1m", Red().Bold().BgWhiteLight().String(), "Red().Bold().BgWhiteLight()"},
		{"\u001b[47;35;3m", Italic().Magenta().BgWhite().String(), "Italic().Magenta().BgWhite()"},
		{"\u001b[34;4m", Blue().Underline().String(), "Blue().Underline()"},
		{"\u001b[107;30m", Black().BgWhiteLight().String(), "Black().BgWhiteLight()"},
	}

	for _, test := range mixedTests {
		if test.Expected != test.Actual {
			t.Error("Expected string is not same as actual string for : " + test.ID)
		}
	}

}
