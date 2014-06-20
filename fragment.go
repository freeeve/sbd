package sbd

type Fragment struct {
	Orig        string
	Next        *Fragment
	EndsSegment bool
	Tokenized   bool
	Prediction  float64
	Label       int // ?
	Features    map[string]int
}
