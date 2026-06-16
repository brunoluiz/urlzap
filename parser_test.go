package urlzap_test

import (
	"os"
	"testing"

	"github.com/brunoluiz/urlzap"
	"github.com/brunoluiz/urlzap/internal/xtest"
)

func TestFromYAML(t *testing.T) {
	r, err := os.Open("./example/config.yml")
	xtest.NoError(t, err)

	c, err := urlzap.FromYAML(r)
	xtest.NoError(t, err)
	xtest.Equal(t, c, urlzap.Config{
		Path: "./output",
		HTTP: urlzap.HTTP{
			BasePath: "/example",
			Address:  ":80",
		},
		URLs: urlzap.URLs{
			"google": "https://google.com",
			"tools": urlzap.URLs{
				"github": "https://github.com",
			},
		},
		DisableMetaFetch: false,
	})
}
