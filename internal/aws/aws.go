package aws

import (
	"bytes"
	"s3-substring-finder/internal/logging"
	"s3-substring-finder/internal/options"
	"strings"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()
}

// Find does the heavy lifting, communicates with the S3 and finds the files
func Find(opts *options.S3SubstringFinderOptions) error {
	// txtChan := make(chan *s3.Object)
	var txtSlice []*s3.Object

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(opts.Region),
		Credentials: credentials.NewStaticCredentials(opts.AccessKey, opts.SecretKey, ""),
	})
	if err != nil {
		return err
	}
	logger.Info("session successfully obtained")

	svc := s3.New(sess)
	listResult, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(opts.BucketName),
	})
	if err != nil {
		return err
	}

	for _, v := range listResult.Contents {
		if strings.Contains(*v.Key, "txt") {
			// logger.Info("adding object to the txtChan", zap.String("key", *v.Key))
			// txtChan <- v
			logger.Info("adding object to the txtSlice", zap.String("key", *v.Key))
			txtSlice = append(txtSlice, v)
		}
	}

	for _, v := range txtSlice {
		getResult, err := svc.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(opts.BucketName),
			Key:    v.Key,
		})
		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(getResult.Body); err != nil {
			return err
		}

		if strings.Contains(buf.String(), opts.Substring) {
			logger.Info("match!", zap.String("key", *v.Key))
		}

		if err := getResult.Body.Close(); err != nil {
			panic(err)
		}
	}

	return nil
}
