package csvmap_test

import (
    "os"
    "testing"
    "github.com/jwhett/csvmap"
)

func BenchmarkNewAndPrint(b *testing.B) {
    for i := 0; i < b.N; i++ {
        infile, _ := os.Open("testdata/test.csv")
        csvMap, _ := csvmap.NewCsvMap(infile)
        csvMap.PrintValuesByCol()
    }
}

func TestGetHeaders(t *testing.T) {
    infile, _ := os.Open("testdata/test.csv")
    csvMap, _ := csvmap.NewCsvMap(infile)
    headers := csvMap.GetHeaders()
    if len(headers) != 3 {
        t.Error("Header length incorrect.")
    }
    if headers["fname"] != 0 {
        t.Error("Headers are out of order.")
    }
}

func TestGetRows(t *testing.T) {
    infile, _ := os.Open("testdata/test.csv")
    csvMap, _ := csvmap.NewCsvMap(infile)
    rows := csvMap.GetRows()
    if len(rows) != 1 {
        t.Error("Failed to get correct number of rows.")
    }
}

func TestHeader(t *testing.T) {
    infile, _ := os.Open("testdata/test.csv")
    csvMap, _ := csvmap.NewCsvMap(infile)
    if csvMap.Header("fname") != 0 {
        t.Error("Failed to get header by name.")
    }
}

func TestGet(t *testing.T) {
    infile, _ := os.Open("testdata/test.csv")
    csvMap, _ := csvmap.NewCsvMap(infile)
    rows := csvMap.GetRows()
    if csvMap.Get("fname", rows[0]) != "john" {
        t.Error("Failed to get row item by header.")
    }
}
