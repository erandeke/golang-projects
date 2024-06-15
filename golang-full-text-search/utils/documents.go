package utils

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// document represents a Wikipedia abstract dump document.
type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

// load documents will take the document path and return slice of documents and error if any

func LoadDocument(path string) ([]document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	gz, err := gzip.NewReader(f)

	if err != nil {
		return nil, err
	}
	defer gz.Close()

	// xml decoder
	dec := xml.NewDecoder(gz)

	//create the dump that actaul has the slice of documents

	dump := struct {
		Documents []document `xml:"doc"`
	}{}

	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	// get the slice of documents
	docs := dump.Documents

	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil

}
