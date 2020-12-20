package urlzap_test

import (
	"os"
	"testing"

	"github.com/brunoluiz/urlzap"
	"github.com/stretchr/testify/require"
)

func TestFromYAML(t *testing.T) {
	r, err := os.Open("./example/config.yml")
	require.NoError(t, err)

	c, err := urlzap.FromYAML(r)
	require.NoError(t, err)
	require.Equal(t, c, urlzap.Config{
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
