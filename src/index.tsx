import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { Route, BrowserRouter, Switch } from 'react-router-dom'
import { Header } from './header'
import { CardSearch02Page } from './card_search/CardSearch02'
import { CardSearch05Page } from './card_search/CardSearch02'

ReactDOM.render(
  <>
    <Header/>
    <BrowserRouter>
      <Switch>
        <Route exact path='/' component={CardSearch02Page} />
        <Route exact path='/search02' component={CardSearch02Page} />
        <Route exact path='/search05' component={CardSearch05Page} />
        <CardSearch02Page/>
      </Switch>
    </BrowserRouter>
  </>,
  document.getElementById('app')
)
