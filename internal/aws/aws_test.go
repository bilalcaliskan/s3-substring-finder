package aws

import (
	"os"
	"s3-substring-finder/internal/options"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/stretchr/testify/assert"
)

// Define a mock struct to be used in your unit tests
type mockS3Client struct {
	s3iface.S3API
}

// ListObjects mocks the S3API ListObjects method
func (m *mockS3Client) ListObjects(obj *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	return &s3.ListObjectsOutput{
		Name:        obj.Bucket,
		Marker:      aws.String(""),
		MaxKeys:     aws.Int64(1000),
		Prefix:      aws.String(""),
		IsTruncated: aws.Bool(false),
		Contents: []*s3.Object{
			{
				ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5449d"),
				Key:          aws.String("../../mock/file1.txt"),
				StorageClass: aws.String("STANDARD"),
			},
			{
				ETag:         aws.String("03c0fe42b7efa3470fc99037a8e54122"),
				Key:          aws.String("../../mock/file2.txt"),
				StorageClass: aws.String("STANDARD"),
			},
			{
				ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5443d"),
				Key:          aws.String("../../mock/file3.txt"),
				StorageClass: aws.String("STANDARD"),
			},
		},
	}, nil
}

// GetObject mocks the S3API GetObject method
func (m *mockS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	bytes, _ := os.Open(*input.Key)

	return &s3.GetObjectOutput{
		AcceptRanges:  aws.String("bytes"),
		Body:          bytes,
		ContentLength: aws.Int64(1000),
		ContentType:   aws.String("text/plain"),
		ETag:          aws.String("d73a503d212d9279e6b2ed8ac6bb81f3"),
	}, nil
}

func TestFind(t *testing.T) {
	mockSvc := &mockS3Client{}

	opts := options.GetS3SubstringFinderOptions()
	opts.Substring = "akqASmLLlK"
	result, errs := Find(mockSvc, opts)
	assert.NotNil(t, result)
	assert.Empty(t, errs)
}
