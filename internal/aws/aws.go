package aws

import (
	"s3-substring-finder/internal/logging"
	"s3-substring-finder/internal/options"

	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()
}

// Find does the heavy lifting, communicates with the S3 and finds the files
func Find(opts *options.S3SubstringFinderOptions) error {
	logger.Info("", zap.Any("opts", opts))
	return nil
}
