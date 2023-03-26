package options

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootOptions = &RootOptions{}

// RootOptions contains frequent command line and application options.
type RootOptions struct {
	// AccessKey is the access key credentials for accessing AWS over client
	AccessKey string
	// SecretKey is the secret key credentials for accessing AWS over client
	SecretKey string
	// BucketName is the name of target bucket
	BucketName string
	// Region is the region of the target bucket
	Region string
	// Substring is the target string to find in a bucket
	Substring string
	// FileExtensions is a comma separated list of file extensions to search on S3 bucket (txt, json etc)
	FileExtensions string
	// VerboseLog is the verbosity of the logging library
	VerboseLog bool
}

// GetRootOptions returns the pointer of S3SubstringFinderOptions
func GetRootOptions() *RootOptions {
	return rootOptions
}

func (opts *RootOptions) SetAccessCredentialsFromEnv(cmd *cobra.Command) error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("aws")
	if err := viper.BindEnv("access_key", "secret_key", "bucket_name", "region"); err != nil {
		return err
	}

	if accessKey := viper.Get("access_key"); accessKey != nil {
		opts.AccessKey = fmt.Sprintf("%v", accessKey)
	} else {
		_ = cmd.MarkFlagRequired("accessKey")
	}

	if secretKey := viper.Get("secret_key"); secretKey != nil {
		opts.SecretKey = fmt.Sprintf("%v", secretKey)
	} else {
		_ = cmd.MarkFlagRequired("secretKey")
	}

	if bucketName := viper.Get("bucket_name"); bucketName != nil {
		opts.BucketName = fmt.Sprintf("%v", bucketName)
	} else {
		_ = cmd.MarkFlagRequired("bucketName")
	}

	if region := viper.Get("region"); region != nil {
		opts.Region = fmt.Sprintf("%v", region)
	} else {
		_ = cmd.MarkFlagRequired("region")
	}

	return nil
}

func (opts *RootOptions) InitFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&opts.BucketName, "bucketName", "", "",
		"name of the target bucket on S3 (default \"\")")
	cmd.Flags().StringVarP(&opts.AccessKey, "accessKey", "", "",
		"access key credential to access S3 bucket (default \"\")")
	cmd.Flags().StringVarP(&opts.SecretKey, "secretKey", "", "",
		"secret key credential to access S3 bucket (default \"\")")
	cmd.Flags().StringVarP(&opts.Region, "region", "", "",
		"region of the target bucket on S3 (default \"\")")
	cmd.Flags().StringVarP(&opts.Substring, "substring", "", "",
		"substring to find on txt files on target bucket (default \"\")")
	cmd.Flags().StringVarP(&opts.FileExtensions, "fileExtensions", "", "txt",
		"comma separated list of file extensions to search on S3 bucket")
	cmd.Flags().BoolVarP(&opts.VerboseLog, "verbose", "", false,
		"verbose output of the logging library (default false)")
}
