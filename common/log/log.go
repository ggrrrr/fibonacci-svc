package log

import (
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func Debug() *zerolog.Event {
	return log.Debug()
}

func Info() *zerolog.Event {
	return log.Info()
}
func Warn() *zerolog.Event {
	return log.Warn()
}

func Error(err error) *zerolog.Event {
	return log.Error().Err(err)
}

func init() {
	out := zerolog.NewConsoleWriter()
	out.NoColor = true
	l := zerolog.New(out).Level(zerolog.TraceLevel)
	log = l
}
