package util

import (
	"strings"
	"unicode"
)

func LowerCamelCase(s string) string {
	i := strings.IndexRune(s, '_')
	if i == 0 {
		return LowerCamelCase(s[i+1:])
	}

	if i > 0 {
		return LowerCamelCase(s[:i]) + UpperCamelCase(s[i+1:])
	}

	var ur = []rune(s)
	if unicode.IsLetter(ur[0]) {
		ur[0] = unicode.ToLower(ur[0])
	}
	return string(ur)
}

func UpperCamelCase(s string) string {
	i := strings.IndexRune(s, '_')
	if i == 0 {
		return UpperCamelCase(s[i+1:])
	}

	if i > 0 {
		return UpperCamelCase(s[:i]) + UpperCamelCase(s[i+1:])
	}

	var ur = []rune(s)
	if unicode.IsLetter(ur[0]) {
		ur[0] = unicode.ToUpper(ur[0])
	}
	return string(ur)
}

func LowerSnakeCase(s string) string {
	i := strings.IndexRune(s, '_')
	if i < 0 {
		var ur = []rune(s)
		for r := range ur {
			if r > 0 && unicode.IsUpper(ur[r]) {
				return LowerSnakeCase(s[:r]) + "_" + LowerSnakeCase(s[r:])
			}
		}
	}

	return strings.ToLower(s)
}

func UpperSnakeCase(s string) string {
	i := strings.IndexRune(s, '_')
	if i < 0 {
		var ur = []rune(s)
		for r := range ur {
			if r > 0 && unicode.IsUpper(ur[r]) {
				return UpperSnakeCase(s[:r]) + "_" + UpperSnakeCase(s[r:])
			}
		}
	}

	return strings.ToUpper(s)
}
