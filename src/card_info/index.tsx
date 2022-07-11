import * as React from 'react'
import { useYugiohCard } from './useYugiohCard'

interface Props {
  name: string
}

export const YugiohCardLink = (props: Props) => {
  const yugiohCard = useYugiohCard(props.name)

  return (
    <a href={yugiohCard.url} target="_blank">{props.name}</a>
  )
}

export const YugiohCardText = (props: Props) => {
  const yugiohCard = useYugiohCard(props.name)

  return (
    <>
      {yugiohCard.text}
    </>
  )
}

export const YugiohCardImage = (props: {name: string, imagescale?: number}) => {
  const yugiohCard = useYugiohCard(props.name)

  return (
    <>
      <img src={"https://" + yugiohCard.imageurl[0]}
           alt={props.name}
           className="img-fluid"
           style={{ pointerEvents: 'none' }}
      />
    </>
  )
}
