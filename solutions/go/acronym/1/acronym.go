package acronym

import "strings"

// Abbreviate creates the acronym of a given string
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-' || r == '_'
	})

	var acr strings.Builder

	for _, w := range words {
		r := []rune(w)
		acr.WriteRune(r[0])
	}
	
	return strings.ToUpper(acr.String())
}
