package options

import (
	"os"
	"testing"

	"github.com/spf13/cobra"

	"github.com/stretchr/testify/assert"
)

// TestGetS3SubstringFinderOptions function tests if GetS3SubstringFinderOptions function running properly
func TestGetS3SubstringFinderOptions(t *testing.T) {
	t.Log("fetching default options.S3SubstringFinderOptions")
	opts := GetS3SubstringFinderOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.S3SubstringFinderOptions, %v\n", opts)
}

func TestS3SubstringFinderOptions_GetAccessCredentialsFromEnv(t *testing.T) {
	rootCmd := &cobra.Command{}

	cases := []struct {
		caseName, envName, envValue string
	}{
		{"case1", "AWS_ACCESS_KEY", "asdasfas"},
		{"case2", "AWS_SECRET_KEY", "asdasfas"},
		{"case3", "AWS_BUCKET_NAME", "asdasfas"},
		{"case4", "AWS_REGION", "asdasfas"},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			opts := GetS3SubstringFinderOptions()
			_ = os.Setenv(tc.envName, tc.envValue)
			if err := opts.SetAccessCredentialsFromEnv(rootCmd); err != nil {
				t.Error(err)
				return
			}
			assert.Equal(t, opts.AccessKey, tc.envValue)
			_ = os.Unsetenv(tc.envName)
		})
	}
}
