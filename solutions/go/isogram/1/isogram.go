package isogram

import (
	"strings"
	"unicode"
)
func IsIsogram(word string) bool {
	letters := make([]string, len(word))
	
	for _, l := range strings.ToLower(word) {
		if !unicode.IsLetter(l) {
			continue
		}
		
		for _, g := range letters {
			if g == string(l) {
				return false
			}
		}
		letters = append(letters, string(l))
	}

	return true
}
