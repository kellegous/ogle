package query

import (
  "io"
)

type Token int

const (
  NONE Token = iota
  IDENT
  LPAREN // (
  RPAREN // )
  TILDE  // ~
  AMP    // &
  PIPE   // |
)

type Parser struct {
  d []byte
  i int
}

func Parse(data []byte) *Parser {
  return &Parser{
    d: data,
  }
}

func eatIdent(p *Parser) string {
}

func (p *Parser) Next() (Token, string, error) {
  if len(p.d) == 0 {
    return NONE, "", io.EOF
  }

  b := p.d[p.i]
  switch b {
  case '&':
    return AMP, "", nil
  case '|':
    return PIPE, "", nil
  case '~':
    return TILDE, "", nil
  case '(':
    return LPAREN, "", nil
  case ')':
    return RPAREN, "", nil
  default:
    return IDENT, eatIdent(p), nil
  }
}
