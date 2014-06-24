package sbd

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TokenizeSuite struct{}

var _ = Suite(&TokenizeSuite{})

func (s *TokenizeSuite) TestUniformQuotes(c *C) {
	// uniform quotes
	c.Assert(uniformQuotes("''hello''"), Equals, "\"hello\"")
	c.Assert(uniformQuotes("``hello``"), Equals, "\"hello\"")
}

func (s *TokenizeSuite) TestSeparatePunctuation(c *C) {
	// separate punctuation (except period) from words
	c.Assert(separatePunctuation("hello:"), Equals, "hello :")
	c.Assert(separatePunctuation("hello:wut"), Equals, "hello : wut")
	c.Assert(separatePunctuation("hello;"), Equals, "hello ;")
	c.Assert(separatePunctuation("hello;wut"), Equals, "hello ; wut")
	c.Assert(separatePunctuation("hello&"), Equals, "hello &")
	c.Assert(separatePunctuation("hello&wut"), Equals, "hello & wut")
	c.Assert(separatePunctuation("hello\\"), Equals, "hello \\")
	c.Assert(separatePunctuation("hello\\wut"), Equals, "hello \\ wut")
	c.Assert(separatePunctuation("hello/"), Equals, "hello /")
	c.Assert(separatePunctuation("hello/wut"), Equals, "hello / wut")
	c.Assert(separatePunctuation("hello("), Equals, "hello (")
	c.Assert(separatePunctuation("hello(wut"), Equals, "hello ( wut")
	c.Assert(separatePunctuation("hello)"), Equals, "hello )")
	c.Assert(separatePunctuation("hello)wut"), Equals, "hello ) wut")
	c.Assert(separatePunctuation("hello#"), Equals, "hello #")
	c.Assert(separatePunctuation("hello#wut"), Equals, "hello # wut")
	c.Assert(separatePunctuation("hello*"), Equals, "hello *")
	c.Assert(separatePunctuation("hello*wut"), Equals, "hello * wut")
	c.Assert(separatePunctuation("hello@"), Equals, "hello @")
	c.Assert(separatePunctuation("hello@wut"), Equals, "hello @ wut")
	c.Assert(separatePunctuation("hello'"), Equals, "hello '")
	c.Assert(separatePunctuation("hello'wut"), Equals, "hello ' wut")
	c.Assert(separatePunctuation("hello{"), Equals, "hello {")
	c.Assert(separatePunctuation("hello{wut"), Equals, "hello { wut")
	c.Assert(separatePunctuation("hello}"), Equals, "hello }")
	c.Assert(separatePunctuation("hello}wut"), Equals, "hello } wut")
	c.Assert(separatePunctuation("hello-"), Equals, "hello -")
	c.Assert(separatePunctuation("hello-wut"), Equals, "hello - wut")
	c.Assert(separatePunctuation("hello%"), Equals, "hello %")
	c.Assert(separatePunctuation("hello%wut"), Equals, "hello % wut")
}

func (s *TokenizeSuite) TestDoubleHyphenToOne(c *C) {
	c.Assert(doubleHyphenToOne("hello--wut"), Equals, "hello -- wut")
}

func (s *TokenizeSuite) TestCommaSpaceFollows(c *C) {
	c.Assert(commaSpaceFollows("hello,wut"), Equals, "hello,wut")
	c.Assert(commaSpaceFollows("hello, wut"), Equals, "hello , wut")
}

func (s *TokenizeSuite) TestCombineDots(c *C) {
	c.Assert(combineDots(". . ."), Equals, "...")
}

func (s *TokenizeSuite) TestSeparateNoNum(c *C) {
	c.Assert(separateNoNum("No.6"), Equals, "No. 6")
}

func (s *TokenizeSuite) TestSeparateWordsFromEllipses(c *C) {
	c.Assert(separateWordsFromEllipses("hello..."), Equals, "hello ...")
}

func (s *TokenizeSuite) TestFixPercentDollarAmpersand(c *C) {
	c.Assert(fixPercentDollarAmpersand("hello& wut"), Equals, "hello & wut")
	c.Assert(fixPercentDollarAmpersand("hello&wut"), Equals, "hello & wut")
}

func (s *TokenizeSuite) TestTokenize(c *C) {
	c.Assert(tokenize("hello--wut"), Equals, "hello -- wut")
	c.Assert(tokenize("hello--wut--foo"), Equals, "hello -- wut -- foo")
	c.Assert(tokenize("60%"), Equals, "60 %")
	c.Assert(tokenize("hello..."), Equals, "hello ...")
}
