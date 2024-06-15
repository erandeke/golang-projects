package main

import (
	"flag"
	"golang-full-text-search/utils"
	"log"
	"time"
)

func main() {

	var dumpPath string
	var queryString string
	flag.StringVar(&queryString, "query", "small wild cat", "search query")
	flag.StringVar(&dumpPath, "dump", "golang-full-text-search/enwiki-20240501-pages-articles-multistream-index1.txt-p1p41242.bz2", "dump")
	flag.Parse()

	// Log the start of the search process.
	log.Println("Full text search query is in prgogress:", queryString)

	// Record the start time of the search process.
	start := time.Now()

	// Load the document from the specified dump file path.
	// The LoadDocument function is assumed to be defined in the utils package.
	docs, err := utils.LoadDocument(dumpPath)

	// Handle any errors that may occur during the document loading process.
	if err != nil {
		log.Fatalf("Error loading document: %v", err)
	}

	//calculate the total time taken to load the document from the dump

	log.Println("Total time taken to load the document:", time.Since(start))

	//start creating the index on the documents

	start_time := time.Now()

	idx := make(utils.Index)

	//add the document to the index
	idx.Add(docs)

	//log how many documents were indexed in the time since the start
	log.Println("Indexed documents %d in %v", len(docs), time.Since(start_time))

	// Now we have the documents indexed .. we will search the query through the indexed docs and get the results
	// matched id's

	matchedIds := idx.Search(queryString)

	//Print the matched documents
	log.Printf("Search found documents %d", len(matchedIds))

	//Range over the matched documents and print the text

	for _, id := range matchedIds {
		doc := docs[id]
		log.Printf("%d%s\n", doc.Id, doc.Text)

	}

}
