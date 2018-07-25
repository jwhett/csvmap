package csvmap

import (
    "os"
    "bufio"
    "encoding/csv"
    "fmt"
    "io"
    "log"
)

type CsvMap struct {
    headers map[string]int
    rows    [][]string
}

//func NewCsvMap(r io.Reader) (*CsvMap, error) {
func NewCsvMap(s string) CsvMap {
    return buildCsvMap(readerFactory(s))

}

func (c *CsvMap) addRow(r []string) {
    // Helper to add rows
    c.rows = append(c.rows, r)
}

func (c *CsvMap) ValuesByCol() {
    // Print values broken up by header
    for header, _ := range c.headers {
        for _, row := range c.rows {
            fmt.Printf("\t%v\n", row[c.headers[header]])
        }
    }
}

func initCsvMap() CsvMap {
    // Initialize a CsvMap
    return CsvMap{headers: make(map[string]int), rows: make([][]string, 0, 1024)}
}

func readerFactory(infile string) *csv.Reader {
    // Create a CSV Reader
    // TODO handle file open error here
    csvFile, _ := os.Open(infile)
    return csv.NewReader(bufio.NewReader(csvFile))
}

func buildHeaders(r []string) map[string]int {
    // Build a map for the headers so we can pull
    // data from rows by header vs index
    headerMap := make(map[string]int)
    for i, headerName := range r {
        headerMap[headerName] = i
    }
    return headerMap
}

//func buildCsvMap(c *csv.Reader) CsvMap {
func buildCsvMap(c *csv.Reader) CsvMap {
    // Create a CsvMap from the CSV file
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

        cm.addRow(row)
    }
    return cm
}
