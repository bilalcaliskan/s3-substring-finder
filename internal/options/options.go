package options

var s3SubstringFinderOptions = &S3SubstringFinderOptions{}

// S3SubstringFinderOptions contains frequent command line and application options.
type S3SubstringFinderOptions struct {
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
}

// GetS3SubstringFinderOptions returns the pointer of S3SubstringFinderOptions
func GetS3SubstringFinderOptions() *S3SubstringFinderOptions {
	return s3SubstringFinderOptions
}
