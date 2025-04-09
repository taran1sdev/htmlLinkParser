package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"io"
	"path/filepath"
	"link/link"
)

// Get the html file from the user

var htmlFile = flag.String("htmlFile", "test/ex1.html", "The html file to parse for links")

// Return the io.Reader type from the file provided

func getReader(fileName string) io.Reader {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open the file.")
	}
	return f
}

func main() {
	flag.Parse()
	
	// Check the file provided by the user is a valid html file
	ext := filepath.Ext(*htmlFile)

	if ext != ".html" {
		log.Fatalf("File must be a html file!")
	}

	
	reader := getReader(*htmlFile)
	
	anchors, err := link.ParseAnchors(reader)
	if err != nil {
		log.Fatalf("Error occured while parsing Anchor tags:", err)
	}
	
	for _, a := range anchors {
		fmt.Println("href:", a.Href, "|", "Text:", a.Text)
	}
}
