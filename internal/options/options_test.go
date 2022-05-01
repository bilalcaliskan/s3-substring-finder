package options

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetOreillyTrialOptions function tests if GetOreillyTrialOptions function running properly
func TestGetS3SubstringFinderOptions(t *testing.T) {
	t.Log("fetching default options.S3SubstringFinderOptions")
	opts := GetS3SubstringFinderOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.S3SubstringFinderOptions, %v\n", opts)
}
