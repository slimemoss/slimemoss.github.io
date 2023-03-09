import * as React from 'react'

import ReactPaginate from 'react-paginate';

import { YugiohCard } from '../card_info/useYugiohCard'
import { CardRow } from './CardRow'

interface Props {
  cards: YugiohCard[]
  limit: (name: string) => number
}
export const CardsDom = (props: Props) => {
  const {cards, limit} = props
  const maxItemPerPage = 30

  const [page, setPage] = React.useState(0)
  const scrollRef = React.useRef<HTMLDivElement>(null)

  React.useEffect(() => {
    setPage(0)
  }, [props.cards])

  const useEffectPage = React.useRef(false)
  React.useEffect(()=>{
    if(useEffectPage.current) {
      scrollRef?.current.scrollIntoView()
    }
    useEffectPage.current = true
  }, [page])

  return (
    <>
      <div ref={scrollRef}/>

      <div className="m-3">
        {cards.length}枚中 {maxItemPerPage * page + 1}〜{maxItemPerPage * page + maxItemPerPage}枚表示
      </div>

      {props.cards.slice(maxItemPerPage * page, maxItemPerPage * page + maxItemPerPage).map((card, index) => (
        <div key={index} className="p-2">
          <CardRow card={card} limit={limit} />
        </div>
      ))}

      <ReactPaginate
        nextLabel="next >"
        onPageChange={(e)=>{
          setPage(e.selected)
        }}
        pageRangeDisplayed={0}
        marginPagesDisplayed={0}
        pageCount={Math.ceil(cards.length / maxItemPerPage)}
        forcePage={page}
        previousLabel="< previous"
        pageClassName="page-item"
        pageLinkClassName="page-link"
        previousClassName="page-item"
        previousLinkClassName="page-link"
        nextClassName="page-item"
        nextLinkClassName="page-link"
        breakLabel="..."
        breakClassName="page-item"
        breakLinkClassName="page-link"
        containerClassName="pagination"
        activeClassName="active"
        renderOnZeroPageCount={null}
      />
    </>
  )

}
