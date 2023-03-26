package options

import (
	"os"
	"testing"

	"github.com/spf13/cobra"

	"github.com/stretchr/testify/assert"
)

// TestGetRootOptions function tests if GetRootOptions function running properly
func TestGetRootOptions(t *testing.T) {
	t.Log("fetching default options.RootOptions")
	opts := GetRootOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.RootOptions, %v\n", opts)
}

func TestRootOptions_GetAccessCredentialsFromEnv(t *testing.T) {
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
			opts := GetRootOptions()
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
