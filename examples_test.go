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
        fmt.Println(row)
        // Output: [John Doe jdoe]
    }

    newRow := make([]string, 3)
    newRow[0] = "Jane"
    newRow[1] = "Doe"
    newRow[2] = "jdoe2"

    cm.AddRow(newRow)

    for _, row := range cm.GetRows() {
        fmt.Println(row)
        // Output: [John Doe jdoe]
        // Output: [Jane Doe jdoe2]
    }
}

func ExampleGetHeaders() {
    headers := cm.GetHeaders()
    for header, _ := range headers {
        fmt.Println(header)
    }
    // Output: firstname
    // Output: lastname
    // Output: id
}

func ExampleGetRows() {
    rows := cm.GetRows()
    for _, row := range rows {
        fmt.Println(row)
    }
    // Output: [John Doe jdoe]
}
