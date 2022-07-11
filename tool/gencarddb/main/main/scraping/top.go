package scraping

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func PackUrls(html string) []string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	var res []string
	doc.Find("div#card_list_1").Each(func(i int, s *goquery.Selection) {
		s.Find("div.pack").Each(func(i int, s *goquery.Selection) {
			v, exists := s.Find("input").Attr("value")
			if exists {
				res = append(res, "https://www.db.yugioh-card.com" + v + "&request_locale=ja")
			}
		})
	})
	doc.Find("div#card_list_2").Each(func(i int, s *goquery.Selection) {
		s.Find("div.pack").Each(func(i int, s *goquery.Selection) {
			v, exists := s.Find("input").Attr("value")
			if exists {
				res = append(res, "https://www.db.yugioh-card.com" + v + "&request_locale=ja")
			}
		})
	})
	return res
}
