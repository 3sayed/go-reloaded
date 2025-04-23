package utils

import (
	"regexp"
	"strings"
)

func Punctuations(p string) string {
	

	// Step 1: Remove spaces before punctuation (like "Hello , world !" -> "Hello, world!")
	re := regexp.MustCompile(`\s+([.,!?;:])`)
	p = re.ReplaceAllString(p, "$1")

	// Step 2: Replace multiple spaces with a single space
	re = regexp.MustCompile(`\s{2,}`)
	p = re.ReplaceAllString(p, " ")

	re = regexp.MustCompile(`,([^\s])`)
    p = re.ReplaceAllString(p, ", $1")

	// Step 3: Trim leading and trailing spaces
	p = strings.TrimSpace(p)
	return p
}

func Qoute(q string ) string{
	re := regexp.MustCompile(`'\s+([^']+)\s+'`)
    q = re.ReplaceAllString(q, "'$1'")

	return q
}