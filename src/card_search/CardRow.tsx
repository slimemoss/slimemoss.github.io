import * as React from 'react'
import { Badge, Col, Row } from 'react-bootstrap'
import { FiExternalLink } from 'react-icons/fi'

import { YugiohCard } from '../card_info/useYugiohCard'
import { yugiohCardColor } from '../card_info/color'
import { limit02 } from '../card_info/env02'

interface Props {
  card: YugiohCard
}

const LimitBadge = (props: {name: string}) => {
  const limit = limit02(props.name)
  return (
    <Badge bg="danger" hidden={limit == 3}>
      {limit == 1 ? '制限' : '準制限'}
    </Badge>
  )
}

export const CardRow = (props: Props) => {
  const {card} = props

  return (
    <>
      <Row>
        <Col style={{display: 'flex', alignItems: 'center',
                     background: yugiohCardColor(card.attribute, card.cardtype)}}
             sm xs="12">
          <div className="p-1"
               style={{fontSize: '1.2em'}}>
            {card.name}
          </div>
          <LimitBadge name={card.name} />
          <div style={{marginLeft: 'auto'}}>
            <a href={card.url} target='_blank'>公式<FiExternalLink/></a>
          </div>
        </Col>

        <Col style={{display: 'flex', flexWrap: 'wrap', alignItems: 'center'}}>
          <div className="mx-2" style={{minWidth: '9em'}}>
            {card.attribute}
            {card.monstertype ? '/' : ''}
            {card.monstertype}
          </div>
          <div className="mx-2" style={{minWidth: '4em'}}>
            {card.level == 0 ? '' : 'レベル' + card.level}
          </div>
          <div className="mx-2">
            {card.atk ? '攻' : ''}
            {card.atk}
            {card.atk ? '/守' : ''}
            {card.def}
          </div>
        </Col>
      </Row>
      <Row style={{borderTop: '2px solid', borderTopColor: '#333333'}}>
        <Col>
          {card.text}
        </Col>
      </Row>
    </>
  )
}
