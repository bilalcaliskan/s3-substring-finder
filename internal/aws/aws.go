package aws

import (
	"bytes"
	"strings"
	"sync"

	"github.com/bilalcaliskan/s3-substring-finder/cmd/root/options"
	"github.com/rs/zerolog"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/schollz/progressbar/v3"
)

// Find does the heavy lifting, communicates with the S3 and finds the files
func Find(svc s3iface.S3API, opts *options.RootOptions, logger zerolog.Logger) ([]string, []error) {
	var errors []error
	var matchedFiles []string
	mu := &sync.Mutex{}

	// fetch all the objects in target bucket
	listResult, err := svc.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(opts.BucketName),
	})
	if err != nil {
		errors = append(errors, err)
		return matchedFiles, errors
	}

	var resultArr []*s3.Object
	var wg sync.WaitGroup

	extensions := strings.Split(opts.FileExtensions, ",")

	// separate the txt files from all of the fetched objects from bucket
	for _, v := range listResult.Contents {
		for _, y := range extensions {
			if strings.HasSuffix(*v.Key, y) {
				logger.Debug().Str("fileName", *v.Key).Msg("found file")
				resultArr = append(resultArr, v)
			}
		}
	}

	bar := progressbar.Default(int64(len(resultArr)))
	// check each txt file individually if it contains provided substring
	for _, obj := range resultArr {
		wg.Add(1)
		go func(obj *s3.Object, wg *sync.WaitGroup) {
			defer wg.Done()
			getResult, err := svc.GetObject(&s3.GetObjectInput{
				Bucket: aws.String(opts.BucketName),
				Key:    obj.Key,
			})

			if err != nil {
				errors = append(errors, err)
				return
			}

			buf := new(bytes.Buffer)
			if _, err := buf.ReadFrom(getResult.Body); err != nil {
				errors = append(errors, err)
				return
			}

			if strings.Contains(buf.String(), opts.Substring) {
				mu.Lock()
				matchedFiles = append(matchedFiles, *obj.Key)
				mu.Unlock()
			}

			defer func() {
				if err := getResult.Body.Close(); err != nil {
					panic(err)
				}
			}()

			_ = bar.Add(1)
		}(obj, &wg)
	}

	// wait for all the goroutines to complete
	wg.Wait()

	return matchedFiles, errors
}
