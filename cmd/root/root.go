package root

import (
	"os"
	"strings"

	"github.com/bilalcaliskan/s3-substring-finder/cmd/root/options"
	"github.com/rs/zerolog"

	"github.com/dimiro1/banner"

	"github.com/bilalcaliskan/s3-substring-finder/internal/version"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bilalcaliskan/s3-substring-finder/internal/aws"
	"github.com/bilalcaliskan/s3-substring-finder/internal/logging"
	"github.com/spf13/cobra"
)

func init() {
	opts = options.GetRootOptions()
	opts.InitFlags(rootCmd)

	if err := opts.SetAccessCredentialsFromEnv(rootCmd); err != nil {
		panic(err)
	}
}

var (
	opts   *options.RootOptions
	ver    = version.Get()
	logger zerolog.Logger
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:     "s3-substring-finder",
		Short:   "Substring finder in files on a S3 bucket",
		Version: ver.GitVersion,
		Long:    `This tool searches the specific substring in files on AWS S3 and returns the file names`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat("build/ci/banner.txt"); err == nil {
				bannerBytes, _ := os.ReadFile("build/ci/banner.txt")
				banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
			}

			if opts.VerboseLog {
				logging.EnableDebugLogging()
			}

			logger = logging.GetLogger(opts)
			logger.Info().
				Str("appVersion", ver.GitVersion).
				Str("goVersion", ver.GoVersion).
				Str("goOS", ver.GoOs).
				Str("goArch", ver.GoArch).
				Str("gitCommit", ver.GitCommit).
				Str("buildDate", ver.BuildDate).
				Msg("s3-substring-finder is started!")

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			sess, err := aws.CreateSession(opts)
			if err != nil {
				logger.Error().
					Str("error", err.Error()).
					Msg("an error occurred while creating session")
				return err
			}

			// obtain S3 client with initialized session
			svc := s3.New(sess)

			logger.Debug().
				Str("fileExtensions", opts.FileExtensions).
				Msg("trying to find files on target bucket")
			matchedFiles, errors := aws.Find(svc, opts, logger)
			if len(errors) != 0 {
				logger.Error().Str("error", err.Error()).Msg("an error occurred while finding target files on target bucket")
				return err
			}

			if len(matchedFiles) == 0 {
				logger.Info().
					Any("matchedFiles", matchedFiles).
					Str("substring", opts.Substring).
					Msg("no matched files on the bucket")
				return nil
			}

			logger.Info().
				Any("matchedFiles", matchedFiles).
				Str("substring", opts.Substring).
				Msg("fetched matching files")
			return nil
		},
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
