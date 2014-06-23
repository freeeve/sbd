package sbd

import (
	"regexp"
)

func uniformQuotes(s string) string {
	// uniform quotes
	s = regexp.MustCompile("''").ReplaceAllString(s, "\"")
	s = regexp.MustCompile("``").ReplaceAllString(s, "\"")
   return s
}

func separatePunctuation(s string) string {
	return s
}
func doubleHyphenToOne(s string) string {
	return s
}

func commaSpaceFollows(s string) string {
	return s
}

func combineDots(s string) string {
	return s
}
func separateNoNum(s string) string {
	return s
}
func separateWordsFromEllipses(s string) string {
	return s
}

func fixPercentDollarAmpersand(s string) string {
	return s
}

func fixNt(s string) string {
	return s
}

func separateCant(s string) string {
	return s
}

func compactSpaces(s string) string {
	return s
}

func tokenize(s string) string {
	s = uniformQuotes(s)
	s = separatePunctuation(s)
	s = doubleHyphenToOne(s)
	s = commaSpaceFollows(s)
	s = combineDots(s)
	s = separateNoNum(s)
	s = separateWordsFromEllipses(s)
	s = fixPercentDollarAmpersand(s)
	s = fixNt(s)
	s = separateCant(s)
	s = compactSpaces(s)
	return s
}

/*
"""
A list of (regexp, repl) pairs applied in sequence.
The resulting string is split on whitespace.
(Adapted from the Punkt Word Tokenizer)
"""

_tokenize_regexps = [

    # uniform quotes
    (re.compile(r'\'\''), r'"'),
    (re.compile(r'\`\`'), r'"'),

    # Separate punctuation (except period) from words:
    (re.compile(r'(^|\s)(\')'), r'\1\2 '),
    (re.compile(r'(?=[\(\"\`{\[:;&\#\*@])(.)'), r'\1 '),

    (re.compile(r'(.)(?=[?!)\";}\]\*:@\'])'), r'\1 '),
    (re.compile(r'(?=[\)}\]])(.)'), r'\1 '),
    (re.compile(r'(.)(?=[({\[])'), r'\1 '),
    (re.compile(r'((^|\s)\-)(?=[^\-])'), r'\1 '),

    # Treat double-hyphen as one token:
    (re.compile(r'([^-])(\-\-+)([^-])'), r'\1 \2 \3'),
    (re.compile(r'(\s|^)(,)(?=(\S))'), r'\1\2 '),

    # Only separate comma if space follows:
    (re.compile(r'(.)(,)(\s|$)'), r'\1 \2\3'),

    # Combine dots separated by whitespace to be a single token:
    (re.compile(r'\.\s\.\s\.'), r'...'),

    # Separate "No.6"
    (re.compile(r'([A-Za-z]\.)(\d+)'), r'\1 \2'),

    # Separate words from ellipses
    (re.compile(r'([^\.]|^)(\.{2,})(.?)'), r'\1 \2 \3'),
    (re.compile(r'(^|\s)(\.{2,})([^\.\s])'), r'\1\2 \3'),
    (re.compile(r'([^\.\s])(\.{2,})($|\s)'), r'\1 \2\3'),

    ## adding a few things here:

    # fix %, $, &
    (re.compile(r'(\d)%'), r'\1 %'),
    (re.compile(r'\$(\.?\d)'), r'$ \1'),
    (re.compile(r'(\w)& (\w)'), r'\1&\2'),
    (re.compile(r'(\w\w+)&(\w\w+)'), r'\1 & \2'),

    # fix (n 't) --> ( n't)
    (re.compile(r'n \'t( |$)'), r" n't\1"),
    (re.compile(r'N \'T( |$)'), r" N'T\1"),

    # treebank tokenizer special words
    (re.compile(r'([Cc])annot'), r'\1an not'),

    (re.compile(r'\s+'), r' '),
*/
