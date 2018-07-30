package csvmap_test

import (
    "fmt"
    "os"
    "github.com/jwhett/csvmap"
)

func ExampleNewCsvMap_file() {
    f, _ := os.Open("test.csv")
    cm, _ := csvmap.NewCsvMap(f)
    cm.PrintValuesByCol()
}

func ExampleCsvMap_AddRow() {
    f, _ := os.Open("test.csv")
    cm, _ := csvmap.NewCsvMap(f)
    newRow := make([]string, 3)
    newRow[0] = "Jane"
    newRow[1] = "Doe"
    newRow[2] = "jdoe2"
    cm.AddRow(newRow)
}

func ExampleCsvMap_GetHeaders() {
    f, _ := os.Open("test.csv")
    cm, _ := csvmap.NewCsvMap(f)
    headers := cm.GetHeaders()
    for header, _ := range headers {
        fmt.Println(header)
    }
}

func ExampleCsvMap_GetRows() {
    f, _ := os.Open("test.csv")
    cm, _ := csvmap.NewCsvMap(f)
    rows := cm.GetRows()
    for _, row := range rows {
        fmt.Println(row)
    }
}
