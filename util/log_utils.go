package util

import (
	"os"

	"github.com/rs/zerolog"
)

func GetTextFileLogger(filepath string) (zerolog.Logger, error) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		return zerolog.Logger{}, err
	}

	return zerolog.New(file).With().Timestamp().Str("from", "auth").Logger(), nil
}
