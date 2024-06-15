package utils

import (
	"strings"
	"unicode"
)

// tokenize returns a slice of tokens for the given text.
func tokenize(text string) []string {
	//Field func splits the string into tokens separated by delimiters .
	// in this case the delimiters are any characters that are not letters , Numbers, punctuation
	return strings.FieldsFunc(text, func(r rune) bool {
		// rune is the single unicode character
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && !unicode.IsPunct(r) && !unicode.IsSpace(r)
	})

}

// analyze function would do the following

/*
 1. takes the input as any string and return the array of strings
    1.1 tokenises the input string
    1.2 lowercase the input string as cat Cat  caT should be equal to cat
    1.3 remove the stopwords from the input string " I should be doing something" stopwards are the common words like "I", "should", "be"
    1.4 stem the input string to correct grammatically for example  everything is like when
    // you say likely , likes, liking
*/
func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = filterStopWords(tokens)
	tokens = stemmerFilter(tokens)
	return tokens

}
