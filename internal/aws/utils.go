package aws

import (
	"s3-substring-finder/internal/options"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

/*func downloadFile(sess *session.Session, key, bucketName string) (int64, error) {
	downloader := s3manager.NewDownloader(sess)

	// Download the item from the bucket. If an error occurs, log it and exit. Otherwise, notify the user that the download succeeded.
	file, err := os.Create(key)
	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	return numBytes, err
}*/

// CreateSession initializes session with provided credentials
func CreateSession(opts *options.S3SubstringFinderOptions) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(opts.Region),
		Credentials: credentials.NewStaticCredentials(opts.AccessKey, opts.SecretKey, ""),
	})

	return sess, err
}
