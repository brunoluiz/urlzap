package urlzap

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type MetaData struct {
	Title string
	Tags  []string
}

// GetMetaData Fetches meta data from URL, such as title, open graph and twitter data
func GetMetaData(html io.ReadCloser) (meta MetaData, err error) {
	defer html.Close()

	doc, err := goquery.NewDocumentFromReader(html)
	if err != nil {
		return meta, err
	}

	meta.Title = doc.Find("title").Text()

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" || strings.Contains(name, "twitter:") {
			if tag, err := goquery.OuterHtml(s); err == nil {
				meta.Tags = append(meta.Tags, tag)
			}
		}

		if property, _ := s.Attr("property"); strings.Contains(property, "og:") || property == "fb:app_id" {
			if tag, err := goquery.OuterHtml(s); err == nil {
				meta.Tags = append(meta.Tags, tag)
			}
		}
	})

	return meta, nil
}
