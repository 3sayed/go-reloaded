package utils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Cap(c string) string {
	// Regular expression to find (cap) or (cap, n) patterns
	re := regexp.MustCompile(`\(cap(?:,\s*(\d+))?\)`)
	
	// Find all matches in the string
	matches := re.FindAllStringSubmatchIndex(c, -1)
	if len(matches) == 0 {
		return c
	}

	// Process from last match to first to avoid position shifting
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		start, end := match[0], match[1]

		// Default to 1 word if no number is specified
		n := 1
		if len(match) > 2 && match[2] != -1 {
			numStr := c[match[2]:match[3]]
			if num, err := strconv.Atoi(numStr); err == nil && num > 0 {
				n = num
			}
		}

		// Get the text before the pattern
		prefix := c[:start]
		words := strings.Fields(prefix)
		if len(words) == 0 {
			continue
		}

		// Determine how many words we can actually capitalize
		if n > len(words) {
			n = len(words)
		}

		// Capitalize the last n words
		for i := len(words) - n; i < len(words); i++ {
			if len(words[i]) > 0 {
				runes := []rune(words[i])
				runes[0] = unicode.ToUpper(runes[0])
				words[i] = string(runes)
			}
		}

		// Rebuild the string
		modifiedPrefix := strings.Join(words, " ")
		c = modifiedPrefix + c[end:]
	}

	return c
}