package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func main() {
	fmt.Println("Hellooooo")
	bytes, err := ioutil.ReadAll(os.Stdin)

	log.Println(err, string(bytes))
	t := table.NewWriter()
	// a row need not be just strings
	t.AppendRow(table.Row{1, "Arya", "Stark", 3000})
	// all rows need not have the same number of columns
	t.AppendRow(table.Row{20, "Jon", "Snow", 2000, "You know nothing, Jon Snow!"})
	// table.Row is just a shorthand for []interface{}
	t.AppendRow([]interface{}{300, "Tyrion", "Lannister", 5000})
	// time to take a peek
	t.SetCaption("Simple Table with 3 Rows.\n")
	t.SetStyle(table.StyleDouble)
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
	fmt.Println(t.Render())

}
