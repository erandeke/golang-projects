package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

// Lowecase filters applies on the tokens and turn them into lower
func lowercaseFilter(tokens []string) []string {
	r := make([]string, 0, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

// filter the stop words

func filterStopWords(tokens []string) []string {
	var stopwords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}

	r := make([]string, 0, len(tokens))

	for _, token := range tokens {
		if _, ok := stopwords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

// stemmerFilter returns a slice of stemmed tokens.
func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
