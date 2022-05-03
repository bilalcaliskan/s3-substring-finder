package cmd

import (
	"os"
	"s3-substring-finder/internal/aws"
	"s3-substring-finder/internal/logging"
	"s3-substring-finder/internal/options"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()
	opts := options.GetS3SubstringFinderOptions()
	rootCmd.PersistentFlags().StringVarP(&opts.AccessKey, "accessKey", "", "",
		"access key to access S3")
	rootCmd.PersistentFlags().StringVarP(&opts.SecretKey, "secretKey", "", "",
		"secret key to access S3")
	rootCmd.PersistentFlags().StringVarP(&opts.BucketName, "bucketName", "", "",
		"name of the target bucket")
	rootCmd.PersistentFlags().StringVarP(&opts.Region, "region", "", "",
		"region of the target bucket")
	rootCmd.PersistentFlags().StringVarP(&opts.Substring, "substring", "", "",
		"substring to find on target bucket")

	// set required flags
	_ = rootCmd.MarkPersistentFlagRequired("accessKey")
	_ = rootCmd.MarkPersistentFlagRequired("secretKey")
	_ = rootCmd.MarkPersistentFlagRequired("bucketName")
	_ = rootCmd.MarkPersistentFlagRequired("region")
	_ = rootCmd.MarkPersistentFlagRequired("substring")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "s3-substring-finder",
	Short: "Substring finder in files on a S3 bucket",
	Long:  `This tool searches the specific substring in files on AWS S3 and returns the file names`,
	Run: func(cmd *cobra.Command, args []string) {
		opts := options.GetS3SubstringFinderOptions()
		matchedFiles, errors := aws.Find(opts)
		if len(errors) != 0 {
			logger.Fatal("fatal error occured", zap.Errors("errors", errors))
		}
		logger.Info("fetched matched files", zap.Any("matchedFiles", matchedFiles))
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
