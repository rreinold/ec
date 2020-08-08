package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/jedib0t/go-pretty/table"
	ptable "github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

func main() {

	csvRaw := readFromStdin()
	table := createTable()
	csvReader := createCSVReader(csvRaw)

	for {
		csvRow, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			continue
		}
		tableRow := make([]interface{}, len(csvRow))
		for i := range tableRow {
			tableRow[i] = csvRow[i]
		}
		table.AppendRow(tableRow)

	}

	fmt.Println(table.Render())

}

func readFromStdin() string {

	if !StdInHasData() {
		fmt.Println("No data piped.")
		return ""
	}
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Failed to read stdin")
		os.Exit(1)
	}
	return string(bytes)

}

func createTable() ptable.Writer {
	t := table.NewWriter()
	t.SetCaption("Simple Table with 3 Rows.\n")
	t.SetStyle(table.StyleLight)
	t.Style().Format = table.FormatOptions{
		Footer: text.FormatLower,
		Header: text.FormatLower,
		Row:    text.FormatUpper,
	}
	t.Style().Options.DrawBorder = true
	t.Style().Options.SeparateColumns = true
	t.Style().Options.SeparateFooter = true
	t.Style().Options.SeparateHeader = true
	t.SetCaption("Table using the style 'funkyStyle'.\n")
	return t
}

func createCSVReader(input string) *csv.Reader {
	r := csv.NewReader(strings.NewReader(input))
	r.FieldsPerRecord = -1
	r.TrimLeadingSpace = true
	return r
}

func StdInHasData() bool {
	file := os.Stdin
	fi, err := file.Stat()

	if err != nil {
		fmt.Println("file.Stat()", err)
		return false
	}
	size := fi.Size()
	if size <= 1 {
		fmt.Println("Stdin is empty")
		return false
	}
	return true
}
