package urlzap_test

import (
	"context"
	"os"
	"testing"

	"github.com/brunoluiz/urlzap"
	"github.com/brunoluiz/urlzap/internal/xtest"
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
	xtest.NoError(t, err)

	f1, err := os.Stat("./output/google/index.html")
	xtest.NoError(t, err)
	xtest.NotEqual(t, f1.Size(), 0)

	f2, err := os.Stat("./output/tools/github/index.html")
	xtest.NoError(t, err)
	xtest.NotEqual(t, f2.Size(), 0)

	err = os.RemoveAll("./output")
	xtest.NoError(t, err)
}
