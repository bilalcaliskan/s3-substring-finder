package aws

import (
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/bilalcaliskan/s3-substring-finder/internal/options"
	"github.com/stretchr/testify/assert"
)

var defaultListObjectsOutput = &s3.ListObjectsOutput{
	Name:        aws.String(""),
	Marker:      aws.String(""),
	MaxKeys:     aws.Int64(1000),
	Prefix:      aws.String(""),
	IsTruncated: aws.Bool(false),
}

// Define a mock struct to be used in your unit tests
type mockS3Client struct {
	s3iface.S3API
}

// ListObjects mocks the S3API ListObjects method
func (m *mockS3Client) ListObjects(obj *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	return defaultListObjectsOutput, nil
}

// GetObject mocks the S3API GetObject method
func (m *mockS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	bytes, err := os.Open(*input.Key)
	if err != nil {
		return nil, err
	}

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
	defaultListObjectsOutput.Contents = []*s3.Object{
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
	}
	opts := options.GetS3SubstringFinderOptions()
	opts.Substring = "akqASmLLlK"
	result, errs := Find(mockSvc, opts)
	assert.NotNil(t, result)
	assert.Empty(t, errs)
}

func TestFindWrongFilePath(t *testing.T) {
	mockSvc := &mockS3Client{}
	defaultListObjectsOutput.Contents = []*s3.Object{
		{
			ETag:         aws.String("03c0fe42b7efa3470fc99037a8e5449d"),
			Key:          aws.String("file1asdfasdf.txt"),
			StorageClass: aws.String("STANDARD"),
		},
	}
	opts := options.GetS3SubstringFinderOptions()
	res, err := Find(mockSvc, opts)
	assert.Nil(t, res)
	assert.NotEmpty(t, err)
}
