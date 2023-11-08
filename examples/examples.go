package main

import (
	"fmt"

	"github.com/vinay03/chalk"
)

func main() {
	// Colors
	fmt.Println("\n---> Colors")
	chalk.BlackLight().Println("BlackLight-Text")
	chalk.Black().Println("Black-Text")
	chalk.Red().Println("Red-Text")
	chalk.RedLight().Println("RedLight-Text")
	chalk.Green().Println("Green-Text")
	chalk.GreenLight().Println("GreenLight-Text")
	chalk.Yellow().Println("Yellow-Text")
	chalk.YellowLight().Println("YellowLight-Text")
	chalk.Blue().Println("Blue-Text")
	chalk.BlueLight().Println("BlueLight-Text")
	chalk.Magenta().Println("Magenta-Text")
	chalk.MagentaLight().Println("MagentaLight-Text")
	chalk.Cyan().Println("Cyan-Text")
	chalk.CyanLight().Println("CyanLight-Text")
	chalk.White().Println("White-Text")
	chalk.WhiteLight().Println("WhiteLight-Text")

	// Style
	fmt.Println("\n---> Text Style")
	chalk.Bold().Println("Bold-Text")
	chalk.Dim().Println("Dim-Text")
	chalk.Italic().Println("Italic-Text")
	chalk.Underline().Println("Underline-Text")
	chalk.Inverse().Println("Inverse-Text")
	chalk.Hidden().Println("Hidden-Text")
	chalk.Strikethrough().Println("Strikethrough-Text")

	// Background Colors
	fmt.Println("\n---> Background Colors")
	chalk.WhiteLight().BgBlack().Println("WhiteLight-Text-on-BgBlack")
	chalk.WhiteLight().BgBlackLight().Println("WhiteLight-Text-on-BgBlackLight")
	chalk.Black().BgRed().Println("Black-Text-on-BgRed")
	chalk.Black().BgRedLight().Println("Black-Text-on-BgRedLight")
	chalk.Black().BgGreen().Println("Black-Text-on-BgGreen")
	chalk.Black().BgGreenLight().Println("Black-Text-on-BgGreenLight")
	chalk.Black().BgYellow().Println("Black-Text-on-BgYellow")
	chalk.Black().BgYellowLight().Println("Black-Text-on-BgYellowLight")
	chalk.Black().BgBlue().Println("Black-Text-on-BgBlue")
	chalk.Black().BgBlueLight().Println("Black-Text-on-BgBlueLight")
	chalk.Black().BgMagenta().Println("Black-Text-on-BgMagenta")
	chalk.Black().BgMagentaLight().Println("Black-Text-on-BgMagentaLight")
	chalk.Black().BgCyan().Println("Black-Text-on-BgCyan")
	chalk.Black().BgCyanLight().Println("Black-Text-on-BgCyanLight")
	chalk.Black().BgWhite().Println("Black-Text-on-BgWhite")
	chalk.Black().BgWhiteLight().Println("Black-Text-on-BgWhiteLight")

	// Mixed
	fmt.Println("\n---> Mixed")
	chalk.Red().Italic().Println("Red-Italic-Text")
	fmt.Println("Plain Text")
	chalk.Green().Strikethrough().Println("Green-Strikethrough-Text")
	chalk.Cyan().Underline().BgBlackLight().Println("Cyan-underline-text-on-BgBlackLight")
	chalk.Yellow().BgRed().Inverse().Println("Yellow-text-on-red-inverted")

	// Reusable Formatting
	fmt.Println("\n---> Reusable Formatting")
	SuccessFormatting := chalk.Green()
	SuccessFormatting.Println("Completed Successfully")
	ErrorFormatting := chalk.Red()
	ErrorFormatting.Println("Failed to execute. Try again!")

	// Reusable formatted strings
	fmt.Println("\n---> Reusable Pre-Formatted Strings")
	SuccessMessagePrefix := chalk.Green().Bold().Sprint("SUCCESS :")
	fmt.Println(SuccessMessagePrefix, "Process completed successfully.")

	WarningMessagePrefix := chalk.YellowLight().Bold().Sprint("WARNING :")
	fmt.Println(WarningMessagePrefix, "Call Deprecated")

	ErrorMessagePrefix := chalk.RedLight().Bold().Sprint("ERROR :")
	fmt.Println(ErrorMessagePrefix, "Fatal error ocurred")

	// Wrap formatting around Pre-existing code
	fmt.Println("\n---> Wrap formatting around Pre-existing code")
	fmt.Print(chalk.Green()) // Starts the formatting
	fmt.Println("All these lines ")
	fmt.Println("will be printed in ")
	fmt.Println("Green")
	fmt.Print(chalk.Reset()) // Resets all the formatting
}
