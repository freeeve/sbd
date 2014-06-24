package sbd

import (
	"fmt"
	"regexp"
	"strings"
)

func uniformQuotes(s string) string {
	//(re.compile(r'\'\''), r'"'),
	//(re.compile(r'\`\`'), r'"'),
	s = regexp.MustCompile("''").ReplaceAllString(s, "\"")
	s = regexp.MustCompile("``").ReplaceAllString(s, "\"")
	return s
}

func separatePunctuation(s string) string {
	// it's supposed to do this. it might not.
	//(re.compile(r'(^|\s)(\')'), r'\1\2 '),
	//(re.compile(r'(?=[\(\"\`{\[:;&\#\*@])(.)'), r'\1 '),

	//(re.compile(r'(.)(?=[?!)\";}\]\*:@\'])'), r'\1 '),
	//(re.compile(r'(?=[\)}\]])(.)'), r'\1 '),
	//(re.compile(r'(.)(?=[({\[])'), r'\1 '),
	//(re.compile(r'((^|\s)\-)(?=[^\-])'), r'\1 '),
	r := strings.NewReader(s)
	retVal := ""
	justSpaced := false
	for c, _, err := r.ReadRune(); err == nil; c, _, err = r.ReadRune() {
		switch c {
		case ':', ';', '&', '\\', '/', '(', ')', '#', '*', '@', '\'', '{', '}', '-', '%':
			retVal = fmt.Sprintf("%s %c", retVal, c)
			justSpaced = true
		default:
			if justSpaced {
				retVal = fmt.Sprintf("%s ", retVal)
				justSpaced = false
			}
			retVal = fmt.Sprintf("%s%c", retVal, c)
		}
	}
	return retVal
}
func doubleHyphenToOne(s string) string {
	//(re.compile(r'([^-])(\-\-+)([^-])'), r'\1 \2 \3'),
	//(re.compile(r'(\s|^)(,)(?=(\S))'), r'\1\2 '),
	s = regexp.MustCompile("([^-])(\\-\\-+)([^-])").ReplaceAllString(s, "$1 $2 $3")
	s = regexp.MustCompile("(\\-)\\s+(\\-)").ReplaceAllString(s, "$1$2")
	return s
}

func commaSpaceFollows(s string) string {
	//Only separate comma if space follows:
	//(re.compile(r'(.)(,)(\s|$)'), r'\1 \2\3'),
	s = regexp.MustCompile("(.)(,)(\\s|$)").ReplaceAllString(s, "$1 $2$3")
	return s
}

func combineDots(s string) string {
	//# Combine dots separated by whitespace to be a single token:
	//(re.compile(r'\.\s\.\s\.'), r'...'),
	s = regexp.MustCompile("\\.\\s\\.\\s\\.").ReplaceAllString(s, "...")
	return s
}
func separateNoNum(s string) string {
	//# Separate "No.6"
	//(re.compile(r'([A-Za-z]\.)(\d+)'), r'\1 \2'),
	s = regexp.MustCompile("([A-Za-z]\\.)(\\d+)").ReplaceAllString(s, "$1 $2")
	return s
}
func separateWordsFromEllipses(s string) string {
	//# Separate words from ellipses
	//(re.compile(r'([^\.]|^)(\.{2,})(.?)'), r'\1 \2 \3'),
	//(re.compile(r'(^|\s)(\.{2,})([^\.\s])'), r'\1\2 \3'),
	//(re.compile(r'([^\.\s])(\.{2,})($|\s)'), r'\1 \2\3'),
	s = regexp.MustCompile("")
	return s
}

func fixPercentDollarAmpersand(s string) string {
	//# fix %, $, &
	//(re.compile(r'(\d)%'), r'\1 %'),
	//(re.compile(r'\$(\.?\d)'), r'$ \1'),
	//(re.compile(r'(\w)& (\w)'), r'\1&\2'),
	//(re.compile(r'(\w\w+)&(\w\w+)'), r'\1 & \2'),
	// is this necessary ?
	s = regexp.MustCompile("(\\w)& (\\w)").ReplaceAllString(s, "$1&$2")
	s = regexp.MustCompile("(\\w\\w+)&(\\w\\w+)").ReplaceAllString(s, "$1 & $2")
	return s
}

func fixNt(s string) string {
	//# fix (n 't) --> ( n't)
	//(re.compile(r'n \'t( |$)'), r" n't\1"),
	//(re.compile(r'N \'T( |$)'), r" N'T\1"),

	return s
}

func separateCant(s string) string {
	//# treebank tokenizer special words
	//(re.compile(r'([Cc])annot'), r'\1an not'),

	return s
}

func compactSpaces(s string) string {
	//# compact spaces into just one
	//(re.compile(r'\s+'), r' '),
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
