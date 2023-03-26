package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/bilalcaliskan/s3-substring-finder/cmd/root/options"
)

// CreateSession initializes session with provided credentials
func CreateSession(opts *options.RootOptions) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(opts.Region),
		Credentials: credentials.NewStaticCredentials(opts.AccessKey, opts.SecretKey, ""),
	})

	return sess, err
}
