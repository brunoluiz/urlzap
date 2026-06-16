package urlzap_test

import (
	"os"
	"testing"

	"github.com/brunoluiz/urlzap"
	"github.com/brunoluiz/urlzap/internal/testtool"
)

func TestFromYAML(t *testing.T) {
	r, err := os.Open("./example/config.yml")
	testtool.NoError(t, err)

	c, err := urlzap.FromYAML(r)
	testtool.NoError(t, err)
	testtool.Equal(t, c, urlzap.Config{
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
