package slug

import "regexp"

// Sub is a map of runes to strings.
// It is used to replace runes with strings.
// For example, if you want to replace "ş" with "s", you can use:
// slug.AddSub(slug.TR, slug.Sub{'ş': "s"})
// You can add multiple runes to string mappings.
type Sub map[rune]string

var (
	// MaxLength is the maximum length of a slug.
	MaxLength int = 50

	// NonAuthorizedChars is a regular expression that matches all non-authorized characters.
	// It is used to replace non-authorized characters with dashes.
	// For example, if you want to replace all non-authorized characters with dashes, you can use:
	// slug.NonAuthorizedChars = regexp.MustCompile("[^a-zA-Z0-9-_]")
	// You can use any regular expression you want.
	NonAuthorizedChars *regexp.Regexp = regexp.MustCompile("[^a-zA-Z0-9-_]")

	// MultipleDashes is a regular expression that matches multiple dashes.
	// It is used to replace multiple dashes with a single dash.
	// For example, if you want to replace multiple dashes with a single dash, you can use:
	// slug.MultipleDashes = regexp.MustCompile("-+")
	// You can use any regular expression you want.
	MultipleDashes *regexp.Regexp = regexp.MustCompile("-+")
)

// Generate a new slug.
// str is the string to be slugified.
// l is the language of the string. Default is English.
// If you want to use a different language, you can use:
// slug.New("Merhaba Dünya", slug.TR)
// You can use any language you want.
func New(str string, l ...Lang) string {
	return new(str, l...)
}

// Check if a string is a slug.
// slug is the string to be checked.
func Is(slug string) bool {
	return is(slug)
}

// Add a new sub.
// lang is the language of the sub.
// sub is the sub to be added.
// If you want to add a new sub, you can use:
// slug.AddSub(slug.TR, slug.Sub{'ş': "s"})
// You can use any language and sub you want.
func AddSub(lang Lang, sub Sub) {
	addSub(lang, sub)
}
