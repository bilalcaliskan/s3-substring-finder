package cmd

import (
	"os"
	"strings"

	"github.com/dimiro1/banner"

	"github.com/bilalcaliskan/s3-substring-finder/internal/version"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/bilalcaliskan/s3-substring-finder/internal/aws"
	"github.com/bilalcaliskan/s3-substring-finder/internal/logging"
	"github.com/bilalcaliskan/s3-substring-finder/internal/options"
	"github.com/spf13/cobra"

	"go.uber.org/zap"
)

var (
	opts   *options.S3SubstringFinderOptions
	logger *zap.Logger
	ver    = version.Get()
)

func init() {
	opts = options.GetS3SubstringFinderOptions()
	logger = logging.GetLogger()
	rootCmd.Flags().StringVarP(&opts.BucketName, "bucketName", "", "",
		"name of the target bucket on S3")
	rootCmd.Flags().StringVarP(&opts.AccessKey, "accessKey", "", "",
		"access key credential to access S3 bucket")
	rootCmd.Flags().StringVarP(&opts.SecretKey, "secretKey", "", "",
		"secret key credential to access S3 bucket")
	rootCmd.Flags().StringVarP(&opts.Region, "region", "", "",
		"region of the target bucket on S3")
	rootCmd.Flags().StringVarP(&opts.Substring, "substring", "", "",
		"substring to find on txt files on target bucket")
	rootCmd.Flags().StringVarP(&opts.FileExtensions, "fileExtensions", "", "txt",
		"comma separated list of file extensions to search on S3 bucket")
	rootCmd.Flags().BoolVarP(&opts.VerboseLog, "verbose", "v", false,
		"verbose output of the logging library (default false)")

	if err := opts.SetAccessCredentialsFromEnv(rootCmd); err != nil {
		panic(err)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "s3-substring-finder",
	Short:   "Substring finder in files on a S3 bucket",
	Version: ver.GitVersion,
	Long:    `This tool searches the specific substring in files on AWS S3 and returns the file names`,
	Run: func(cmd *cobra.Command, args []string) {
		if opts.VerboseLog {
			logging.Atomic.SetLevel(zap.DebugLevel)
		}

		if _, err := os.Stat("build/ci/banner.txt"); err == nil {
			bannerBytes, _ := os.ReadFile("build/ci/banner.txt")
			banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
		}

		logger.Info("s3-substring-finder is started",
			zap.String("appVersion", ver.GitVersion),
			zap.String("goVersion", ver.GoVersion),
			zap.String("goOS", ver.GoOs),
			zap.String("goArch", ver.GoArch),
			zap.String("gitCommit", ver.GitCommit),
			zap.String("buildDate", ver.BuildDate))

		sess, err := aws.CreateSession(opts)
		if err != nil {
			logger.Fatal("fatal error occurred", zap.Error(err))
		}

		// obtain S3 client with initialized session
		svc := s3.New(sess)

		logger.Debug("trying to find files on bucket", zap.String("fileExtensions", opts.FileExtensions),
			zap.String("bucketName", opts.BucketName), zap.String("region", opts.Region))
		matchedFiles, errors := aws.Find(svc, opts)
		if len(errors) != 0 {
			logger.Fatal("fatal error occurred", zap.Errors("errors", errors))
		}

		if len(matchedFiles) == 0 {
			logger.Info("no matched files on the bucket", zap.Any("matchedFiles", matchedFiles),
				zap.String("bucket", opts.BucketName), zap.String("region", opts.Region),
				zap.String("substring", opts.Substring))
			return
		}

		logger.Info("fetched matched files", zap.Any("matchedFiles", matchedFiles),
			zap.String("bucket", opts.BucketName), zap.String("region", opts.Region),
			zap.String("substring", opts.Substring))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
