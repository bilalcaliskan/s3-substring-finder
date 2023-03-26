package logging

import (
	"os"

	"github.com/bilalcaliskan/s3-substring-finder/cmd/root/options"

	"github.com/rs/zerolog"
)

var (
	logger zerolog.Logger
	Level  = zerolog.InfoLevel
)

func init() {
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger = zerolog.New(consoleWriter).With().Timestamp().Logger().Level(Level)
}

func GetLogger(opts *options.RootOptions) zerolog.Logger {
	logger = logger.With().
		Str("bucketName", opts.BucketName).
		Str("region", opts.Region).
		Logger()

	return logger
}

func EnableDebugLogging() {
	logger = logger.Level(zerolog.DebugLevel)
}
