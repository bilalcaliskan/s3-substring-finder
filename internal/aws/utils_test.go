package aws

import (
	"testing"

	"github.com/bilalcaliskan/s3-substring-finder/cmd/root/options"

	"github.com/stretchr/testify/assert"
)

func TestCreateSession(t *testing.T) {
	opts := options.GetRootOptions()
	opts.Region = "us-east-1"
	opts.AccessKey = "asdfasdfasdfasdfasdfadsf"
	opts.SecretKey = "asdfadsfadsfasdfasdfadsfa"
	sess, err := CreateSession(opts)
	assert.NotNil(t, sess)
	assert.Nil(t, err)
}
