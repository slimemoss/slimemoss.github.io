package scraping

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func readTop(t *testing.T) string {
	htmlPath := "./top_test/top.html"
	html, err := ioutil.ReadFile(htmlPath)
	if err != nil {
		t.Fatal("")
	}
	return string(html)
}

func TestPackUrls(t *testing.T) {
	html := readTop(t)
	data := PackUrls(html)
	urlPattern := regexp.MustCompile(`https://www\.db\.yugioh-card\.com/yugiohdb/card_search\.action\?ope=1&sess=1&pid=.*&rp=99999&request_locale=ja`)
	for _, v := range data {
		if ! urlPattern.MatchString(v) {
			t.Fatal("")
		}
	}
	if len(data) < 1000 {
		t.Fatalf("検出したパックの数が少なすぎます %d < %d", len(data), 1000)
	}
}
