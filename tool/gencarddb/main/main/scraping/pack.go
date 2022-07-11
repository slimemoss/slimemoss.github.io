package scraping

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func CardUrls(html string) []string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	var res []string
	doc.Find("div#card_list").Find("div.t_row").Each(func(i int, s *goquery.Selection) {
		if url, exists := s.Find("input.link_value").Attr("value"); exists {
			res = append(res, "https://www.db.yugioh-card.com" + url + "&request_locale=ja")
		}
	})
	return res
}

func Date(html string) (time.Time, error) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
	data := doc.Find("header#broad_title").Find("p").Text()
	data = removeSpace(data)
	data = strings.ReplaceAll(data, "(公開日:", "")
	data = strings.ReplaceAll(data, "日)", "")
	data = strings.ReplaceAll(data, "年", "-")
	data = strings.ReplaceAll(data, "月", "-")
	date, err := time.Parse("2006-01-02", data)
	if err != nil {
		return time.Now(), err
	}
	return date, err
}
