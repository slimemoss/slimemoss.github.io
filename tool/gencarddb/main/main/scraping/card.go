package scraping

import (
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"main/card"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func removeSpace(data string) string {
	return regexp.MustCompile(`\s+`).ReplaceAllString(data, "")
}

func getBase64(body io.ReadCloser, err error) (string, error) {
	if err != nil {
		return "", err
	}
	res, err := ioutil.ReadAll(body)
	body.Close()
	return base64.StdEncoding.EncodeToString(res), err
}

func ScrapingDocument(html string) (*goquery.Document, error) {
	return goquery.NewDocumentFromReader(strings.NewReader(html))
}

func Card(doc *goquery.Document, url string) (card.Card, error) {
	return scrapingCard(doc, url)
}

func scrapingCard(doc *goquery.Document, url string) (card.Card, error) {
	levelstr := strings.Trim(getItemboxData(doc, "レベル"), "レベル")
	levelstr = strings.TrimSpace(levelstr)
	level, err := strconv.Atoi(levelstr)
	if err != nil {level = 0}

	res := card.Card{
		ImageSrcUrl: imageUrl(doc),
		Url:         url,
		Text:        text(doc),
		Atk:         getItemboxData(doc, "攻撃力"),
		Def:         getItemboxData(doc, "守備力"),
		Attribute:   attribute(doc),
		MonsterType: monsterType(doc),
		Level:       level,
		CardType:    cardType(doc),
	}

	if res.Name, err = name(doc); err != nil { return card.Card{}, err }
	if res.Ruby, err = ruby(doc); err != nil { return card.Card{}, err }
	if res.ReleaseDate, err = releaseDate(doc); err != nil { return card.Card{}, err }

	return res, nil
}

func name(doc *goquery.Document) (string, error) {
	span := doc.Find("div#CardSet").Find("div#cardname").Find("span")

	if len(span.Nodes) == 0 {
		return "", errors.New("Can not scraping name")
	}
	name := span.Nodes[0].NextSibling.Data
	return strings.TrimSpace(name), nil
}

func ruby(doc *goquery.Document) (string, error) {
	span := doc.Find("div#CardSet").Find("div#cardname").Find("span")
	
	if len(span.Nodes) == 0 {
		return "", errors.New("Can not scraping ruby")
	}

	ruby := span.Nodes[0].FirstChild.Data
	return ruby, nil
}

func imageUrl(doc *goquery.Document) []string {
	thumbnail := doc.Find("div#thumbnail")
	var res []string
	thumbnail.Find("img").Each(func(i int, s *goquery.Selection) {
		if src, exists := s.Attr("src"); exists {
			res = append(res, "https://www.db.yugioh-card.com" + src + "&request_locale=ja")
		}
	})
	return res
}

func attribute(doc *goquery.Document) string {
	val := getItemboxData(doc, "属性")
	if val == "" {		
		val = getItemboxData(doc, "効果")
	}
	return val
}

func cardType(doc *goquery.Document) []string {
	sp := doc.Find("p.species").Text()
	sp = removeSpace(sp)
	return strings.Split(sp, "／")[1:]
}

func releaseDate(doc *goquery.Document) (string, error) {
	packs := doc.Find("div#update_list").Find("div.t_row")
	row := packs.Last()
	if 0 == len(packs.Nodes) {
		return "", errors.New("packs for releaseDate is 0")
	}

	div := row.Find("div.time")

	if len(div.Nodes) == 0 {
		return "", errors.New("Can not scraping releaseDate")
	}
	dateString := div.Nodes[0].FirstChild.Data
	dateString = strings.TrimSpace(dateString)

	return dateString, nil
}

func monsterType(doc *goquery.Document) string {
	sp := doc.Find("p.species").Text()
	sp = removeSpace(sp)
	return strings.Split(sp, "／")[0]
}

func text(doc *goquery.Document) string {
	res := doc.Find("div.item_box_text").Text()
	res = removeSpace(res)
	res = strings.ReplaceAll(res, "カードテキスト", "")
	return res
}

func getItemboxData(doc *goquery.Document, title string) string {
	table := doc.Find("div#CardTextSet")
	
	value := ""
	table.Find(".item_box").Each(func(i int, s *goquery.Selection) {
		t := s.Find("span.item_box_title").Text()
		t = strings.TrimSpace(t)

		if t == title {
			value = s.Find("span.item_box_value").Text()
			value = strings.TrimSpace(value)
		}
	})

	return strings.TrimSpace(value)
}
