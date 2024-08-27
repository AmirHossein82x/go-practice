package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger() *zerolog.Logger {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND| os.O_WRONLY, 0666)
	if err != nil {
		log.Print(
			"no such file",
		)
		return nil
	} else {
		logger := zerolog.New(file).With().Str("appName", "myApp").Timestamp().Logger()
		return &logger
	}

}
