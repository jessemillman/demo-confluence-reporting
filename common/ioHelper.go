package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
)

// FileWriter handles which type of report to write
func FileWriter(r []ReportLine, s string) {
	switch s {
	case "csv":
		tryWriteCSV(r)
	case "json":
		tryWriteJSON(r)
	default:
		log.Fatal("File format not recognised")
	}
}

// TryWriteCSV tries to write the report to /output/results.csv as a CSV
func tryWriteCSV(r []ReportLine) {
	output := "/output/results.csv"

	// marshall the object to a csv string
	s, err := gocsv.MarshalString(r)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to marshall csv string"))
	}

	file, err := os.Create(output)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to write CSV file"))
	} else {
		file.WriteString(s)
	}
	file.Close()
	fmt.Println("Successfully Wrote file to /output/results.csv")
}

// TryWriteJSON tries to write the report to /output as JSON
func tryWriteJSON(r []ReportLine) {
	output := "/output/results.json"

	// marshall the object to a json string
	s, err := json.Marshal(r)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to marshall json string"))
	}

	fileErr := ioutil.WriteFile(output, s, 0644)
	if fileErr != nil {
		log.Fatal(errors.Wrap(err, "failed to write json file"))
	}
	fmt.Println("Successfully Wrote file to /output/results.json")
}
