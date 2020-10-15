package urlzap

import (
	"html/template"
)

const redirectHTMLTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <title>{{.URL}}</title>
    <link rel="canonical" href="{{.URL}}"/>
    <meta name="robots" content="noindex">
    <meta charset="utf-8" />
    <meta http-equiv="refresh" content="0; url={{.URL}}" />
  </head>
</html>
`

// Generate generate static files with HTML redirects
func Generate(c Config) error {
	tmpl, err := template.New("redirect").Parse(redirectHTMLTemplate)
	if err != nil {
		return err
	}

	return Read("", c.URLs, HTMLFileCallback(c.Path, tmpl))
}
