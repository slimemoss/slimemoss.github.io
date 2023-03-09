import * as React from 'react'

import * as Scroll from 'react-scroll'

import { YugiohCard } from '../card_info/useYugiohCard'
import { CardsDom } from './CardsDom'
import { Button, Col, Container, FloatingLabel, Form, Row, Spinner, ToggleButton } from 'react-bootstrap'
import { AiOutlineArrowUp, AiOutlineArrowDown } from 'react-icons/ai'
import { FiChevronDown, FiChevronUp } from 'react-icons/fi'

const parseAtkDef = (data: string): number => {
  if(!data) { return 0 }
  if(!parseInt(data)) { return 0 }
  return parseInt(data)
}

const monstertypeList = [
  "アンデット族",
  "サイキック族",
  "サイバース族",
  "ドラゴン族",
  "悪魔族",
  "炎族",
  "海竜族",
  "岩石族",
  "機械族",
  "魚族",
  "恐竜族",
  "幻神獣族",
  "幻竜族",
  "昆虫族",
  "獣戦士族",
  "獣族",
  "植物族",
  "水族",
  "戦士族",
  "創造神族",
  "鳥獣族",
  "天使族",
  "魔法使い族",
  "雷族",
  "爬虫類族",
]

const cardtypeList = [
  "通常",
  "効果",
  "儀式",
  "融合",
  "特殊召喚",
  "リバース",
]

const attributeMons = [
  "闇属性",
  "光属性",
  "地属性",
  "水属性",
  "炎属性",
  "風属性",
  "神属性",
]
const attributeMagic = [
  "通常魔法",
  "速攻魔法",
  "永続魔法",
  "フィールド魔法",
  "装備魔法",
  "儀式魔法",
]
const attributeTrap = [
  "通常罠",
  "カウンター罠",
  "永続罠",
]

interface SortableElem {
  Name: string
  CompareKey: (a: YugiohCard) => number | string
  IsDescending: boolean
}

const ListOfSortableElemDefault = {Name: "名前", CompareKey: card => card.name, IsDescending: false}
const ListOfSortableElem: SortableElem[] = [
  {Name: "名前", CompareKey: card => card.name, IsDescending: true},
  {Name: "攻撃力", CompareKey: card => parseAtkDef(card.atk), IsDescending: true},
  {Name: "守備力", CompareKey: card => parseAtkDef(card.def), IsDescending: true},
  {Name: "レベル", CompareKey: card => card.level, IsDescending: true},
]

const isIncludeAll = (big: string[], small: string[]): boolean => {
  for(const e of small) {
    if (!big.includes(e)) {
      return false
    }
  }
  return true
}

const isIncludeSome = (big: string[], small: string[]): boolean => {
  for(const e of small) {
    if (big.includes(e)) {
      return true
    }
  }
  return false
}

const removeAll = (big: string[], small: string[]): string[] => {
  for(const e of small) {
    big = big.filter(b => b != e)
  }
  return big
}

const pushAll = (pushto: string[], thisone: string[]): string[] => {
  for(const e of thisone) {
    if(!pushto.includes(e)) {
      pushto.push(e)
    }
  }
  return pushto
}


interface ToggleAttributeProps {
  attributeList: string[]
  attribute: string[]
  setAttribute: React.Dispatch<React.SetStateAction<string[]>>
  text: string
}
const ToggleAttribute = (props: ToggleAttributeProps) => {
  const {attributeList, attribute, setAttribute, text} = props

  return (
    <>
      <Button variant="dark" className="mr-3 p-1" size="sm"
              style={{width: '6em'}}
              onClick={() => {
                if(isIncludeAll(attribute, attributeList)) {
                  setAttribute((prev) => (removeAll(prev.concat(), attributeList)))
                } else {
                  setAttribute((prev) => (pushAll(prev.concat(), attributeList)))
                }
              }}>
        {text}</Button>
      {attributeList.map((att, index) => (
        <ToggleElemButton key={index} list={attribute} elem={att}
                               setList={setAttribute} />
      ))}
    </>
  )
}

interface ToggleElemButtonProps {
  list: string[]
  elem: string,
  setList: React.Dispatch<React.SetStateAction<string[]>>
}
const ToggleElemButton = (props: ToggleElemButtonProps) => {
  const {list, elem, setList} = props
  return (
    <ToggleButton className="m-1" value="1" variant="outline-secondary"
                  id={elem} type="checkbox" size="sm"
                  checked={list.includes(elem)}
                  onChange={(e) => {
                    if(e.currentTarget.checked) {
                      setList((prev) => {
                        const d = prev.concat()
                        d.push(elem)
                        return d
                      })
                    } else {
                      setList((prev) => (
                        prev.concat().filter((a) => (a != elem)))
                      )
                    }
                  }}>
      {elem}</ToggleButton>
  )
}

interface Props {
  title: string
  CardData: any
  limit: (name: string) => number
}
export const CardSearchPage = (props: Props) => {
  const {title, CardData, limit} = props

  React.useEffect(() => {
    document.title = title
  }, [title]);

  const [cards, setCards] = React.useState<YugiohCard[]>([])

  const [atkMax, setAtkMax] = React.useState("5000")
  const [atkMin, setAtkMin] = React.useState("0")
  const [defMax, setDefMax] = React.useState("5000")
  const [defMin, setDefMin] = React.useState("0")
  const [levelMax, setLevelMax] = React.useState("10")
  const [levelMin, setLevelMin] = React.useState("0")
  const [attribute, setAttribute] = React.useState<string[]>([])
  const [searchText, setSearchText] = React.useState<string>("")
  const [monstertype, setMonstertype] = React.useState<string[]>([])
  const [monsterAttribute, setMonsterAttribute] = React.useState<string[]>([])

  const [sortElem, setSortElem] = React.useState<SortableElem>(ListOfSortableElemDefault)

  const filterAndSort = (cards: YugiohCard[]): (YugiohCard[]) => {
    if(attribute.length != 0) {
      cards = cards.filter((card) => {
        if(attribute.includes(card.attribute)){
          return true
        }
        for(const c of card.cardtype){
          if(attribute.includes(c)){
            return true
          }
        }
        return false
      })
    }

    // モンスターのフィルター
    cards = cards.filter((card) => {
      if(attributeMagic.includes(card.attribute)){
        return true
      }
      if(attributeTrap.includes(card.attribute)){
        return true
      }

      const statusF = (parseAtkDef(card.atk) <= parseAtkDef(atkMax)) &&
                      (parseAtkDef(card.atk) >= parseAtkDef(atkMin)) &&
                      (parseAtkDef(card.def) <= parseAtkDef(defMax)) &&
                      (parseAtkDef(card.def) >= parseAtkDef(defMin)) &&
                      (card.level <= parseAtkDef(levelMax)) &&
                      (card.level >= parseAtkDef(levelMin))

      let monstertypeF = true
      if(monstertype.length != 0) {
        monstertypeF = monstertype.includes(card.monstertype)
      }

      let monsterattrF = true
      if(monsterAttribute.length != 0) {
        monsterattrF = monsterAttribute.includes(card.attribute)
      }

      return statusF && monstertypeF && monsterattrF
    })

    for(const elem of searchText.split(/(\s)/)) {
      cards = cards.filter((card) => 
        card.name.includes(elem) || card.text.includes(elem)
      )
    }

    cards = cards.sort((a, b) => (
      (sortElem.CompareKey(a) > sortElem.CompareKey(b)) == sortElem.IsDescending ? -1 : 1
    ))

    return cards
  }

  const setCardsWithFilterAndSort = () => {
    setCards(filterAndSort(CardData))
  }

  React.useEffect(() => {
    if(!isIncludeSome(attribute, cardtypeList) && attribute.length != 0){
      setAtkMax("5000"); setAtkMin("0"); setDefMax("5000"); setDefMin("0");
      setLevelMax("10"); setLevelMin("0");
      setMonsterAttribute([]); setMonstertype([]);
    }
  }, [attribute])

  React.useEffect(() => {
    setCardsWithFilterAndSort()
  }, [CardData, sortElem])

  const [showMonsterAttribute, setShowMonsterAttribute] = React.useState(false)
  const [showMonsterType, setShowMonsterType] = React.useState(false)

  return (
    <Container style={{fontSize: '0.8em'}}>

      <div className="m-3">
        <Form>
          <Form.Label className="m-0" column="lg">フィルタ</Form.Label>
          <Row>
            <Col>
              <ToggleAttribute attributeList={attributeMagic} attribute={attribute}
                               setAttribute={setAttribute} text="魔法" />
            </Col>
          </Row>
          <Row>
            <Col>
              <ToggleAttribute attributeList={attributeTrap} attribute={attribute}
                               setAttribute={setAttribute} text="罠" />
            </Col>
          </Row>
          <Row>
            <Col>
              <ToggleAttribute attributeList={cardtypeList} attribute={attribute}
                               setAttribute={setAttribute} text="モンスター" />
            </Col>
          </Row>
          
          <Row hidden={!isIncludeSome(attribute, cardtypeList) && attribute.length != 0 }>
            <Col style={{display: 'flex', flexWrap: 'wrap'}}>
              <div className="mx-1">
                <Form.Label className="m-0">攻撃力</Form.Label>
                <div style={{display: 'flex', alignItems: 'center'}}>
                  <Form.Control className="py-0 px-1"
                                type="text" value={atkMin}
                                onChange={(e) => {setAtkMin(e.target.value)}}
                                style={{width: '3.5em'}}/>
                  <div className="mx-1">〜</div>
                  <Form.Control className="py-0 px-1"
                                type="text" value={atkMax}
                                onChange={(e) => {setAtkMax(e.target.value)}}
                                style={{width: '3.5em'}}/>
                </div>
              </div>
              <div style={{marginRight: '5em'}}>
                <Form.Label className="m-0">守備力</Form.Label>
                <div style={{display: 'flex', alignItems: 'center'}}>
                  <Form.Control className="py-0 px-1"
                                type="text" value={defMin}
                                onChange={(e) => {setDefMin(e.target.value)}}
                                style={{width: '3.5em'}}/>
                  <div className="mx-1">〜</div>
                  <Form.Control className="py-0 px-1"
                                type="text" value={defMax}
                                onChange={(e) => {setDefMax(e.target.value)}}
                                style={{width: '3.5em'}}/>
                </div>
              </div>
              <div className="mx-1">
                <Form.Label className="m-0">レベル</Form.Label>
                <div style={{display: 'flex', alignItems: 'center'}}>
                  <Button variant="dark" size="sm" className="m-1"
                          onClick={() => {
                            setLevelMin('0')
                            setLevelMax('4')
                          }}>下級</Button>
                  <Button variant="dark" size="sm" className="m-1"
                          onClick={() => {
                            setLevelMin('5')
                            setLevelMax('10')
                          }}>上級</Button>
                  <Form.Control className="py-0 px-1"
                                type="text" value={levelMin}
                                onChange={(e) => {setLevelMin(e.target.value)}}
                                style={{width: '3.5em'}}/>
                  <div className="mx-1">〜</div>
                  <Form.Control className="py-0 px-1"
                                type="text" value={levelMax}
                                onChange={(e) => {setLevelMax(e.target.value)}}
                                style={{width: '3.5em'}}/>
                </div>
              </div>
            </Col>
          </Row>

          <Row className="mx my-2" 
               hidden={!isIncludeSome(attribute, cardtypeList) && attribute.length != 0 }>
            <Col>
              <Form.Label className="m-0"
                          onClick={() => setShowMonsterAttribute(prev => !prev) }>
                属性
                <span hidden={showMonsterAttribute}><FiChevronUp/></span>
                <span hidden={!showMonsterAttribute}><FiChevronDown/></span>
              </Form.Label>
              <div hidden={!showMonsterAttribute}>
                {attributeMons.map((mt, index) => (
                  <ToggleElemButton key={index} list={monsterAttribute} elem={mt}
                                    setList={setMonsterAttribute} />
                ))}
              </div>
            </Col>
          </Row>

          <Row className="mx my-2" 
               hidden={!isIncludeSome(attribute, cardtypeList) && attribute.length != 0 }>
            <Col>
              <Form.Label className="m-0"
                          onClick={() => setShowMonsterType(prev => !prev) }>
                種族
                <span hidden={showMonsterType}><FiChevronUp/></span>
                <span hidden={!showMonsterType}><FiChevronDown/></span>
              </Form.Label>
              <div hidden={!showMonsterType}>
                {monstertypeList.map((mt, index) => (
                  <ToggleElemButton key={index} list={monstertype} elem={mt}
                                    setList={setMonstertype} />
                ))}
              </div>
            </Col>
          </Row>

          <Row>
            <Col>
              <div>
                <FloatingLabel controlId="floatingInput"
                               label="テキスト検索(完全一致)"
                               className="mb-3">
                  <Form.Control type="text" value={searchText}
                                onChange={e => setSearchText(e.target.value)} />
                </FloatingLabel>
              </div>
            </Col>
          </Row>

          <Row>
            <Col>
              <Button className="mx-3"
                      onClick={() => {setCardsWithFilterAndSort()}}>検索</Button>
              <Button onClick={() => {
                setAtkMax("5000"); setAtkMin("0"); setDefMax("5000"); setDefMin("0");
                setLevelMax("10"); setLevelMin("0"); setAttribute([]); setSearchText("");
                setMonstertype([]); setMonsterAttribute([]);
                setSortElem(ListOfSortableElemDefault)
                if(CardData){
                  setCards(
                    CardData.sort((a, b) => (
                      (ListOfSortableElemDefault.CompareKey(a) > ListOfSortableElemDefault.CompareKey(b)) == ListOfSortableElemDefault.IsDescending ? -1 : 1
                    ))
                  )
                }
              }}>リセット</Button>
            </Col>
          </Row>

          <Row>
            <Col>
              <Form.Label className="m-0" column="lg">ソート</Form.Label>
              <div style={{display: 'flex', alignItems: 'center', flexWrap: 'wrap'}}>
                {ListOfSortableElem.map((elem, index) => (
                  <div key={index} className="mx-1">
                    <Button size="sm"
                            onClick={() => {
                              if(elem.Name == sortElem.Name){
                                setSortElem((prev) => ({
                                  ...prev,
                                  IsDescending: !prev.IsDescending
                                }))
                              } else {
                                setSortElem(elem)
                              }
                            }}>
                      {elem.Name}
                      {elem.Name != sortElem.Name ? "" :
                       (sortElem.IsDescending ? <AiOutlineArrowDown/> : <AiOutlineArrowUp/>)}
                    </Button>
                  </div>
                ))}
              </div>
            </Col>
          </Row>

        </Form>
              </div>

              <div className="mx-3">
                該当枚数 {cards.length} / {CardData.length} 枚
              </div>

              <div className="m-3">
                <CardsDom cards={cards} limit={limit} />
              </div>

              <div style={{position: 'fixed', bottom: '1em', right: '1em'}}
                   onClick={() => {Scroll.animateScroll.scrollToTop({
                     smooth: 'linear',
                     delay: 0,
                     duration: 0,
                     isDynamic: false,
                   })}}>
                <FiChevronUp size="3em" />
              </div>

    </Container>
  )
}
