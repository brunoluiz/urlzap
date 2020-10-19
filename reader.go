package urlzap

import (
	"errors"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// Callback Function which would be called once a key/value is read.
type Callback func(path, url string) error

// Mux defines interface used to register HTTP routes on a mux.
type Mux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

// HTMLFileCallback creates static files based on a HTML template.
func HTMLFileCallback(outputPath string, tmpl *template.Template) Callback {
	outputPath = strings.TrimSuffix(outputPath, "/")

	return func(path, url string) error {
		fullpath := outputPath + "/" + path
		if err := os.MkdirAll(fullpath, 0700); err != nil {
			return err
		}

		w, err := os.Create(fullpath + "/index.html")
		if err != nil {
			return err
		}

		return tmpl.Execute(w, struct {
			URL string
		}{url})
	}
}

// HTTPMuxCallback register redirect routes in a mux.
func HTTPMuxCallback(parent string, r Mux) Callback {
	parent = strings.TrimSuffix(parent, "/")

	return func(path, url string) error {
		route := parent + "/" + path
		r.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, url, 301)
		})
		return nil
	}
}

// Read parses all URLs and calls the callback once found a key (string) and value (url)
func Read(path string, urls URLs, cb Callback) error {
	for k, v := range urls {
		id, ok := k.(string)
		if !ok {
			return errors.New("malformated document")
		}

		switch val := v.(type) {
		case string:
			if err := cb(path+id, val); err != nil {
				return err
			}
		case URLs:
			if err := Read(path+id+"/", val, cb); err != nil {
				return err
			}
		default:
			continue
		}
	}

	return nil
}
