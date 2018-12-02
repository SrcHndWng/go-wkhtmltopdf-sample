package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	const html = `<!doctype html>
	<style>
		.line {width:100px; border:#000000 solid 1px}
	</style>
	<html>
			<head>
				<title>PRINT PDF TEST</title>
				<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
			</head>
		<body>
			<div class="line">あああ</div>
			<div class="line">いいい</div>
			<div class="line">ううう</div>
		</body>
	</html>`

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
