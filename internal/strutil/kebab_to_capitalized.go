package strutil

import "strings"

// KebabToCapitalized converts a kebab-case string to a capitalized sentence.
func KebabToCapitalized(str string) string {
	if str == "" {
		return ""
	}

	words := strings.Split(str, "-")
	for i, word := range words {
		if word == "" {
			continue
		}
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}

	return strings.Join(words, " ")
}
