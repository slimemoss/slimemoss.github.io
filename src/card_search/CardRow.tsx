import * as React from 'react'
import { Badge, Col, Row } from 'react-bootstrap'
import { FiExternalLink } from 'react-icons/fi'

import { YugiohCard } from '../card_info/useYugiohCard'
import { yugiohCardColor } from '../card_info/color'

interface LimitBadgeProps {
  name: string
  limit: (name: string) => number
}
const LimitBadge = (props: LimitBadgeProps) => {
  const limit = props.limit(props.name)

  const str = (limit: number): string => {
    switch(limit) {
      case 0:
        return '禁止'
      case 1:
        return '制限'
      case 2:
        return '準制限'
      default:
    }
    return ''
  }

  return (
    <Badge bg="danger" hidden={limit == 3}>
      {str(limit)}
    </Badge>
  )
}

interface Props {
  card: YugiohCard
  limit: (name: string) => number
}

export const CardRow = (props: Props) => {
  const {card, limit} = props

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
          <LimitBadge name={card.name} limit={limit} />
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
