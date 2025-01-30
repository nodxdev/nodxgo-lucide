package strutil

import "strings"

// kebabToCapitalized converts a kebab-case string to a capitalized sentence.
func kebabToCapitalized(str string) string {
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
