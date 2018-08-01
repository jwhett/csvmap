// Package csvmap allows you to interact with CSV
// data by key vs index position. CSV --> CsvMap
package csvmap

import (
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "log"
)

// CsvMap provides a conventient interface for interacting
// with CSV data.
type CsvMap struct {
    headers map[string]int
    rows    [][]string
}

// NewCsvMap is your entry point for creating a map
// from CSV data.
func NewCsvMap(r io.Reader) (*CsvMap, error) {
    return buildCsvMap(readerFactory(r)), nil
}

// AddRow is a method for adding rows
// of data to the CsvMap struct.
func (c *CsvMap) AddRow(r []string) {
    c.rows = append(c.rows, r)
}

// GetHeaders returns a map of headers with
// index.
func (c CsvMap) GetHeaders() map[string]int {
    return c.headers
}

// GetRows returns a slice of slices (rows)
func (c CsvMap) GetRows() [][]string {
    return c.rows
}

// Header returns index location
func (c CsvMap) Header(h string) int {
    return c.headers[h]
}

// Get retuns row value that matches key
func (c CsvMap) Get(key string, row []string) string {
    return row[c.Header(key)]
}

// PrintValuesByCol is exposed for visually testing the mapped data.
func (c *CsvMap) PrintValuesByCol() {
    for header := range c.headers {
        fmt.Printf("%v:\n", header)
        for _, row := range c.rows {
            fmt.Printf("\t%v\n", row[c.headers[header]])
        }
    }
}

// Initialize a CsvMap with empty data.
func initCsvMap() *CsvMap {
    return &CsvMap{headers: make(map[string]int), rows: make([][]string, 0, 1024)}
}

// readerFactory will build and return CSV Readers.
func readerFactory(infile io.Reader) *csv.Reader {
    return csv.NewReader(bufio.NewReader(infile))
}

// buildHeaders will build a map for the header row so
// we can reference data from each row by header.
func buildHeaders(r []string) map[string]int {
    headerMap := make(map[string]int)
    for i, headerName := range r {
        headerMap[headerName] = i
    }
    return headerMap
}

// Create a CsvMap from the CSV file
func buildCsvMap(c *csv.Reader) *CsvMap {
    isFirst := true
    cm := initCsvMap()

    for {
        row, err := c.Read()
        if err == io.EOF {
            // End of file
            break
        }
        if err != nil {
            // We hit an error
            log.Fatal(err)
        }
        if isFirst {
            // we're at the header
            header := buildHeaders(row)
            cm.headers = header
            isFirst = false
            continue
        }

        cm.AddRow(row)
    }
    return cm
}
