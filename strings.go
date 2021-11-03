package goutils

import "strings"

// TitleCasedName
//example： is-extis-link to IsExtisLink
func TitleCasedName(name string) string {
	newStr := make([]rune, 0)
	upNextChar := true

	name = strings.ToLower(name)

	for _, chr := range name {
		switch {
		case upNextChar:
			upNextChar = false
			if 'a' <= chr && chr <= 'z' {
				chr -= ('a' - 'A')
			}
		case chr == '-':
			upNextChar = true
			continue
		}

		newStr = append(newStr, chr)
	}

	return string(newStr)
}
