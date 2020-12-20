package urlzap

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"
	"text/template"

	"golang.org/x/sync/errgroup"
)

// Callback Function which would be called once a key/value is read.
type Callback func(path, url string) error

// Mux defines interface used to register HTTP routes on a mux.
type Mux interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

// HTMLFileCallback creates static files based on a HTML template.
func HTMLFileCallback(c Config, tmpl *template.Template) Callback {
	outputPath := strings.TrimSuffix(c.Path, "/")

	return func(path, url string) error {
		fullpath := outputPath + "/" + path
		if err := os.MkdirAll(fullpath, 0700); err != nil {
			return err
		}

		w, err := os.Create(fullpath + "/index.html")
		if err != nil {
			return err
		}

		if c.DisableMetaFetch {
			return tmpl.Execute(w, redirectHTMLArgs{Title: url, URL: url})
		}

		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		meta, err := GetMetaData(resp.Body)
		if err != nil {
			return err
		}

		return tmpl.Execute(w, redirectHTMLArgs{
			URL:      url,
			Title:    meta.Title,
			MetaTags: meta.Tags,
		})
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

func read(ctx context.Context, eg *errgroup.Group, path string, urls URLs, cb Callback) {
	for k, v := range urls {
		kk, vv := k, v
		eg.Go(func() error {
			id, ok := kk.(string)
			if !ok {
				return errors.New("malformated document")
			}

			switch val := vv.(type) {
			case string:
				if err := cb(path+id, val); err != nil {
					return err
				}
			case URLs:
				read(ctx, eg, path+id+"/", val, cb)
			}

			return nil
		})
	}
}

// Read parses all URLs and calls the callback once found a key (string) and value (url)
func Read(ctx context.Context, path string, urls URLs, cb Callback) error {
	eg, ctx := errgroup.WithContext(ctx)
	read(ctx, eg, path, urls, cb)

	return eg.Wait()
}
