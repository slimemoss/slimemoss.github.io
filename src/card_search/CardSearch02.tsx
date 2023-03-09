import * as React from 'react'
import { CardSearchPage } from './CardSearchPage';
import { limit02 } from '../card_info/env02';
import { limit05 } from '../card_info/env05';

import CardData from './data.json'
import CardData05 from './data05.json'

export const CardSearch02Page = () => {
  return <CardSearchPage title='02環境 カード検索 | slimemoss'
                         CardData={CardData}
                         limit={limit02}
         />
}

export const CardSearch05Page = () => {
  return <CardSearchPage title='02環境 カード検索 | slimemoss'
                         CardData={CardData05}
                         limit={limit05}
         />
}
