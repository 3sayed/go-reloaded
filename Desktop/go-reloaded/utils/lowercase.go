package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// // Lower processes a string to convert words preceding (low) patterns to lowercase
// func Low(input string) string {
// 	// Compile the regular expression to find (low) patterns
// 	re := regexp.MustCompile(`\(low(?:,\s*(\d+))?\)`)

// 	// Find all matches in the input string
// 	matches := re.FindAllStringSubmatchIndex(input, -1)
// 	if len(matches) == 0 {
// 		return input
// 	}

// 	// Process matches from last to first to avoid position shifting
// 	for i := len(matches) - 1; i >= 0; i-- {
// 		match := matches[i]
// 		patternStart := match[0]
// 		patternEnd := match[1]

// 		// Default to converting 1 word if no number is specified
// 		wordCount := 1
// 		if len(match) > 2 && match[2] != -1 {
// 			// Extract the number from (low, n)
// 			numStr := input[match[2]:match[3]]
// 			if num, err := strconv.Atoi(numStr); err == nil && num > 0 {
// 				wordCount = num
// 			}
// 		}

// 		// Get the text before the pattern
// 		prefix := input[:patternStart]
// 		words := strings.Fields(prefix)
// 		if len(words) == 0 {
// 			// No words before the pattern, just remove the pattern
// 			input = input[patternEnd:]
// 			continue
// 		}

// 		// Determine how many words we can actually convert
// 		if wordCount > len(words) {
//             log.Fatalf("ERROR: Cannot lowercase %d words - only %d words exist before (low) at position %d",
//                 wordCount, len(words), patternStart)
//             // Program will exit here
//         }

// 		// Convert the specified number of words to lowercase
// 		for j := len(words) - wordCount; j < len(words); j++ {
// 			words[j] = strings.ToLower(words[j])
// 		}

// 		// Rebuild the string
// 		modifiedPrefix := strings.Join(words, " ")
// 		input = modifiedPrefix + input[patternEnd:]
// 	}

// 	return input
// }

func Low(input string) string {
    re := regexp.MustCompile(`\(low(?:,\s*(\d+))?\)`)
    matches := re.FindAllStringSubmatchIndex(input, -1)
    if len(matches) == 0 {
        return input
    }

    for i := len(matches) - 1; i >= 0; i-- {
        match := matches[i]
        patternStart := match[0]
        patternEnd := match[1]

        wordCount := 1
        if len(match) > 2 && match[2] != -1 {
            numStr := input[match[2]:match[3]]
            num, err := strconv.Atoi(numStr);
			if err != nil || num <= 0 {
                log.Fatalf("ERROR: Invalid number format in (low, %s)", numStr)
            }
			wordCount = num
        }

        prefix := input[:patternStart]
        words := strings.Fields(prefix)
        if len(words) == 0 {
            input = input[patternEnd:]
            continue
        }

        // Critical change: Stop execution if wordCount exceeds available words
        if wordCount > len(words) {
			log.Fatalf("ERROR: Invalid number, formated number is larger than range")
            // Program will exit here
        }

        for j := len(words) - wordCount; j < len(words); j++ {
			if strings.Contains(words[j], ",") {
                log.Fatalf("ERROR: Invalid number, formated number is out of the rage")
            }else{
				words[j] = strings.ToLower(words[j])
			}
        }

        input = strings.Join(words, " ") + input[patternEnd:]
    }
    return input
}