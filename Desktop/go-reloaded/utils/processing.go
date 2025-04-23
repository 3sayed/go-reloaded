package utils

import (
	"strings"
)

func ProcessText(txt string) string {
    processed := txt
    
    // 1. First convert numbers (these might affect word counts)
    if strings.Contains(processed, "(hex)") {
        processed = Hex(processed)
    }
    if strings.Contains(processed, "(bin)") {
        processed = Bin(processed)
    }

    // // 2. Handle case modifications
    if strings.Contains(processed, "(low") {
        processed = Low(processed)
    }
    if strings.Contains(processed, "(up") {
        processed = Up(processed)
    }
    if strings.Contains(processed, "(cap") {
        processed = Cap(processed)
    }

    // // 3. Fix articles (a → an before vowels)
    processed = Vowel(processed)

    // // 4. Handle punctuation spacing
    processed = Punctuations(processed)

    // // 5. Clean quotes and final spaces
    processed = Qoute(processed)

    return processed
}