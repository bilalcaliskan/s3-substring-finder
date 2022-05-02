package aws

import (
	"bytes"
	"s3-substring-finder/internal/options"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/schollz/progressbar/v3"
)

// Find does the heavy lifting, communicates with the S3 and finds the files
func Find(opts *options.S3SubstringFinderOptions) ([]string, []error) {
	var errors []error
	var matchedFiles []string

	// initialize session with provided credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(opts.Region),
		Credentials: credentials.NewStaticCredentials(opts.AccessKey, opts.SecretKey, ""),
	})
	if err != nil {
		errors = append(errors, err)
		return matchedFiles, errors
	}

	// obtain S3 client with initialized session
	svc := s3.New(sess)

	// fetch all the objects in target bucket
	listResult, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(opts.BucketName),
	})
	if err != nil {
		errors = append(errors, err)
		return matchedFiles, errors
	}

	var txtArr []*s3.Object
	var wg sync.WaitGroup

	// separate the txt files from all of the fetched objects from bucket
	for _, v := range listResult.Contents {
		if strings.HasSuffix(*v.Key, "txt") {
			txtArr = append(txtArr, v)
		}
	}

	bar := progressbar.Default(int64(len(txtArr)))
	// check each txt file individually if it contains provided substring
	for _, obj := range txtArr {
		wg.Add(1)
		go func(obj *s3.Object, wg *sync.WaitGroup) {
			defer wg.Done()
			getResult, err := svc.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(opts.BucketName),
				Key:    obj.Key,
			})
			if err != nil {
				errors = append(errors, err)
			}

			buf := new(bytes.Buffer)
			if _, err := buf.ReadFrom(getResult.Body); err != nil {
				errors = append(errors, err)
			}

			if strings.Contains(buf.String(), opts.Substring) {
				matchedFiles = append(matchedFiles, *obj.Key)
			}

			if err := getResult.Body.Close(); err != nil {
				panic(err)
			}

			_ = bar.Add(1)
		}(obj, &wg)
	}

	// wait for all the goroutines to complete
	wg.Wait()

	return matchedFiles, errors
}
