import * as React from 'react'
import { Nav, Navbar, Container, Image } from 'react-bootstrap'
import Favicon from 'react-favicon'

import IconImage from './icon2.png'

export const Header = () => {
  return (
    <>
      <Favicon url={IconImage} />

      <Navbar bg="dark" variant="dark">
        <Container>
          <Navbar.Brand href="/">
            <Image src={IconImage} height="30" className="d-inline-block align-top"/>
          </Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="me-auto">
              <Nav.Link href="/search02">02環境</Nav.Link>
              <Nav.Link href="/search05">05環境</Nav.Link>
            </Nav>
          </Navbar.Collapse>
        </Container>
      </Navbar>
    </>
  )
}
