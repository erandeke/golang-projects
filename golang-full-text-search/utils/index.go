package utils

// type is the base interface for all data types in golang .
// This means that all data types in golang implements the type interface

// Index is an inverted index . It maps tokens to documents Ids
// exmple "I love India" --> I :document[1], Love:document[2], India:document[3]
type Index map[string]int

// function to add indexing to slice of documents
func (idx Index) Add(doc []document) {

}
