package scraping

import (
	"io/ioutil"
	"testing"
	"time"
)

func readPack(file string, t *testing.T) string {
	htmlPath := "./pack_test/" + file + ".html"
	html, err := ioutil.ReadFile(htmlPath)
	if err != nil {
		t.Fatalf("Can not find test file : %v", htmlPath)
	}
	return string(html)
}

func equal(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true	
}

func TestCardUrls(t *testing.T) {
	html := readPack("Vol1", t)
	data := CardUrls(html)

	if 0 == len(data) {
		t.Fatal("no card detected")
	}

	if data[0] != "https://www.db.yugioh-card.com/yugiohdb/card_search.action?ope=2&cid=4009&request_locale=ja" {
		t.Fatal("wrong card url")
	}
	if len(data) != 40 {
		t.Fatalf("detected cards: %d != 40", len(data))
	}
}

func TestDate(t *testing.T) {
	html := readPack("Vol1", t)
	data, err := Date(html)
	if err != nil {
		t.Fatal("can not parse date")
	}
	if !data.Equal(time.Date(1999, 2, 4, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("date is wrong")
	}
}
