package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	t := template.Must(template.ParseFiles("sample.tpl"))

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, nil); err != nil {
		log.Fatal(err)
	}
	html := tpl.String()

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
