package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	const html = `<!doctype html><html><head><title>PRINT PDF TEST</title></head><body><div>aaa</div><div>bbb</div><div>ccc</div></body></html>`

	pdfg := wkhtmltopdf.NewPDFPreparer()
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	pdfg.Dpi.Set(600)

	jsonBytes, err := pdfg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	pdfgFromJSON, err := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfgFromJSON.WriteFile("./sample.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("print done, pdf size %d bytes", pdfgFromJSON.Buffer().Len())
}
