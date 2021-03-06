package main

import (
	"encoding/csv"
	"flag"
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
	var filepath *string
	if len(os.Args) == 2 {
		filepath = &os.Args[1]
	} else {
		filepath = flag.String("f", "", "filename")

	}
	flag.Parse()

	table, err := read(*filepath)
	if err != nil {
		return
	}
	fmt.Println(table.Render())

}

func read(path string) (ptable.Writer, error) {
	var r *csv.Reader

	if path != "" {
		file, OSErr := os.Open(path)
		if OSErr != nil {
			fmt.Println(OSErr)
			return nil, OSErr
		}
		r = csv.NewReader(file)
	} else {
		stdin := readFromStdin()
		input := strings.NewReader(stdin)
		r = csv.NewReader(input)

	}

	r = configureCSVReader(r)
	r.FieldsPerRecord = -1
	r.TrimLeadingSpace = true

	table := createTable()
	for {
		csvRow, err := r.Read()
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
	return table, nil

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
	t.SetStyle(table.StyleLight)
	t.Style().Format = table.FormatOptions{
		Footer: text.FormatLower,
		Header: text.FormatLower,
	}
	t.Style().Options.DrawBorder = true
	t.Style().Options.SeparateColumns = true
	t.Style().Options.SeparateFooter = true
	t.Style().Options.SeparateHeader = true
	return t
}

func configureCSVReader(r *csv.Reader) *csv.Reader {

	// Enable variable number of columns per entry
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
