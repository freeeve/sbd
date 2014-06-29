package sbd

import (
	. "gopkg.in/check.v1"
)

type FragmentSuite struct{}

var _ = Suite(&FragmentSuite{})

func (s *FragmentSuite) TestString(c *C) {
	f := Fragment{Orig: "hello"}
	c.Assert(f.String(), Equals, "hello")
	f.EndsSegment = true
	c.Assert(f.String(), Equals, "hello <EOS> ")
}
