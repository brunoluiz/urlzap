package urlzap

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetMetaData Fetches meta data from URL, such as title, open graph and twitter data
func GetMetaData(url string) (title string, meta []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return title, meta, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return title, meta, err
	}

	title = doc.Find("title").Text()
	if title == "" {
		title = url
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" || strings.Contains(name, "twitter:") {
			if tag, err := goquery.OuterHtml(s); err == nil {
				meta = append(meta, tag)
			}
		}

		if property, _ := s.Attr("property"); strings.Contains(property, "og:") || property == "fb:app_id" {
			if tag, err := goquery.OuterHtml(s); err == nil {
				meta = append(meta, tag)
			}
		}
	})

	return title, meta, nil
}
