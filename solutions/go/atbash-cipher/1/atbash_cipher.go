package atbashcipher

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	plain := "abcdefghijklmnopqrstuvwxyz"
	cypher := []rune("zyxwvutsrqponmlkjihgfedcba")
	var r []rune 
	for _, c := range []rune(strings.ToLower(s)) {
		idx := strings.IndexRune(plain, c)

		if idx != -1 {
			r = append(r, cypher[idx])
			continue
		}

		if unicode.IsDigit(c) {
			r = append(r, c)
		}
	}
	return splitByFive(string(r))
}

func splitByFive(s string) string {
	var splited []string

	for len(s) > 5 {
		splited = append(splited, s[:5])
		s = s[5:]
	}

	if len(s) > 0 {
		splited = append(splited, s)
	}

	return strings.Join(splited, " ")
}
