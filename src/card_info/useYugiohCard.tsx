import * as React from 'react'
import useAxios from 'axios-hooks'

export interface YugiohCard {
  imageurl: string[],
  url: string,
  text: string,
  name: string,
  atk: string,
  def: string,
  attribute: string,
  monstertype: string,
  level: number,
  releasedate: string,
  cardtype: string[]
}

export const YugiohCardDefault: YugiohCard = {
  imageurl: [""],
  url: "https://www.db.yugioh-card.com/yugiohdb/",
  text: "N/A",
  name: "N/A",
  atk: "N/A",
  def: "N/A",
  attribute: "N/A",
  monstertype: "N/A",
  level: 0,
  releasedate: "N/A",
  cardtype: []
}

// datestring: 2002-5-1
// 2002/5/1以前かどうか
export const is02 = (datestring: string, name: string) => {
  if(name == 'オシリスの天空竜') { return true}
  if(name == 'オベリスクの巨神兵') { return true}
  if(name == 'ラーの翼神竜') { return true}

  const [year, month, day] = datestring.split('-')
  const date = new Date(parseInt(year), parseInt(month) - 1, parseInt(day))
  return date <= new Date(2002, 4, 1)
}

// 2006/3/23依然かどうか
export const is04 = (datestring: string) => {
  const [year, month, day] = datestring.split('-')
  const date = new Date(parseInt(year), parseInt(month) - 1, parseInt(day))
  return date <= new Date(2006, 3, 23)
}

export const useYugiohCard = (name: string): YugiohCard => {
  const [yugiohCard, setYugiohCard] = React.useState<YugiohCard>({
      ...YugiohCardDefault,
      name,
  })
  const [resp, _] = useAxios({
    url: "/api/carddb-app/card",
    method: "get",
    params: {name}
  }, {
    manual: false
  })

  React.useEffect(() => {
    if(!resp.loading && !resp.error && resp.data) {
      setYugiohCard(resp.data)
    }
  }, [resp])

  return yugiohCard
}
