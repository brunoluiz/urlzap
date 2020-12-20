package urlzap_test

import (
	"os"
	"testing"

	"github.com/brunoluiz/urlzap"
	"github.com/stretchr/testify/require"
)

func TestGetMetaData(t *testing.T) {
	r, err := os.Open("./example/page.html")
	require.NoError(t, err)

	meta, err := urlzap.GetMetaData(r)
	require.NoError(t, err)
	require.Equal(t, urlzap.MetaData{
		Title: "Bruno Luiz Silva",
		Tags: []string{
			"<meta name=\"description\" content=\"A collection of random software engineering thoughts\"/>",
			"<meta property=\"og:title\" content=\"Bruno Luiz Silva\"/>",
			"<meta property=\"og:description\" content=\"A collection of random software engineering thoughts\"/>",
			"<meta property=\"og:type\" content=\"website\"/>",
			"<meta property=\"og:url\" content=\"https://brunoluiz.net/\"/>",
			"<meta property=\"og:image\" content=\"https://brunoluiz.net/android-chrome-512x512.png\"/>",
			"<meta property=\"og:updated_time\" content=\"2020-09-28T19:00:00+00:00\"/>",
			"<meta property=\"og:site_name\" content=\"Bruno Luiz Silva\"/>",
			"<meta name=\"twitter:card\" content=\"summary_large_image\"/>",
			"<meta name=\"twitter:image\" content=\"https://brunoluiz.net/android-chrome-512x512.png\"/>",
			"<meta name=\"twitter:title\" content=\"Bruno Luiz Silva\"/>",
			"<meta name=\"twitter:description\" content=\"A collection of random software engineering thoughts\"/>",
		},
	}, meta)
}
