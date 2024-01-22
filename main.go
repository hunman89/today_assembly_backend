package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func main() {
	// Replace "input.pdf" with the path to your PDF file.
	inputPath := "data/test.pdf"

	// Create a new PDF reader.
	pdfReader, f, err := model.NewPdfReaderFromFile(inputPath, nil)
	defer f.Close()
	if err != nil {
		log.Fatalf("Error creating PDF reader: %v", err)
	}

	// Get the number of pages in the PDF.
	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		log.Fatalf("Error getting number of pages: %v", err)
	}

	// Extract text from each page and concatenate it into a single string.
	var textBuilder strings.Builder
	for pageNum := 1; pageNum <= numPages; pageNum++ {
		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			log.Fatalf("Error getting page %d: %v", pageNum, err)
		}

		extractor, err := extractor.New(page)
		if err != nil {
			log.Fatalf("Error extracting New: %v", err)
		}
		text, err := extractor.ExtractText()
		if err != nil {
			log.Fatalf("Error extracting text from page %d: %v", pageNum, err)
		}

		textBuilder.WriteString(text)
	}

	// Print the extracted text.
	fmt.Println(textBuilder.String())
}
