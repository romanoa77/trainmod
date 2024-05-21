package simplelogger

import (
	"os"

	"github.com/rs/zerolog"
)

type LogWrite struct {
	Level   string `json:"level"`
	Id      int    `json:"Id"`
	Size    int    `json:"size"`
	Name    string `json:"name"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

func LogWriteFile(dest string, filename string, i int, size int, fname string) {
	file, err := os.OpenFile(
		dest+filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	logger := zerolog.New(file).With().Timestamp().Logger()

	logger.Info().Int("id", i).Int("size", size).Str("name", fname).Msg("Write file")
}
