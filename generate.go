package urlzap

import (
	"context"
	"text/template"
)

type redirectHTMLArgs struct {
	Title    string
	URL      string
	MetaTags []string
}

const redirectHTMLTemplate = `<!DOCTYPE html>
<html>
	<head>
		<title>{{.Title}}</title>
		<link rel="canonical" href="{{.URL}}"/>
		<meta name="robots" content="noindex">
		<meta charset="utf-8" />
		<meta http-equiv="refresh" content="0; url={{.URL}}" />
		{{ range $key, $value := .MetaTags }}
		{{ $value }}{{ end }}
	</head>
</html>
`

// Generate generate static files with HTML redirects
func Generate(ctx context.Context, c Config) error {
	tmpl, err := template.New("redirect").Parse(redirectHTMLTemplate)
	if err != nil {
		return err
	}

	return Read(ctx, "", c.URLs, HTMLFileCallback(c, tmpl))
}
