package util

import (
	"regexp"
	"strings"
)

func GenerateSlug(text string) string {
	nonAuthorizedChars := regexp.MustCompile("[^a-zA-Z0-9-_]")
	multipleDashes := regexp.MustCompile("-+")

	text = nonAuthorizedChars.ReplaceAllString(text, "-")
	text = multipleDashes.ReplaceAllString(text, "-")
	text = strings.Trim(strings.ToLower(text), "-_")

	return text
}

// IsSlug returns True if provided text does not contain white characters,
// punctuation, all letters are lower case and only from ASCII range.
// It could contain `-` and `_` but not at the beginning or end of the text.
// It should be in range of the MaxLength var if specified.
// All output from slug.Make(text) should pass this test.
func IsSlug(text string) bool {
	if text == "" {
		return false
	}

	for _, c := range text {
		if (c < 'a' || c > 'z') && c != '-' && c != '_' && (c < '0' || c > '9') {
			return false
		}
	}

	return true
}
