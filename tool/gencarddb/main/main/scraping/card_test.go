package scraping

import (
	"io/ioutil"
	"main/card"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)


func readCard(file string) (string, error) {
	htmlPath := "./card_test/" + file + ".html"
	html, err := ioutil.ReadFile(htmlPath)
	if err != nil {
		return "", err
	}
	return string(html), nil
}

func TestCard(t *testing.T) {
	testdatas := []struct{
		file string
		ans card.Card
	}{
		{file: "GaiaTheFierceKnight",ans: card.Card{
			Name:        "暗黒騎士ガイア",
			Ruby:        "あんこくきしガイア",
			ImageSrcUrl: []string{"https://www.db.yugioh-card.com/yugiohdb/get_image.action?type=1&cid=4044&ciid=1&enc=gcAVxqWH8Pmf7O2WyTpN3A&request_locale=ja", "https://www.db.yugioh-card.com/yugiohdb/get_image.action?type=1&cid=4044&ciid=2&enc=gcAVxqWH8Pmf7O2WyTpN3A&request_locale=ja"},
			ReleaseDate: "1999-02-04",
			Url:         "",
			Text:        "風よりも速く走る馬に乗った騎士。突進攻撃に注意。",
			Atk:         "2300",
			Def:         "2100",
			Attribute:   "地属性",
			MonsterType: "戦士族",
			Level:       7,
			CardType:    []string{"通常"},
		}},
		{file: "ALegendaryOcean",ans: card.Card{
			Name:        "伝説の都 アトランティス",
			Ruby:        "でんせつのみやこ　アトランティス",
			ImageSrcUrl: []string{"https://www.db.yugioh-card.com/yugiohdb/get_image.action?type=1&cid=5387&ciid=1&enc=yvVk4ycBO-1hcRNunz_y_w&request_locale=ja"},
			ReleaseDate: "2001-11-29",
			Url:         "",
			Text:        "このカード名はルール上「海」として扱う。①：フィールドの水属性モンスターの攻撃力・守備力は２００アップする。②：このカードがフィールドゾーンに存在する限り、お互いの手札・フィールドの水属性モンスターのレベルは１つ下がる。",
			Atk:         "",
			Def:         "",
			Attribute:   "フィールド魔法",
			MonsterType: "",
			Level:       0,
			CardType:    []string{},
		}},
		{file: "Relinquished",ans: card.Card{
			Name:        "サクリファイス",
			Ruby:        "サクリファイス",
			ImageSrcUrl: []string{"https://www.db.yugioh-card.com/yugiohdb/get_image.action?type=1&cid=4737&ciid=1&enc=VlA4aUF9QIvwfFX9W-udAQ&request_locale=ja"},
			ReleaseDate: "2000-04-20",
			Url:         "",
			Text:        "「イリュージョンの儀式」により降臨。①：１ターンに１度、相手フィールドのモンスター１体を対象として発動できる。その相手モンスターを装備カード扱いとしてこのカードに装備する（１体のみ装備可能）。②：このカードの攻撃力・守備力は、このカードの効果で装備したモンスターのそれぞれの数値になり、このカードが戦闘で破壊される場合、代わりに装備したそのモンスターを破壊する。③：このカードの効果でモンスターを装備したこのカードの戦闘で自分が戦闘ダメージを受けた時、相手も同じ数値分の効果ダメージを受ける。",
			Atk:         "0",
			Def:         "0",
			Attribute:   "闇属性",
			MonsterType: "魔法使い族",
			Level:       1,
			CardType:    []string{"儀式", "効果"},
		}},
	}

	for _, testdata := range testdatas {
		html, err := readCard(testdata.file)
		if err != nil {t.Fatal(err)}
		doc, err := ScrapingDocument(html)
		if err != nil {t.Fatal(err)}

		card, err := scrapingCard(doc, "")
		if err != nil {t.Fatal(err)}
		if !reflect.DeepEqual(card, testdata.ans) {
			t.Fatalf("\n%s", pretty.Compare(card, testdata.ans))
		}
	}
}
