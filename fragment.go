package sbd

import "fmt"

type Fragment struct {
	Orig        string
	EndsSegment bool
	Tokenized   bool
	Prediction  float64
	Label       int // ?
	Features    map[string]int
}

func (f Fragment) String() string {
	if f.EndsSegment {
		return fmt.Sprintf("%s <EOS> ", f.Orig)
	}
	return fmt.Sprintf("%s", f.Orig)
}
