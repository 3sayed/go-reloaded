package utils

import "strings"

func Vowel(v string) string {
	var builder strings.Builder
	builder.Grow(len(v)) 

	for i:=0; i < len(v); i++ {
		current := v[i]

		if current == 'a' || current == 'A' && i+2 < len(v) {
			next := v[i+1]
			nextOfnext := v[i+2]
			
			if next == ' ' && isVowel(nextOfnext) {
				builder.WriteByte(current)
				builder.WriteByte('n')
				continue
			}
			
		}
		builder.WriteByte(current)
	}
	return builder.String()
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u', 'h':
		return true
	default:
		return false	
	}
    
}