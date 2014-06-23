package sbd

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TokenizeSuite struct{}

var _ = Suite(&TokenizeSuite{})

func (s *TokenizeSuite) TestTokenize(c *C) {
	// uniform quotes
	c.Assert(tokenize("''hello''"), Equals, "\"hello\"")
	c.Assert(tokenize("``hello``"), Equals, "\"hello\"")

	// separate punctuation (except period) from words
	c.Assert(tokenize("hello:"), Equals, "hello :")
	c.Assert(tokenize("hello:wut"), Equals, "hello : wut")
}
