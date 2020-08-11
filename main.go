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
	var csvFilepath = flag.String("f", "", "filename")

	flag.Parse()
	fmt.Println(*csvFilepath)

	table := read(*csvFilepath)
	fmt.Println(table.Render())

}

func read(path string) ptable.Writer {
	var r *csv.Reader

	if path != "" {
		file, OSErr := os.Open(path)
		if OSErr != nil {
			// TODO Throw
			log.Fatalln("Couldn't open the csv file", OSErr)
			r = csv.NewReader(file)
		}
	} else {
		stdin := readFromStdin()
		input := strings.NewReader(stdin)
		r = csv.NewReader(input)

	}
	r = configureCSVReader(r)

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
	return table

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
		Row:    text.FormatUpper,
	}
	t.Style().Options.DrawBorder = true
	t.Style().Options.SeparateColumns = true
	t.Style().Options.SeparateFooter = true
	t.Style().Options.SeparateHeader = true
	t.SetCaption("Table name'.\n")
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
