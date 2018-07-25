package csvmap

import (
    "io"
)

type CsvMap struct {
    headers map[string]int
    rows    [][]string
}

/*
func NewCsvMap(r io.Reader) (*CsvMap, error) {

}
*/

func (c *CsvMap) addRow(r []string) {
    // Helper to add rows
    c.rows = append(c.rows, r)
}

func (c *CsvMap) valuesByCol() {
    // Print values broken up by header
    for header, _ := range c.headers {
        fmt.Println("Header:", header)
        for _, row := range c.rows {
            fmt.Printf("\t%v\n", row[c.headers[header]])
        }
    }
}

func newCsvMap() CsvMap {
    // Initialize a CsvMap
    return CsvMap{headers: make(map[string]int), rows: make([][]string, 0, 1024)}
}

func readerFactory(infile string) *csv.Reader {
    // Create a CSV Reader
    // TODO handle file open error here
    csvFile, _ := os.Open(infile)
    return csv.NewReader(bufio.NewReader(csvFile))
}

func buildCsvMap(c *csv.Reader) CsvMap {
    // Create a CsvMap from the CSV file
    isFirst := true
    cm := newCsvMap()

    fmt.Println("Building CsvMap...")

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
