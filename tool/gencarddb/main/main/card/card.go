package card

type Card struct {
	Name string `bson:"name" json:"name"`
	Ruby string `bson:"ruby" json:"ruby"`
	ImageSrcUrl []string `bson:"imagesrcurl" json:"imagesrcurl"`
	ImageUrl []string `bson:"imageurl" json:"imageurl"`
	ReleaseDate string `bson:"releasedate" json:"releasedate"`
	Url string `bson:"url" json:"url"`
	Text string `bson:"text" json:"text"`
	Atk string `bson:"atk" json:"atk"`
	Def string `bson:"def" json:"def"`
	// 属性・魔法罠の種類
	Attribute string `bson:"attribute" json:"attribute"`
	MonsterType string `bson:"monstertype" json:"monstertype"`
	Level int `bson:"level" json:"level"`
	// シンクロ・効果とか
	CardType []string `bson:"cardtype" json:"cardtype"`
}
