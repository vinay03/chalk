package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/vinay03/chalk"
)

var formattedLevels = map[string]string{
	"INFO":  chalk.BgBlue().Bold().White().Sprint("INFO "),
	"WARN":  chalk.BgYellowLight().Black().Sprint("WARN "),
	"ERROR": chalk.White().BgRed().Sprint("ERROR"),
	"DEBUG": chalk.Black().BgWhite().Sprint("DEBUG"),
	"TRACE": chalk.White().BgRed().Sprint("TRACE"),
	"PANIC": chalk.White().BgRed().Sprint("PANIC"),
	"FATAL": chalk.White().BgRed().Sprint("FATAL"),
}

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	output.FormatLevel = func(i interface{}) string {
		return formattedLevels[strings.ToUpper(fmt.Sprintf("%s", i))]
	}
	output.FormatTimestamp = func(i interface{}) string {
		return chalk.Green().Sprintf("%s", i)
	}

	log := zerolog.New(output).With().Timestamp().Logger()

	log.Info().Msg("This is an info level message")
	log.Debug().Msg("This is an debug level message")
	log.Warn().Msg("This is an warning level message")
	log.Error().Msg("This is an error level message")
	log.Panic().Msg("This is an panic level message")
}
