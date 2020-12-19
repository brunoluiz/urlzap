package urlzap_test

import (
	"context"
	"os"
	"testing"

	"github.com/brunoluiz/urlzap"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	err := urlzap.Generate(context.Background(), urlzap.Config{
		Path: "./output",
		URLs: urlzap.URLs{
			"google": "https://google.com",
			"tools": urlzap.URLs{
				"github": "https://github.com",
			},
		},
	})
	require.NoError(t, err)

	f1, err := os.Stat("./output/google/index.html")
	require.NoError(t, err)
	require.NotEqual(t, f1.Size(), 0)

	f2, err := os.Stat("./output/tools/github/index.html")
	require.NoError(t, err)
	require.NotEqual(t, f2.Size(), 0)

	err = os.RemoveAll("./output")
	require.NoError(t, err)
}
