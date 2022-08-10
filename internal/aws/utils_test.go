package aws

import (
	"testing"

	"github.com/bilalcaliskan/s3-substring-finder/internal/options"
	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	opts := options.GetS3SubstringFinderOptions()
	opts.Region = "us-east-1"
	opts.AccessKey = "asdfasdfasdfasdfasdfadsf"
	opts.SecretKey = "asdfadsfadsfasdfasdfadsfa"
	sess, err := CreateSession(opts)
	assert.NotNil(t, sess)
	assert.Nil(t, err)
}
