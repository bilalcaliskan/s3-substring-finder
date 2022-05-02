package cmd

import (
	"io/ioutil"
	"os"
	"s3-substring-finder/internal/aws"
	"s3-substring-finder/internal/logging"
	"s3-substring-finder/internal/options"
	"strings"

	"github.com/dimiro1/banner"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
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
		if err := aws.Find(opts); err != nil {
			logging.GetLogger().Fatal("fatal error occured", zap.Error(err))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
