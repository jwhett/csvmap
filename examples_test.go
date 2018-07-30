package csvmap_test

import (
    "fmt"
    "os"
    "github.com/jwhett/csvmap"
)

func ExampleNewCsvMap_file() {
    f, _ := os.Open("test.csv")
    cm, _ := csvmap.NewCsvMap(f)
}

func ExampleAddRow() {
    for _, row := range cm.GetRows() {
        fmt.Println("Row:", row)
        // Output: "Row: [John Doe jdoe]
    }

    newRow := make([]string, 3)
    newRow[0] = "Jane"
    newRow[1] = "Doe"
    newRow[2] = "jdoe2"

    cm.AddRow(newRow)

    for _, row := range cm.GetRows() {
        fmt.Println("Row:", row)
        // Output: "Row: [John Doe jdoe]
        // Output: "Row: [Jane Doe jdoe2]
    }
}

func ExampleGetHeaders() {
    headers := cm.GetHeaders()
    for header, _ := range headers {
        fmt.Println("Header:", header)
    }
    // Output: "Header: firstname"
    // Output: "Header: lastname"
    // Output: "Header: id"
}

func ExampleGetRows() {
    rows := cm.GetRows()
    for _, row := range rows {
        fmt.Println("Row:", row)
    }
    // Output: "Row: [John Doe jdoe]"
}
