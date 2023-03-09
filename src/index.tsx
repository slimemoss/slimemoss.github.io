import * as React from 'react'
import * as ReactDOM from 'react-dom'

import {Tab, Tabs} from 'react-bootstrap'

import { CardSearch02Page } from './card_search/CardSearch02'
import { CardSearch05Page } from './card_search/CardSearch02'

ReactDOM.render(
  <>
    <Tabs
      defaultActiveKey="s02"
      id="tab"
      className="mb-3"
    >
      <Tab eventKey="s02" title="02環境">
        <CardSearch02Page />
      </Tab>
      <Tab eventKey="s05" title="05環境">
        <CardSearch05Page />
      </Tab>
    </Tabs>
  </>,
  document.getElementById('app')
)
