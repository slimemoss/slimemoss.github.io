package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"main/card"
	"main/myhttp"
	"main/scraping"
	"os"
	"time"
)

func main() {
	log.Printf("[info] %s Start Update", time.Now().String())

	toppage := "https://www.db.yugioh-card.com/yugiohdb/card_list.action?request_locale=ja"
	html, err := getHtml(myhttp.Curl(toppage, "", false, true))
	if err != nil {
		log.Fatalf("E Update: can not get toppage  %s", err)
	}
	packUrls := scraping.PackUrls(html)

	cards := []card.Card{}
	for packIndex, packUrl := range packUrls {
		log.Printf("%d / %d", packIndex + 1, len(packUrls))

		html, err := getHtml(myhttp.Curl(packUrl, "", true, true))
		if err != nil {
			log.Println(err.Error())
			continue
		}
		cardUrls := scraping.CardUrls(html)
		packDate, err := scraping.Date(html)
		if !is02(packDate) { continue }
		log.Printf("date: %s, cards: %d", packDate,len(cardUrls))

		for _, cardUrl := range cardUrls {
			html, err := getHtml(myhttp.Curl(cardUrl, "", true, true))
			if err != nil {
				log.Println(err.Error())
				continue
			}

			doc, err := scraping.ScrapingDocument(html)
			if err != nil {
				log.Println(err.Error())
				continue
			}

			card, err := scraping.Card(doc, cardUrl)
			if err != nil {
				log.Printf("[Error] Skip register card to DB (%s): %s", cardUrl, err)
				continue
			}
			if !contains(cards, card) {
				cards = append(cards, card)
			}
		}
	}

	log.Printf("cards: %d", len(cards))

	d, err := json.Marshal(cards)
	if err != nil {
		log.Fatal(err)
	}
	err = writeData(d)
	if err != nil {
		log.Fatalf("%s", err)
	}

	time.Sleep(10 * 60 * time.Second)
}

func contains(cards []card.Card, t card.Card) bool {
	for _, c := range cards {
		if c.Name == t.Name {
			return true
		}
	}
	return false
}

func writeData(data []byte) error {
	f, err := os.Create("./out/data.json")
	if err != nil {
		return err
	}

	log.Printf("[info] Write to %s", f.Name())
	f.Write(data)
	return nil
}

func getHtml(body io.ReadCloser, err error) (string, error) {
	if err != nil {
		return "", err
	}
	res, err := ioutil.ReadAll(body)
	body.Close()
	return string(res), err
}

func is02(d time.Time) bool {
	th := time.Date(2002, 3, 22, 0, 0, 0, 0, time.UTC)
	return d.Before(th)
}
