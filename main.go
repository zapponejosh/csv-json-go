package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type CSVImport [][]string

type JSONExport []map[string]string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFile(args []string) CSVImport {
	filename := args[0]

	raw, err := os.ReadFile(filename)
	check(err)
	reader := csv.NewReader(strings.NewReader(string(raw)))

	if len(args) > 1 { // default to comma if no arg
		switch delimeter := args[1]; delimeter {
		case "tab":
			reader.Comma = '\t'
		case "pipe":
			reader.Comma = '|'
		default:
		}
	}

	records, e := reader.ReadAll()
	check(e)

	return records
}

func processCSV(data CSVImport) JSONExport {
	processedData := make(JSONExport, len(data)-1)
	for i, row := range data {
		if i > 0 {
			newMap := make(map[string]string)
			for k, val := range row {
				newMap[data[0][k]] = val
			}
			processedData[i-1] = newMap
		}

	}

	return processedData
}

func main() {
	fmt.Println("running app...")
	args := os.Args[1:]
	csvData := getFile(args)

	data := processCSV(csvData)
	json, err := json.Marshal(data)
	check(err)

	os.WriteFile("output.json", json, 0644)
	fmt.Println("complete")

}
