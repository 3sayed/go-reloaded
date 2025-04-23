package utils

import (
	"regexp"
	"strconv"
	"strings"
)

// Upper converts preceding words to uppercase based on (up) and (up, n) patterns
func Up(input string) string {
	// Compile the regular expression to find (up) patterns
	re := regexp.MustCompile(`\(up(?:,\s*(\d+))?\)`)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatchIndex(input, -1)
	if len(matches) == 0 {
		return input
	}

	// Process matches from last to first to avoid position shifting
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		patternStart := match[0]
		patternEnd := match[1]

		// Default to converting 1 word if no number is specified
		wordCount := 1
		if len(match) > 2 && match[2] != -1 {
			// Extract the number from (up, n)
			numStr := input[match[2]:match[3]]
			if num, err := strconv.Atoi(numStr); err == nil && num > 0 {
				wordCount = num
			}
		}

		// Get the text before the pattern
		prefix := input[:patternStart]
		words := strings.Fields(prefix)
		if len(words) == 0 {
			// No words before the pattern, just remove the pattern
			input = input[patternEnd:]
			continue
		}

		// Determine how many words we can actually convert
		if wordCount > len(words) {
			wordCount = len(words)
		}

		// Convert the specified number of words to uppercase
		for j := len(words) - wordCount; j < len(words); j++ {
			words[j] = strings.ToUpper(words[j])
		}

		// Rebuild the string
		modifiedPrefix := strings.Join(words, " ")
		input = modifiedPrefix + input[patternEnd:]
	}

	return input
}