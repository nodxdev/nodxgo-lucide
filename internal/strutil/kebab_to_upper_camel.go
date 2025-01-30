package strutil

import "strings"

// KebabToUpperCamel converts a kebab-case string to a UpperCamelCase string.
func KebabToUpperCamel(str string) string {
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

	return strings.Join(words, "")
}
