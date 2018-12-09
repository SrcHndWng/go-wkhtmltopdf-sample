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
	.table {
		border-collapse:collapse;
		border-spacing:0px; 
		border:1px solid #FF0000;
	}

	.table tr td:nth-child(1) {
		width:300px;
	}

	.table tr td {
		border:1px solid #000000;
	}
	</style>
	<html>
		<head>
			<title>PRINT PDF TEST</title>
			<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		</head>
		<body>
		<table class="table">
			<tr>
				<td>品名</td>
				<td>口径</td>
			</tr>
			<tr>
				<td>M16</td>
				<td>5.56</td>
			</tr>
			<tr>
				<td>AK47</td>
				<td>7.72</td>
			</tr>
			<tr>
				<td>MP5</td>
				<td>9.00</td>
			</tr>
	  </table>
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
